package skeleton

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
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
		return err
	}

	tmpl, err := ioutil.ReadAll(tmplHTTPFile)
	if err != nil {
		return err
	}

	outputPath, err := renderPath(t.OutputPathTmpl, data)
	if err != nil {
		return err
	}

	outputDir := filepath.Dir(outputPath)
	_, err = os.Stat(outputDir)
	if err != nil {
		log.Printf("dir created: %s", outputDir)
		_ = os.MkdirAll(outputDir, 0755)
	}

	fp, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	filepath.Dir(t.OutputPathTmpl)

	if err := render(string(tmpl), fp, data); err != nil {
		return err
	}

	if err := fp.Close(); err != nil {
		return err
	}

	return nil
}

func renderPath(tmpl string, data interface{}) (string, error) {
	var outputPathBuf bytes.Buffer
	if err := render(tmpl, &outputPathBuf, data); err != nil {
		return "", err
	}
	return outputPathBuf.String(), nil
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
