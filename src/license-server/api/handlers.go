package api

import (
	"fmt"
	"license-server/api/apiutils"
	"license-server/license"
	"license-server/utils/logger"
	"net/http"
	"strings"
)

const maxMemory = 10 * (1 << 20) // 10MB
const tempPrefix = "dc_lic-"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Simple Licensing Server")
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxMemory); logger.LogOnError(err) {
		apiutils.ReturnWithStatus(w, http.StatusRequestEntityTooLarge)
		return
	}
	licenseFile, header, err := r.FormFile("license")
	if logger.LogOnError(err) {
		apiutils.ReturnWithStatus(w, http.StatusBadRequest)
		return
	}
	defer licenseFile.Close()

	orgName := strings.TrimSuffix(header.Filename, ".license.dat")
	// tmpFile, err := ioutil.TempFile(os.TempDir(), tempPrefix)
	// if logger.LogOnError(err) {
	// 	apiutils.ReturnWithStatus(w, http.StatusBadRequest)
	// 	return
	// }

	// defer os.Remove(tmpFile.Name())
	// if _, err := io.Copy(tmpFile, licenseFile); logger.LogOnError(err) {
	// 	apiutils.ReturnWithStatus(w, http.StatusBadRequest)
	// 	return
	// }

	if license.Verify(orgName, licenseFile) {
		logger.Info.Println("License for", orgName, "has PASSED integrity check.")
		return
	}
	logger.Info.Println("License for", orgName, "has FAILED integrity check.")
	apiutils.ReturnWithStatus(w, http.StatusUnauthorized)
	return

}
