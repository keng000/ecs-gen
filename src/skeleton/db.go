package skeleton

import "path/filepath"

// DB is
func (s *Skeleton) DB(executable *DBExecutable) error {
	for _, tmpl := range replicateDBTemplate {
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(executable); err != nil {
			return err
		}
	}
	return nil
}
