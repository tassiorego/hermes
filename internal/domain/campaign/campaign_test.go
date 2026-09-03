package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	expectedName := "Test Campaign"
	expectedContent := "This is a test campaign"
	expectedContacts := []string{"test@example.com"}

	// Act
	campaign := New(expectedName, expectedContent, expectedContacts)

	// Assert
	assert.NotNil(campaign)
	assert.NotEmpty(campaign.ID)
	assert.Equal(expectedName, campaign.Name)
	assert.Equal(expectedContent, campaign.Content)
	assert.Equal(len(expectedContacts), len(campaign.Contacts))
	assert.Equal(expectedContacts[0], campaign.Contacts[0].Email)
	assert.NotNil(campaign.CreatedAt)
}

func Test_NewCampaign_CreatedAtIsSet(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	expectedName := "Test Campaign"
	expectedContent := "This is a test campaign"
	expectedContacts := []string{"test@example.com"}

	// Act
	campaign := New(expectedName, expectedContent, expectedContacts)

	// Assert
	assert.NotNil(campaign.CreatedAt)
	assert.Greater(time.Now(), campaign.CreatedAt)
}
