package campaign

import (
	"testing"

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
	assert.Equal("1", campaign.ID)
	assert.Equal(expectedName, campaign.Name)
	assert.Equal(expectedContent, campaign.Content)
	assert.Equal(len(expectedContacts), len(campaign.Contacts))
	assert.Equal(expectedContacts[0], campaign.Contacts[0].Email)

}
