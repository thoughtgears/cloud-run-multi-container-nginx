package config

type Config struct {
	ProjectID string `envconfig:"PROJECT_ID" required:"true"`
}
