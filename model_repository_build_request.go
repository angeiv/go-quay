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

// Description of a new repository build.
type RepositoryBuildRequest struct {
	// Subdirectory in which the Dockerfile can be found. You can only specify this or dockerfile_path
	Subdirectory string `json:"subdirectory,omitempty"`
	// The URL of the .tar.gz to build. Must start with \"http\" or \"https\".
	ArchiveUrl string `json:"archive_url,omitempty"`
	// The tags to which the built images will be pushed. If none specified, \"latest\" is used.
	DockerTags []string `json:"docker_tags,omitempty"`
	// Username of a Quay robot account to use as pull credentials
	PullRobot string `json:"pull_robot,omitempty"`
	// The file id that was generated when the build spec was uploaded
	FileId string `json:"file_id,omitempty"`
	// Pass in the context for the dockerfile. This is optional.
	Context string `json:"context,omitempty"`
	// Path to a dockerfile. You can only specify this or subdirectory.
	DockerfilePath string `json:"dockerfile_path,omitempty"`
}
