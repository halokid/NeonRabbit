package pkg

import (
    "io/ioutil"

    "github.com/halokid/ColorfulRabbit"
    "gopkg.in/yaml.v3"
)

type Env struct {
    Broker struct {
        Adapter string
        Server  string
    }
}

var EnvGlobal *Env

func init() {
    env, _ := ReadEnv()
    EnvGlobal = env
}

func ReadEnv() (*Env, error) {
    buf, err := ioutil.ReadFile("../config.yaml")
    ColorfulRabbit.CheckFatal(err, "")
    env := &Env{}
    err = yaml.Unmarshal(buf, env)
    ColorfulRabbit.CheckError(err, "")
    return env, nil
}
