// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpUdp

import (
	"net"

	"github.com/omec-project/pfcp"
)

type Message struct {
	RemoteAddr  *net.UDPAddr
	PfcpMessage *pfcp.Message
	EventData   interface{}
}

func NewMessage(remoteAddr *net.UDPAddr, pfcpMessage *pfcp.Message, eventData interface{}) (msg Message) {
	msg = Message{}
	msg.RemoteAddr = remoteAddr
	msg.PfcpMessage = pfcpMessage
	msg.EventData = eventData
	return
}
