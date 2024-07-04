package main

type SubjectConfig struct {
	Config struct {
		ExamType string   `json:"exam_type"`
		Subject  string   `json:"subject"`
		Topics   []string `json:"topics"`
	} `json:"config"`
}

type SubjectQuestionBank struct {
	ExamType string           `json:"exam_type"`
	Subject  string           `json:"subject"`
	Topics   []TopicQuestions `json:"topics"`
}

type TopicQuestions struct {
	Topic     string `json:"topic"`
	Questions string `json:"questions"`
}
