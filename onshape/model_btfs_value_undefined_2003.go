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

// BTFSValueUndefined2003 struct for BTFSValueUndefined2003
type BTFSValueUndefined2003 struct {
	BTFSValue1888
	BtType *string `json:"btType,omitempty"`
}

// NewBTFSValueUndefined2003 instantiates a new BTFSValueUndefined2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueUndefined2003() *BTFSValueUndefined2003 {
	this := BTFSValueUndefined2003{}
	return &this
}

// NewBTFSValueUndefined2003WithDefaults instantiates a new BTFSValueUndefined2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueUndefined2003WithDefaults() *BTFSValueUndefined2003 {
	this := BTFSValueUndefined2003{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSValueUndefined2003) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueUndefined2003) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSValueUndefined2003) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSValueUndefined2003) SetBtType(v string) {
	o.BtType = &v
}

func (o BTFSValueUndefined2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTFSValue1888, errBTFSValue1888 := json.Marshal(o.BTFSValue1888)
	if errBTFSValue1888 != nil {
		return []byte{}, errBTFSValue1888
	}
	errBTFSValue1888 = json.Unmarshal([]byte(serializedBTFSValue1888), &toSerialize)
	if errBTFSValue1888 != nil {
		return []byte{}, errBTFSValue1888
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueUndefined2003 struct {
	value *BTFSValueUndefined2003
	isSet bool
}

func (v NullableBTFSValueUndefined2003) Get() *BTFSValueUndefined2003 {
	return v.value
}

func (v *NullableBTFSValueUndefined2003) Set(val *BTFSValueUndefined2003) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueUndefined2003) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueUndefined2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueUndefined2003(val *BTFSValueUndefined2003) *NullableBTFSValueUndefined2003 {
	return &NullableBTFSValueUndefined2003{value: val, isSet: true}
}

func (v NullableBTFSValueUndefined2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueUndefined2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
