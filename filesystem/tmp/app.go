package tmp

import (
	"strings"

	"github.com/kyaxcorp/go-helper/config"
	"github.com/kyaxcorp/go-helper/errors2/define"
	"github.com/kyaxcorp/go-helper/filesystem"
	fsPath "github.com/kyaxcorp/go-helper/filesystem/path"
	"github.com/kyaxcorp/go-helper/folder"
)

func GetAppTmpPath(paths ...string) (string, error) {
	var _err error = nil
	itemPath := config.GetConfig().TempPath
	itemPath, _err = fsPath.GenRealPath(itemPath, true)

	if _err != nil {
		//log.Println(err)
		return "", _err
	}

	if len(paths) > 0 {
		itemPath = itemPath + strings.Join(paths, filesystem.DirSeparator())
	}

	if !folder.Exists(itemPath) {
		if !folder.MkDir(itemPath) {
			return "", define.Err(0, "failed to create path -> ", itemPath)
		}
	}

	return itemPath, nil
}
