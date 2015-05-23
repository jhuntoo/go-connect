package oauth2

type ClientType string
type AuthenticationType string

const (
	PUBLIC_CLIENT       ClientType = "public"
	CONFIDENTIAL_CLIENT            = "confidential"
)

const (
	BASIC_SECRET AuthenticationType = "basic"
	POST_SECRET                     = "post"
)

type Client struct {
	Id                            string
	Description                   string
	Secret                        string
	Type                          ClientType
	AuthType                      AuthenticationType
	Scopes                        []string
	GrantTypes                    []string
	RedirectUris                  []string
	AccessTokenValidityInSeconds  int
	RefreshTokenValidityInSeconds int
}
