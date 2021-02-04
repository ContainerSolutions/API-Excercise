package endpoint

import "gitlab.com/DGuedes/API-Exercise/storage/postgres"

// ServerConfig is an abstraction for the server configuration.
// To build a new instance, use `endpoint.BuildSvrConfig(env)`.
type ServerConfig struct {
	PostgresConfig postgres.Config
	Port           string
	TemplatePath   string
}

// BuildSvrConfig builds a config to be used by the server depending on the
// given environment. If "env" is production, it returns a config to be used on
// Kubernetes. If "dev", it returns a config to be used in a docker-compose
// environment.
func BuildSvrConfig(env string) ServerConfig {
	return ServerConfig{
		PostgresConfig: postgres.BuildConfig(env),
		Port:           "8080",
		TemplatePath:   getTemplatePath(env),
	}
}

// Docker-compose mounts this repo as a volume under /opt, so that changes to
// the templates are reflected without needing any image rebuilding. However,
// `prod` uses /src as a result of the Dockerfile copy instruction.
func getTemplatePath(env string) string {
	switch env {
	case "prod":
		return "/src/"

	default:
		return "/opt/"
	}
}
