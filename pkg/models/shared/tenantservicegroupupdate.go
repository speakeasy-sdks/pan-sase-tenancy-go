// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TenantServiceGroupUpdateVertical - A token that identifies the business vertical supported by the SASE
// products managed by this TSG.
type TenantServiceGroupUpdateVertical string

const (
	TenantServiceGroupUpdateVerticalHighTech                      TenantServiceGroupUpdateVertical = "High Tech"
	TenantServiceGroupUpdateVerticalEducation                     TenantServiceGroupUpdateVertical = "Education"
	TenantServiceGroupUpdateVerticalManufacturing                 TenantServiceGroupUpdateVertical = "Manufacturing"
	TenantServiceGroupUpdateVerticalHospitality                   TenantServiceGroupUpdateVertical = "Hospitality"
	TenantServiceGroupUpdateVerticalProfessionalAndLegalServices  TenantServiceGroupUpdateVertical = "Professional & Legal Services"
	TenantServiceGroupUpdateVerticalWholesaleAndRetail            TenantServiceGroupUpdateVertical = "Wholesale & Retail"
	TenantServiceGroupUpdateVerticalFinance                       TenantServiceGroupUpdateVertical = "Finance"
	TenantServiceGroupUpdateVerticalTelecommunications            TenantServiceGroupUpdateVertical = "Telecommunications"
	TenantServiceGroupUpdateVerticalStateAndLocalGovernment       TenantServiceGroupUpdateVertical = "State & Local Government"
	TenantServiceGroupUpdateVerticalTransportationAndLogistics    TenantServiceGroupUpdateVertical = "Transportation & Logistics"
	TenantServiceGroupUpdateVerticalFederalGovernment             TenantServiceGroupUpdateVertical = "Federal Government"
	TenantServiceGroupUpdateVerticalMediaAndEntertainment         TenantServiceGroupUpdateVertical = "Media & Entertainment"
	TenantServiceGroupUpdateVerticalNonclassifiableEstablishments TenantServiceGroupUpdateVertical = "Nonclassifiable Establishments"
	TenantServiceGroupUpdateVerticalHealthcare                    TenantServiceGroupUpdateVertical = "Healthcare"
	TenantServiceGroupUpdateVerticalUtilitiesAndEnergy            TenantServiceGroupUpdateVertical = "Utilities & Energy"
	TenantServiceGroupUpdateVerticalInsurance                     TenantServiceGroupUpdateVertical = "Insurance"
	TenantServiceGroupUpdateVerticalAgriculture                   TenantServiceGroupUpdateVertical = "Agriculture"
	TenantServiceGroupUpdateVerticalPharmaAndLifeSciences         TenantServiceGroupUpdateVertical = "Pharma & Life Sciences"
	TenantServiceGroupUpdateVerticalConstruction                  TenantServiceGroupUpdateVertical = "Construction"
	TenantServiceGroupUpdateVerticalAerospaceAndDefense           TenantServiceGroupUpdateVertical = "Aerospace & Defense"
	TenantServiceGroupUpdateVerticalRealEstate                    TenantServiceGroupUpdateVertical = "Real Estate"
	TenantServiceGroupUpdateVerticalRestaurantFoodIndustry        TenantServiceGroupUpdateVertical = "Restaurant/Food Industry"
	TenantServiceGroupUpdateVerticalOther                         TenantServiceGroupUpdateVertical = "Other"
)

func (e TenantServiceGroupUpdateVertical) ToPointer() *TenantServiceGroupUpdateVertical {
	return &e
}

func (e *TenantServiceGroupUpdateVertical) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "High Tech":
		fallthrough
	case "Education":
		fallthrough
	case "Manufacturing":
		fallthrough
	case "Hospitality":
		fallthrough
	case "Professional & Legal Services":
		fallthrough
	case "Wholesale & Retail":
		fallthrough
	case "Finance":
		fallthrough
	case "Telecommunications":
		fallthrough
	case "State & Local Government":
		fallthrough
	case "Transportation & Logistics":
		fallthrough
	case "Federal Government":
		fallthrough
	case "Media & Entertainment":
		fallthrough
	case "Nonclassifiable Establishments":
		fallthrough
	case "Healthcare":
		fallthrough
	case "Utilities & Energy":
		fallthrough
	case "Insurance":
		fallthrough
	case "Agriculture":
		fallthrough
	case "Pharma & Life Sciences":
		fallthrough
	case "Construction":
		fallthrough
	case "Aerospace & Defense":
		fallthrough
	case "Real Estate":
		fallthrough
	case "Restaurant/Food Industry":
		fallthrough
	case "Other":
		*e = TenantServiceGroupUpdateVertical(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TenantServiceGroupUpdateVertical: %v", v)
	}
}

type TenantServiceGroupUpdate struct {
	// The tenant service group's display name.
	//
	DisplayName *string `json:"display_name,omitempty"`
	// The email address of the person or organization that should
	// be contacted for support of this TSG.
	//
	SupportContact *string `json:"support_contact,omitempty"`
	// A token that identifies the business vertical supported by the SASE
	// products managed by this TSG.
	//
	Vertical *TenantServiceGroupUpdateVertical `json:"vertical,omitempty"`
}