package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	expectedName     = "Test Campaign"
	expectedContent  = "This is a test campaign"
	expectedContacts = []string{"test@example.com"}
)

func TestNewCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	campaign, _ := New(expectedName, expectedContent, expectedContacts)

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

	// Act
	campaign, _ := New(expectedName, expectedContent, expectedContacts)

	// Assert
	assert.NotNil(campaign.CreatedAt)
	assert.Greater(time.Now(), campaign.CreatedAt)
}

func Test_NewCampaign_EmptyName(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	_, err := New("", expectedContent, expectedContacts)

	// Assert
	assert.ErrorContains(err, "Name is required")
}

func Test_NewCampaign_EmptyContent(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	_, err := New(expectedName, "", expectedContacts)

	// Assert
	assert.ErrorContains(err, "Content is required")
}
func Test_NewCampaign_EmptyContacts(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	_, err := New(expectedName, expectedContent, []string{})

	// Assert
	assert.ErrorContains(err, "At least one contact is required")
}
