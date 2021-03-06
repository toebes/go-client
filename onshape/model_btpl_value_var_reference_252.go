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

// BTPLValueVarReference252 struct for BTPLValueVarReference252
type BTPLValueVarReference252 struct {
	BTPLValue249
	BtType *string `json:"btType,omitempty"`
	Name *BTPIdentifier8 `json:"name,omitempty"`
}

// NewBTPLValueVarReference252 instantiates a new BTPLValueVarReference252 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLValueVarReference252() *BTPLValueVarReference252 {
	this := BTPLValueVarReference252{}
	return &this
}

// NewBTPLValueVarReference252WithDefaults instantiates a new BTPLValueVarReference252 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLValueVarReference252WithDefaults() *BTPLValueVarReference252 {
	this := BTPLValueVarReference252{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLValueVarReference252) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueVarReference252) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLValueVarReference252) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLValueVarReference252) SetBtType(v string) {
	o.BtType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPLValueVarReference252) GetName() BTPIdentifier8 {
	if o == nil || o.Name == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueVarReference252) GetNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPLValueVarReference252) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given BTPIdentifier8 and assigns it to the Name field.
func (o *BTPLValueVarReference252) SetName(v BTPIdentifier8) {
	o.Name = &v
}

func (o BTPLValueVarReference252) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPLValue249, errBTPLValue249 := json.Marshal(o.BTPLValue249)
	if errBTPLValue249 != nil {
		return []byte{}, errBTPLValue249
	}
	errBTPLValue249 = json.Unmarshal([]byte(serializedBTPLValue249), &toSerialize)
	if errBTPLValue249 != nil {
		return []byte{}, errBTPLValue249
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTPLValueVarReference252 struct {
	value *BTPLValueVarReference252
	isSet bool
}

func (v NullableBTPLValueVarReference252) Get() *BTPLValueVarReference252 {
	return v.value
}

func (v *NullableBTPLValueVarReference252) Set(val *BTPLValueVarReference252) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLValueVarReference252) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLValueVarReference252) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLValueVarReference252(val *BTPLValueVarReference252) *NullableBTPLValueVarReference252 {
	return &NullableBTPLValueVarReference252{value: val, isSet: true}
}

func (v NullableBTPLValueVarReference252) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLValueVarReference252) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
