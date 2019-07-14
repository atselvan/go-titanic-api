package backend

import (
	"log"
	"os"
	"regexp"
)

func logger(message string) {
	logger := log.New(os.Stdout, "[INFO]: ", log.LstdFlags)
	logger.Println(message)
}

func HandleError(err error, errorMessage string) {
	if err != nil {
		logger(errorMessage)
		logger(err.Error())
	}
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
