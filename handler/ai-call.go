package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/TranThang-2804/ai-blog-writer/backend/ai"
	"github.com/TranThang-2804/ai-blog-writer/backend/shared/log"
)

type AIRequestDTO struct {
	OutputFormat ai.OutputType `json:"output_format"`
	Prompt       string        `json:"prompt"`
}

func CallAI(w http.ResponseWriter, r *http.Request) {
	var aiRequestDTO AIRequestDTO
	// Unmarshal the JSON body
	body, err := io.ReadAll(r.Body)
	err = json.Unmarshal(body, &aiRequestDTO)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
    log.Logger.Error("Error unmarshalling request", "error", err)
		return
	}

	response, err := ai.GetAiResponseChain(aiRequestDTO.Prompt, aiRequestDTO.OutputFormat)
	if err != nil {
		log.Logger.Error("Error getting AI response", "error", err)
	}

	// Wait for the client to disconnect
	w.Write([]byte(response))
}
