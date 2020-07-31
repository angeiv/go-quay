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

// Information about the user or team to which the rule grants access
type NewPrototypeDelegate struct {
	// Whether the delegate is a user or a team
	Kind string `json:"kind"`
	// The name for the delegate team or user
	Name string `json:"name"`
}
