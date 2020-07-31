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

// Description of updates for an existing organization
type UpdateOrg struct {
	// Whether the organization desires to receive emails for invoices
	InvoiceEmail bool `json:"invoice_email,omitempty"`
	// Organization contact email
	Email string `json:"email,omitempty"`
	// The number of seconds for tag expiration
	TagExpirationS int32 `json:"tag_expiration_s,omitempty"`
}