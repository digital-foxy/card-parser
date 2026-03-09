package character

import "github.com/digital-foxy/card-parser/property"

// Asset asset structure of a V3 chara card
type Asset struct {
	Type      property.String `json:"type"`
	URI       property.String `json:"uri"`
	Name      property.String `json:"name"`
	Extension property.String `json:"ext"`
}
