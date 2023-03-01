package pkg

import (
	"io/ioutil"
	"log"
	"path"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Env struct {
	Broker struct {
		AppPort string		`yaml:"app_port"`
		Adapter string
		Server  string
		Topic   string
		GroupId string		`yaml:"group_id"`
	}
}

var EnvGlobal *Env

func EnvInit() {
	log.Println("-->>> Pkg Env init")
	env, _ := ReadEnv()
	EnvGlobal = env
}

func ReadEnv() (*Env, error) {
	_, filePath, _, _ := runtime.Caller(1)
	envFile := path.Join(path.Dir(filePath), "../config.yaml")
	//envFile := "../config.yaml"
	// absFilePath, _ := filepath.Abs(envFile)
	//absFilePath, _ := filepath.Rel(envFile)
	// log.Println(absFilePath)

	// buf, err := ioutil.ReadFile("../config.yaml")
	buf, err := ioutil.ReadFile(envFile)
	Logger.Infof("-->>> Env file read err: %+v, path: %+v", err, envFile)
	env := &Env{}
	err = yaml.Unmarshal(buf, env)
	Logger.Infof("-->>> Env file Unmarshal err: %+v, path: %+v", err, envFile)
	return env, nil
}
