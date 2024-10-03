package realpath

import (
	"bytes"
	"github.com/kyaxcorp/go-core/core/helpers/str"
	"os"
	"path/filepath"
)

// Realpath returns the real path of a given file in the os
// It should be an existing!
// And relative paths work only from executable path!!!
func Realpath(fPath string) (string, error) {

	if len(fPath) == 0 {
		return "", os.ErrInvalid
	}

	if !filepath.IsAbs(fPath) {
		var pwd string
		var err error
		if str.GetChar(fPath, 0) == "." && str.GetChar(fPath, 1) == "/" {
			// Working directory
			pwd, err = os.Getwd()
			if err != nil {
				return "", err
			}
		} else {
			// Check directory
			pwd, err = os.Executable()
			if err != nil {
				return "", err
			}
			pwd = filepath.Dir(pwd)
		}
		fPath = filepath.Join(pwd, fPath)
	}

	path := []byte(fPath)
	nlinks := 0
	start := 1
	prev := 1
	for start < len(path) {
		c := nextComponent(path, start)
		cur := c[start:]

		switch {

		case len(cur) == 0:
			copy(path[start:], path[start+1:])
			path = path[0 : len(path)-1]

		case len(cur) == 1 && cur[0] == '.':
			if start+2 < len(path) {
				copy(path[start:], path[start+2:])
			}
			path = path[0 : len(path)-2]

		case len(cur) == 2 && cur[0] == '.' && cur[1] == '.':
			copy(path[prev:], path[start+2:])
			path = path[0 : len(path)+prev-(start+2)]
			prev = 1
			start = 1

		default:

			fi, err := os.Lstat(string(c))
			if err != nil {
				return "", err
			}
			if isSymlink(fi) {

				nlinks++
				if nlinks > 16 {
					return "", os.ErrInvalid
				}

				var link string
				link, err = os.Readlink(string(c))
				after := string(path[len(c):])

				// switch symlink component with its real path
				path = switchSymlinkCom(path, start, link, after)

				prev = 1
				start = 1
			} else {
				// Directories
				prev = start
				start = len(c) + 1
			}
		}
	}

	for len(path) > 1 && path[len(path)-1] == os.PathSeparator {
		path = path[0 : len(path)-1]
	}
	return string(path), nil

}

// test if a link is symbolic link
func isSymlink(fi os.FileInfo) bool {
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}

// switch a symbolic link component to its real path
func switchSymlinkCom(path []byte, start int, link, after string) []byte {

	if link[0] == os.PathSeparator {
		// Absolute links
		return []byte(filepath.Join(link, after))
	}

	// Relative links
	return []byte(filepath.Join(string(path[0:start]), link, after))
}

// get the next component
func nextComponent(path []byte, start int) []byte {
	v := bytes.IndexByte(path[start:], os.PathSeparator)
	if v < 0 {
		return path
	}
	return path[0 : start+v]
}
