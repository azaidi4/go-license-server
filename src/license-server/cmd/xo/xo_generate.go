package main

import (
	"fmt"
	"license-server/db/dbutils"
	"os"
	"os/exec"
	"path/filepath"
)

func generateModel() {
	URL := dbutils.BuildDBURL("pgsql")
	path, _ := filepath.Abs("../../db/models")

	fmt.Println("Generating xo code...")
	err := exec.Command("xo", URL, "-o", path).Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Done!")
}

func main() {
	generateModel()
}
