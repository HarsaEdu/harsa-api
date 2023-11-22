package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CreateThreadRequestToUserChatTopicDomain(request *web.CreateThreadRequest, threadId string) *domain.UserChatTopic {
	return &domain.UserChatTopic{
		ID: threadId,
		Topic: request.Topic,
	}
}