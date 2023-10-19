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
    AppPort string `yaml:"app_port"`
    Adapter string
    Server  string
    Topic   string
    GroupId string `yaml:"group_id"`
  }

  Cb struct {
    AppPort string `yaml:"app_port"`
  }

  Sn struct {
    AppPort string `yaml:"app_port"`
  }

  Sp struct {
    AppPort string `yaml:"app_port"`
  }

  Schedule struct {
    AppPort string `yaml:"app_port"`
  }
}

// var EnvGlobal *Env

func EnvInit(pkgx *Pkg) {
  log.Println("-->>> Pkg Env init", pkgx)
  env, _ := ReadEnv(pkgx)
  // EnvGlobal = env
  pkgx.Env = env
}

func ReadEnv(pkgx *Pkg) (*Env, error) {
  _, filePath, _, _ := runtime.Caller(1)
  envFile := path.Join(path.Dir(filePath), "../config.yaml")
  buf, err := ioutil.ReadFile(envFile)
  pkgx.Logger.L.Infof("-->>> Env file read err: %+v, path: %+v", err, envFile)
  env := &Env{}
  err = yaml.Unmarshal(buf, env)
  pkgx.Logger.L.Infof("-->>> Env file Unmarshal err: %+v, path: %+v", err, envFile)
  return env, nil
}
