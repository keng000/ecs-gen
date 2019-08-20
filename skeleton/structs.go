package skeleton

// Skeleton stores meta data of skeleton
type Skeleton struct {
	// Path is where skeleton is generated.
	Path string
}

// Executable store the executable meta information
type Executable struct {
	// Project is the name of the ecs project
	Project string

	// Region is the aws region where the project will deployed
	Region string

	// APIName is the name for auto scale generate
	APIName string
}

// DumpExecutable is
type DumpExecutable struct {
	// Project is the name of the ecs project
	Project string

	// Region is the aws region where the project will deployed
	Region string

	// APIName is the name for auto scale generate
	APIName []string
}
