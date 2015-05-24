package registration

import "encoding/json"

type Request struct {
	ClientName              string           `json:"client_name"`
	ClientUri               string           `json:"client_uri"`
	LogoUri                 string           `json:"logo_uri"`
	Contacts                string           `json:"contacts"`
	TermsOfServiceUri       string           `json:"tos_uri"`
	PolicyUri               string           `json:"policy_uri"`
	TokenEndPointAuthMethod string           `json:"token_endpoint_auth_method"`
	Scopes                  string           `json:"scope"`
	GrantTypes              []string         `json:"grant_types"`
	ResponseTypes           []string         `json:"response_types"`
	RedirectUris            []string         `json:"redirect_uris"`
	JwksUri                 string           `json:"jwks_uri"`
	JwksRaw                 *json.RawMessage `json:"jwks"`
	SoftwareId              string           `json:"software_id"`
	SoftwareVersion         string           `json:"software_version"`
	SoftwareStatement       string           `json:"software_statement"`
}
