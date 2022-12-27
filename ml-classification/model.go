package classification

// Model represents a machine learning model
type Model interface {
	// Predict makes a prediction using the model
	Predict(input interface{}) (interface{}, error)
}
