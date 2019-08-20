package skeleton

import "path/filepath"

// API is
func (s *Skeleton) API(data map[string]string) error {
	for _, tmpl := range autoScaleAPITemplates {
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(data); err != nil {
			return err
		}
	}
	return nil
}
