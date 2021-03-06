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

// BTNullableQuantityRange1340 struct for BTNullableQuantityRange1340
type BTNullableQuantityRange1340 struct {
	BTQuantityRange181
	BtType *string `json:"btType,omitempty"`
	HasDefaultValue *bool `json:"hasDefaultValue,omitempty"`
	HasMaxValue *bool `json:"hasMaxValue,omitempty"`
	HasMinValue *bool `json:"hasMinValue,omitempty"`
}

// NewBTNullableQuantityRange1340 instantiates a new BTNullableQuantityRange1340 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNullableQuantityRange1340() *BTNullableQuantityRange1340 {
	this := BTNullableQuantityRange1340{}
	return &this
}

// NewBTNullableQuantityRange1340WithDefaults instantiates a new BTNullableQuantityRange1340 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNullableQuantityRange1340WithDefaults() *BTNullableQuantityRange1340 {
	this := BTNullableQuantityRange1340{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTNullableQuantityRange1340) SetBtType(v string) {
	o.BtType = &v
}

// GetHasDefaultValue returns the HasDefaultValue field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetHasDefaultValue() bool {
	if o == nil || o.HasDefaultValue == nil {
		var ret bool
		return ret
	}
	return *o.HasDefaultValue
}

// GetHasDefaultValueOk returns a tuple with the HasDefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetHasDefaultValueOk() (*bool, bool) {
	if o == nil || o.HasDefaultValue == nil {
		return nil, false
	}
	return o.HasDefaultValue, true
}

// HasHasDefaultValue returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasHasDefaultValue() bool {
	if o != nil && o.HasDefaultValue != nil {
		return true
	}

	return false
}

// SetHasDefaultValue gets a reference to the given bool and assigns it to the HasDefaultValue field.
func (o *BTNullableQuantityRange1340) SetHasDefaultValue(v bool) {
	o.HasDefaultValue = &v
}

// GetHasMaxValue returns the HasMaxValue field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetHasMaxValue() bool {
	if o == nil || o.HasMaxValue == nil {
		var ret bool
		return ret
	}
	return *o.HasMaxValue
}

// GetHasMaxValueOk returns a tuple with the HasMaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetHasMaxValueOk() (*bool, bool) {
	if o == nil || o.HasMaxValue == nil {
		return nil, false
	}
	return o.HasMaxValue, true
}

// HasHasMaxValue returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasHasMaxValue() bool {
	if o != nil && o.HasMaxValue != nil {
		return true
	}

	return false
}

// SetHasMaxValue gets a reference to the given bool and assigns it to the HasMaxValue field.
func (o *BTNullableQuantityRange1340) SetHasMaxValue(v bool) {
	o.HasMaxValue = &v
}

// GetHasMinValue returns the HasMinValue field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetHasMinValue() bool {
	if o == nil || o.HasMinValue == nil {
		var ret bool
		return ret
	}
	return *o.HasMinValue
}

// GetHasMinValueOk returns a tuple with the HasMinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetHasMinValueOk() (*bool, bool) {
	if o == nil || o.HasMinValue == nil {
		return nil, false
	}
	return o.HasMinValue, true
}

// HasHasMinValue returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasHasMinValue() bool {
	if o != nil && o.HasMinValue != nil {
		return true
	}

	return false
}

// SetHasMinValue gets a reference to the given bool and assigns it to the HasMinValue field.
func (o *BTNullableQuantityRange1340) SetHasMinValue(v bool) {
	o.HasMinValue = &v
}

func (o BTNullableQuantityRange1340) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQuantityRange181, errBTQuantityRange181 := json.Marshal(o.BTQuantityRange181)
	if errBTQuantityRange181 != nil {
		return []byte{}, errBTQuantityRange181
	}
	errBTQuantityRange181 = json.Unmarshal([]byte(serializedBTQuantityRange181), &toSerialize)
	if errBTQuantityRange181 != nil {
		return []byte{}, errBTQuantityRange181
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.HasDefaultValue != nil {
		toSerialize["hasDefaultValue"] = o.HasDefaultValue
	}
	if o.HasMaxValue != nil {
		toSerialize["hasMaxValue"] = o.HasMaxValue
	}
	if o.HasMinValue != nil {
		toSerialize["hasMinValue"] = o.HasMinValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTNullableQuantityRange1340 struct {
	value *BTNullableQuantityRange1340
	isSet bool
}

func (v NullableBTNullableQuantityRange1340) Get() *BTNullableQuantityRange1340 {
	return v.value
}

func (v *NullableBTNullableQuantityRange1340) Set(val *BTNullableQuantityRange1340) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNullableQuantityRange1340) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNullableQuantityRange1340) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNullableQuantityRange1340(val *BTNullableQuantityRange1340) *NullableBTNullableQuantityRange1340 {
	return &NullableBTNullableQuantityRange1340{value: val, isSet: true}
}

func (v NullableBTNullableQuantityRange1340) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNullableQuantityRange1340) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
