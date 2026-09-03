package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	campaign := New("Test Campaign", "This is a test campaign", []string{"test@example.com"})

	if campaign.ID != "1" {
		t.Errorf("Expected campaign ID to be '1', got '%s'", campaign.ID)
	}

	if campaign.Name != "Test Campaign" {
		t.Errorf("Expected campaign name to be 'Test Campaign', got '%s'", campaign.Name)
	}

	if campaign.Content != "This is a test campaign" {
		t.Errorf("Expected campaign content to be 'This is a test campaign', got '%s'", campaign.Content)
	}

	if len(campaign.Contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(campaign.Contacts))
	} else if campaign.Contacts[0].Email != "test@example.com" {
		t.Errorf("Expected contact email to be 'test@example.com', got '%s'", campaign.Contacts[0].Email)
	}

}
