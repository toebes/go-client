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

// BTParameterSpecAppearance1740 struct for BTParameterSpecAppearance1740
type BTParameterSpecAppearance1740 struct {
	BTParameterSpec6
	BtType *string `json:"btType,omitempty"`
}

// NewBTParameterSpecAppearance1740 instantiates a new BTParameterSpecAppearance1740 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecAppearance1740() *BTParameterSpecAppearance1740 {
	this := BTParameterSpecAppearance1740{}
	return &this
}

// NewBTParameterSpecAppearance1740WithDefaults instantiates a new BTParameterSpecAppearance1740 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecAppearance1740WithDefaults() *BTParameterSpecAppearance1740 {
	this := BTParameterSpecAppearance1740{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecAppearance1740) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecAppearance1740) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecAppearance1740) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecAppearance1740) SetBtType(v string) {
	o.BtType = &v
}

func (o BTParameterSpecAppearance1740) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecAppearance1740 struct {
	value *BTParameterSpecAppearance1740
	isSet bool
}

func (v NullableBTParameterSpecAppearance1740) Get() *BTParameterSpecAppearance1740 {
	return v.value
}

func (v *NullableBTParameterSpecAppearance1740) Set(val *BTParameterSpecAppearance1740) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecAppearance1740) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecAppearance1740) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecAppearance1740(val *BTParameterSpecAppearance1740) *NullableBTParameterSpecAppearance1740 {
	return &NullableBTParameterSpecAppearance1740{value: val, isSet: true}
}

func (v NullableBTParameterSpecAppearance1740) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecAppearance1740) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
