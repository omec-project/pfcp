// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalFailedRuleID(t *testing.T) {
	testData := FailedRuleID{
		RuleIdType:  2,
		RuleIdValue: []byte{12, 34},
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{2, 12, 34}, buf)
}

func TestUnmarshalFailedRuleID(t *testing.T) {
	buf := []byte{2, 12, 34}
	var testData FailedRuleID
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := FailedRuleID{
		RuleIdType:  2,
		RuleIdValue: []byte{12, 34},
	}
	assert.Equal(t, expectData, testData)
}
