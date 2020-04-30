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

// BTFSValueString1422 struct for BTFSValueString1422
type BTFSValueString1422 struct {
	BTFSValue1888
	BtType *string `json:"btType,omitempty"`
	Value *string `json:"value,omitempty"`
	ValueObject *string `json:"valueObject,omitempty"`
}

// NewBTFSValueString1422 instantiates a new BTFSValueString1422 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueString1422() *BTFSValueString1422 {
	this := BTFSValueString1422{}
	return &this
}

// NewBTFSValueString1422WithDefaults instantiates a new BTFSValueString1422 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueString1422WithDefaults() *BTFSValueString1422 {
	this := BTFSValueString1422{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSValueString1422) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueString1422) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSValueString1422) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSValueString1422) SetBtType(v string) {
	o.BtType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueString1422) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueString1422) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueString1422) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTFSValueString1422) SetValue(v string) {
	o.Value = &v
}

// GetValueObject returns the ValueObject field value if set, zero value otherwise.
func (o *BTFSValueString1422) GetValueObject() string {
	if o == nil || o.ValueObject == nil {
		var ret string
		return ret
	}
	return *o.ValueObject
}

// GetValueObjectOk returns a tuple with the ValueObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueString1422) GetValueObjectOk() (*string, bool) {
	if o == nil || o.ValueObject == nil {
		return nil, false
	}
	return o.ValueObject, true
}

// HasValueObject returns a boolean if a field has been set.
func (o *BTFSValueString1422) HasValueObject() bool {
	if o != nil && o.ValueObject != nil {
		return true
	}

	return false
}

// SetValueObject gets a reference to the given string and assigns it to the ValueObject field.
func (o *BTFSValueString1422) SetValueObject(v string) {
	o.ValueObject = &v
}

func (o BTFSValueString1422) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueObject != nil {
		toSerialize["valueObject"] = o.ValueObject
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueString1422 struct {
	value *BTFSValueString1422
	isSet bool
}

func (v NullableBTFSValueString1422) Get() *BTFSValueString1422 {
	return v.value
}

func (v *NullableBTFSValueString1422) Set(val *BTFSValueString1422) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueString1422) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueString1422) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueString1422(val *BTFSValueString1422) *NullableBTFSValueString1422 {
	return &NullableBTFSValueString1422{value: val, isSet: true}
}

func (v NullableBTFSValueString1422) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueString1422) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
