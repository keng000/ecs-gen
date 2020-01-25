package skeleton

type (
	// Skeleton stores meta data of skeleton
	Skeleton struct {
		// Path is where skeleton is generated.
		Path string
	}

	// InitExecutable is
	InitExecutable struct {
		Project string
	}

	// APIExecutable is
	APIExecutable struct {
		Project string

		APIName string
	}

	// DeployExecutable is
	DeployExecutable struct {
		Project string

		Region string
	}
)
