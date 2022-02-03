// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

type DeactivatePredefinedRules struct {
	PredefinedRulesName []byte
}

func (d *DeactivatePredefinedRules) MarshalBinary() (data []byte, err error) {
	// Octet 5 to (n+4)
	data = d.PredefinedRulesName

	return data, nil
}

func (d *DeactivatePredefinedRules) UnmarshalBinary(data []byte) error {
	// Octet 5 to (n+4)
	d.PredefinedRulesName = data

	return nil
}
