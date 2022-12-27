package integrators

import (
	"encoding/json"
	"feed-processor/feedback"
	"feed-processor/integrators/models"
	classification "feed-processor/ml-classification"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	layout     = "2006-01-02"
	likeOffset = 45
)

var (
	// sensitiveInfoRegex is a regular expression that matches sensitive information
	sensitiveInfoRegex = regexp.MustCompile(`(?i)(password|credit card|social security|personal identification)`)
)

// DiscourseIntegrator represents an integrator for the Discourse feedback source.
type DiscourseIntegrator struct {
	BaseURL    string
	TenantID   int64
	TenantName string
}

// FetchData fetches feedback records from the Discourse source.
func (d *DiscourseIntegrator) FetchData(startTime time.Time, endTime time.Time) (interface{}, error) {
	// make HTTP GET request to the Discourse API to search for posts in the specified time range
	url := fmt.Sprintf("%s/search.json?page=1&q=after%%3A%s+before%%3A%s", d.BaseURL,
		startTime.Format(layout), endTime.Format(layout))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse response body into searchResult struct
	var searchResult models.SearchResult
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return nil, err
	}
	// create a slice to store the fetched feedback records
	var feedbacks []*models.DiscoursePostData
	// iterate over the search result posts and fetch the actual post using the post IDs
	for _, post := range searchResult.Posts {
		// make HTTP GET request to the Discourse API to get the individual post
		url := fmt.Sprintf("%s/t/%d/posts.json?post_ids[]=%d", d.BaseURL, post.TopicID, post.ID)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		// read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
		// parse response body into postResult struct
		var postResult models.PostResult
		if err := json.Unmarshal(body, &postResult); err != nil {
			return nil, err
		}
		// get the first post in the post result (should be the post we requested)
		postData := postResult.PostStream.UserPosts[0]
		fb := models.ToDiscourseDataModel(post, postData)
		// add the feedback record to the slice
		feedbacks = append(feedbacks, fb)
	}
	return feedbacks, nil
}

// ProcessData processes the raw data from the FetchData method and transforms it into a uniform internal structure
func (d *DiscourseIntegrator) ProcessData(rawData interface{}) ([]*feedback.Feedback, error) {
	// assert that the raw data is a slice of pointers to Feedback structs
	discourseData, ok := rawData.([]*models.DiscoursePostData)
	if !ok {
		return nil, fmt.Errorf("invalid raw data type: %T", rawData)
	}
	// create a map to store the processed feedback records
	processedFeedbacks := make(map[int64]*feedback.Feedback)
	// iterate over the slice of feedback records and apply the transformation
	for _, data := range discourseData {
		fb := &feedback.Feedback{
			ID:          data.ID,
			Username:    data.Username,
			Description: data.Description,
			Source:      d.getSource(),
			Tenant:      d.getTenant(),
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		}
		// add source-specific metadata
		fb.Metadata["topic_id"] = strconv.Itoa(int(data.TopicID))
		fb.Type = getType(data)
		// detect the language of the feedback using a language detection library or API
		language, err := detectLanguage(fb.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to detect language: %w", err)
		}
		fb.Language = language
		// remove sensitive information from the feedback description
		fb.Description = removeSensitiveInfo(fb.Description)
		fb.Data = []byte(fb.Description)

		// enrich the feedback with additional information from external sources
		user, err := getUser(data.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to get user: %w", err)
		}
		fb.Metadata["user"] = user

		topic, err := getTopic(fb.Metadata["topic_id"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to get topic: %w", err)
		}
		fb.Metadata["topic"] = topic

		// classify the feedback into categories using a machine learning model or a rules-based approach
		categories, err := classification.ClassifyFeedback(fb.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to classify feedback: %w", err)
		}
		fb.Categories = categories

		// add the processed feedback record to the map
		processedFeedbacks[data.ID] = fb
	}

	// convert the map of processed feedbacks to a slice
	result := make([]*feedback.Feedback, 0, len(processedFeedbacks))
	for _, fb := range processedFeedbacks {
		result = append(result, fb)
	}

	return result, nil
}

// StoreData stores the processed feedback records in the database.
func (d *DiscourseIntegrator) StoreData(records []*feedback.Feedback, db *database.DB) error {
	// TODO: implement this method
	// insert the processed feedback records into the database
	// return any error
	return nil
}

func (d *DiscourseIntegrator) getSource() *feedback.Source {
	return &feedback.Source{
		ID:   int64(Discourse),
		Name: Discourse.String(),
	}
}

func (d *DiscourseIntegrator) getTenant() *feedback.Tenant {
	return &feedback.Tenant{
		ID:   d.TenantID,
		Name: d.TenantName,
	}
}

// determine the type of feedback based on the number of likes
func getType(data *models.DiscoursePostData) string {
	if data.LikeCount > likeOffset {
		return "Popular"
	} else if data.LikeCount > 0 {
		return "Engaging"
	}
	return "Post"
}

// detectLanguage is a placeholder function that detects the language
// of the given text using a language detection library or API
func detectLanguage(text string) (string, error) {
	// TODO: Implement language detection using a library or API
	// for now, just return "en" (English) as the detected language
	return "en", nil
}

// removeSensitiveInfo removes sensitive information from the feedback description
func removeSensitiveInfo(description string) string {
	// use a regular expression to find and replace sensitive information
	return sensitiveInfoRegex.ReplaceAllString(description, "[redacted]")
}

// getUser retrieves additional information about the user from an external source
func getUser(username string) (interface{}, error) {
	// make an API call to the external source to retrieve the user information
	response, err := http.Get(fmt.Sprintf("https://meta.discourse.org/users/%s.json", username))
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	defer response.Body.Close()
	// decode the response into a map
	var user map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}
	return user["user"], nil
}

// getTopic retrieves additional information about the topic from an external source
func getTopic(topicID string) (map[string]interface{}, error) {
	// make an API call to the external source to retrieve the topic information
	response, err := http.Get(fmt.Sprintf("https://meta.discourse.org/t/%s.json", topicID))
	if err != nil {
		return nil, fmt.Errorf("failed to get topic: %w", err)
	}
	defer response.Body.Close()
	// decode the response into a map
	var topic map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&topic); err != nil {
		return nil, fmt.Errorf("failed to decode topic: %w", err)
	}
	return topic, nil
}
