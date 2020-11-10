package utils

import (
	"github.com/i-open/beego/core/utils"
)

// GetGOPATHs returns all paths in GOPATH variable.
func GetGOPATHs() []string {
	return utils.GetGOPATHs()
}
