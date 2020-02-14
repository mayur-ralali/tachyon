package objects

type OauthUserProfileObject struct {
	Success OauthUserProfileDetailObject `json:"success"`
}

type OauthUserProfileDetailObject struct {
	Code        string                 `json:"code"`
	Label       string                 `json:"label"`
	Message     string                 `json:"message"`
	UserSession OauthUserSessionObject `json:"user_session"`
}

type OauthUserSessionObject struct {
	User  OauthUserObject      `json:"user"`
	Token OauthUserTokenObject `json:"token"`
}

type OauthUserObject struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Email     string `json:"email"`
	Handphone string `json:"handphone"`
}

type OauthUserTokenObject struct {
	ClientID         string `json:"client_id"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	RefreshTokenUsed bool   `json:"refresh_token_used"`
	Logout           bool   `json:"logout"`
}
