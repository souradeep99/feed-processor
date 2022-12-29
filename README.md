# Feedback Ingestion Project

This project is designed to ingest feedback records from various sources and store them in a uniform internal structure. The system supports heterogeneous feedback sources, such as Intercom, Playstore, Twitter, and Discourse Posts, and allows for both push and pull integration models. The system also supports metadata ingestion, with each source having different types of metadata values, such as app-version from Playstore or country from Twitter.

## Features

- Ingests feedback records from multiple sources
- Supports both push and pull integration models
- Ingests source-specific metadata
- Transforms incoming records into a uniform internal structure
- Stores records in a persistent storage solution
- Implements multi-tenancy support to isolate records from different tenants

## Requirements

- Go 1.x
- Dependencies (specified in [go.mod](go.mod))


## Sources

The following sources are currently supported:

- Intercom
- Playstore
- Twitter
- Discourse Posts
