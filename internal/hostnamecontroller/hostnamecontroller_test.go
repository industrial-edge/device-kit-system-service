/*
 * Copyright © Siemens 2024 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package hostnamecontroller

import (
	"errors"
	"testing"

	"systemservice/internal/common/mocks"
	"systemservice/internal/hostnamecontroller/hostnameservice"

	"github.com/stretchr/testify/assert"
)

func initializeController() (*mocks.MUtil, *mocks.MFS, *HostnameController) {
	tUtil := new(mocks.MUtil)
	tMFS := new(mocks.MFS)
	tService := hostnameservice.NewHostnameService(tUtil, tMFS)
	tController := NewHostnameController(*tService)
	return tUtil, tMFS, tController
}

func TestGetHostname(t *testing.T) {
	tUtil, _, controller := initializeController()

	tUtil.HostnameVal = "test-hostname"
	tUtil.HostnameErr = nil

	hostname, err := controller.GetHostname()
	assert.NoError(t, err)
	assert.Equal(t, "test-hostname", hostname)
}

func TestGetHostname_Error(t *testing.T) {
	tUtil, _, controller := initializeController()

	tUtil.HostnameErr = errors.New("failed to get hostname")

	hostname, err := controller.GetHostname()
	assert.Error(t, err)
	assert.Empty(t, hostname)
}

func TestUpdateHostname(t *testing.T) {
	tUtil, tFS, controller := initializeController()

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

	err := controller.UpdateHostname("new-hostname")
	assert.NoError(t, err)
}

func TestUpdateHostname_Invalid(t *testing.T) {
	_, _, controller := initializeController()

	err := controller.UpdateHostname("invalid_hostname")
	assert.Error(t, err)
}

func TestUpdateHostname_Error(t *testing.T) {
	tUtil, tFS, controller := initializeController()

	tFS.ReadFileList = []mocks.ReadFileContainer{
		{ReadFileVal: []byte("original content"), ReadFileErr: nil},
	}
	tFS.WriteFileList = []mocks.WriteFileContainer{
		{WriteFileErr: nil},
	}
	tUtil.CommandList = []mocks.CmdContainer{
		{CommandVal: []byte(""), CommandErr: errors.New("failed to write to /etc/hostname")},
	}

	err := controller.UpdateHostname("new-hostname")
	assert.Error(t, err)
}

func TestValidateHostname(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		hostname string
		wantErr  bool
	}{
		{"ValidHostname", "valid-hostname", false},
		{"InvalidHostnameSpace1", "invalid hostname", true},
		{"InvalidHostnameSpace2", "inv    ali   d  host name", true},
		{"InvalidHostnameSpace3", "invalid	host name", true},
		{"InvalidHostnameTab", "invalid		hostname", true},
		{"InvalidHostnameUnderscore", "invalid_hostname", true},
		{"EmptyHostname", "", true},
		{"ExceedMaxLengthHostname", "a-very-long-hostname-that-exceeds-the-maximum-length-limit-which-is-253-characters-a-very-long-hostname-that-exceeds-the-maximum-length-limit-which-is-253-characters-a-very-long-hostname-that-exceeds-the-maximum-length-limit-which-is-253-characters-a-very-long-hostname-that-exceeds-the-maximum-length-limit-which-is-253-characters", true},
		{"ValidHostnameWithNumbers", "hostname123", false},
		{"ValidHostnameWithUppercase", "Valid-Hostname", false},
		{"InvalidHostnameLeadingHyphen", "-hostname", true},
		{"InvalidHostnameTrailingHyphen", "hostname-", true},
		{"InvalidHostnameConsecutiveDots", "host..name", true},
		{"InvalidHostnameConsecutiveHyphens", "host--name", true},
		{"InvalidHostnameLeadingDot", ".hostname", true},
		{"InvalidHostnameTrailingDot", "hostname.", true},
		{"ValidHostnameWithSingleCharacter", "a", false},
		{"ValidHostnameWithMaxLength", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aa", false},
		{"ValidHostnameWithMaxLengthNoDot", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"ValidSimpleHostnameMaxLength", "a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-thatis", false},
		{"InvalidSimpleHostnameExceedsMaxLength", "a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters", true},
		{"ValidFQDNMaxLength", "a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"InvalidFQDNExceedsMaxLength", "a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-62-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
		{"InvalidFQDNLabelExceedsMaxLength", "a-64-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-61-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
		{"ValidHostnameWithHyphen", "valid-hostname-with-hyphen", false},
		{"ValidHostnameWithDot", "valid.hostname.with.dot", false},
		{"InvalidHostnameWithConsecutiveHyphens", "invalid--hostname", true},
		{"InvalidHostnameWithConsecutiveDots", "invalid..hostname", true},
		{"InvalidHostnameWithLeadingHyphen", "-invalidhostname", true},
		{"InvalidHostnameWithTrailingHyphen", "invalidhostname-", true},
		{"InvalidHostnameWithLeadingDot", ".invalidhostname", true},
		{"InvalidHostnameWithTrailingDot", "invalidhostname.", true},
		{"SimpleHostnameMaxLength", "a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-that-is-exactly-255-characters-long-a-very-long-hostname-thatis", false},
		{"SimpleHostnameExceedsMaxLength", "a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters", true},
		{"FQDNMaxLength", "a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"FQDNExceedsMaxLength", "a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-62-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
		{"FQDNLabelExceedsMaxLength", "a-64-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-61-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
		{"InvalidSimpleHostnameWithUnderscore", "invalid_hostname", true},
		{"InvalidSimpleHostnameWithConsecutiveHyphens", "invalid--hostname", true},
		{"InvalidFQDNWithUnderscore", "invalid_hostname.example.com", true},
		{"InvalidFQDNWithConsecutiveHyphens", "invalid--hostname.example.com", true},
		{"SimpleHostnameExceedsMaxLengthError", "a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters-a-very-long-hostname-that-exceeds-255-characters", true},
		{"FQDNLabelExceedsMaxLengthError", "a-64-character-long-label-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.a-63-characte", true},
		{"InvalidHostnameWithSpecialChars", "invalid@hostname", true},
		{"InvalidHostnameWithSlash", "invalid/hostname", true},
		{"InvalidHostnameWithBackslash", "invalid\\hostname", true},
		{"InvalidHostnameWithColon", "invalid:hostname", true},
		{"InvalidHostnameWithSemicolon", "invalid;hostname", true},
		{"InvalidHostnameWithComma", "invalid,hostname", true},
		{"InvalidHostnameWithAsterisk", "invalid*hostname", true},
		{"InvalidHostnameWithExclamation", "invalid!hostname", true},
		{"InvalidHostnameWithQuestionMark", "invalid?hostname", true},
		{"InvalidHostnameWithPercent", "invalid%hostname", true},
		{"InvalidHostnameWithDollar", "invalid$hostname", true},
		{"InvalidHostnameWithAmpersand", "invalid&hostname", true},
		{"InvalidHostnameWithParentheses", "invalid(hostname)", true},
		{"InvalidHostnameWithBrackets", "invalid[hostname]", true},
		{"InvalidHostnameWithBraces", "invalid{hostname}", true},
		{"InvalidHostnameWithAngleBrackets", "invalid<hostname>", true},
		{"InvalidHostnameWithPipe", "invalid|hostname", true},
		{"InvalidHostnameWithCaret", "invalid^hostname", true},
		{"InvalidHostnameWithTilde", "invalid~hostname", true},
		{"InvalidHostnameWithGraveAccent", "invalid`hostname", true},
		{"InvalidHostnameWithPlus", "invalid+hostname", true},
		{"InvalidHostnameWithEqual", "invalid=hostname", true},
		{"InvalidHostnameWithDoubleQuote", "invalid\"hostname", true},
		{"InvalidHostnameWithSingleQuote", "invalid'hostname", true},
		{"InvalidHostnameWithSpace", "invalid hostname", true},
		{"InvalidHostnameWithTab", "invalid\thostname", true},
		{"InvalidHostnameWithNewline", "invalid\nhostname", true},
		{"InvalidHostnameWithCarriageReturn", "invalid\rhostname", true},
		{"InvalidHostnameWithFormFeed", "invalid\fhostname", true},
		{"InvalidHostnameWithVerticalTab", "invalid\vhostname", true},
		{"InvalidHostnameWithNullChar", "invalid\000hostname", true},
		{"InvalidHostnameWithControlChar", "invalid\001hostname", true},
		{"InvalidHostnameWithUnicodeChar", "invalid\u202Ehostname", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateHostname(tt.hostname)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateHostname() error = %v, wantErr %v, name = %v", err, tt.wantErr, tt.name)
			}
		})
	}
}
