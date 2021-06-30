// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalTrafficEndpointID(t *testing.T) {
	testData := TrafficEndpointID{
		TrafficEndpointIdValue: 21,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{21}, buf)
}

func TestUnmarshalTrafficEndpointID(t *testing.T) {
	buf := []byte{21}
	var testData TrafficEndpointID
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := TrafficEndpointID{
		TrafficEndpointIdValue: 21,
	}
	assert.Equal(t, expectData, testData)
}
