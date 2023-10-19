package unify

import (
  "log"

  "github.com/halokid/NeonRabbit/pkg"
  "github.com/halokid/NeonRabbit/vo"
)

type Unify struct {
  Pkg *pkg.Pkg
  Vo  *vo.Vo
}

var Unifyx *Unify

func init() {
  unifyTmp := &Unify{}
  unifyTmp.Pkg = pkg.Pkgx
  unifyTmp.Vo = vo.Vox
  Unifyx = unifyTmp
  log.Printf("Unify init -->>> %+v", Unifyx)
}
