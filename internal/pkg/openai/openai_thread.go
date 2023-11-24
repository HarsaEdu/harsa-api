package openai

import (
	"context"
	"fmt"
	"time"

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

func (openAi *OpenAiImpl) GetAllMessagesInThread(threadId string, limit int, after string, before string) (openai.MessagesList, error) {
	ctx := context.Background()

	order := "desc"

	messages, err := openAi.Client.ListMessage(ctx, threadId, &limit, &order, &after, &before)
	if err != nil {
		return openai.MessagesList{}, err
	}

	return messages, nil
}

func (openAi *OpenAiImpl) CreateMessageInThread(ctx context.Context, threadId string, message string) (openai.Message, error) {
	messageContent := openai.MessageRequest{
		Role: openai.ChatMessageRoleUser,
		Content: message,
	}

	userMessage, err := openAi.Client.CreateMessage(ctx, threadId, messageContent)
	if err != nil {
		return openai.Message{}, err
	}

	return userMessage, nil
}

func (openAi *OpenAiImpl) RunMessageInThread(ctx context.Context, assistant openai.Assistant, threadId string) (openai.Run, error) {
	request := openai.RunRequest{
		AssistantID: assistant.ID,
		Model: &assistant.Model,
		Instructions: assistant.Instructions,
	}

	runRequest, err := openAi.Client.CreateRun(ctx, threadId, request)
	if err != nil {
		return openai.Run{}, err
	}

	var run openai.Run

	for {
		run, err := openAi.Client.RetrieveRun(ctx, runRequest.ThreadID, runRequest.ID)
		if err != nil {
			return openai.Run{}, err
		}

		if run.Status != "in_progress" {
			break
		}

		time.Sleep(2 * time.Second)
	}

	return run, nil
}

func (openAi *OpenAiImpl) GetResponseMessageInThread(ctx context.Context, threadId string) (openai.Message, error) {
	limit := 1
	order := "desc"
	after := ""
	before := ""

	responseMessages, err := openAi.Client.ListMessage(ctx, threadId, &limit, &order, &after, &before)
	if err != nil {
		return openai.Message{}, err
	}

	responseMessage, err := openAi.Client.RetrieveMessage(ctx, threadId, *responseMessages.FirstID)
	if err != nil {
		return openai.Message{}, err
	}

	return responseMessage, nil
}