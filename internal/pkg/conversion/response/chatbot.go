package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/sashabaranov/go-openai"
)

func UserChatDomainToGetAllThreadByUserIdResponse(topics []domain.UserChatTopic) []web.GetAllThreadByUserIdResponse {
	var topicsResponse []web.GetAllThreadByUserIdResponse

	for _, topic := range topics {
		topicResponse := web.GetAllThreadByUserIdResponse{
			ID: topic.ID,
			UserId: int(topic.UserID),
			Topic: topic.Topic,
			CreatedAt: topic.CreatedAt,
			UpdatedAt: topic.UpdatedAt,
		}

		topicsResponse = append(topicsResponse, topicResponse)
	}

	return topicsResponse
}

func OpenAiThreadMessagesResponseToGetAllMessageInThreadResponse(messages openai.MessagesList) []web.GetMessageInThreadResponse {
	var messagesResponse []web.GetMessageInThreadResponse

	for idx, message := range messages.Messages {
		if idx != len(messages.Messages)-1 {
			messageResponse := web.GetMessageInThreadResponse{
				ID:        message.ID,
				Role:      message.Role,
				Message:   message.Content[0].Text.Value,
				CreatedAt: message.CreatedAt,
			}
			messagesResponse = append(messagesResponse, messageResponse)
		}

	}

	return messagesResponse
}