package token

import "github.com/kyaxcorp/go-helper/str"

// AutoGenerated -> will create a random token with prefix auto-generated-xxxxxxxxxxxxxxxxxxxxx
func AutoGenerated(length int) string {
	return "auto-generated-" + str.RandomGenerate(length)
}
