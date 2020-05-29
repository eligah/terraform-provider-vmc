/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// UpdateServiceRolesRequest struct for UpdateServiceRolesRequest
type UpdateServiceRolesRequest struct {
	ServiceDefinitionId string `json:"serviceDefinitionId"`
	// Roles to remove
	RolesToRemove []ServiceRoleBindingDto `json:"rolesToRemove,omitempty"`
	// Roles to add
	RolesToAdd []ServiceRoleBindingDto `json:"rolesToAdd,omitempty"`
}
