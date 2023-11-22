package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (openAi *OpenAiImpl) CreateThread(studentName string, topic string) (string, error) {
	ctx := context.Background()

	content := fmt.Sprintf("Hi! I, %s, am seeking assistance on the topic of %s. Can you please provide guidance and support?", studentName, topic)

	threadTopic := openai.ThreadMessage{
		Role: openai.ChatMessageRoleUser,
		Content: content,
	}

	thread, err := openAi.Client.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			threadTopic,
		},
	})
	if err != nil {
		return "", err
	}


	return thread.ID, nil
}