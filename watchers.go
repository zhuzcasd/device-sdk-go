// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
//
// SPDX-License-Identifier: Apache-2.0
//
package device

import (
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"sync"
)

type Watchers struct {
	devices map[string]models.Device
}

var (
	wcOnce   sync.Once
	watchers *Watchers
)

// Create a singleton WatcherCache instance
func newWatchers() *Watchers {

	wcOnce.Do(func() {
		watchers = &Watchers{}
	})

	return watchers
}