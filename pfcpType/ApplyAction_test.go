// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalApplyAction(t *testing.T) {
	testData := ApplyAction{
		Dupl: true,
		Nocp: false,
		Buff: true,
		Forw: false,
		Drop: true,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{21}, buf)
}

func TestUnmarshalApplyAction(t *testing.T) {
	buf := []byte{21}
	var testData ApplyAction
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := ApplyAction{
		Dupl: true,
		Nocp: false,
		Buff: true,
		Forw: false,
		Drop: true,
	}
	assert.Equal(t, expectData, testData)
}
