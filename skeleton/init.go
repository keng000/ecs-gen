package skeleton

import (
	"path/filepath"
)

// Init render the base tmpl file and generate required files.
func (s *Skeleton) Init() error {
	for _, tmpl := range baseTemplates {
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(s.Executable); err != nil {
			return err
		}
	}
	return nil
}
