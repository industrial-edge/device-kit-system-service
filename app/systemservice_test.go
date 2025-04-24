/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package app

import (
	"context"
	"errors"
	"testing"

	v1 "systemservice/api/siemens_iedge_dmapi_v1"
	clientfct "systemservice/internal/clientfactory"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"time"
)

type tSystemController struct {
}

type tSystemController2 struct {
}

func (s tSystemController2) RestartDevice() error {
	return nil
}
func (s tSystemController2) ShutdownDevice() error {
	return nil
}
func (s tSystemController2) HardReset(ctx context.Context, pack *clientfct.ClientPack) error {
	return nil
}
func (s tSystemController) RestartDevice() error {
	return errors.New("Failed RestartDevice")
}

func (s tSystemController) ShutdownDevice() error {
	return errors.New("Failed ShutdownDevice")
}

func (s tSystemController) HardReset(context.Context, *clientfct.ClientPack) error {
	return errors.New("Failed HardReset")
}

type tSystemInfo struct {
}

func (s tSystemInfo) GetLimits() (*v1.Limits, error) {
	return &v1.Limits{}, errors.New("Failed GetLimits")
}

func (s tSystemInfo) GetModelNumber() (*v1.ModelNumber, error) {
	return &v1.ModelNumber{}, errors.New("Failed GetModelNumber")
}

func (s tSystemInfo) GetFirmwareInfo() (*v1.FirmwareInfo, error) {
	return &v1.FirmwareInfo{}, errors.New("Failed GetFirmwareInfo")
}

func (s tSystemInfo) GetResourceStats() (*v1.Stats, error) {
	return &v1.Stats{}, errors.New("Failed GetResourceStats")
}
func (s tSystemInfo) GetLogFile(request *v1.LogRequest) (*v1.LogResponse, error) {
	return &v1.LogResponse{LogPath: request.SaveFolderPath + "/logs.tar.gz"}, nil
}

type MockHostnameManager struct {
	hostname   string
	shouldFail bool
}

func (m *MockHostnameManager) GetHostname() (string, error) {
	if m.shouldFail {
		return "", errors.New("Failed to get hostname")
	}
	return m.hostname, nil
}

func (m *MockHostnameManager) UpdateHostname(hostname string) error {
	if m.shouldFail {
		return errors.New("Failed to update hostname")
	}
	m.hostname = hostname
	return nil
}

func Test_VerifyArgsForStartGRPC_WithLessArgs(t *testing.T) {
	t.Parallel()
	//Create App to use
	tApp := CreateServiceApp(clientfct.ClientFactoryImpl{})

	//Test with 0 argument
	tArgs := []string{}
	tErr := tApp.StartGRPC(tArgs)

	assert.NotNil(t, tErr, "Did not get expected result. Wanted: %q, got: %q", "parameter not supported!", tErr)

	//Test with 1 argument
	tArgs = []string{"ntpserver"}
	tErr = tApp.StartGRPC(tArgs)

	assert.NotNil(t, tErr, "Did not get expected result. Wanted: %q, got: %q", "parameter not supported!", tErr)

	//Test with 2 arguments
	tArgs = []string{"ntpserver", "unix"}
	tErr = tApp.StartGRPC(tArgs)

	assert.NotNil(t, tErr, "Did not get expected result. Wanted: %q, got: %q", "parameter not supported!", tErr.Error())
}

func Test_VerifyArgsForStartGRPC_WithInappropriateArgs(t *testing.T) {
	t.Parallel()
	//Create App to use
	tApp := CreateServiceApp(clientfct.ClientFactoryImpl{})

	tApp.StartApp()

	tArgs := []string{"ntpserver", "dummy", "11111"}
	tErr := tApp.StartGRPC(tArgs)

	//Kill the goroutine
	tApp.done <- true

	//Connection failure is expected with dummy sock
	assert.Equal(t, "parameter not supported: dummy", tErr.Error(), "Did not get expected result. Wanted: %q, got: %q", "parameter not supported: dummy", tErr.Error())
}

