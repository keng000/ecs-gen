package skeleton

import (
	"fmt"
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
		s := `
	"{{ .APIName }}.enable_autoscale"      = false
	"{{ .APIName }}.health-check-interval" = 20
	"{{ .APIName }}.health-check-timeout"  = 19
	"{{ .APIName }}.healthy-threshold"     = 2
	"{{ .APIName }}.unhealthy-threshold"   = 2
	"{{ .APIName }}.min-capacity"          = 1
	"{{ .APIName }}.max-capacity"          = 5
	"{{ .APIName }}.cpu-high-statistic"    = "Average"
	"{{ .APIName }}.cpu-low-statistic"     = "Average"
	"{{ .APIName }}.cpu-high-threshold"    = 30
	"{{ .APIName }}.cpu-low-threshold"     = 10
	"{{ .APIName }}.scale-up-cooldown"     = 180
	"{{ .APIName }}.scale-down-cooldown"   = 300
		`
		rule, err := renderString(s, struct{ APIName string }{APIName: api})
		if err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("Failed render lbRule: %s", api))
		}
		chunk += rule
	}
	return chunk, nil
}
