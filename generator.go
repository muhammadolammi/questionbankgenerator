package main

import (
	"context"
	"log"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func (config *Config) generate(subjectConfigs []SubjectConfig) {
	// log.Println(subjectConfigs)
	ctx := context.Background()

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	// Lets generate for each subject
	// for _, subjectConfig := range subjectConfigs {

	// 	subjectQuestionBank := generateQuestionBank(subjectConfig, ctx, model)
	// 	config.saveQuestionBank(subjectConfig, subjectQuestionBank)

	// }

	var wg sync.WaitGroup

	// Generate for each subject concurrently
	for _, subjectConfig := range subjectConfigs {
		wg.Add(1)
		go func(subjectConfig SubjectConfig) {
			defer wg.Done()
			subjectQuestionBank := generateQuestionBank(subjectConfig, ctx, model)
			config.saveQuestionBank(subjectConfig, subjectQuestionBank)
		}(subjectConfig)
	}

	// Wait for all go routines to finish
	wg.Wait()

}
