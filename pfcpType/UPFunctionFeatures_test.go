// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUPFunctionFeatures(t *testing.T) {
	testData := UPFunctionFeatures{
		SupportedFeatures: UpFunctionFeaturesTrst | UpFunctionFeaturesUdbc,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{8, 4, 0, 0, 0, 0}, buf)
}

func TestUnmarshalUPFunctionFeatures(t *testing.T) {
	buf := []byte{8, 4}
	var testData UPFunctionFeatures
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := UPFunctionFeatures{
		SupportedFeatures: UpFunctionFeaturesTrst | UpFunctionFeaturesUdbc,
	}
	assert.Equal(t, expectData, testData)
}
