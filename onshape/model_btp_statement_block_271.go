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

// BTPStatementBlock271 struct for BTPStatementBlock271
type BTPStatementBlock271 struct {
	BTPStatement269
	BtType *string `json:"btType,omitempty"`
	SpaceAfterOpen *BTPSpace10 `json:"spaceAfterOpen,omitempty"`
}

// NewBTPStatementBlock271 instantiates a new BTPStatementBlock271 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementBlock271() *BTPStatementBlock271 {
	this := BTPStatementBlock271{}
	return &this
}

// NewBTPStatementBlock271WithDefaults instantiates a new BTPStatementBlock271 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementBlock271WithDefaults() *BTPStatementBlock271 {
	this := BTPStatementBlock271{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementBlock271) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementBlock271) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementBlock271) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementBlock271) SetBtType(v string) {
	o.BtType = &v
}

// GetSpaceAfterOpen returns the SpaceAfterOpen field value if set, zero value otherwise.
func (o *BTPStatementBlock271) GetSpaceAfterOpen() BTPSpace10 {
	if o == nil || o.SpaceAfterOpen == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterOpen
}

// GetSpaceAfterOpenOk returns a tuple with the SpaceAfterOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementBlock271) GetSpaceAfterOpenOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterOpen == nil {
		return nil, false
	}
	return o.SpaceAfterOpen, true
}

// HasSpaceAfterOpen returns a boolean if a field has been set.
func (o *BTPStatementBlock271) HasSpaceAfterOpen() bool {
	if o != nil && o.SpaceAfterOpen != nil {
		return true
	}

	return false
}

// SetSpaceAfterOpen gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterOpen field.
func (o *BTPStatementBlock271) SetSpaceAfterOpen(v BTPSpace10) {
	o.SpaceAfterOpen = &v
}

func (o BTPStatementBlock271) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPStatement269, errBTPStatement269 := json.Marshal(o.BTPStatement269)
	if errBTPStatement269 != nil {
		return []byte{}, errBTPStatement269
	}
	errBTPStatement269 = json.Unmarshal([]byte(serializedBTPStatement269), &toSerialize)
	if errBTPStatement269 != nil {
		return []byte{}, errBTPStatement269
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SpaceAfterOpen != nil {
		toSerialize["spaceAfterOpen"] = o.SpaceAfterOpen
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementBlock271 struct {
	value *BTPStatementBlock271
	isSet bool
}

func (v NullableBTPStatementBlock271) Get() *BTPStatementBlock271 {
	return v.value
}

func (v *NullableBTPStatementBlock271) Set(val *BTPStatementBlock271) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementBlock271) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementBlock271) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementBlock271(val *BTPStatementBlock271) *NullableBTPStatementBlock271 {
	return &NullableBTPStatementBlock271{value: val, isSet: true}
}

func (v NullableBTPStatementBlock271) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementBlock271) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
