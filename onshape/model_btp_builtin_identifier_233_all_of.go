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

// BTPBuiltinIdentifier233AllOf struct for BTPBuiltinIdentifier233AllOf
type BTPBuiltinIdentifier233AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// NewBTPBuiltinIdentifier233AllOf instantiates a new BTPBuiltinIdentifier233AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPBuiltinIdentifier233AllOf() *BTPBuiltinIdentifier233AllOf {
	this := BTPBuiltinIdentifier233AllOf{}
	return &this
}

// NewBTPBuiltinIdentifier233AllOfWithDefaults instantiates a new BTPBuiltinIdentifier233AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPBuiltinIdentifier233AllOfWithDefaults() *BTPBuiltinIdentifier233AllOf {
	this := BTPBuiltinIdentifier233AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPBuiltinIdentifier233AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPBuiltinIdentifier233AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPBuiltinIdentifier233AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPBuiltinIdentifier233AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *BTPBuiltinIdentifier233AllOf) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPBuiltinIdentifier233AllOf) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *BTPBuiltinIdentifier233AllOf) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *BTPBuiltinIdentifier233AllOf) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o BTPBuiltinIdentifier233AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableBTPBuiltinIdentifier233AllOf struct {
	value *BTPBuiltinIdentifier233AllOf
	isSet bool
}

func (v NullableBTPBuiltinIdentifier233AllOf) Get() *BTPBuiltinIdentifier233AllOf {
	return v.value
}

func (v *NullableBTPBuiltinIdentifier233AllOf) Set(val *BTPBuiltinIdentifier233AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPBuiltinIdentifier233AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPBuiltinIdentifier233AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPBuiltinIdentifier233AllOf(val *BTPBuiltinIdentifier233AllOf) *NullableBTPBuiltinIdentifier233AllOf {
	return &NullableBTPBuiltinIdentifier233AllOf{value: val, isSet: true}
}

func (v NullableBTPBuiltinIdentifier233AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPBuiltinIdentifier233AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
