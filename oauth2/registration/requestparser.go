package registration

import (
	"fmt"
	"strings"
	"encoding/json"
	. "github.com/jhuntoo/go-connect/oauth2"
	jose "github.com/square/go-jose"
	"errors"
)


type Parser interface{
	Parse(json []byte) (*Client, error)
}

type RequestParser struct {
	
}

func NewRequestParser() *RequestParser {
	return & RequestParser{}
}


func (p * RequestParser)  Parse(jsonBody []byte) (*Client, error) {
	request := &Request{}
	
	if err := json.Unmarshal(jsonBody, &request); err != nil {
		return nil, errors.New("Failed to parse JSON in Registration Request")
	}
	
	
	client, err := mapToClient(request)
	
	return client, err
}

func mapToClient(request *Request) (*Client, error) {
	scopes := toArray(request.Scopes)
	
	var keyArray []jose.JsonWebKey
	var err error
	if keyArray, err = GetKeyArray(request.JwksRaw); err != nil {
		return nil, err
	}
	
	
	 
	return &Client {
		Name: request.ClientName,
		Uri: request.ClientUri,
		RedirectUris: request.RedirectUris,
		LogoUri: request.LogoUri,
		Contacts: request.Contacts,
		TermsOfServiceUri: request.TermsOfServiceUri,
		PolicyUri: request.PolicyUri,
		TokenEndPointAuthMethod: AuthMethod(request.TokenEndPointAuthMethod),
		Scopes: scopes,
		GrantTypes: request.GrantTypes,
		ResponseTypes: request.ResponseTypes,
		JwksUri: request.JwksUri,
		Jwks: keyArray,
		SoftwareId: request.SoftwareId,
		SoftwareVersion: request.SoftwareVersion,
		SoftwareStatementRawToken: request.SoftwareStatement, 
		
	}, nil
}

func GetKeyArray(jwkRawMessage *json.RawMessage) ([]jose.JsonWebKey, error) {
	if jwkRawMessage == nil { return make([]jose.JsonWebKey, 0), nil}
	var keyArrayMap map[string][]interface{}
	 var keyArray []interface{}
	 
	 
	if arrayJson, err := json.Marshal(jwkRawMessage); err != nil {
		return nil, errors.New("Failed to parse jwk set1")
	} else {
		if err := json.Unmarshal([]byte(arrayJson), &keyArrayMap); err != nil {
			fmt.Printf("Value: %v", string(arrayJson))
			return nil, errors.New("Failed to parse jwk set2")
		}
	}
	keyArray = keyArrayMap["keys"]
	keyTypedArray := make([]jose.JsonWebKey, 0)
	
	for _, key := range keyArray {
		if keyAsString, err := json.Marshal(key); err != nil {
			return nil, errors.New("Failed to parse jwk set3")
		} else {
			var jwk jose.JsonWebKey
			err := jwk.UnmarshalJSON([]byte(keyAsString))
			if err != nil {
				return  nil, errors.New("Failed parsing jwk set4")
			}
		
		
			keyTypedArray = append(keyTypedArray, jwk)
		}
		
	}
	return keyTypedArray, nil
}

func toArray(scopes string) []string {
	return strings.Split(scopes, " ")
}

