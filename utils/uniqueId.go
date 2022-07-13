package utils

import "github.com/docker/distribution/uuid"

func GetUniqueId() (id string) {
	return uuid.Generate().String()
}
