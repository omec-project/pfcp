// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalCause(t *testing.T) {
	testData := Cause{
		CauseValue: CauseRequestRejected,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{CauseRequestRejected}, buf)
}

func TestUnmarshalCause(t *testing.T) {
	buf := []byte{CauseRequestRejected}
	var testData Cause
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := Cause{
		CauseValue: CauseRequestRejected,
	}
	assert.Equal(t, expectData, testData)
}
