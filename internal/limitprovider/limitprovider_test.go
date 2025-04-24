/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package limitprovider

import (
	"errors"
	"log"
	"systemservice/internal/common/mocks"
	"testing"

	systemapi "systemservice/api/siemens_iedge_dmapi_v1"

	"github.com/stretchr/testify/assert"
)

const testContentProperJSON = `
{
"Limits":
{
	"MaxInstalledApplications":100,
	"MaxRunningApplications":100,
	"MaxMemoryUsageInGB":4.0,
	"MaxStorageUsageInGB":20.0,
	"MaxCpuUsagePerecentage":60
},
"MonitoredStorage":
{
	"Path":"/dummymount"
}
}
`

const testContentNonProperJSON = `
{
"Limits":
{
	"MaxInstalledApplications":100,
	"MaxRunningApplications":100,
	"MaxMemoryUsageInGB":4.0,
	"MaxStorageUsageInGB":20.0,
	"MaxCpuUsagePerecentage":60
},
"MonitoredStorage":
{
	"Path":"/dummymount"
`

func initialize() (*mocks.MFS, *mocks.MUtil, *LimitProvider) {
	tFs := new(mocks.MFS)
	tUtil := new(mocks.MUtil)
	tlimit := CreateLimitProvider(tFs, tUtil)
	log.Println(tlimit)

	return tFs, tUtil, tlimit
}

func Test_WithNotProperLimitJsonContent_ErrorExpected(t *testing.T) {
	t.Parallel()
	//Prepare Content
	tFs, _, tlimit := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(testContentNonProperJSON), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	cont, err := tlimit.GetLimitContent()
	t.Log("content: ", cont, " err: ", err)

	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_ContentReadFailure_ErrorExpected(t *testing.T) {
	t.Parallel()
	//Prepare Content
	tFs, _, tlimit := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(testContentProperJSON), ReadFileErr: errors.New("Failed to Read")}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	cont, err := tlimit.GetLimitContent()
	t.Log("content: ", cont, " err: ", err)

	assert.NotNil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
}

func Test_WithProperLimitJsonContent_ProperLimitExpected(t *testing.T) {
	t.Parallel()
	//Prepare Content
	tFs, _, tlimit := initialize()

	tFs.ReadFileList = make([]mocks.ReadFileContainer, 0)
	r1 := mocks.ReadFileContainer{ReadFileVal: []byte(testContentProperJSON), ReadFileErr: nil}
	tFs.ReadFileList = append(tFs.ReadFileList, r1)

	cont, err := tlimit.GetLimitContent()
	t.Log(cont, err)
	expectedLimits := &systemapi.Limits{MaxInstalledApplications: 100,
		MaxRunningApplications: 100,
		MaxMemoryUsageInGB:     4,
		MaxStorageUsageInGB:    20,
		MaxCpuUsagePerecentage: 60}

	t.Log(expectedLimits, err)
	assert.Nil(t, err, "Did not get expected result. Wanted: %q, got: %q", nil, err)
	assert.Equal(t, cont.MaxInstalledApplications, expectedLimits.MaxInstalledApplications, "Content: %q Expected: %q", cont.MaxInstalledApplications, expectedLimits.MaxInstalledApplications)
	assert.Equal(t, cont.MaxRunningApplications, expectedLimits.MaxRunningApplications, "Content: %q Expected: %q", cont.MaxRunningApplications, expectedLimits.MaxRunningApplications)
	assert.Equal(t, cont.MaxMemoryUsageInGB, expectedLimits.MaxMemoryUsageInGB, "Content: %q Expected: %q", cont.MaxMemoryUsageInGB, expectedLimits.MaxMemoryUsageInGB)
	assert.Equal(t, cont.MaxStorageUsageInGB, expectedLimits.MaxStorageUsageInGB, "Content: %q Expected: %q", cont.MaxStorageUsageInGB, expectedLimits.MaxStorageUsageInGB)
	assert.Equal(t, cont.MaxCpuUsagePerecentage, expectedLimits.MaxCpuUsagePerecentage, "Content: %q Expected: %q", cont.MaxCpuUsagePerecentage, expectedLimits.MaxCpuUsagePerecentage)
}
