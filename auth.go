package mvola

type AuthService service

type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type AuthRequest struct {
	GrantType string `url:"grant_type"`
	Scope     string `url:"scope"`
}

func (s *AuthService) GenerateToken(consumerKey, consumerSecret string) (*Auth, error) {
	var err error
	params := AuthRequest{
		GrantType: "client_credentials",
		Scope:     "EXT_INT_MVOLA_SCOPE",
	}
	res := new(Auth)
	s.client.base.SetBasicAuth(consumerKey, consumerSecret)
	s.client.base.Post("/token").BodyForm(params).Receive(res, err)
	return res, err
}
