package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/sashabaranov/go-openai"
)

func OpenAiThreadMessagesResponseToGetAllMessageInThreadResponse(messages openai.MessagesList) []web.GetMessageInThreadResponse {
	var messagesResponse []web.GetMessageInThreadResponse

	for _, message := range messages.Messages {
		messageResponse := web.GetMessageInThreadResponse{
			ID:        message.ID,
			Role:      message.Role,
			Message:   message.Content[0].Text.Value,
			CreatedAt: message.CreatedAt,
		}

		messagesResponse = append(messagesResponse, messageResponse)

	}

	return messagesResponse
}