package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	expectedName     = "Test Campaign"
	expectedContent  = "This is a test campaign"
	expectedContacts = []string{"test@example.com"}
	fake             = faker.New()
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := New(expectedName, expectedContent, expectedContacts)

	assert.NotNil(campaign)
	assert.NotEmpty(campaign.ID)
	assert.Equal(expectedName, campaign.Name)
	assert.Equal(expectedContent, campaign.Content)
	assert.Equal(len(expectedContacts), len(campaign.Contacts))
	assert.Equal(expectedContacts[0], campaign.Contacts[0].Email)
	assert.NotNil(campaign.CreatedAt)
}

func Test_NewCampaign_CreatedAtIsSet(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := New(expectedName, expectedContent, expectedContacts)

	assert.NotNil(campaign.CreatedAt)
	assert.Greater(time.Now(), campaign.CreatedAt)
}

func Test_NewCampaign_EmptyName(t *testing.T) {
	assert := assert.New(t)

	_, err := New("", expectedContent, expectedContacts)

	assert.ErrorContains(err, "Name is required")
}

func Test_NewCampaign_ShortName(t *testing.T) {
	assert := assert.New(t)

	_, err := New("abc", expectedContent, expectedContacts)

	assert.ErrorContains(err, "Name must be at least 5 characters long")
}

func Test_NewCampaign_LongName(t *testing.T) {
	assert := assert.New(t)

	_, err := New(fake.Lorem().Text(120), expectedContent, expectedContacts)

	assert.ErrorContains(err, "Name must be at most 100 characters long")
}

func Test_NewCampaign_EmptyContent(t *testing.T) {
	assert := assert.New(t)

	_, err := New(expectedName, "", expectedContacts)

	assert.ErrorContains(err, "Content is required")
}

func Test_NewCampaign_ShortContent(t *testing.T) {
	assert := assert.New(t)

	_, err := New(expectedName, "short", expectedContacts)

	assert.ErrorContains(err, "Content must be at least 10 characters long")
}

func Test_NewCampaign_LongContent(t *testing.T) {
	assert := assert.New(t)

	_, err := New(expectedName, fake.Lorem().Text(1040), expectedContacts)

	assert.ErrorContains(err, "Content must be at most 1024 characters long")
}

func Test_NewCampaign_EmptyContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := New(expectedName, expectedContent, []string{})

	assert.ErrorContains(err, "Contacts must be at least 1 characters long")
}

func Test_NewCampaign_InvalidContactEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := New(expectedName, expectedContent, []string{"invalid-email"})

	assert.ErrorContains(err, "Email must be a valid email")
}
