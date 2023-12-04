package systemcontroller

import (
	"errors"
	"testing"

	"systemservice/internal/common/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_RestartDeviceSuccess(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.RestartDevice()

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_ShutdownDeviceFailure(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: errors.New("Failed Shutdown")}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.ShutdownDevice()

	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_ShutdownDevice_Success(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.ShutdownDevice()

	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_RemoveContent_Failure_Due_To_docker_info(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: errors.New("docker info returns exit status 2")}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	err := controller.RemoveContent()
	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)

}

func Test_RemoveContent_Failure_Due_To_rm(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("/dummy/path/var/lib/docker"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)
	s2 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: errors.New("rm exit status 1")}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	err := controller.RemoveContent()
	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)

}

func Test_RemoveContent_Success(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte("/dummy/path/var/lib/docker"), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)
	s2 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)
	s3 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s3)
	s4 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s4)

	err := controller.RemoveContent()
	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_Remove_Success(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s1)
	s2 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: nil}
	tUtil.CommandList = append(tUtil.CommandList, s2)

	removeFileListDuringHardReset := []string{"fake/file", "fake/file2"}

	err := controller.RemoveFiles(removeFileListDuringHardReset)
	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_Remove_ErrorOnDeletingFiles(t *testing.T) {
	t.Parallel()
	tUtil := new(mocks.MUtil)
	tFs := new(mocks.MFS)
	controller := NewSystemController(tFs, tUtil)

	tUtil.CommandList = make([]mocks.CmdContainer, 0)
	s1 := mocks.CmdContainer{CommandVal: []byte(""), CommandErr: errors.New("fake/file does not exit.")}
	tUtil.CommandList = append(tUtil.CommandList, s1)

	removeFileListDuringHardReset := []string{"fake/file", "fake/file2"}

	err := controller.RemoveFiles(removeFileListDuringHardReset)
	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}
