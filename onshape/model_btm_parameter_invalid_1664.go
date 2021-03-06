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

// BTMParameterInvalid1664 struct for BTMParameterInvalid1664
type BTMParameterInvalid1664 struct {
	BTMParameter1
	BtType *string `json:"btType,omitempty"`
}

// NewBTMParameterInvalid1664 instantiates a new BTMParameterInvalid1664 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterInvalid1664() *BTMParameterInvalid1664 {
	this := BTMParameterInvalid1664{}
	return &this
}

// NewBTMParameterInvalid1664WithDefaults instantiates a new BTMParameterInvalid1664 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterInvalid1664WithDefaults() *BTMParameterInvalid1664 {
	this := BTMParameterInvalid1664{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterInvalid1664) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterInvalid1664) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterInvalid1664) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterInvalid1664) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMParameterInvalid1664) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMParameter1, errBTMParameter1 := json.Marshal(o.BTMParameter1)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	errBTMParameter1 = json.Unmarshal([]byte(serializedBTMParameter1), &toSerialize)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterInvalid1664 struct {
	value *BTMParameterInvalid1664
	isSet bool
}

func (v NullableBTMParameterInvalid1664) Get() *BTMParameterInvalid1664 {
	return v.value
}

func (v *NullableBTMParameterInvalid1664) Set(val *BTMParameterInvalid1664) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterInvalid1664) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterInvalid1664) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterInvalid1664(val *BTMParameterInvalid1664) *NullableBTMParameterInvalid1664 {
	return &NullableBTMParameterInvalid1664{value: val, isSet: true}
}

func (v NullableBTMParameterInvalid1664) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterInvalid1664) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
