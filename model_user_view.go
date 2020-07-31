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

// Describes a user
type UserView struct {
	// Information about the organizations in which the user is a member
	Organizations []interface{} `json:"organizations,omitempty"`
	// Whether the user's email address has been verified
	Verified bool `json:"verified,omitempty"`
	// Avatar data representing the user's icon
	Avatar *interface{} `json:"avatar"`
	// true if this user data represents a guest user
	Anonymous bool `json:"anonymous"`
	// The list of external login providers against which the user has authenticated
	Logins []interface{} `json:"logins,omitempty"`
	// Whether the user has permission to create repositories
	CanCreateRepo bool `json:"can_create_repo,omitempty"`
	// If true, the user's namespace is the preferred namespace to display
	PreferredNamespace bool `json:"preferred_namespace,omitempty"`
	// The user's email address
	Email string `json:"email,omitempty"`
}