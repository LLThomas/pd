// Copyright 2022 TiKV Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"time"

	"github.com/stretchr/testify/require"
)

const (
	defaultWaitFor       = time.Second * 20
	defaultSleepInterval = time.Millisecond * 100
)

// WaitOp represents available options when execute Eventually.
type WaitOp struct {
	waitFor       time.Duration
	sleepInterval time.Duration
}

// WaitOption configures WaitOp
type WaitOption func(op *WaitOp)

// WithSleepInterval specify the sleep duration
func WithSleepInterval(sleep time.Duration) WaitOption {
	return func(op *WaitOp) { op.sleepInterval = sleep }
}

// WithWaitFor specify the max wait for duration
func WithWaitFor(waitFor time.Duration) WaitOption {
	return func(op *WaitOp) { op.waitFor = waitFor }
}

// Eventually asserts that given condition will be met in a period of time.
func Eventually(re *require.Assertions, condition func() bool, opts ...WaitOption) {
	option := &WaitOp{
		waitFor:       defaultWaitFor,
		sleepInterval: defaultSleepInterval,
	}
	for _, opt := range opts {
		opt(option)
	}
	re.Eventually(
		condition,
		option.waitFor,
		option.sleepInterval,
	)
}
