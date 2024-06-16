package handlers

import (
    "context"
    "net/http"
    "os"

    openai "github.com/sashabaranov/go-openai"
)

// 実際のAPIキーなしでテストするための模擬レスポンス
func MockChatGPTHandler(w http.ResponseWriter, r *http.Request) {
    mockResponse := "This is a mocked response from ChatGPT"
    w.Write([]byte(mockResponse))
}

// // handleChatgpt は ChatGPT API へのリクエストを処理します。
func ChatGPTHandler(w http.ResponseWriter, r *http.Request) {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        MockChatGPTHandler(w, r)
        return
    }

    client := openai.NewClient(apiKey)

    prompt := "Hello, how can I assist you today?"

    resp, err := client.Completion(context.Background(), openai.CompletionRequest{
        Model:  "text-davinci-003",
        Prompt: prompt,
    })

    if err != nil {
        http.Error(w, "Failed to get response from ChatGPT", http.StatusInternalServerError)
        return
    }

    w.Write([]byte(resp.Choices[0].Text))
}
