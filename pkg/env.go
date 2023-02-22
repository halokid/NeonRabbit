package pkg

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Env struct {
	Broker struct {
		Adapter string
		Server  string
	}
}

var EnvGlobal *Env

func EnvInit() {
	log.Println("-->>> Pkg Env init")
	env, _ := ReadEnv()
	EnvGlobal = env
}

func ReadEnv() (*Env, error) {
	nowFolder := runtime.Caller(1)
	envFile := "../config.yaml"
	// absFilePath, _ := filepath.Abs(envFile)
	absFilePath, _ := filepath.Rel(envFile)
	// log.Println(absFilePath)

	// buf, err := ioutil.ReadFile("../config.yaml")
	buf, err := ioutil.ReadFile(absFilePath)
	Logger.Errorf("-->>> Env file read err: %+v, path: %+v", err, absFilePath)
	env := &Env{}
	err = yaml.Unmarshal(buf, env)
	Logger.Errorf("-->>> Env file Unmarshal err: %+v, path: %+v", err, absFilePath)
	return env, nil
}
