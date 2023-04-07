package utils

import (
	"fmt"
	"os"
)

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
