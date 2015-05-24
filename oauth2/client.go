package oauth2

import (
	jose "github.com/square/go-jose"
)

type ClientType string
type AuthMethod string

const (
	PUBLIC_CLIENT       ClientType = "public"
	CONFIDENTIAL_CLIENT            = "confidential"
)

const (
	NONE                AuthMethod = "none" // The client is public
	CLIENT_SECRET_BASIC AuthMethod = "client_secret_basic"
	CLIENT_SECRET_POST  AuthMethod = "client_secret_post"
)

type Client struct {
	Id                            string
	Secret                        string
	Name                          string
	Uri                           string
	LogoUri                       string
	Contacts                      string
	TermsOfServiceUri             string
	PolicyUri                     string
	TokenEndPointAuthMethod       AuthMethod
	Scopes                        []string
	GrantTypes                    []string
	ResponseTypes                 []string
	RedirectUris                  []string
	JwksUri                       string
	Jwks                          []jose.JsonWebKey
	SoftwareId                    string
	SoftwareVersion               string
	SoftwareStatementRawToken     string
	AccessTokenValidityInSeconds  int
	RefreshTokenValidityInSeconds int
}
