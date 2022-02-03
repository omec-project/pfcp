// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcp

type EventType uint8

const (
	ReceiveResendRequest EventType = iota
	ReceiveValidResponse
)
