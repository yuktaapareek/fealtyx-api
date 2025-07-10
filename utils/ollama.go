package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"fealtyx/models"
)

func GenerateSummary(student models.Student) (string, error) {
	url := "http://localhost:11434/api/generate"
	payload := map[string]interface{}{
		"model": "llama3",
		"prompt": "Generate a short and friendly summary for the following student:\n" +
			"Name: " + student.Name + "\n" +
			"Age: " + strconv.Itoa(student.Age) + "\n" +
			"Email: " + student.Email,
	}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode != 200 {
		return "", errors.New("Ollama API error")
	}

	if out, ok := result["response"].(string); ok {
		return out, nil
	}
	return "", errors.New("Invalid Ollama response")
}
