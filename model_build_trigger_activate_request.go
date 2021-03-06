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

type BuildTriggerActivateRequest struct {
	// The name of the robot that will be used to pull images.
	PullRobot string `json:"pull_robot,omitempty"`
	// Arbitrary json.
	Config *interface{} `json:"config"`
}
