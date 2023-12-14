package config

// Config holds the common resource config fields
type Config struct {
	Tags interface{} `json:"tags"`
	Name string      `json:"name"`
}

// AWSResourceConfig helps define type and name for sub-resources if needed
type AWSResourceConfig struct {
	Resource interface{}
	Metadata map[string]interface{}
	Name     string
	Type     string
}
