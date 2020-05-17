package providers

import (
	"context"
	"fmt"

	"github.com/oauth2-proxy/oauth2-proxy/pkg/apis/sessions"
	"github.com/oauth2-proxy/oauth2-proxy/pkg/requests"
)

// NextcloudProvider represents an Nextcloud based Identity Provider
type NextcloudProvider struct {
	*ProviderData
}

var _ Provider = (*NextcloudProvider)(nil)

// NewNextcloudProvider initiates a new NextcloudProvider
func NewNextcloudProvider(p *ProviderData) *NextcloudProvider {
	p.ProviderName = "Nextcloud"
	return &NextcloudProvider{ProviderData: p}
}

// GetEmailAddress returns the Account email address
func (p *NextcloudProvider) GetEmailAddress(ctx context.Context, s *sessions.SessionState) (string, error) {
	json, err := requests.New(p.ValidateURL.String()).
		WithContext(ctx).
		WithHeaders(getOIDCHeader(s.AccessToken)).
		Do().
		UnmarshalJSON()
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}

	email, err := json.Get("ocs").Get("data").Get("email").String()
	return email, err
}
