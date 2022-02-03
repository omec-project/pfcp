// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalNodeID(t *testing.T) {
	testData := NodeID{
		NodeIdType:  NodeIdTypeFqdn,
		NodeIdValue: []byte("free5gc.local"),
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{2, 102, 114, 101, 101, 53, 103, 99, 46, 108, 111, 99, 97, 108}, buf)
}

func TestUnmarshalNodeID(t *testing.T) {
	buf := []byte{2, 102, 114, 101, 101, 53, 103, 99, 46, 108, 111, 99, 97, 108}
	var testData NodeID
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := NodeID{
		NodeIdType:  NodeIdTypeFqdn,
		NodeIdValue: []byte("free5gc.local"),
	}
	assert.Equal(t, expectData, testData)
}
