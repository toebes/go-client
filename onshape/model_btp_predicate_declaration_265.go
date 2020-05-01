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

// BTPPredicateDeclaration265 struct for BTPPredicateDeclaration265
type BTPPredicateDeclaration265 struct {
	BTPFunctionOrPredicateDeclaration247
	BtType *string `json:"btType,omitempty"`
}

// NewBTPPredicateDeclaration265 instantiates a new BTPPredicateDeclaration265 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPPredicateDeclaration265() *BTPPredicateDeclaration265 {
	this := BTPPredicateDeclaration265{}
	return &this
}

// NewBTPPredicateDeclaration265WithDefaults instantiates a new BTPPredicateDeclaration265 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPPredicateDeclaration265WithDefaults() *BTPPredicateDeclaration265 {
	this := BTPPredicateDeclaration265{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPPredicateDeclaration265) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPPredicateDeclaration265) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPPredicateDeclaration265) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPPredicateDeclaration265) SetBtType(v string) {
	o.BtType = &v
}

func (o BTPPredicateDeclaration265) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPFunctionOrPredicateDeclaration247, errBTPFunctionOrPredicateDeclaration247 := json.Marshal(o.BTPFunctionOrPredicateDeclaration247)
	if errBTPFunctionOrPredicateDeclaration247 != nil {
		return []byte{}, errBTPFunctionOrPredicateDeclaration247
	}
	errBTPFunctionOrPredicateDeclaration247 = json.Unmarshal([]byte(serializedBTPFunctionOrPredicateDeclaration247), &toSerialize)
	if errBTPFunctionOrPredicateDeclaration247 != nil {
		return []byte{}, errBTPFunctionOrPredicateDeclaration247
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTPPredicateDeclaration265 struct {
	value *BTPPredicateDeclaration265
	isSet bool
}

func (v NullableBTPPredicateDeclaration265) Get() *BTPPredicateDeclaration265 {
	return v.value
}

func (v *NullableBTPPredicateDeclaration265) Set(val *BTPPredicateDeclaration265) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPPredicateDeclaration265) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPPredicateDeclaration265) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPPredicateDeclaration265(val *BTPPredicateDeclaration265) *NullableBTPPredicateDeclaration265 {
	return &NullableBTPPredicateDeclaration265{value: val, isSet: true}
}

func (v NullableBTPPredicateDeclaration265) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPPredicateDeclaration265) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
