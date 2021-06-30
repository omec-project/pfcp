// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalOuterHeaderRemoval(t *testing.T) {
	testData := OuterHeaderRemoval{
		OuterHeaderRemovalDescription: OuterHeaderRemovalGtpUUdpIpv4,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{OuterHeaderRemovalGtpUUdpIpv4}, buf)
}

func TestUnmarshalOuterHeaderRemoval(t *testing.T) {
	buf := []byte{OuterHeaderRemovalGtpUUdpIpv4}
	var testData OuterHeaderRemoval
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := OuterHeaderRemoval{
		OuterHeaderRemovalDescription: OuterHeaderRemovalGtpUUdpIpv4,
	}
	assert.Equal(t, expectData, testData)
}
