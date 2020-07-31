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

// Description of a new prototype
type NewPrototype struct {
	ActivatingUser *NewPrototypeActivatingUser `json:"activating_user,omitempty"`
	// Role that should be applied to the delegate
	Role string `json:"role"`
	Delegate *NewPrototypeDelegate `json:"delegate"`
}