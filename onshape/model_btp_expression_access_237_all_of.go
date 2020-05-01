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

// BTPExpressionAccess237AllOf struct for BTPExpressionAccess237AllOf
type BTPExpressionAccess237AllOf struct {
	Accessor *BTPPropertyAccessor23 `json:"accessor,omitempty"`
	Base *BTPExpression9 `json:"base,omitempty"`
	BtType *string `json:"btType,omitempty"`
	SpaceInAccessor *BTPSpace10 `json:"spaceInAccessor,omitempty"`
}

// NewBTPExpressionAccess237AllOf instantiates a new BTPExpressionAccess237AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionAccess237AllOf() *BTPExpressionAccess237AllOf {
	this := BTPExpressionAccess237AllOf{}
	return &this
}

// NewBTPExpressionAccess237AllOfWithDefaults instantiates a new BTPExpressionAccess237AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionAccess237AllOfWithDefaults() *BTPExpressionAccess237AllOf {
	this := BTPExpressionAccess237AllOf{}
	return &this
}

// GetAccessor returns the Accessor field value if set, zero value otherwise.
func (o *BTPExpressionAccess237AllOf) GetAccessor() BTPPropertyAccessor23 {
	if o == nil || o.Accessor == nil {
		var ret BTPPropertyAccessor23
		return ret
	}
	return *o.Accessor
}

// GetAccessorOk returns a tuple with the Accessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionAccess237AllOf) GetAccessorOk() (*BTPPropertyAccessor23, bool) {
	if o == nil || o.Accessor == nil {
		return nil, false
	}
	return o.Accessor, true
}

// HasAccessor returns a boolean if a field has been set.
func (o *BTPExpressionAccess237AllOf) HasAccessor() bool {
	if o != nil && o.Accessor != nil {
		return true
	}

	return false
}

// SetAccessor gets a reference to the given BTPPropertyAccessor23 and assigns it to the Accessor field.
func (o *BTPExpressionAccess237AllOf) SetAccessor(v BTPPropertyAccessor23) {
	o.Accessor = &v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *BTPExpressionAccess237AllOf) GetBase() BTPExpression9 {
	if o == nil || o.Base == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionAccess237AllOf) GetBaseOk() (*BTPExpression9, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *BTPExpressionAccess237AllOf) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given BTPExpression9 and assigns it to the Base field.
func (o *BTPExpressionAccess237AllOf) SetBase(v BTPExpression9) {
	o.Base = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionAccess237AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionAccess237AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionAccess237AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionAccess237AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetSpaceInAccessor returns the SpaceInAccessor field value if set, zero value otherwise.
func (o *BTPExpressionAccess237AllOf) GetSpaceInAccessor() BTPSpace10 {
	if o == nil || o.SpaceInAccessor == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInAccessor
}

// GetSpaceInAccessorOk returns a tuple with the SpaceInAccessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionAccess237AllOf) GetSpaceInAccessorOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInAccessor == nil {
		return nil, false
	}
	return o.SpaceInAccessor, true
}

// HasSpaceInAccessor returns a boolean if a field has been set.
func (o *BTPExpressionAccess237AllOf) HasSpaceInAccessor() bool {
	if o != nil && o.SpaceInAccessor != nil {
		return true
	}

	return false
}

// SetSpaceInAccessor gets a reference to the given BTPSpace10 and assigns it to the SpaceInAccessor field.
func (o *BTPExpressionAccess237AllOf) SetSpaceInAccessor(v BTPSpace10) {
	o.SpaceInAccessor = &v
}

func (o BTPExpressionAccess237AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessor != nil {
		toSerialize["accessor"] = o.Accessor
	}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SpaceInAccessor != nil {
		toSerialize["spaceInAccessor"] = o.SpaceInAccessor
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionAccess237AllOf struct {
	value *BTPExpressionAccess237AllOf
	isSet bool
}

func (v NullableBTPExpressionAccess237AllOf) Get() *BTPExpressionAccess237AllOf {
	return v.value
}

func (v *NullableBTPExpressionAccess237AllOf) Set(val *BTPExpressionAccess237AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionAccess237AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionAccess237AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionAccess237AllOf(val *BTPExpressionAccess237AllOf) *NullableBTPExpressionAccess237AllOf {
	return &NullableBTPExpressionAccess237AllOf{value: val, isSet: true}
}

func (v NullableBTPExpressionAccess237AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionAccess237AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
