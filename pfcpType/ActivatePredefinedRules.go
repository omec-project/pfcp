// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

type ActivatePredefinedRules struct {
	PredefinedRulesName []byte
}

func (a *ActivatePredefinedRules) MarshalBinary() (data []byte, err error) {
	// Octet 5 to (n+4)
	data = a.PredefinedRulesName

	return data, nil
}

func (a *ActivatePredefinedRules) UnmarshalBinary(data []byte) error {
	// Octet 5 to (n+4)
	a.PredefinedRulesName = data

	return nil
}
