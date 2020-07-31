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

// Configuration for an export logs operation
type ExportLogs struct {
	// The e-mail address at which to e-mail a link to the exported logs
	CallbackEmail string `json:"callback_email,omitempty"`
	// The callback URL to invoke with a link to the exported logs
	CallbackUrl string `json:"callback_url,omitempty"`
}
