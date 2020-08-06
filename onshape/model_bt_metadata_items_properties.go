/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMetadataItemsProperties This is the same as BTWorkflowPropertyInfo with the addition of the multivalued field and dropping the dirty, initialValue, isApproverProperty and is NotifierProperty fields
type BTMetadataItemsProperties struct {
	BTMetadataItemsPropertiesInterface interface { GetValueType() string }
}

func (s BTMetadataItemsProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.BTMetadataItemsPropertiesInterface)
}

func (s *BTMetadataItemsProperties) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaled map[string]interface{}
	err = json.Unmarshal(src, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["valueType"]; ok {
		switch v {
			case "BOOL":
				var result *BTMetadataCommonBool = &BTMetadataCommonBool{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonBool":
				var result *BTMetadataCommonBool = &BTMetadataCommonBool{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonCategory":
				var result *BTMetadataCommonCategory = &BTMetadataCommonCategory{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonComputed":
				var result *BTMetadataCommonComputed = &BTMetadataCommonComputed{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonDate":
				var result *BTMetadataCommonDate = &BTMetadataCommonDate{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonEnum":
				var result *BTMetadataCommonEnum = &BTMetadataCommonEnum{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonObject":
				var result *BTMetadataCommonObject = &BTMetadataCommonObject{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonString":
				var result *BTMetadataCommonString = &BTMetadataCommonString{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "BTMetadataCommonUser":
				var result *BTMetadataCommonUser = &BTMetadataCommonUser{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "CATEGORY":
				var result *BTMetadataCommonCategory = &BTMetadataCommonCategory{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "COMPUTED":
				var result *BTMetadataCommonComputed = &BTMetadataCommonComputed{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "DATE":
				var result *BTMetadataCommonDate = &BTMetadataCommonDate{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "ENUM":
				var result *BTMetadataCommonEnum = &BTMetadataCommonEnum{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "OBJECT":
				var result *BTMetadataCommonObject = &BTMetadataCommonObject{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "STRING":
				var result *BTMetadataCommonString = &BTMetadataCommonString{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			case "USER":
				var result *BTMetadataCommonUser = &BTMetadataCommonUser{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.BTMetadataItemsPropertiesInterface = result
				return nil
			default:
				return fmt.Errorf("No oneOf model has 'valueType' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'valueType' not found in unmarshaled payload: %+v", unmarshaled)
	}
}
type NullableBTMetadataItemsProperties struct {
	value *BTMetadataItemsProperties
	isSet bool
}

func (v NullableBTMetadataItemsProperties) Get() *BTMetadataItemsProperties {
	return v.value
}

func (v *NullableBTMetadataItemsProperties) Set(val *BTMetadataItemsProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataItemsProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataItemsProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataItemsProperties(val *BTMetadataItemsProperties) *NullableBTMetadataItemsProperties {
	return &NullableBTMetadataItemsProperties{value: val, isSet: true}
}

func (v NullableBTMetadataItemsProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataItemsProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
