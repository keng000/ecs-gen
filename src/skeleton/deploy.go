package skeleton

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

// Deploy render
func (s *Skeleton) Deploy(executable *DeployExecutable) error {
	rule, err := lbRule(executable)
	if err != nil {
		return err
	}

	executable.LBRule = rule
	for _, tmpl := range deploymentTemplates {
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(executable); err != nil {
			return err
		}
	}
	return nil
}

func lbRule(executable *DeployExecutable) (string, error) {
	chunk := ""
	for _, api := range executable.APIName {
		tmplVarFile, err := Assets.Open("/resource/tmpl/terraform/modules/alb/variables-elements.tmpl")
		if err != nil {
			return "", errors.Wrap(err, "failed to read template from Assets")
		}

		s, err := ioutil.ReadAll(tmplVarFile)
		if err != nil {
			return "", errors.Wrap(err, "failed to read Assets reader")
		}

		rule, err := renderString(string(s), struct{ APIName string }{APIName: api})
		if err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("Failed render lbRule: %s", api))
		}
		chunk += rule
	}
	return chunk, nil
}
