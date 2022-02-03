// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

type ApplicationID struct {
	ApplicationIdentifier []byte
}

func (a *ApplicationID) MarshalBinary() (data []byte, err error) {
	// Octet 5 to (n+4)
	data = a.ApplicationIdentifier

	return data, nil
}

func (a *ApplicationID) UnmarshalBinary(data []byte) error {
	// Octet 5 to (n+4)
	a.ApplicationIdentifier = data

	return nil
}
