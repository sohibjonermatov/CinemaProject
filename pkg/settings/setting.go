package settings

import (
	"encoding/json"
	"log"
	"os"
	"secondProject/models"
)

var Config models.Config

func Setup(F string) {
	byteValue, err := os.ReadFile(F)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	err = json.Unmarshal(byteValue, &Config)

	if err != nil {
		log.Fatalf("%v", err)
		return
	}
}
