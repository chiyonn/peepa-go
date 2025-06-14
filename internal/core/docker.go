package core

import (
	"os"
)

func MustReadSecret(key string) string {
	path := "/run/secrets/" + key

	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read secret: " + err.Error())
	}

	return string(data)
}

