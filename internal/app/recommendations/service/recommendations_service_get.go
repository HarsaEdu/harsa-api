package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (recommendationsService *RecommendationsServiceImpl) GetRecommendations(request *web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error) {
	err := recommendationsService.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	existingUser, err := recommendationsService.UserRepository.UserAvailableByID(request.UserId)
	if err != nil {
		return nil, fmt.Errorf("error when get user : %s", err.Error())
	}
	
	if existingUser == nil {
		return nil, fmt.Errorf("error when get user : user not found")
	}


	response, err := recommendationsService.RecommendationsApi.GetRecommendations(request)
	if err != nil {
		return nil, fmt.Errorf("error when get recommendations : %s", err.Error())
	}

	return response, nil
}

func (recommendationsService *RecommendationsServiceImpl) GetRecommendationsForInstructor() error {
	topInterest, _ := recommendationsService.InterestRepository.GetTopInterests(5)
	if topInterest == nil {
		return fmt.Errorf("error when get top interests : no interests found")
	}

	instructors, _ := recommendationsService.UserRepository.GetUsersRegistrationToken(2)
	if instructors == nil {
		return fmt.Errorf("error when get instructor : no instructor found")
	}

	var interests string

	for _, interest := range topInterest {
		interests += fmt.Sprintf("%s, ", interest.Category.Name)
	}

	message := fmt.Sprintf("Top interests this week : %s", interests)

	systemInstruction := fmt.Sprintf("You are the HarsaEdu system to providing a weekly notification to instructors.Recommend an instructor to create a course title based on the following data: %s. Encourage them to design a course that aligns with the interests and needs of our users, also only put HarsaEdu Team as a sender and dont put any Subject", interests)

	chatResponse, err := recommendationsService.OpenAi.GetChatCompletion(message, systemInstruction)
	if err != nil {
		return fmt.Errorf("error when get chat completion : %s", err.Error())
	}

	notification := web.NotificationMultiCast{
		Title: "Recommendation This Week",
		Message: chatResponse,
	}

	var userIDs []uint

	for _, instrutor := range instructors {
		notification.RegistrationToken = append(notification.RegistrationToken, instrutor.RegistrationToken)
		userIDs = append(userIDs, instrutor.UserID)
	}

	notificationDomain := conversion.NotificationMultiCastRequestToNotificationDomain(notification, userIDs)

	err = recommendationsService.NotificationRepository.CreateMany(notificationDomain)

	if notification.RegistrationToken == nil {
		return fmt.Errorf("error when get instructor registration token : no instructor registration token found")
	}
	
	recommendationsService.Firebase.SendNotificationMulticast(&notification)

	return nil 
}