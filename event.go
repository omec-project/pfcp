// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcp

type EventType uint8

const (
	ReceiveResendRequest EventType = iota
	ReceiveValidResponse
)
