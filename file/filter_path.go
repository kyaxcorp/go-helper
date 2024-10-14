package file

import (
	"strings"

	"github.com/kyaxcorp/go-helper/filesystem"
)

func FilterPath(path string) string {
	p := strings.ReplaceAll(path, `\\`, filesystem.DirSeparator())
	p = strings.ReplaceAll(p, "//", filesystem.DirSeparator())
	p = strings.ReplaceAll(p, "/", filesystem.DirSeparator())
	p = strings.ReplaceAll(p, `\`, filesystem.DirSeparator())
	return p
}
