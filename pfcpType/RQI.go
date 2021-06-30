// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RQI struct {
	RQI bool
}

func (r *RQI) MarshalBinary() ([]byte, error) {
	var buf = &bytes.Buffer{}

	if err := buf.WriteByte(btou(r.RQI)); err != nil {
		return nil, fmt.Errorf("marshal RQI fail: " + err.Error())
	}

	return buf.Bytes(), nil
}

func (r *RQI) UnmarshalBinary(data []byte) error {
	var buf = bytes.NewBuffer(data)
	var tmpByte byte

	if err := binary.Read(buf, binary.BigEndian, &tmpByte); err != nil {
		return fmt.Errorf("unmarshal RQI fail: " + err.Error())
	}

	r.RQI = utob(tmpByte)

	return nil
}
