package openai

import "context"

func (openai *OpenAiImpl) ChatWithAssistant(threadId string, message string) (string, error) {
	ctx := context.Background()

	assistant, err := openai.Client.RetrieveAssistant(ctx, openai.OpenAIConfig.AssistantID)
	if err != nil {
		return "", err
	}

	_, err = openai.CreateMessageInThread(ctx, threadId, message)
	if err != nil {
		return "", err
	}

	_, err = openai.RunMessageInThread(ctx, assistant, threadId)
	if err != nil {
		return "", err
	}

	responseMessage, err := openai.GetResponseMessageInThread(ctx, threadId)
	if err != nil {
		return "", err
	}

	return responseMessage.Content[0].Text.Value, nil
}