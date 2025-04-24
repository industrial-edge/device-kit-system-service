/*
 * Copyright © Siemens 2024 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package hostnamecontroller

import (
	"errors"
	"regexp"
	"strings"
	"systemservice/internal/hostnamecontroller/hostnameservice"
)

const (
	maxHostnameLength         = 255
	underscorePattern         = `_`
	consecutiveHyphensPattern = `--`
	maxLabelLength            = 63
)

// Valid hostname regex: starts and ends with an alphanumeric character, allows hyphens and dots in between, no consecutive hyphens or dots
var validHostnameRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)*[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)
var validateHostnameWithoutDot = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,253}[a-zA-Z0-9])?$`)

type HostnameController struct {
	service hostnameservice.HostnameService
}

// NewHostnameController creates a new instance of HostnameController
func NewHostnameController(service hostnameservice.HostnameService) *HostnameController {
	return &HostnameController{service: service}
}

func (m *HostnameController) GetHostname() (string, error) {
	return m.service.Get()
}

func (m *HostnameController) UpdateHostname(newHostname string) error {
	if err := validateHostname(newHostname); err != nil {
		return err
	}

	if err := m.service.Update(newHostname); err != nil {
		return err
	}

	return nil
}

// validateHostname validates the given hostname string based on specific rules.
// It trims any leading or trailing whitespace from the hostname and checks the following:
// - The hostname length does not exceed the maximum allowed length.
// - If the hostname is a simple hostname (without dots), it ensures the length does not exceed 255 characters.
// - If the hostname is a fully qualified domain name (FQDN), it ensures each label does not exceed 63 characters.
// - The hostname matches the required pattern and does not contain invalid characters such as underscores or consecutive hyphens.
//
// Parameters:
// - hostname: The hostname string to validate.
//
// Returns:
// - An error if the hostname is invalid, otherwise nil.
func validateHostname(hostname string) error {
	hostname = strings.TrimSpace(hostname)
	if len(hostname) > maxHostnameLength {
		return errors.New("hostname exceeds maximum length of 255 characters")
	}

	labels := strings.Split(hostname, ".")
	if len(labels) == 1 {
		return validateSimpleHostname(hostname)
	}
	return validateFQDN(hostname, labels)
}

func validateSimpleHostname(hostname string) error {
	if !validateHostnameWithoutDot.MatchString(hostname) ||
		containsPattern(hostname, underscorePattern) ||
		containsPattern(hostname, consecutiveHyphensPattern) {
		return errors.New("hostname contains invalid characters")
	}
	return nil
}

func validateFQDN(hostname string, labels []string) error {
	for _, label := range labels {
		if len(label) > maxLabelLength {
			return errors.New("hostname label exceeds maximum length of 63 characters")
		}
	}
	if !validHostnameRegex.MatchString(hostname) ||
		containsPattern(hostname, underscorePattern) ||
		containsPattern(hostname, consecutiveHyphensPattern) {
		return errors.New("hostname contains invalid characters")
	}
	return nil
}

func containsPattern(hostname, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(hostname)
}
