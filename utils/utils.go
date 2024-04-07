package utils

import "github.com/aidarkhanov/nanoid/v2"

func GenerateId() string {
	id, _ := nanoid.New()
	return id
}