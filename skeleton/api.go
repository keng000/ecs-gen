package skeleton

import "path/filepath"

// API is
func (s *Skeleton) API() error {
	for _, tmpl := range autoScaleAPITemplates {
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(s.Executable); err != nil {
			return err
		}
	}
	return nil
}
