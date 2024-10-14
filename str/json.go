package str

import "github.com/kyaxcorp/go-helper/json"

func IsJSON(str string) bool {
	return json.IsJSON(str)
}
