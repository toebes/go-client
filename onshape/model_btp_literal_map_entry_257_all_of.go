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

// BTPLiteralMapEntry257AllOf struct for BTPLiteralMapEntry257AllOf
type BTPLiteralMapEntry257AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Key *BTPPropertyAccessor23 `json:"key,omitempty"`
	Value *BTPExpression9 `json:"value,omitempty"`
}

// NewBTPLiteralMapEntry257AllOf instantiates a new BTPLiteralMapEntry257AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLiteralMapEntry257AllOf() *BTPLiteralMapEntry257AllOf {
	this := BTPLiteralMapEntry257AllOf{}
	return &this
}

// NewBTPLiteralMapEntry257AllOfWithDefaults instantiates a new BTPLiteralMapEntry257AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLiteralMapEntry257AllOfWithDefaults() *BTPLiteralMapEntry257AllOf {
	this := BTPLiteralMapEntry257AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLiteralMapEntry257AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257AllOf) GetKey() BTPPropertyAccessor23 {
	if o == nil || o.Key == nil {
		var ret BTPPropertyAccessor23
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257AllOf) GetKeyOk() (*BTPPropertyAccessor23, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257AllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given BTPPropertyAccessor23 and assigns it to the Key field.
func (o *BTPLiteralMapEntry257AllOf) SetKey(v BTPPropertyAccessor23) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTPLiteralMapEntry257AllOf) GetValue() BTPExpression9 {
	if o == nil || o.Value == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralMapEntry257AllOf) GetValueOk() (*BTPExpression9, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTPLiteralMapEntry257AllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTPExpression9 and assigns it to the Value field.
func (o *BTPLiteralMapEntry257AllOf) SetValue(v BTPExpression9) {
	o.Value = &v
}

func (o BTPLiteralMapEntry257AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTPLiteralMapEntry257AllOf struct {
	value *BTPLiteralMapEntry257AllOf
	isSet bool
}

func (v NullableBTPLiteralMapEntry257AllOf) Get() *BTPLiteralMapEntry257AllOf {
	return v.value
}

func (v *NullableBTPLiteralMapEntry257AllOf) Set(val *BTPLiteralMapEntry257AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLiteralMapEntry257AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLiteralMapEntry257AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLiteralMapEntry257AllOf(val *BTPLiteralMapEntry257AllOf) *NullableBTPLiteralMapEntry257AllOf {
	return &NullableBTPLiteralMapEntry257AllOf{value: val, isSet: true}
}

func (v NullableBTPLiteralMapEntry257AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLiteralMapEntry257AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
