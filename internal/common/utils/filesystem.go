/*
 * Copyright (c) Siemens 2021
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package utils

import (
	"io/ioutil"
	"os"
)

// FileSystem Interface has the wrapper of filesystem calls
type FileSystem interface {
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
	Stat(name string) (os.FileInfo, error)
}

// OsFS struct for wrappers
type OsFS struct{}

// OpenFile is a wrapper func
func (OsFS) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

// ReadFile is a wrapper func
func (OsFS) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// WriteFile is a wrapper func
func (OsFS) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

// Stat is a wrapper func
func (OsFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
