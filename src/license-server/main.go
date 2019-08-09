package main

import (
	_ "license-server/db"
	"license-server/license"
)

func main() {
	println(license.Verify("Datachat"))
}
