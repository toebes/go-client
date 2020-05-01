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

// BTPBuiltinIdentifier233 struct for BTPBuiltinIdentifier233
type BTPBuiltinIdentifier233 struct {
	BTPNode7
	BtType *string `json:"btType,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// NewBTPBuiltinIdentifier233 instantiates a new BTPBuiltinIdentifier233 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPBuiltinIdentifier233() *BTPBuiltinIdentifier233 {
	this := BTPBuiltinIdentifier233{}
	return &this
}

// NewBTPBuiltinIdentifier233WithDefaults instantiates a new BTPBuiltinIdentifier233 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPBuiltinIdentifier233WithDefaults() *BTPBuiltinIdentifier233 {
	this := BTPBuiltinIdentifier233{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPBuiltinIdentifier233) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPBuiltinIdentifier233) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPBuiltinIdentifier233) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPBuiltinIdentifier233) SetBtType(v string) {
	o.BtType = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *BTPBuiltinIdentifier233) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPBuiltinIdentifier233) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *BTPBuiltinIdentifier233) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *BTPBuiltinIdentifier233) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o BTPBuiltinIdentifier233) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPNode7, errBTPNode7 := json.Marshal(o.BTPNode7)
	if errBTPNode7 != nil {
		return []byte{}, errBTPNode7
	}
	errBTPNode7 = json.Unmarshal([]byte(serializedBTPNode7), &toSerialize)
	if errBTPNode7 != nil {
		return []byte{}, errBTPNode7
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableBTPBuiltinIdentifier233 struct {
	value *BTPBuiltinIdentifier233
	isSet bool
}

func (v NullableBTPBuiltinIdentifier233) Get() *BTPBuiltinIdentifier233 {
	return v.value
}

func (v *NullableBTPBuiltinIdentifier233) Set(val *BTPBuiltinIdentifier233) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPBuiltinIdentifier233) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPBuiltinIdentifier233) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPBuiltinIdentifier233(val *BTPBuiltinIdentifier233) *NullableBTPBuiltinIdentifier233 {
	return &NullableBTPBuiltinIdentifier233{value: val, isSet: true}
}

func (v NullableBTPBuiltinIdentifier233) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPBuiltinIdentifier233) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}