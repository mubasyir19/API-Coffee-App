package helpers

import "github.com/google/uuid"

func GenerateCodeOrder() string {
	return "BB-" + uuid.New().String()
}
