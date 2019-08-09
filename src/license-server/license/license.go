// Package license serves as the licensing library for Datchat Licese Server.
// Licenses are signed using Elliptic Curve crytography (ECC).
// See blog: (https://arstechnica.com/information-technology/2013/10/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/)
// Algorithm uses P-384 Curve (https://en.wikipedia.org/wiki/P-384)
// License struct used is from license-server/db/models/dclicense.xo.go
package license

import (
	"encoding/json"
	"io/ioutil"
	"license-server/db"
	"license-server/db/models"
	"license-server/utils/logger"
	"mime/multipart"

	"github.com/hyperboloide/lk"
)

const licensePrefix = ".license.dat"

// generatePrivateKey returns a private key for signing the license
func generatePrivateKey() (*lk.PrivateKey, error) {
	return lk.NewPrivateKey()
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// SignAndSave signs the license and saves it as <LicenseOrg>.license.dat
func SignAndSave(license *models.DcLicense) error {
	docBytes, err := json.Marshal(license)
	if err != nil {
		return err
	}
	pk, err := generatePrivateKey()
	if err != nil {
		return err
	}
	signedLicense, err := lk.NewLicense(pk, docBytes)
	if err != nil {
		return err
	}
	signedLicenseBytes, err := signedLicense.ToBytes()
	if err != nil {
		return err
	}
	fName := license.LicenseOrg + licensePrefix
	ioutil.WriteFile(fName, signedLicenseBytes, 0444)

	lkey, err := pk.ToB64String()
	if err != nil {
		return err
	}
	license.LicenseKey = lkey

	if err := license.Save(db.DBConn); err != nil {
		return err
	}
	logger.Debug.Printf("License object saved: %s\n", prettyPrint(license))
	return nil
}

// Verify validates the license file, ensures it hasn't been tampered
func Verify(orgName string, license multipart.File) bool {
	privateKey, err := extractPrivateKey(orgName)
	if logger.LogOnError(err) {
		return false
	}
	publicKey := privateKey.GetPublicKey()

	if licenseBytes, err := ioutil.ReadAll(license); logger.LogOnError(err) {
		return false
	} else if license, err := lk.LicenseFromBytes(licenseBytes); logger.LogOnError(err) {
		return false
	} else if ok, err := license.Verify(publicKey); logger.LogOnError(err) {
		return false
	} else {
		return ok
	}
}

func extractPrivateKey(orgName string) (*lk.PrivateKey, error) {
	if licenseDB, err := models.DcLicenseByLicenseOrg(db.DBConn, orgName); err != nil {
		return nil, err
	} else if privateKey, err := lk.PrivateKeyFromB64String(licenseDB.LicenseKey); err != nil {
		return nil, err
	} else {
		return privateKey, err
	}
}
