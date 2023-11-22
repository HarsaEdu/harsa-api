package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (openAi *OpenAiImpl) CreateThread(studentName string, topic string) (string, error) {
	ctx := context.Background()

	content := fmt.Sprintf("Hi! I, %s, am seeking assistance on the topic of %s. Can you please provide guidance and support?", studentName, topic)

	threadTopic := openai.MessageRequest{
		Role: openai.ChatMessageRoleUser,
		Content: content,
	}

	thread, err := openAi.Client.CreateThread(ctx, openai.ThreadRequest{})
	if err != nil {
		return "", err
	}

	fmt.Println(thread)

	_, err = openAi.Client.CreateMessage(ctx, thread.ID, threadTopic)
	if err != nil {
		return "", err
	}


	return thread.ID, nil
}