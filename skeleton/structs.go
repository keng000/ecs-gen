package skeleton

// Skeleton stores meta data of skeleton
type Skeleton struct {
	// Path is where skeleton is generated.
	Path string
}

// InitExecutable is
type InitExecutable struct {
	Project string
}

// APIExecutable is
type APIExecutable struct {
	Project string

	APIName string
}

// DumpExecutable is
type DumpExecutable struct {
	// Project is the name of the ecs project
	Project string

	// APIName is the name for auto scale generate
	APIName []string
}
