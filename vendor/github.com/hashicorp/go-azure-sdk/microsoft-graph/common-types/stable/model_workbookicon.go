package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookIcon struct {
	// Represents the index of the icon in the given set.
	Index *int64 `json:"index,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the set that the icon is part of. The possible values are: Invalid, ThreeArrows, ThreeArrowsGray,
	// ThreeFlags, ThreeTrafficLights1, ThreeTrafficLights2, ThreeSigns, ThreeSymbols, ThreeSymbols2, FourArrows,
	// FourArrowsGray, FourRedToBlack, FourRating, FourTrafficLights, FiveArrows, FiveArrowsGray, FiveRating, FiveQuarters,
	// ThreeStars, ThreeTriangles, FiveBoxes.
	Set *string `json:"set,omitempty"`
}