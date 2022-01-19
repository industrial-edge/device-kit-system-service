package app

import (
	"context"
	"errors"
	"testing"

	//systemapi "systemservice/api/siemens_iedge_dmapi_v1"
	v1 "systemservice/api/siemens_iedge_dmapi_v1"
	clientfct "systemservice/internal/clientfactory"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	"time"
)

type tSystemController struct {
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

func Test_ShutdownFailure(t *testing.T) {
	t.Parallel()

	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.ShutdownDevice(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to Shutdown", "Did not get expected result. expected: %q got: %q", "Failed to Shutdown", err.Error())
}

func Test_HardResetFailuret(t *testing.T) {
	t.Parallel()

	tApp, dummyctx, dummyEmpty := initializeRPCDatas()

	_, err := tApp.serverInstance.HardReset(dummyctx, dummyEmpty)

	//Kill the goroutine
	tApp.done <- true

	assert.Contains(t, err.Error(), "Failed to HardReset", "Did not get expected result. expected: %q got: %q", "Failed to HardReset", err.Error())
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
