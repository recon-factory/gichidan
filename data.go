// Copyright 2017 hIMEI

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	//"time"
)

// Host struct is a basic data type
type Host struct {
	// HostUrl is an url of host
	HostUrl string
	// AddDate is an date in which host was added to Ichidan index
	AddDate string
	// Services on host
	Services []*Service
}

// Service contains all info about found Host
type Service struct {
	Name        string
	Port        string
	Protocol    string
	State       string
	Version     string
	ServDetails string
}

// NewService is a constructor for Service struct
func NewService(fields []string) *Service {
	service := &Service{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]}

	return service
}

// NewHost is a constructor for Host struct
func NewHost(fields []string, services []*Service) *Host {
	host := &Host{fields[0], fields[1], services}

	return host
}

// String is a Stringer implementation for Service to output
func (s *Service) String() string {
	return fmt.Sprintf("%s\n %s\n %s\n %s\n %s\n %s\n",
		s.Name, s.Port, s.Protocol, s.State, s.Version, s.ServDetails)
}

// String is a Stringer implementation for Host to output
func (h *Host) String() string {
	var servs string

	for _, s := range h.Services {
		servs += s.String() + "\n"
	}

	return fmt.Sprintf("%s\n %s\n %s\n", h.HostUrl, h.AddDate, servs)
}
