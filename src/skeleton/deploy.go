package skeleton

import "path/filepath"

// Deploy render
func (s *Skeleton) Deploy(executable DeployExecutable) error {
	for _, tmpl := range deploymentTemplates {
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(executable); err != nil {
			return err
		}
	}
	return nil
}
