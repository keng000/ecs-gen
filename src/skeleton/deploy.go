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
	"{{ .APIName }}.health_check_interval" = 20
	"{{ .APIName }}.health_check_timeout"  = 19
	"{{ .APIName }}.healthy_threshold"     = 2
	"{{ .APIName }}.unhealthy_threshold"   = 2
	"{{ .APIName }}.min_capacity"          = 1
	"{{ .APIName }}.max_capacity"          = 5
	"{{ .APIName }}.cpu_high_statistic"    = "Average"
	"{{ .APIName }}.cpu_low_statistic"     = "Average"
	"{{ .APIName }}.cpu_high_threshold"    = 30
	"{{ .APIName }}.cpu_low_threshold"     = 10
	"{{ .APIName }}.scale_up_cooldown"     = 180
	"{{ .APIName }}.scale_down_cooldown"   = 300
	"{{ .APIName }}.deregistration_delay"  = 20
		`
		rule, err := renderString(s, struct{ APIName string }{APIName: api})
		if err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("Failed render lbRule: %s", api))
		}
		chunk += rule
	}
	return chunk, nil
}
