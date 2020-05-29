/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// UserAndRolesDto struct for UserAndRolesDto
type UserAndRolesDto struct {
	// Organization Roles
	OrganizationRoles []OrganizationRoleBindingDto `json:"organizationRoles"`
	// The user ID
	UserId string `json:"userId,omitempty"`
	// Last Name
	LastName string `json:"lastName"`
	// The account identifier
	Acct string `json:"acct,omitempty"`
	// The user's email
	Email string `json:"email,omitempty"`
	// Username (Deprecated)
	Username string `json:"username"`
	// First Name
	FirstName string `json:"firstName"`
}
