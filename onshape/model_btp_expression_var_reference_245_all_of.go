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

// BTPExpressionVarReference245AllOf struct for BTPExpressionVarReference245AllOf
type BTPExpressionVarReference245AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Name *BTPName261 `json:"name,omitempty"`
}

// NewBTPExpressionVarReference245AllOf instantiates a new BTPExpressionVarReference245AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionVarReference245AllOf() *BTPExpressionVarReference245AllOf {
	this := BTPExpressionVarReference245AllOf{}
	return &this
}

// NewBTPExpressionVarReference245AllOfWithDefaults instantiates a new BTPExpressionVarReference245AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionVarReference245AllOfWithDefaults() *BTPExpressionVarReference245AllOf {
	this := BTPExpressionVarReference245AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionVarReference245AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionVarReference245AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionVarReference245AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionVarReference245AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPExpressionVarReference245AllOf) GetName() BTPName261 {
	if o == nil || o.Name == nil {
		var ret BTPName261
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionVarReference245AllOf) GetNameOk() (*BTPName261, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPExpressionVarReference245AllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given BTPName261 and assigns it to the Name field.
func (o *BTPExpressionVarReference245AllOf) SetName(v BTPName261) {
	o.Name = &v
}

func (o BTPExpressionVarReference245AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionVarReference245AllOf struct {
	value *BTPExpressionVarReference245AllOf
	isSet bool
}

func (v NullableBTPExpressionVarReference245AllOf) Get() *BTPExpressionVarReference245AllOf {
	return v.value
}

func (v *NullableBTPExpressionVarReference245AllOf) Set(val *BTPExpressionVarReference245AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionVarReference245AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionVarReference245AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionVarReference245AllOf(val *BTPExpressionVarReference245AllOf) *NullableBTPExpressionVarReference245AllOf {
	return &NullableBTPExpressionVarReference245AllOf{value: val, isSet: true}
}

func (v NullableBTPExpressionVarReference245AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionVarReference245AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}