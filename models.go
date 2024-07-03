package main

type SubjectConfig struct {
	Config struct {
		Subject string   `json:"subject"`
		Topics  []string `json:"topics"`
	} `json:"config"`
}

type SubjectQuestionBank struct {
	Subject string           `json:"subject"`
	Topics  []TopicQuestions `json:"topics"`
}

type TopicQuestions struct {
	Topic     string `json:"topic"`
	Questions string `json:"questions"`
}
