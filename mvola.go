package mvola

import "github.com/dghubble/sling"

const SANDBOX_URL = "https://devapi.mvola.mg"
const PRODUCTION_URL = "https://api.mvola.mg"

type Client struct {
	base        *sling.Sling
	common      service
	Auth        *AuthService
	Transaction *TransactionService
}

type service struct {
	client *Client
}

func NewClient(baseUrl string) *Client {
	base := sling.New().Base(baseUrl)
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Auth = (*AuthService)(&c.common)
	c.Transaction = (*TransactionService)(&c.common)
	return c
}

func SetAccessToken(c *Client, accessToken string) {
	c.base.Set("Authorization", "Bearer "+accessToken)
}

func SetOptions(c *Client, opt Options) {
	c.base.Set("Version", opt.Version)
	c.base.Set("X-CorrelationID", opt.CorrelationID)
	/*
		c.base.Set("UserLanguage", opt.UserLanguage)
		c.base.Set("UserAccountIdentifier", opt.UserAccountIdentifier)
		c.base.Set("PartnerName", opt.PartnerName)
	*/
	if opt.CallbackURL != nil {
		c.base.Set("X-Callback-URL", *opt.CallbackURL)
	}
	c.base.Set("Cache-Control", "no-cache")
}
