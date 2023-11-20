package service

import "github.com/HarsaEdu/harsa-api/internal/model/web"

func (chatbotService *ChatbotServiceImpl) GetResponse(request *web.GetResponseRequest) (string, error) {
	err := chatbotService.Validate.Struct(request)
	if err != nil {
		return "", err
	}

	response, err := chatbotService.OpenAiClient.CreateCompletion(request.Message)
	if err != nil {
		return "", err
	}

	return response, nil
}