// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalBARID(t *testing.T) {
	testData := BARID{
		BarIdValue: 21,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{21}, buf)
}

func TestUnmarshalBARID(t *testing.T) {
	buf := []byte{21}
	var testData BARID
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := BARID{
		BarIdValue: 21,
	}
	assert.Equal(t, expectData, testData)
}
