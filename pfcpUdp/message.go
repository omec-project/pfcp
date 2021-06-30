// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpUdp

import (
	"net"

	"github.com/free5gc/pfcp"
)

type Message struct {
	RemoteAddr  *net.UDPAddr
	PfcpMessage *pfcp.Message
}

func NewMessage(remoteAddr *net.UDPAddr, pfcpMessage *pfcp.Message) (msg Message) {
	msg = Message{}
	msg.RemoteAddr = remoteAddr
	msg.PfcpMessage = pfcpMessage
	return
}
