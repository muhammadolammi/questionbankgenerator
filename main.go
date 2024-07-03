package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
)

type Config struct {
	API_KEY string
	WD      string
}

// TRY saving the generated question bank in memory
var generatedQuestionBanks = map[string]SubjectQuestionBank{}
var mutex sync.Mutex

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: questionbankgenerator {test|dev}")
	}
	err := godotenv.Load()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			log.Println("You do not have an .env file in this dir, make sure your linux enviroment have GEMINI_API_KEY/GEMINI_API_KEY_TEST exported, or create an env file ")
		}
		log.Println(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("error getting working  dir. err : %v", err)
		return
	}

	mode := os.Args[1]
	fmt.Println(mode)

	var apiKey string
	if mode == "dev" {
		apiKey = os.Getenv("GEMINI_API_KEY_DEV")
	} else if mode == "dep" {
		apiKey = os.Getenv("GEMINI_API_KEY")
	} else {
		// Actually this is redunctant or useless, since my run.sh wont run without either of the two correct mode
		log.Println("Unknown Mode provided")
		return
	}
	if apiKey == "" {
		log.Println("empty API_KEY")
		return

	}

	subjectConfigFiles, err := getSubjectConfigsFile()
	if err != nil {
		log.Printf("error getting config files. err : %v", err)
		return
	}
	if len(subjectConfigFiles) == 0 {
		log.Println("No config file in directory:")
		return

	}
	config := Config{
		API_KEY: apiKey,
		WD:      wd,
	}
	// Set up signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Received signal: %v. Saving generated data...", sig)
		config.saveQuestionBanks()
		os.Exit(1)
	}()
	subjectConfigs, err := config.getSubjectConfigsFromFiles(subjectConfigFiles)
	if err != nil {
		log.Printf("error getting subject configs. err : %v", err)
		return
	}
	fmt.Println(subjectConfigs[0:1])

	config.generate(subjectConfigs[0:1])

}
