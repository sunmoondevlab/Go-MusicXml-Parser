package testutil

import (
	"os"
	"path/filepath"
)

func GetProjRootAbs() string {
	wd, _ := os.Getwd()
	gmb := "go.mod"
	gm := filepath.Join(wd, gmb)
	for {
		if FileExists(gm) {
			break
		}
		wd = filepath.Dir(wd)
		gm = filepath.Join(wd, gmb)
	}
	return wd
}
