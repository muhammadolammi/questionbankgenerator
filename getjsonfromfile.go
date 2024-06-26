package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func getSubjectConfigFromFile(filePath string) (SubjectConfig, error) {

	data, err := os.ReadFile(filePath)

	if err != nil {
		return SubjectConfig{}, err
	}

	subjectConfig := SubjectConfig{}
	err = json.Unmarshal(data, &subjectConfig)
	if err != nil {
		return SubjectConfig{}, err
	}
	if subjectConfig.Config.Subject == "" {
		return SubjectConfig{}, errors.New("Your config file must include the subject")
	}
	if len(subjectConfig.Config.Topics) == 0 {
		return SubjectConfig{}, errors.New("Your config file must include one topic")
	}

	return subjectConfig, nil

}

func (config *Config) getSubjectConfigsFromFiles(filePaths []fs.DirEntry) ([]SubjectConfig, error) {
	subjectConfigs := []SubjectConfig{}
	for _, filePath := range filePaths {
		subjectConfig, err := getSubjectConfigFromFile(fmt.Sprintf("%s/configs/%s", config.WD, filePath.Name()))
		if err != nil {
			return []SubjectConfig{}, err
		}
		subjectConfigs = append(subjectConfigs, subjectConfig)
	}

	return subjectConfigs, nil
}
