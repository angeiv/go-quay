/*
 * Quay Frontend
 *
 * This API allows you to perform many of the operations required to work with Quay repositories, users, and organizations. You can find out more at <a href=\"https://quay.io\">Quay</a>.
 *
 * API version: v1
 * Contact: support@quay.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package quay

// Description of a team
type TeamDescription struct {
	// Org wide permissions that should apply to the team
	Role string `json:"role"`
	// Markdown description for the team
	Description string `json:"description,omitempty"`
}
