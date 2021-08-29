package config

import (
	"log"
	"os"
	"path/filepath"
)

const (
	secretsJSONFile = ".secrets.json"
)

// Setup Инициализировать конфигурации с помощью переменных окружения.
func Setup() {
	if err := readEnv(DB, App, Server, Logger); err != nil {
		log.Fatalln(err)
	}
}

// loadEnv Выгрузить переменные оркужения из .secrets.json файда в текущую сессию.
// Перезаписываются уже существуюшие переменные окружения.
func LoadEnv() {
	var (
		dir string
		err error
	)

	if dir, err = os.Getwd(); err != nil {
		log.Fatalln(err)
	}

	if err = loadEnvJSON(filepath.Join(dir, secretsJSONFile)); err != nil {
		log.Println(err)
	}
}
