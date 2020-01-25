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
