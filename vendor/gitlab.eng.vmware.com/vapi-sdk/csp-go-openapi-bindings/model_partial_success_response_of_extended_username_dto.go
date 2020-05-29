/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// PartialSuccessResponseOfExtendedUsernameDto struct for PartialSuccessResponseOfExtendedUsernameDto
type PartialSuccessResponseOfExtendedUsernameDto struct {
	// The entities for which the operation fails
	Failed []ExtendedUsernameDto `json:"failed,omitempty"`
	// The entities for which the operation is successful
	Succeeded []ExtendedUsernameDto `json:"succeeded,omitempty"`
}
