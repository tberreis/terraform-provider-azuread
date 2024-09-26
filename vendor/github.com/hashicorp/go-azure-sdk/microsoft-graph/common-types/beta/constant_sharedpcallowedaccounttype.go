package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCAllowedAccountType string

const (
	SharedPCAllowedAccountType_Domain        SharedPCAllowedAccountType = "domain"
	SharedPCAllowedAccountType_Guest         SharedPCAllowedAccountType = "guest"
	SharedPCAllowedAccountType_NotConfigured SharedPCAllowedAccountType = "notConfigured"
)

func PossibleValuesForSharedPCAllowedAccountType() []string {
	return []string{
		string(SharedPCAllowedAccountType_Domain),
		string(SharedPCAllowedAccountType_Guest),
		string(SharedPCAllowedAccountType_NotConfigured),
	}
}

func (s *SharedPCAllowedAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCAllowedAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCAllowedAccountType(input string) (*SharedPCAllowedAccountType, error) {
	vals := map[string]SharedPCAllowedAccountType{
		"domain":        SharedPCAllowedAccountType_Domain,
		"guest":         SharedPCAllowedAccountType_Guest,
		"notconfigured": SharedPCAllowedAccountType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCAllowedAccountType(input)
	return &out, nil
}