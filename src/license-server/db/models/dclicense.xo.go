// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// DcLicense represents a row from 'public.dc_licenses'.
type DcLicense struct {
	ID               int       `json:"id"`                 // id
	LicenseExpiry    time.Time `json:"license_expiry"`     // license_expiry
	LicenseUserLimit int       `json:"license_user_limit"` // license_user_limit
	LicenseKey       int64     `json:"license_key"`        // license_key

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DcLicense exists in the database.
func (dl *DcLicense) Exists() bool {
	return dl._exists
}

// Deleted provides information if the DcLicense has been deleted from the database.
func (dl *DcLicense) Deleted() bool {
	return dl._deleted
}

// Insert inserts the DcLicense to the database.
func (dl *DcLicense) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dl._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.dc_licenses (` +
		`license_expiry, license_user_limit, license_key` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, dl.LicenseExpiry, dl.LicenseUserLimit, dl.LicenseKey)
	err = db.QueryRow(sqlstr, dl.LicenseExpiry, dl.LicenseUserLimit, dl.LicenseKey).Scan(&dl.ID)
	if err != nil {
		return err
	}

	// set existence
	dl._exists = true

	return nil
}

// Update updates the DcLicense in the database.
func (dl *DcLicense) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dl._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.dc_licenses SET (` +
		`license_expiry, license_user_limit, license_key` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, dl.LicenseExpiry, dl.LicenseUserLimit, dl.LicenseKey, dl.ID)
	_, err = db.Exec(sqlstr, dl.LicenseExpiry, dl.LicenseUserLimit, dl.LicenseKey, dl.ID)
	return err
}

// Save saves the DcLicense to the database.
func (dl *DcLicense) Save(db XODB) error {
	if dl.Exists() {
		return dl.Update(db)
	}

	return dl.Insert(db)
}

// Upsert performs an upsert for DcLicense.
//
// NOTE: PostgreSQL 9.5+ only
func (dl *DcLicense) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if dl._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.dc_licenses (` +
		`id, license_expiry, license_user_limit, license_key` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, license_expiry, license_user_limit, license_key` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.license_expiry, EXCLUDED.license_user_limit, EXCLUDED.license_key` +
		`)`

	// run query
	XOLog(sqlstr, dl.ID, dl.LicenseExpiry, dl.LicenseUserLimit, dl.LicenseKey)
	_, err = db.Exec(sqlstr, dl.ID, dl.LicenseExpiry, dl.LicenseUserLimit, dl.LicenseKey)
	if err != nil {
		return err
	}

	// set existence
	dl._exists = true

	return nil
}

// Delete deletes the DcLicense from the database.
func (dl *DcLicense) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dl._exists {
		return nil
	}

	// if deleted, bail
	if dl._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.dc_licenses WHERE id = $1`

	// run query
	XOLog(sqlstr, dl.ID)
	_, err = db.Exec(sqlstr, dl.ID)
	if err != nil {
		return err
	}

	// set deleted
	dl._deleted = true

	return nil
}

// DcLicenseByID retrieves a row from 'public.dc_licenses' as a DcLicense.
//
// Generated from index 'dc_licenses_pkey'.
func DcLicenseByID(db XODB, id int) (*DcLicense, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, license_expiry, license_user_limit, license_key ` +
		`FROM public.dc_licenses ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	dl := DcLicense{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dl.ID, &dl.LicenseExpiry, &dl.LicenseUserLimit, &dl.LicenseKey)
	if err != nil {
		return nil, err
	}

	return &dl, nil
}
