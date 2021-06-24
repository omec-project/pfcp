// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalDestinationInterface(t *testing.T) {
	testData := DestinationInterface{
		InterfaceValue: DestinationInterfaceCore,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{DestinationInterfaceCore}, buf)
}

func TestUnmarshalDestinationInterface(t *testing.T) {
	buf := []byte{DestinationInterfaceCore}
	var testData DestinationInterface
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := DestinationInterface{
		InterfaceValue: DestinationInterfaceCore,
	}
	assert.Equal(t, expectData, testData)
}
