package models

type Function struct {
	Name string `json:"name,omitempty"`
	// This will detremine type of function like SOURCE_URL, SOURCE_BLOB,
	// DOCKER_IMAGE, INBUILT etc
	Type string `json:"type,omitempty"`

	// source file: This is final function file cached from url or
	SourceFile string `json:"sourcefile,omitempty"`

	// URL for source file, that will be downloaded
	SourceURL string `json:"sourceurl,omitempty"`

	// Source code as strng blob
	SourceBlob []byte `json:"sourceblob,omitempty"`

	// Language identifier example golang, python, php etc
	SourceLang string `json:"sourcelang,omitempty"`

	// Docker image which will be use to build execute functions
	BaseImage string `json:"baseimage,omitempty"`

	// Agrgument to be passed during build command
	BuildArgs []string `json:"buildargs,omitempty"`

	// Additonal params to be passed during function execution
	RunParams []string `json:"runparams,omitempty"`

	// Include folders
	IncludeDir []string `json:"includedir,omitempty"`

	// Target Directory for mounting source & include
	TargetDir string `json:"targetdir,omitempty"`

	// CacheDir: directory on docker host in which sourceblob/url will be cached.
	CacheDir string `json:"cachedir,omitempty"`

	//Namespace : logical namespace name to isolate functions for user/projects
	Namespace string `json:"namespace,omitempty"`

	// version of function
	Version string `json:"version,omitempty"`
}

const (
	FUNCTION_TYPE_BLOB   = "FUNCTION_TYPE_BLOB"   // Blob to gen source file
	FUNCTION_TYPE_URL    = "FUNCTION_TYPE_URL"    // download from url to sourcefile
	FUNCTION_TYPE_FILE   = "FUNTION_TYPE_FILE"    // Assume file exist locally??
	FUNCTION_TYPE_DOCKER = "FUNCTION_TYPE_DOCKER" // Exec docker not source
)

type FunctionResponse struct {
	// Exit code of docker run
	ExitCode int `json: "exitcode,omitempty"`
	// std out from container run
	StdOut string `json: "stdout,omitempty"`
	// std err from container run
	StdErr string `json: "stderr,omitempty"`
	// Error from connector
	Error error `json:"stderr,omitempty"`
	// Message where above not relevant
	Message string `json:"message,omitempty"`
}
