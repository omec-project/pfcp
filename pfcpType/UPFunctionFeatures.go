// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

import (
	"encoding/binary"
	"fmt"
)

const (
	UpFunctionFeaturesBucp  uint16 = 1
	UpFunctionFeaturesDdnd  uint16 = 1 << 1
	UpFunctionFeaturesDlbd  uint16 = 1 << 2
	UpFunctionFeaturesTrst  uint16 = 1 << 3
	UpFunctionFeaturesFtup  uint16 = 1 << 4
	UpFunctionFeaturesPfdm  uint16 = 1 << 5
	UpFunctionFeaturesHeeu  uint16 = 1 << 6
	UpFunctionFeaturesTreu  uint16 = 1 << 7
	UpFunctionFeaturesEmpu  uint16 = 1 << 8
	UpFunctionFeaturesPdiu  uint16 = 1 << 9
	UpFunctionFeaturesUdbc  uint16 = 1 << 10
	UpFunctionFeaturesQuoac uint16 = 1 << 11
	UpFunctionFeaturesTrace uint16 = 1 << 12
	UpFunctionFeaturesFrrt  uint16 = 1 << 13
	UpFunctionFeaturesPfde  uint16 = 1 << 14
	UpFunctionFeaturesEpfar uint16 = 1 << 15
)

//Supported Feature-1
const (
	UpFunctionFeatures1Dprda uint16 = 1
	UpFunctionFeatures1Adpdp uint16 = 1 << 1
	UpFunctionFeatures1Ueip  uint16 = 1 << 2
	UpFunctionFeatures1Sset  uint16 = 1 << 3
	UpFunctionFeatures1Mnop  uint16 = 1 << 4
	UpFunctionFeatures1Mte   uint16 = 1 << 5
	UpFunctionFeatures1Bundl uint16 = 1 << 6
	UpFunctionFeatures1Gcom  uint16 = 1 << 7
	UpFunctionFeatures1Mpas  uint16 = 1 << 8
	UpFunctionFeatures1Rttl  uint16 = 1 << 9
	UpFunctionFeatures1Vtime uint16 = 1 << 10
	UpFunctionFeatures1Norp  uint16 = 1 << 11
	UpFunctionFeatures1Iptv  uint16 = 1 << 12
	UpFunctionFeatures1Ip6pl uint16 = 1 << 13
	UpFunctionFeatures1Tscu  uint16 = 1 << 14
	UpFunctionFeatures1Mptcp uint16 = 1 << 15
)

//Supported Feature-2
const (
	UpFunctionFeatures2Atsssll uint16 = 1
	UpFunctionFeatures2Qfqm    uint16 = 1 << 1
	UpFunctionFeatures2Gpqm    uint16 = 1 << 2
	UpFunctionFeatures2Mtedt   uint16 = 1 << 3
	UpFunctionFeatures2Ciot    uint16 = 1 << 4
	UpFunctionFeatures2Ethar   uint16 = 1 << 5
	UpFunctionFeatures2Ddds    uint16 = 1 << 6
	UpFunctionFeatures2Rds     uint16 = 1 << 7
	UpFunctionFeatures2Rttwp   uint16 = 1 << 8
)

type UPFunctionFeatures struct {
	SupportedFeatures uint16
	SupportedFeatures1 uint16
	SupportedFeatures2 uint16
}

func (u *UPFunctionFeatures) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 6
	data = make([]byte, 6)
	binary.LittleEndian.PutUint16(data[idx:], u.SupportedFeatures)

	// Octet 7 to 8 : Additional Supported-Features 1
	idx = 2
	binary.LittleEndian.PutUint16(data[idx:], u.SupportedFeatures1)

	/// Octet 9 to 10 : Additional Supported-Features 2
	idx = 4
	binary.LittleEndian.PutUint16(data[idx:], u.SupportedFeatures2)
	return data, nil
}

func (u *UPFunctionFeatures) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 6
	if length < idx+2 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	//Additional Supported-Features
	if length >= 2 {
		u.SupportedFeatures = binary.LittleEndian.Uint16(data[idx : idx+2])
	}

	if length >= 4 {
		//Additional Supported-Features 1
		idx += 2
		u.SupportedFeatures1 = binary.LittleEndian.Uint16(data[idx : idx+2])
	}

	if length == 6 {
		//Additional Supported-Features 2
		idx += 2
		u.SupportedFeatures2 = binary.LittleEndian.Uint16(data[idx : idx+2])
	}
	return nil
}