func Test_VerifyArgsForStartGRPC_WithDummySocketForUnix(t *testing.T) {
	t.Parallel()
	//Create App to use
	tApp := CreateServiceApp(clientfct.ClientFactoryImpl{})

	//wait until system is up and goroutines are running
	tApp.StartApp()
	time.Sleep(time.Second * 2)

	tArgs := []string{"ntpserver", "unix", "/dummy/unix/path.sock"}
	tErr := tApp.StartGRPC(tArgs)

	//Kill the goroutine
	tApp.done <- true

	//Connection failure is expected with dummy sock
	assert.NotNil(t, tErr, "Did not get expected result. got: %q", tErr)
}

func Test_chownSocketFailure(t *testing.T) {
	t.Parallel()
	//Fail the function with Non existing file path
	err := chownSocket("Non/existing/Path", "root", "root")

	assert.NotNil(t, err, "Did not get expected result. got: %q", err)
}

func initializeRPCDatas() (*MainApp, context.Context, *emptypb.Empty) {
	//Prepare Test Data
	var dummyctx context.Context
	var dummyEmpty *emptypb.Empty

	//Create App to use
	tApp := CreateServiceApp(clientfct.ClientFactoryImpl{})

	//inject new struct
	tApp.serverInstance.IsysController = tSystemController{}
	tApp.serverInstance.IsysInfo = tSystemInfo{}
	tApp.serverInstance.IhostnameController = &MockHostnameManager{hostname: "initial-hostname"}

	tApp.StartApp()

	return tApp, dummyctx, dummyEmpty
}

