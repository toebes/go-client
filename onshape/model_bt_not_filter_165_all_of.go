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

// BTNotFilter165AllOf struct for BTNotFilter165AllOf
type BTNotFilter165AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Operand *BTQueryFilter183 `json:"operand,omitempty"`
}

// NewBTNotFilter165AllOf instantiates a new BTNotFilter165AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNotFilter165AllOf() *BTNotFilter165AllOf {
	this := BTNotFilter165AllOf{}
	return &this
}

// NewBTNotFilter165AllOfWithDefaults instantiates a new BTNotFilter165AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNotFilter165AllOfWithDefaults() *BTNotFilter165AllOf {
	this := BTNotFilter165AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTNotFilter165AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotFilter165AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTNotFilter165AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTNotFilter165AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetOperand returns the Operand field value if set, zero value otherwise.
func (o *BTNotFilter165AllOf) GetOperand() BTQueryFilter183 {
	if o == nil || o.Operand == nil {
		var ret BTQueryFilter183
		return ret
	}
	return *o.Operand
}

// GetOperandOk returns a tuple with the Operand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotFilter165AllOf) GetOperandOk() (*BTQueryFilter183, bool) {
	if o == nil || o.Operand == nil {
		return nil, false
	}
	return o.Operand, true
}

// HasOperand returns a boolean if a field has been set.
func (o *BTNotFilter165AllOf) HasOperand() bool {
	if o != nil && o.Operand != nil {
		return true
	}

	return false
}

// SetOperand gets a reference to the given BTQueryFilter183 and assigns it to the Operand field.
func (o *BTNotFilter165AllOf) SetOperand(v BTQueryFilter183) {
	o.Operand = &v
}

func (o BTNotFilter165AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Operand != nil {
		toSerialize["operand"] = o.Operand
	}
	return json.Marshal(toSerialize)
}

type NullableBTNotFilter165AllOf struct {
	value *BTNotFilter165AllOf
	isSet bool
}

func (v NullableBTNotFilter165AllOf) Get() *BTNotFilter165AllOf {
	return v.value
}

func (v *NullableBTNotFilter165AllOf) Set(val *BTNotFilter165AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNotFilter165AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNotFilter165AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNotFilter165AllOf(val *BTNotFilter165AllOf) *NullableBTNotFilter165AllOf {
	return &NullableBTNotFilter165AllOf{value: val, isSet: true}
}

func (v NullableBTNotFilter165AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNotFilter165AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
