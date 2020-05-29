/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// AssignRolesToMembersRequest struct for AssignRolesToMembersRequest
type AssignRolesToMembersRequest struct {
	// IDs of the members to whom the new roles will be assigned
	Ids []string `json:"ids,omitempty"`
	// The service roles that will be assigned to the member
	ServiceRoles []ServiceRolesWithIdDto `json:"serviceRoles,omitempty"`
	// The organization roles that will be assigned to the member
	OrganizationRoles []BaseRoleBindingDto `json:"organizationRoles,omitempty"`
}
