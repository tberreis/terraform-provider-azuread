package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOrganizationalScope struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// List of groups to which the custom detection rule applies.
	ScopeNames *[]string `json:"scopeNames,omitempty"`

	ScopeType *SecurityScopeType `json:"scopeType,omitempty"`
}