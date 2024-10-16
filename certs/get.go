package certs

import (
	"log"

	"github.com/kyaxcorp/go-helper/config"
	"github.com/kyaxcorp/go-helper/filesystem"
	fsPath "github.com/kyaxcorp/go-helper/filesystem/path"
	"github.com/kyaxcorp/go-helper/folder"
)

func GetCertsFullPath() string {
	var err error = nil
	certsPath := config.GetConfig().CertsPath
	certsPath, err = fsPath.GenRealPath(certsPath, true)

	if err != nil {
		log.Println(err)
	}

	if !folder.Exists(certsPath) {
		folder.MkDir(certsPath)
	}

	return certsPath
}

func GetCertsFullPathByScope(scope string) string {
	certsPath := GetCertsFullPath()
	if certsPath == "" {
		return ""
	}

	certsPath = certsPath + filesystem.DirSeparator() + scope
	if !folder.Exists(certsPath) {
		folder.MkDir(certsPath)
	}
	return certsPath
}
