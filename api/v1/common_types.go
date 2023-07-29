package v1

// +structType=atomic
type Build struct {
	// Git is a reference to a git repository that will be built within the cluster.
	// Built image will be set in the .spec.image field.
	Git *BuildGit `json:"git,omitempty"`
	// Upload is a signal that a local tar of the directory should be uploaded to be built as an image.
	Upload *BuildUpload `json:"upload,omitempty"`
}

// +structType=atomic
type BuildUpload struct {
	// Md5Checksum is the md5 checksum of the tar'd repo root requested to be uploaded and built.
	// +kubebuilder:validation:MaxLength=32
	// +kubebuilder:validation:MinLength=32
	// +kubebuilder:validation:Pattern="^[a-fA-F0-9]{32}$"
	Md5Checksum string `json:"md5checksum,omitempty"`
}

// +structType=atomic
type BuildGit struct {
	// URL to the git repository.
	// Example: https://github.com/my-username/my-repo
	URL string `json:"url"`
	// Path within the git repository referenced by url.
	Path string `json:"path,omitempty"`

	// Tag is the git tag to use. Choose either tag or branch.
	Tag string `json:"tag,omitempty"`
	// Branch is the git branch to use. Choose either branch or tag.
	Branch string `json:"branch,omitempty"`
}

type BuildStatus struct {
	// the Signed upload URL
	UploadURL string `json:"uploadURL,omitempty"`
	// Md5Checksum is the last md5 checksum that resulted in the successful creation of an UploadURL.
	// +kubebuilder:validation:MaxLength=32
	// +kubebuilder:validation:MinLength=32
	// +kubebuilder:validation:Pattern="^[a-fA-F0-9]{32}$"
	Md5Checksum string `json:"md5checksum,omitempty"`
}

type ObjectRef struct {
	// Name of Kubernetes object.
	Name string `json:"name"`

	// FUTURE: Possibly allow for cross-namespace references.
	// FUTURE: Possibly allow for cross-cluster references.
}

type Resources struct {
	//+kubebuilder:default:=2
	// CPU resources.
	CPU int64 `json:"cpu,omitempty"`

	//+kubebuilder:default:=10
	// Disk size in Gigabytes.
	Disk int64 `json:"disk,omitempty"`

	//+kubebuilder:default:=10
	// Memory is the amount of RAM in Gigabytes.
	Memory int64 `json:"memory,omitempty"`

	// GPU resources.
	GPU *GPUResources `json:"gpu,omitempty"`
}

type GPUType string

const (
	GPUTypeNvidiaTeslaT4 = GPUType("nvidia-tesla-t4")
	GPUTypeNvidiaL4      = GPUType("nvidia-l4")
)

type GPUResources struct {
	// Type of GPU.
	Type GPUType `json:"type,omitempty"`
	// Count is the number of GPUs.
	Count int64 `json:"count,omitempty"`
}

type ArtifactsStatus struct {
	URL string `json:"url,omitempty"`
}
