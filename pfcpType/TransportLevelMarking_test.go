// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalTransportLevelMarking(t *testing.T) {
	testData := TransportLevelMarking{
		TosTrafficClass: []byte{43, 21},
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{43, 21}, buf)
}

func TestUnmarshalTransportLevelMarking(t *testing.T) {
	buf := []byte{43, 21}
	var testData TransportLevelMarking
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := TransportLevelMarking{
		TosTrafficClass: []byte{43, 21},
	}
	assert.Equal(t, expectData, testData)
}
