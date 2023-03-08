package pkg

import (
	"log"

	"go.uber.org/zap"
)

// TODO: one package only have on init(), need to control to init() process and unify entry
// TODO: this is just a point, not the `point of struct`
var Pkgx *Pkg

func init() {
	log.Printf("-->>> Pkgx init %+v", Pkgx)
	pkgxTmp := Pkg{}
	LoggerInit(&pkgxTmp)
	EnvInit(&pkgxTmp)

	// Logger.Infof("EnvGlobal -->>> %+v", EnvGlobal)
	Pkgx = &pkgxTmp
	log.Printf("-->>> Pkgx init assign %+v", Pkgx)
	Pkgx.Logger.Infof("EnvGlobal -->>> %+v", Pkgx.Env)
}

type Pkg struct {
	Logger *zap.SugaredLogger
	Env    *Env
}
