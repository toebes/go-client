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

// BTStringMinimumLengthPattern895 struct for BTStringMinimumLengthPattern895
type BTStringMinimumLengthPattern895 struct {
	BTStringFormatCondition683
	BtType *string `json:"btType,omitempty"`
	MinimumLength *int32 `json:"minimumLength,omitempty"`
}

// NewBTStringMinimumLengthPattern895 instantiates a new BTStringMinimumLengthPattern895 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringMinimumLengthPattern895() *BTStringMinimumLengthPattern895 {
	this := BTStringMinimumLengthPattern895{}
	return &this
}

// NewBTStringMinimumLengthPattern895WithDefaults instantiates a new BTStringMinimumLengthPattern895 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringMinimumLengthPattern895WithDefaults() *BTStringMinimumLengthPattern895 {
	this := BTStringMinimumLengthPattern895{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringMinimumLengthPattern895) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMinimumLengthPattern895) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringMinimumLengthPattern895) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringMinimumLengthPattern895) SetBtType(v string) {
	o.BtType = &v
}

// GetMinimumLength returns the MinimumLength field value if set, zero value otherwise.
func (o *BTStringMinimumLengthPattern895) GetMinimumLength() int32 {
	if o == nil || o.MinimumLength == nil {
		var ret int32
		return ret
	}
	return *o.MinimumLength
}

// GetMinimumLengthOk returns a tuple with the MinimumLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMinimumLengthPattern895) GetMinimumLengthOk() (*int32, bool) {
	if o == nil || o.MinimumLength == nil {
		return nil, false
	}
	return o.MinimumLength, true
}

// HasMinimumLength returns a boolean if a field has been set.
func (o *BTStringMinimumLengthPattern895) HasMinimumLength() bool {
	if o != nil && o.MinimumLength != nil {
		return true
	}

	return false
}

// SetMinimumLength gets a reference to the given int32 and assigns it to the MinimumLength field.
func (o *BTStringMinimumLengthPattern895) SetMinimumLength(v int32) {
	o.MinimumLength = &v
}

func (o BTStringMinimumLengthPattern895) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTStringFormatCondition683, errBTStringFormatCondition683 := json.Marshal(o.BTStringFormatCondition683)
	if errBTStringFormatCondition683 != nil {
		return []byte{}, errBTStringFormatCondition683
	}
	errBTStringFormatCondition683 = json.Unmarshal([]byte(serializedBTStringFormatCondition683), &toSerialize)
	if errBTStringFormatCondition683 != nil {
		return []byte{}, errBTStringFormatCondition683
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.MinimumLength != nil {
		toSerialize["minimumLength"] = o.MinimumLength
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringMinimumLengthPattern895 struct {
	value *BTStringMinimumLengthPattern895
	isSet bool
}

func (v NullableBTStringMinimumLengthPattern895) Get() *BTStringMinimumLengthPattern895 {
	return v.value
}

func (v *NullableBTStringMinimumLengthPattern895) Set(val *BTStringMinimumLengthPattern895) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringMinimumLengthPattern895) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringMinimumLengthPattern895) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringMinimumLengthPattern895(val *BTStringMinimumLengthPattern895) *NullableBTStringMinimumLengthPattern895 {
	return &NullableBTStringMinimumLengthPattern895{value: val, isSet: true}
}

func (v NullableBTStringMinimumLengthPattern895) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringMinimumLengthPattern895) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
