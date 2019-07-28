package skeleton

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/keng000/ecs-gen/utils"
)

const envFile string = ".ecs-gen.json"

// Skeleton stores meta data of skeleton
type Skeleton struct {
	// Path is where skeleton is generated.
	Path string

	Executable *Executable
}

// Executable store the executable meta information
type Executable struct {
	// Project is the name of the ecs project
	Project string

	// Region is the aws region where the project will deployed
	Region string
}

// BaseTemplates is a list of template files when it will render in init
var BaseTemplates = []Template{
	{"skeleton/resource/tmpl/terraform/modules/vpc/main.tf.tmpl", "modules/vpc/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl", "modules/vpc/outputs.tf"},
	{"skeleton/resource/tmpl/terraform/modules/vpc/variables.tf.tmpl", "modules/vpc/variables.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/main.tf.tmpl", "modules/sg/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl", "modules/sg/outputs.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/variables.tf.tmpl", "modules/sg/variables.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecs/main.tf.tmpl", "modules/ecs/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl", "modules/ecs/outputs.tf"},
}

// Base render the base tmpl file and generate required files.
func (s *Skeleton) Base() error {
	for _, tmpl := range BaseTemplates {
		tmpl.OutputPathTmpl = filepath.Join(s.Path, tmpl.OutputPathTmpl)
		if err := tmpl.Exec(s.Executable); err != nil {
			return err
		}
	}

	if err := writeEnvInfo(s); err != nil {
		return err
	}
	return nil
}

func writeEnvInfo(s *Skeleton) error {
	environments := filepath.Join(s.Path, envFile)
	bytesData, err := json.Marshal(&s.Executable)
	if err != nil {
		return err
	}

	var data map[string]string
	if err := json.Unmarshal(bytesData, &data); err != nil {
		return err
	}
	if err := utils.AppendWrite(environments, data); err != nil {
		return err
	}

	return nil
}

func searchEnv() (string, error) {
	curPath, err := filepath.Abs("./")
	if err != nil {
		return "", err
	}

	for {
		exist, err := exists(envFile, curPath)
		if err != nil {
			return "", err
		}
		if exist {
			// .ecs-gen.json file found
			break
		}
		if curPath == "/" {
			return "", errors.New(".ecs-gen.json not found. please run ecs-gen init before")
		}

		curPath = filepath.Dir(curPath)
	}

	envFilePath := filepath.Join(curPath, envFile)
	if err != nil {
		panic(err)
	}
	return envFilePath, nil
}

func exists(filename, dstPath string) (bool, error) {
	files, err := ioutil.ReadDir(dstPath)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if file.Name() == filename {
			return true, nil
		}
	}

	return false, nil
}
