package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
)

func generateQuestionBank(subjectConfig SubjectConfig, ctx context.Context, model *genai.GenerativeModel) QuestionBank {
	questions := make(map[int]string)
	timespent := time.Now()

	for i := 0; i < 10000; i++ {
		topic := subjectConfig.Config.Topics[i%len(subjectConfig.Config.Topics)]

		log.Printf("generating.... for subject %v, for topic %v", subjectConfig.Config.Subject, topic)
		question := generateQuestionOnTopic(model, topic, ctx)
		appendToQuestion(questions, question)

		// time.Sleep(time.Second * 10)
		log.Printf("done generating.... for subject %v, for topic %v", subjectConfig.Config.Subject, topic)
		log.Printf("%v questions generated", i+1)

	}
	log.Println("generated in", time.Since(timespent))

	questionBank := QuestionBank{
		Subject:   subjectConfig.Config.Subject,
		Questions: questions,
	}

	return questionBank

}
func (config *Config) saveQuestionBank(subjectConfig SubjectConfig, questionBank QuestionBank) {
	savedir := fmt.Sprintf("%s/jsonoutput", config.WD)
	if _, err := os.Stat(savedir); os.IsNotExist(err) {
		if err := os.MkdirAll(savedir, 0755); err != nil {
			log.Fatalf("Error creating save directory: %s", err)
		}
	}
	ouputdir := fmt.Sprintf("%s/output", config.WD)
	// Make sure the output dir is created fro migration
	if _, err := os.Stat(ouputdir); os.IsNotExist(err) {
		if err := os.MkdirAll(ouputdir, 0755); err != nil {
			log.Fatalf("Error creating save directory: %s", err)
		}
	}
	filePath := fmt.Sprintf("%s/%s.json", savedir, subjectConfig.Config.Subject)
	filedata, err := json.Marshal(questionBank)
	if err != nil {
		log.Println("Error marshalling file data:", err)
		return
	}
	log.Printf("saving.... for subject %v", subjectConfig.Config.Subject)

	if err := os.WriteFile(filePath, filedata, 0644); err != nil {
		log.Println("Error saving chat:", err)
		return
	}
	log.Printf("done saving.... for subject %v", subjectConfig.Config.Subject)
}

func appendToQuestion(m map[int]string, value string) {
	maxKey := -1
	for k := range m {
		if k > maxKey {
			maxKey = k
		}
	}
	nextKey := maxKey + 1
	m[nextKey] = value
}

func formatResponse(resp *genai.GenerateContentResponse) string {
	var formattedContent strings.Builder
	if resp != nil && resp.Candidates != nil {
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					formattedContent.WriteString(fmt.Sprintf("%v", part))
				}
			}
		}
	}

	return formattedContent.String()
}

func generateQuestionOnTopic(model *genai.GenerativeModel, topic string, ctx context.Context) string {

	prompt := fmt.Sprintf("Generate a multiple choice question on %s. The question should have four options (a, b, c, d), specify the correct answer, and provide a brief explanation for the correct answer.", topic)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Println(err)
	}
	formattedResponse := formatResponse(resp)
	return (formattedResponse)

}
