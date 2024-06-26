package main

type SubjectConfig struct {
	Config struct {
		Subject string   `json:"subject"`
		Topics  []string `json:"topics"`
	} `json:"config"`
}

type QuestionBank struct {
	Subject   string         `json:"subject"`
	Questions map[int]string `json:"questions"`
}
