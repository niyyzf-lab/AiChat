package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Data struct {
	Model      string  `json:"model"`
	Created_at string  `json:"created_at"`
	Message    Message `json:"message"`
	Done       string  `json:"done"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type MessageData struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Greet returns a greeting for the given name
func (a *App) PostOllma(model string, prompt string) string {
	var input []Message
	fmt.Println(prompt)
	_ = json.Unmarshal([]byte(prompt), &input)
	messageData := MessageData{
		Model:    model,
		Messages: input,
	}
	runtime.EventsEmit(a.ctx, "onData", "%$%"+input[len(input)-1].Content)
	data, _ := json.Marshal(messageData)
	//post请求http://localhost:11434/api/generate
	res, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(data))
	if err != nil {
	}
	defer res.Body.Close()
	//流式读取
	var buffer [1024]byte
	var returnValue string
	for {
		n, err := res.Body.Read(buffer[0:])
		var payload Data
		_ = json.Unmarshal(buffer[0:n], &payload)
		t := payload.Message.Content
		returnValue += t
		runtime.EventsEmit(a.ctx, "onData", t)
		//读取done字段
		if err != nil && err.Error() == "EOF" {
			break
		} else if err != nil {
			break
		}
	}
	return returnValue
}
