package registration_test

import (
//	"encoding/json"
//	"fmt"
	. "github.com/jhuntoo/go-connect/oauth2/registration"
	. "github.com/jhuntoo/go-connect/oauth2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Requestparser", func() {
	
		var (
			parser *RequestParser
		)
		
		BeforeEach(func() {
			parser = NewRequestParser()
		})
		
		Describe("Given a JSON request document with all fields", func() {
			
			reqestBody := `{
		      "redirect_uris":["https://client.example.org/callback","https://client.example.org/callback2"],
		      "client_name":"My Example Client",
			  "client_name#ja-Jpan-JP": "\u30AF\u30E9\u30A4\u30A2\u30F3\u30C8\u540D",
			  "client_uri":"https://client.example.org",
			  "tos_uri":"https://client.example.org/tos",
			  "policy_uri":"https://client.example.org/policy",
			  "scope": "a b",
			  "grant_types":["implicit","password"],
		      "response_types":["token","code"],
		      "token_endpoint_auth_method":"client_secret_basic",
		      "logo_uri":"https://client.example.org/logo.png",
		      "jwks_uri":"https://client.example.org/my_public_keys.jwks",
			  "software_id":"1",
		      "software_version": "1.0.0",
			  "jwks":{"keys":[{ "e": "AQAB",  "n": "nj", "kty" : "RSA"}]},			 
			  "software_statement":"eyJhbGciOiJSUzI1NiJ9.eyJzb2Z0d2FyZV9pZCI6IjROUkIxLTBYWkFCWkk5RTYtNVNNM1IiLCJjbGllbnRfbmFtZSI6IkV4YW1wbGUgU3RhdGVtZW50LWJhc2VkIENsaWVudCIsImNsaWVudF91cmkiOiJodHRwczovL2NsaWVudC5leGFtcGxlLm5ldC8ifQ.GHfL4QNIrQwL18BSRdE595T9jbzqa06R9BT8w409x9oIcKaZo_mt15riEXHazdISUvDIZhtiyNrSHQ8K4TvqWxH6uJgcmoodZdPwmWRIEYbQDLqPNxREtYn05X3AR7ia4FRjQ2ojZjk5fJqJdQ-JcfxyhK-P8BAWBd6I2LLA77IG32xtbhxYfHX7VhuU5ProJO8uvu3Ayv4XRhLZJY4yKfmyjiiKiPNe-Ia4SMy_d_QSWxskU5XIQl5Sa2YRPMbDRXttm2TfnZM1xx70DoYi8g6czz-CPGRi4SW_S2RKHIJfIjoI3zTJ0Y2oe0_EJAiXbL6OyF9S5tKxDXV8JIndSA"
			  
		     }`
						
			request, err := parser.Parse([]byte(reqestBody))
			
			It("should not error", func() {
				Expect(err).To(BeNil())
			})
			It("should parse RedirectUris", func() {
				Expect(request.RedirectUris).To(HaveLen(2))
				Expect(request.RedirectUris[0]).To(Equal("https://client.example.org/callback"))
				Expect(request.RedirectUris[1]).To(Equal("https://client.example.org/callback2"))
			})
			It("should parse client_name", func() {
				Expect(request.ClientName).To(Equal("My Example Client"))
			})
			It("should parse client_uri", func() {
				Expect(request.ClientUri).To(Equal("https://client.example.org"))
			})
			It("should parse tos_uri", func() {
				Expect(request.TermsOfServiceUri).To(Equal("https://client.example.org/tos"))
			})
			It("should parse policy_uri", func() {
				Expect(request.PolicyUri).To(Equal("https://client.example.org/policy"))
			})
			It("should parse scope", func() {
				Expect(request.Scopes).To(Equal("a b"))
			})
			It("should parse grant_types", func() {
				Expect(request.GrantTypes).To(HaveLen(2))
				Expect(request.GrantTypes[0]).To(Equal("implicit"))
				Expect(request.GrantTypes[1]).To(Equal("password"))
			})
			It("should parse token_endpoint_auth_method", func() {
				Expect(request.TokenEndPointAuthMethod).To(Equal(string(CLIENT_SECRET_BASIC)))
			})
			It("should parse logo_uri", func() {
				Expect(request.LogoUri).To(Equal("https://client.example.org/logo.png"))
			})
			It("should parse jwks_uri", func() {
				Expect(request.JwksUri).To(Equal("https://client.example.org/my_public_keys.jwks"))
			})
			It("should parse software_id", func() {
				Expect(request.SoftwareId).To(Equal("1"))
			})
			It("should parse software_version", func() {
				Expect(request.SoftwareVersion).To(Equal("1.0.0"))
			})
			It("should parse software_statement", func() {
				Expect(request.SoftwareStatement).To(Equal("eyJhbGciOiJSUzI1NiJ9.eyJzb2Z0d2FyZV9pZCI6IjROUkIxLTBYWkFCWkk5RTYtNVNNM1IiLCJjbGllbnRfbmFtZSI6IkV4YW1wbGUgU3RhdGVtZW50LWJhc2VkIENsaWVudCIsImNsaWVudF91cmkiOiJodHRwczovL2NsaWVudC5leGFtcGxlLm5ldC8ifQ.GHfL4QNIrQwL18BSRdE595T9jbzqa06R9BT8w409x9oIcKaZo_mt15riEXHazdISUvDIZhtiyNrSHQ8K4TvqWxH6uJgcmoodZdPwmWRIEYbQDLqPNxREtYn05X3AR7ia4FRjQ2ojZjk5fJqJdQ-JcfxyhK-P8BAWBd6I2LLA77IG32xtbhxYfHX7VhuU5ProJO8uvu3Ayv4XRhLZJY4yKfmyjiiKiPNe-Ia4SMy_d_QSWxskU5XIQl5Sa2YRPMbDRXttm2TfnZM1xx70DoYi8g6czz-CPGRi4SW_S2RKHIJfIjoI3zTJ0Y2oe0_EJAiXbL6OyF9S5tKxDXV8JIndSA"))
			})
//			It("should parse jwks", func() {
//				str, _ := json.Marshal(client.Jwks[0])
//				Expect(client.Jwks).To(HaveLen(1))
//				fmt.Printf("Key: %v", client.Jwks[0].Key)
//				fmt.Printf("KeyId: %v", client.Jwks[0].KeyID)
//				fmt.Printf("alg: %v", client.Jwks[0].Algorithm)
//				fmt.Printf("Key: %v", client.Jwks[0].Key)
//				fmt.Printf("Str: %v", string(str))
//				Expect(client.Jwks[0].KeyID).To(HaveLen(1))
//			})
			
		})
})
