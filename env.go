package env

import (
	"github.com/joho/godotenv"
	"log"
)

var envs map[string]string

func Load(filenames ...string) {
	data, err := godotenv.Read(filenames...)
	if err != nil {
		log.Fatal("Error loading env file")
		return
	}

	envs = data
}

func Get(k string) string {
	if v, ok := envs[k]; ok {
		return v
	}

	return ""
}

func Set(k, v string) {
	envs[k] = v
}
