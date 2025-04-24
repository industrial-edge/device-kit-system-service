/*
 * Copyright © Siemens 2024 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package hostnameservice

import (
	"errors"
	"testing"

	"systemservice/internal/common/mocks"

	"github.com/stretchr/testify/assert"
)

func initialize() (*mocks.MUtil, *mocks.MFS, *HostnameService) {
	tUtil := new(mocks.MUtil)
	tFS := new(mocks.MFS)
	tService := NewHostnameService(tUtil, tFS)
	return tUtil, tFS, tService
}

func TestGet(t *testing.T) {
	tUtil, _, service := initialize()

	tUtil.HostnameVal = "test-hostname"
	tUtil.HostnameErr = nil

	hostname, err := service.Get()
	assert.NoError(t, err)
	assert.Equal(t, "test-hostname", hostname)
}

func TestGet_Error(t *testing.T) {
	tUtil, _, service := initialize()

	tUtil.HostnameErr = errors.New("failed to get hostname")

	hostname, err := service.Get()
	assert.Error(t, err)
	assert.Empty(t, hostname)
}

func TestUpdate(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
		{WriteFileErr: nil},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
		{CommandVal: []byte(""), CommandErr: nil},
		{CommandVal: []byte(""), CommandErr: nil},
	}
	tUtil.SetenvErr = nil

	err := service.Update("new-hostname")
	assert.NoError(t, err)
}

func TestUpdate_Error(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: errors.New("failed to write to /etc/hostname")},
	}

	err := service.Update("new-hostname")
	assert.Error(t, err)
}

func TestUpdateWithHostname_Error(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
		{CommandVal: []byte(""), CommandErr: errors.New("failed to set hostname with hostname command")},
	}

	err := service.Update("new-hostname")
	assert.Error(t, err)
}

func TestUpdateWithHostname_UpdateEtcHosts_Error(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
		{WriteFileErr: errors.New("failed to update /etc/hosts")},
		{WriteFileErr: nil}, // Simulate successful rollback
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
		{CommandVal: []byte(""), CommandErr: nil},
	}

	err := service.Update("new-hostname")
	assert.Error(t, err)
}

func TestWriteToFile_Error(t *testing.T) {
	_, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: nil, ReadFileErr: errors.New("failed to read /etc/hostname")},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
	}

	err := service.writeToFile("/etc/hostname", "new-hostname")
	assert.Error(t, err)
	assert.Equal(t, "failed to read /etc/hostname: failed to read /etc/hostname", err.Error())
}

func TestWriteToFile_ClearError(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: errors.New("failed to clear /etc/hostname")},
	}

	err := service.writeToFile("/etc/hostname", "new-hostname")
	assert.Error(t, err)
	assert.Equal(t, "failed to clear /etc/hostname: failed to clear /etc/hostname", err.Error())
}

func TestUpdateHostsFile_ReadFileError(t *testing.T) {
	_, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: nil, ReadFileErr: errors.New("failed to read /etc/hosts")},
	}

	err := service.updateHostsFile()
	assert.Error(t, err)
	assert.Equal(t, "failed to read /etc/hosts: failed to read /etc/hosts", err.Error())
}

func TestUpdateHostsFile_UpdateHostsEntry_Error(t *testing.T) {
	_, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: errors.New("failed to update /etc/hosts")},
		{WriteFileErr: nil}, // Simulate successful rollback
	}

	err := service.updateHostsFile()
	assert.Error(t, err)
	assert.Equal(t, "failed to update /etc/hosts: failed to update /etc/hosts", err.Error())
}

func TestWriteToFile_Success(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
	}

	err := service.writeToFile("/etc/hostname", "new-hostname")
	assert.NoError(t, err)
}

func TestUpdateHostsFile_Success(t *testing.T) {
	_, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
	}

	err := service.updateHostsFile()
	assert.NoError(t, err)
}

func TestWriteToFile_WriteError(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: errors.New("failed to write to /etc/hostname")},
		{WriteFileErr: nil}, // Simulate successful rollback
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
	}

	err := service.writeToFile("/etc/hostname", "new-hostname")
	assert.Error(t, err)
	assert.Equal(t, "failed to write to /etc/hostname: failed to write to /etc/hostname", err.Error())
}

func TestUpdateHostsFile_UpdateError(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: errors.New("failed to update /etc/hosts")},
		{WriteFileErr: nil}, // Simulate successful rollback
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
	}

	err := service.updateHostsFile()
	assert.Error(t, err)
	assert.Equal(t, "failed to update /etc/hosts: failed to update /etc/hosts", err.Error())
}

func TestWriteToFile_WriteErrorWithRollbackFailure(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: errors.New("failed to write to /etc/hostname")},
		{WriteFileErr: errors.New("failed to rollback")},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
	}

	err := service.writeToFile("/etc/hostname", "new-hostname")
	assert.Error(t, err)
	assert.Equal(t, "failed to write to /etc/hostname: failed to write to /etc/hostname; additionally, failed to rollback: failed to rollback", err.Error())
}

func TestUpdateHostsFile_UpdateErrorWithRollbackFailure(t *testing.T) {
	tUtil, tFS, service := initialize()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: errors.New("failed to update /etc/hosts")},
		{WriteFileErr: errors.New("failed to rollback")},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: nil},
	}

	err := service.updateHostsFile()
	assert.Error(t, err)
	assert.Equal(t, "failed to update /etc/hosts: failed to update /etc/hosts; additionally, failed to rollback: failed to rollback", err.Error())
}
