package skeleton

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Template stores meta data of template
type Template struct {
	// Path is the path to this template.
	Path string

	// OutputPathTmpl is the template for outputPath.
	OutputPathTmpl string
}

// Exec is
func (t *Template) Exec(data interface{}) error {
	tmplHTTPFile, err := Assets.Open(t.Path)
	if err != nil {
		return errors.Wrap(err, "failed to read template from Assets")
	}

	tmpl, err := ioutil.ReadAll(tmplHTTPFile)
	if err != nil {
		return errors.Wrap(err, "failed to read Assets reader")
	}

	outputPath, err := renderString(t.OutputPathTmpl, data)
	if err != nil {
		return err
	}

	outputDir := filepath.Dir(outputPath)
	_, err = os.Stat(outputDir)
	if err != nil {
		log.Infof("dir created: %s", outputDir)
		_ = os.MkdirAll(outputDir, 0755)
	}

	fp, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	if err := render(string(tmpl), fp, data); err != nil {
		return err
	}

	if err := fp.Close(); err != nil {
		return err
	}

	return nil
}

func renderString(tmpl string, data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := render(tmpl, &buf, data); err != nil {
		return "", errors.Wrap(err, "failed to render into string")
	}
	return buf.String(), nil
}

func render(tmpl string, fp io.Writer, data interface{}) error {
	name := ""
	tmplEx, err := template.New(name).Parse(tmpl)
	if err != nil {
		return err
	}

	if err := tmplEx.Execute(fp, data); err != nil {
		return err
	}

	return nil
}
