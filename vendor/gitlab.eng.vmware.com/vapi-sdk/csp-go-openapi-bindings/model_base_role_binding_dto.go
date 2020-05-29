/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// BaseRoleBindingDto struct for BaseRoleBindingDto
type BaseRoleBindingDto struct {
	// Membership type of the member in the organization. DIRECT: if the member roles were assigned directly. GROUP: if the member roles were assigned through custom or enterprise group.
	MembershipType string `json:"membershipType,omitempty"`
	// The role name
	Name string `json:"name,omitempty"`
}
