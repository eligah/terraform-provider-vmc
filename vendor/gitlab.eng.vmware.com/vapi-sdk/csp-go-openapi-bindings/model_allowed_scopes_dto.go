/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// AllowedScopesDto struct for AllowedScopesDto
type AllowedScopesDto struct {
	// Service scopes grouped by service definition
	ServicesScopes []ServiceScopesDto `json:"servicesScopes,omitempty"`
	// General scopes (openid for example)
	GeneralScopes []string `json:"generalScopes,omitempty"`
	OrganizationScopes OrganizationScopesDto `json:"organizationScopes,omitempty"`
}
