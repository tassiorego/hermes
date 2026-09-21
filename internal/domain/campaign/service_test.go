package campaign

import (
	"errors"
	"hermes/internal/contract"
	internalerrors "hermes/internal/internal-errors"
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

func (mockedRepository *MockRepository) Get() ([]Campaign, error) {
	args := mockedRepository.Called()
	return args.Get(0).([]Campaign), args.Error(1)
}

var (
	newCampaignDTO = contract.CreateCampaignDTO{
		Name:     "Test Campaign",
		Content:  "This is a test campaign",
		Contacts: []string{"test@example.com"},
	}
	mockRepository = new(MockRepository)
	service        = Service{
		Repository: mockRepository,
	}
)

func TestService_CreateCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	mockRepository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
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

	// Act
	id, err := service.Create(newCampaignDTO)
	// Assert
	assert.NoError(err)
	assert.NotEmpty(id)
	mockRepository.AssertCalled(t, "Save", mock.AnythingOfType("*campaign.Campaign"))
}

func TestService_CreateCampaign_RepositoryError(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	newCampaignDTO.Name = ""
	// Act
	id, err := service.Create(newCampaignDTO)

	// Assert
	assert.NotNil(err)
	assert.EqualError(err, "Name is required")
	assert.Empty(id)
}

func TestService_CreateCampaign_RepositorySaveError(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	mockRepository = new(MockRepository)
	newCampaignDTO.Name = "Test Campaign"
	mockRepository.On("Save", mock.Anything).Return(errors.New("Error saving campaign"))
	service.Repository = mockRepository

	// Act
	id, err := service.Create(newCampaignDTO)

	// Assert
	assert.NotNil(err)
	assert.EqualError(err, internalerrors.InternalServerError.Error())
	assert.True(errors.Is(internalerrors.InternalServerError, err))
	assert.Empty(id)
}
