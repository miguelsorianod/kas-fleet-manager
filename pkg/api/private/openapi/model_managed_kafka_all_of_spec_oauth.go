/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ManagedKafkaAllOfSpecOauth struct for ManagedKafkaAllOfSpecOauth
type ManagedKafkaAllOfSpecOauth struct {
	ClientId               string `json:"clientId,omitempty"`
	ClientSecret           string `json:"clientSecret,omitempty"`
	TokenEndpointURI       string `json:"tokenEndpointURI,omitempty"`
	JwksEndpointURI        string `json:"jwksEndpointURI,omitempty"`
	ValidIssuerEndpointURI string `json:"validIssuerEndpointURI,omitempty"`
	UserNameClaim          string `json:"userNameClaim,omitempty"`
	TlsTrustedCertificate  string `json:"tlsTrustedCertificate,omitempty"`
	CustomClaimCheck       string `json:"customClaimCheck,omitempty"`
}
