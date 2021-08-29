package config

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// loadEnvJSON Выгрузить переменные оркужения из json файда в текущую сессию.
// Перезаписываются уже существуюшие переменные окружения.
func loadEnvJSON(filePath string) (err error) {
	var data []byte

	if data, err = os.ReadFile(filePath); err != nil {
		return
	}

	var envs map[string]string
	if err = json.Unmarshal(data, &envs); err != nil {
		return
	}

	for key, value := range envs {
		if err = os.Setenv(key, value); err != nil {
			return
		}
	}

	return
}

// readEnv Заполнить струкруры переменными окружения.
func readEnv(configs ...interface{}) (err error) {
	for _, config := range configs {
		if err = envconfig.Process("", config); err != nil {
			return
		}
	}

	return
}
