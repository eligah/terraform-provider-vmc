/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// BaseUserWithProfileDto struct for BaseUserWithProfileDto
type BaseUserWithProfileDto struct {
	// The user ID
	UserId string `json:"userId,omitempty"`
	// The account identifier
	Acct string `json:"acct,omitempty"`
	UserProfile UserLocalePreferencesDto `json:"userProfile,omitempty"`
	// The domain of the user. (will be vmwareid incase the user is VMware ID user)
	Domain string `json:"domain,omitempty"`
	// The user's email
	Email string `json:"email,omitempty"`
	// The last name of the user
	LastName string `json:"lastName,omitempty"`
	// The user's Identity Provider ID
	IdpId string `json:"idpId,omitempty"`
	// Deprecated. The user's username (email address)
	Username string `json:"username,omitempty"`
	// User is inaccessible if it will not be able to login with it's VMware ID credentials. Since the user is federated and will be redirected to it's Identity Provider
	Accessible bool `json:"accessible,omitempty"`
	// The first name of the user
	FirstName string `json:"firstName,omitempty"`
}
