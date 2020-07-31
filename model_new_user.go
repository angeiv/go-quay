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

// Fields which must be specified for a new user.
type NewUser struct {
	// The user's username
	Username string `json:"username"`
	// The user's password
	Password string `json:"password"`
	// The user's email address
	Email string `json:"email,omitempty"`
	// The optional invite code
	InviteCode string `json:"invite_code,omitempty"`
	// The (may be disabled) recaptcha response code for verification
	RecaptchaResponse string `json:"recaptcha_response,omitempty"`
}