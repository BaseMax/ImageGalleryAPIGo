package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Logger(r *http.Request) {
	log.Printf("IP: %s | URL: %s | User Agent: %s | Method: %s", r.RemoteAddr, r.URL.Path, r.UserAgent(), r.Method)
}

func CheckEnvVars(vars []string) (map[string]string, error) {
	envs := make(map[string]string)
	for _, v := range vars {
		val, exists := os.LookupEnv(v)
		if !exists {
			return nil, fmt.Errorf("environment variable '%s' is not set", v)
		}
		envs[v] = val
	}
	return envs, nil
}