func Test_RestartFailure(t *testing.T) {
	t.Parallel()

	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.RestartDevice(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to Restart", "Did not get expected result. expected: %q got: %q", "Failed to Restart", err.Error())
}

func Test_RestartSuccess(t *testing.T) {
	t.Parallel()
	var dummyctx context.Context
	var dummyEmpty *emptypb.Empty
	tApp := CreateServiceApp(clientfct.ClientFactoryImpl{})
	tApp.serverInstance.IsysController = tSystemController2{}
	tApp.serverInstance.IsysInfo = tSystemInfo{}
	tApp.StartApp()

	_, err := tApp.serverInstance.RestartDevice(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Nil(t, err)
}

func Test_ShutdownFailure(t *testing.T) {
	t.Parallel()

	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.ShutdownDevice(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to Shutdown", "Did not get expected result. expected: %q got: %q", "Failed to Shutdown", err.Error())
}

func Test_ShutdownSuccess(t *testing.T) {
	t.Parallel()
	var dummyctx context.Context
	var dummyEmpty *emptypb.Empty
	tApp := CreateServiceApp(clientfct.ClientFactoryImpl{})
	tApp.serverInstance.IsysController = tSystemController2{}
	tApp.serverInstance.IsysInfo = tSystemInfo{}
	tApp.StartApp()

	_, err := tApp.serverInstance.ShutdownDevice(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Nil(t, err)
}

func Test_HardResetFailuret(t *testing.T) {
	t.Parallel()

	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.HardReset(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to HardReset", "Did not get expected result. expected: %q got: %q", "Failed to HardReset", err.Error())
}

func Test_HardResetSuccess(t *testing.T) {
	t.Parallel()
	var dummyctx context.Context
	var dummyEmpty *emptypb.Empty
	tApp := CreateServiceApp(clientfct.ClientFactoryImpl{})
	tApp.serverInstance.IsysController = tSystemController2{}
	tApp.serverInstance.IsysInfo = tSystemInfo{}
	tApp.StartApp()

	_, err := tApp.serverInstance.HardReset(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Nil(t, err)
}

func Test_GetLimitsFailuret(t *testing.T) {
	t.Parallel()

	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.GetLimits(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to GetLimits", "Did not get expected result. expected: %q got: %q", "Failed to GetLimits", err.Error())
}

func Test_GetModelNumberFailuret(t *testing.T) {
	t.Parallel()
	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.GetModelNumber(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to GetModelNumber", "Did not get expected result. expected: %q got: %q", "Failed to GetModelNumber", err.Error())
}

func Test_GetFirmwareInfoFailure(t *testing.T) {
	t.Parallel()
	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.GetFirmwareInfo(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to GetFirmwareInfo", "Did not get expected result. expected: %q got: %q", "Failed to GetFirmwareInfo", err.Error())
}

func Test_GetResourceStatsFailure(t *testing.T) {
	t.Parallel()
	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.GetResourceStats(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to GetResourceStats", "Did not get expected result. expected: %q got: %q", "Failed to GetResourceStats", err.Error())
}

func Test_GetCustomSettings(t *testing.T) {
	t.Parallel()
	tApp, dummyCtx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.GetCustomSettings(dummyCtx, dummyEmpty)
	if err != nil {
		t.Log("FAILED GetCustomSettings")
		t.Fail()
	}

	// Kill the goroutine
	tApp.done <- true
}

func Test_ApplyCustomSettings(t *testing.T) {
	t.Parallel()
	tApp, dummyCtx, _ := initializeRPCDatas()

	jsonTxt := []byte(`{ "key" : "value", "key2" : "value2" }`)

	anyMessage := anypb.Any{
		Value: jsonTxt,
	}

	_, err := tApp.serverInstance.ApplyCustomSettings(dummyCtx, &anyMessage)
	if err != nil {
		t.Log("FAILED ApplyCustomSettings")
		t.Fail()
	}
	//Kill the goroutine
	tApp.done <- true
}

func TestGetHostname(t *testing.T) {
	t.Parallel()
	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	resp, err := tApp.serverInstance.GetHostname(dummyctx, dummyEmpty)
	if err != nil {
		t.Fatalf("GetHostname() error = %v", err)
	}

	assert.Equal(t, "initial-hostname", resp.Name, "GetHostname() = %v, want %v", resp.Name, "initial-hostname")

	// Kill the goroutine
	tApp.done <- true
}

func TestSetHostname(t *testing.T) {
	t.Parallel()
	tApp, dummyctx, _ := initializeRPCDatas()

	newHostname := "new-hostname"
	req := &v1.Hostname{Name: newHostname}

	_, err := tApp.serverInstance.UpdateHostname(dummyctx, req)
	if err != nil {
		t.Fatalf("SetHostname() error = %v", err)
	}

	// Check if the new hostname was set correctly
	currentHostname, err := tApp.serverInstance.IhostnameController.GetHostname()
	if err != nil {
		t.Fatalf("failed to get current hostname: %v", err)
	}

	assert.Equal(t, newHostname, currentHostname, "current hostname = %v, want %v", currentHostname, newHostname)

	// Kill the goroutine
	tApp.done <- true
}

func TestGetHostnameError(t *testing.T) {
	t.Parallel()
	tApp, dummyctx, dummyEmpty := initializeRPCDatas()
	tApp.serverInstance.IhostnameController = &MockHostnameManager{shouldFail: true}

	_, err := tApp.serverInstance.GetHostname(dummyctx, dummyEmpty)
	assert.NotNil(t, err, "Expected error but got nil")
	assert.Contains(t, err.Error(), "failed to get hostname", "Did not get expected result. expected: %q got: %q", "failed to get hostname", err.Error())

	// Kill the goroutine
	tApp.done <- true
}

func TestSetHostnameError(t *testing.T) {
	t.Parallel()
	tApp, dummyctx, _ := initializeRPCDatas()
	tApp.serverInstance.IhostnameController = &MockHostnameManager{shouldFail: true}

	newHostname := "new-hostname"
	req := &v1.Hostname{Name: newHostname}

	_, err := tApp.serverInstance.UpdateHostname(dummyctx, req)
	assert.NotNil(t, err, "Expected error but got nil")
	assert.Contains(t, err.Error(), "failed to update hostname", "Did not get expected result. expected: %q got: %q", "failed to update hostname", err.Error())

	// Kill the goroutine
	tApp.done <- true
}
