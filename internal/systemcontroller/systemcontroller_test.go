package systemcontroller

import (
	"context"
	"errors"
	"systemservice/internal/clientfactory"
	"testing"

	"systemservice/internal/common/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_RestartDeviceSuccess(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	controller := NewSystemController(nil, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.RestartDevice()

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_ShutdownDeviceFailure(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	controller := NewSystemController(nil, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: errors.New("failed Shutdown")}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.ShutdownDevice()

	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", err, nil)
}

func Test_ShutdownDevice_Success(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	controller := NewSystemController(nil, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.ShutdownDevice()

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_HardReset_Failure_Due_To_GetDockerRootDir_Error(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	controller := NewSystemController(nil, tUtil)
	var dummyCtx context.Context
	var dummyClientPack *clientfactory.ClientPack

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: errors.New("docker info returns exit status 2")}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.HardReset(dummyCtx, dummyClientPack)
	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", err, nil)
}

func Test_HardReset_Failure_Due_To_RemoveFile_Error(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	controller := NewSystemController(nil, tUtil)
	var dummyCtx context.Context
	var dummyClientPack *clientfactory.ClientPack

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("/dummy/path/var/lib/docker"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)
	s2 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: errors.New("rm exit status 1")}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	err := controller.HardReset(dummyCtx, dummyClientPack)
	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", err, nil)

}

func Test_HardReset_Success(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)
	var dummyCtx context.Context
	var dummyClientPack *clientfactory.ClientPack

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("/dummy/path/var/lib/docker"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)
	s2 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)
	s3 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s3)
	s4 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s4)
	s5 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s5)
	s6 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s6)

	err := controller.HardReset(dummyCtx, dummyClientPack)
	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}
