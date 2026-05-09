package app

import (
	"fmt"
)

var (
	version     = "0.2.0"
	release     = "alpha"
	license     = "MIT"
	description = "XJTUPortal v%s %s, a web portal authentication manager for XJTU iHarbor campus network."
	copyright   = "Developed by Anonymous@XJTUANA, under %s License."
	contact     = "Wechat Official: XJTUPortal; Repo: https://github.com/noking-xdj/xjtuportal"
)

func ProgramInfo() string {
	return fmt.Sprintf("%s\n%s\n%s\n",
		fmt.Sprintf(description, version, release),
		fmt.Sprintf(copyright, license),
		contact,
	)
}
