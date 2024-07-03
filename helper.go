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

func generateQuestionBank(subjectConfig SubjectConfig, ctx context.Context, model *genai.GenerativeModel) {
	startTime := time.Now()
	var topics []TopicQuestions

	for _, topic := range subjectConfig.Config.Topics {
		log.Printf("Generating questions for subject %v, topic %v", subjectConfig.Config.Subject, topic)

		topicQuestions := generateQuestionOnTopic(model, subjectConfig.Config.Subject, topic, ctx)
		time.Sleep(time.Second * 10) // Consider adding a constant or a parameter for this delay

		log.Printf("Done generating questions for subject %v, topic %v", subjectConfig.Config.Subject, topic)

		topicQuestionsData := TopicQuestions{
			Topic:     topic,
			Questions: topicQuestions,
		}
		topics = append(topics, topicQuestionsData)

		currentQuestionBank := SubjectQuestionBank{
			Subject: subjectConfig.Config.Subject,
			Topics:  topics,
		}

		// Save the current question bank to memory
		mutex.Lock()
		generatedQuestionBanks[subjectConfig.Config.Subject] = currentQuestionBank
		mutex.Unlock()
	}

	log.Println("Total time spent:", time.Since(startTime))
}

func (config *Config) saveQuestionBanks() {
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

	for _, questionBank := range generatedQuestionBanks {
		filePath := fmt.Sprintf("%s/%s.json", savedir, questionBank.Subject)
		filedata, err := json.MarshalIndent(questionBank, "", "   ")
		if err != nil {
			log.Println("Error marshalling file data:", err)
			return
		}
		log.Printf("saving.... for subject %v", questionBank.Subject)

		if err := os.WriteFile(filePath, filedata, 0644); err != nil {
			log.Println("Error saving chat:", err)
			return
		}
		log.Printf("done saving.... for subject %v", questionBank.Subject)
	}

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

func generateQuestionOnTopic(model *genai.GenerativeModel, subject, topic string, ctx context.Context) string {

	prompt := fmt.Sprintf(`Following the JAMB %v syllabus Generate a 200  multiple choice question bank on topic %s. Each question should have four options (a, b, c, d), specify the correct answer, and provide a brief explanation for the correct answer. it must be a 200 question bank `, subject, topic)
	log.Println(prompt)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Println(err)
	}
	formattedResponse := formatResponse(resp)
	log.Println(formattedResponse)
	return (formattedResponse)

}

// func (config *Config) saveGeneratedData() {
// 	savedir := fmt.Sprintf("%s/jsonoutput", config.WD)
// 	if _, err := os.Stat(savedir); os.IsNotExist(err) {
// 		if err := os.MkdirAll(savedir, 0755); err != nil {
// 			log.Fatalf("Error creating save directory: %s", err)
// 		}
// 	}

// 	log.Println(generatedQuestionBanks)
// 	for subject, questionBank := range generatedQuestionBanks {
// 		filePath := fmt.Sprintf("%s/%s.json", savedir, subject)
// 		filedata, err := json.Marshal(questionBank)
// 		if err != nil {
// 			log.Println("Error marshalling file data:", err)
// 			continue
// 		}
// 		log.Printf("saving.... for subject %v", subject)

// 		if err := os.WriteFile(filePath, filedata, 0644); err != nil {
// 			log.Println("Error saving chat:", err)
// 			continue
// 		}
// 		log.Printf("done saving.... for subject %v", subject)
// 	}
// }
