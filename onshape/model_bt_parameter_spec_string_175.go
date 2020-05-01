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
)

// BTParameterSpecString175 struct for BTParameterSpecString175
type BTParameterSpecString175 struct {
	BTParameterSpec6
	BtType *string `json:"btType,omitempty"`
	FormatConditions *[]BTStringFormatCondition683 `json:"formatConditions,omitempty"`
}

// NewBTParameterSpecString175 instantiates a new BTParameterSpecString175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecString175() *BTParameterSpecString175 {
	this := BTParameterSpecString175{}
	return &this
}

// NewBTParameterSpecString175WithDefaults instantiates a new BTParameterSpecString175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecString175WithDefaults() *BTParameterSpecString175 {
	this := BTParameterSpecString175{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecString175) SetBtType(v string) {
	o.BtType = &v
}

// GetFormatConditions returns the FormatConditions field value if set, zero value otherwise.
func (o *BTParameterSpecString175) GetFormatConditions() []BTStringFormatCondition683 {
	if o == nil || o.FormatConditions == nil {
		var ret []BTStringFormatCondition683
		return ret
	}
	return *o.FormatConditions
}

// GetFormatConditionsOk returns a tuple with the FormatConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecString175) GetFormatConditionsOk() (*[]BTStringFormatCondition683, bool) {
	if o == nil || o.FormatConditions == nil {
		return nil, false
	}
	return o.FormatConditions, true
}

// HasFormatConditions returns a boolean if a field has been set.
func (o *BTParameterSpecString175) HasFormatConditions() bool {
	if o != nil && o.FormatConditions != nil {
		return true
	}

	return false
}

// SetFormatConditions gets a reference to the given []BTStringFormatCondition683 and assigns it to the FormatConditions field.
func (o *BTParameterSpecString175) SetFormatConditions(v []BTStringFormatCondition683) {
	o.FormatConditions = &v
}

func (o BTParameterSpecString175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterSpec6, errBTParameterSpec6 := json.Marshal(o.BTParameterSpec6)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	errBTParameterSpec6 = json.Unmarshal([]byte(serializedBTParameterSpec6), &toSerialize)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FormatConditions != nil {
		toSerialize["formatConditions"] = o.FormatConditions
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecString175 struct {
	value *BTParameterSpecString175
	isSet bool
}

func (v NullableBTParameterSpecString175) Get() *BTParameterSpecString175 {
	return v.value
}

func (v *NullableBTParameterSpecString175) Set(val *BTParameterSpecString175) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecString175) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecString175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecString175(val *BTParameterSpecString175) *NullableBTParameterSpecString175 {
	return &NullableBTParameterSpecString175{value: val, isSet: true}
}

func (v NullableBTParameterSpecString175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecString175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}