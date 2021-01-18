package utils

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/alexNgari/flooksTest/models"
)

// ReadJSONFile reads a json file and returns the object
func ReadJSONFile(path string) (data *models.JSONData, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer jsonFile.Close()

	jsonObject, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	json.Unmarshal(jsonObject, data)

	return data, nil
}


// WriteToJSONFile writes data to a json file
func WriteToJSONFile(path string, borrower *models.Borrower, mutex *sync.Mutex) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	scoresFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	defer scoresFile.Close()

	scoresJSON, err := ioutil.ReadAll(scoresFile)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	var scores models.CreditScores

	json.Unmarshal(scoresJSON, &scores)
	
	creditScore := CalculateScore(borrower)
	
	found := false

	for _, score := range scores.CreditScores {
		if score.BorrowerID == creditScore.BorrowerID {
			score.CreditScore = creditScore.CreditScore
			found = true
		}
	}

	if !found {
		scores.CreditScores = append(scores.CreditScores, *creditScore)
	}

	return nil
}