package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	imgData1, err := os.ReadFile("../../images/turtle1.png")
	if err != nil {
		log.Fatal(err)
	}

	imgData2, err := os.ReadFile("../../images/turtle2.png")
	if err != nil {
		log.Fatal(err)
	}

	parts := []llms.ContentPart{
		llms.BinaryPart("image/png", imgData1),
		llms.BinaryPart("image/png", imgData2),
		llms.TextPart("Describe the difference between these two pictures, with scientific detail"),
	}

	content := []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: parts,
		},
	}

	resp, err := llm.GenerateContent(ctx, content, llms.WithModel("gemini-pro-vision"))
	if err != nil {
		log.Fatal(err)
	}

	bs, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(bs))
}
