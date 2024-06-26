package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	API_KEY string
	WD      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			log.Println("You do not have an .env file in this dir, make sure your linux enviroment have GEMINI_API_KEY exported, or create an env file ")
		}
		log.Println(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("error getting working  dir. err : %v", err)
		return
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Println("empty API_KEY")
		return

	}

	subjectConfigFiles, err := getSubjectConfigsFile()
	if err != nil {
		log.Printf("error getting config files. err : %v", err)
		return
	}
	config := Config{
		API_KEY: apiKey,
		WD:      wd,
	}
	subjectConfigs, err := config.getSubjectConfigsFromFiles(subjectConfigFiles)
	if err != nil {
		log.Printf("error getting subject configs. err : %v", err)
		return
	}
	fmt.Println(subjectConfigs[0:1])

	config.generate(subjectConfigs[0:1])

}
