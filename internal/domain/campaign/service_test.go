package campaign

import (
	"hermes/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mockedRepository *MockRepository) Save(campaign *Campaign) error {
	args := mockedRepository.Called(campaign)
	return args.Error(0)
}

func TestService_CreateCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	newCampaignDTO := contract.CreateCampaignDTO{
		Name:     "Test Campaign",
		Content:  "This is a test campaign",
		Contacts: []string{"test@example.com"},
	}
	repositoryMock := new(MockRepository)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaignDTO.Name {
			return false
		} else if campaign.Content != newCampaignDTO.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaignDTO.Contacts) {
			return false
		} else if campaign.Contacts[0].Email != newCampaignDTO.Contacts[0] {
			return false
		}
		return true
	})).Return(nil)

	service := Service{
		repository: repositoryMock,
	}

	// Act
	id, err := service.Create(newCampaignDTO)
	// Assert
	assert.NoError(err)
	assert.NotEmpty(id)
	repositoryMock.AssertCalled(t, "Save", mock.AnythingOfType("*campaign.Campaign"))
}
