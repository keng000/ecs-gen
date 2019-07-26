package skeleton

import "path/filepath"

// Skeleton stores meta data of skeleton
type Skeleton struct {
	// Path is where skeleton is generated.
	Path string

	Executable *Executable
}

// Executable store the executable meta information
type Executable struct {
	// Project is the name of the ecs project
	Project string

	// Region is the aws region where the project will deployed
	Region string
}

// BaseTemplates is
var BaseTemplates = []Template{
	{"skeleton/resource/tmpl/terraform/modules/vpc/main.tf.tmpl", "modules/vpc/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl", "modules/vpc/outputs.tf"},
	{"skeleton/resource/tmpl/terraform/modules/vpc/variables.tf.tmpl", "modules/vpc/variables.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/main.tf.tmpl", "modules/sg/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl", "modules/sg/outputs.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/variables.tf.tmpl", "modules/sg/variables.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecs/main.tf.tmpl", "modules/ecs/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl", "modules/ecs/outputs.tf"},
}

// Base render the base tmpl file and generate required files.
func (s *Skeleton) Base() error {
	for _, tmpl := range BaseTemplates {
		// println("----------------------")
		// println(tmpl.Path)
		// println(tmpl.OutputPathTmpl)
		// fmt.Printf("%+v\n", s.Executable)
		// println("----------------------")
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(s.Executable); err != nil {
			return err
		}
	}
	return nil
}

// func (s *Skeleton) generate(templates []Template, data interface{}) error {

// }
