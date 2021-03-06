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

// BTStringFormatBlockPattern1755 struct for BTStringFormatBlockPattern1755
type BTStringFormatBlockPattern1755 struct {
	BTStringFormatCondition683
	BtType *string `json:"btType,omitempty"`
	RegExpToBlock *string `json:"regExpToBlock,omitempty"`
}

// NewBTStringFormatBlockPattern1755 instantiates a new BTStringFormatBlockPattern1755 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringFormatBlockPattern1755() *BTStringFormatBlockPattern1755 {
	this := BTStringFormatBlockPattern1755{}
	return &this
}

// NewBTStringFormatBlockPattern1755WithDefaults instantiates a new BTStringFormatBlockPattern1755 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringFormatBlockPattern1755WithDefaults() *BTStringFormatBlockPattern1755 {
	this := BTStringFormatBlockPattern1755{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringFormatBlockPattern1755) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatBlockPattern1755) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringFormatBlockPattern1755) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringFormatBlockPattern1755) SetBtType(v string) {
	o.BtType = &v
}

// GetRegExpToBlock returns the RegExpToBlock field value if set, zero value otherwise.
func (o *BTStringFormatBlockPattern1755) GetRegExpToBlock() string {
	if o == nil || o.RegExpToBlock == nil {
		var ret string
		return ret
	}
	return *o.RegExpToBlock
}

// GetRegExpToBlockOk returns a tuple with the RegExpToBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatBlockPattern1755) GetRegExpToBlockOk() (*string, bool) {
	if o == nil || o.RegExpToBlock == nil {
		return nil, false
	}
	return o.RegExpToBlock, true
}

// HasRegExpToBlock returns a boolean if a field has been set.
func (o *BTStringFormatBlockPattern1755) HasRegExpToBlock() bool {
	if o != nil && o.RegExpToBlock != nil {
		return true
	}

	return false
}

// SetRegExpToBlock gets a reference to the given string and assigns it to the RegExpToBlock field.
func (o *BTStringFormatBlockPattern1755) SetRegExpToBlock(v string) {
	o.RegExpToBlock = &v
}

func (o BTStringFormatBlockPattern1755) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTStringFormatCondition683, errBTStringFormatCondition683 := json.Marshal(o.BTStringFormatCondition683)
	if errBTStringFormatCondition683 != nil {
		return []byte{}, errBTStringFormatCondition683
	}
	errBTStringFormatCondition683 = json.Unmarshal([]byte(serializedBTStringFormatCondition683), &toSerialize)
	if errBTStringFormatCondition683 != nil {
		return []byte{}, errBTStringFormatCondition683
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.RegExpToBlock != nil {
		toSerialize["regExpToBlock"] = o.RegExpToBlock
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringFormatBlockPattern1755 struct {
	value *BTStringFormatBlockPattern1755
	isSet bool
}

func (v NullableBTStringFormatBlockPattern1755) Get() *BTStringFormatBlockPattern1755 {
	return v.value
}

func (v *NullableBTStringFormatBlockPattern1755) Set(val *BTStringFormatBlockPattern1755) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringFormatBlockPattern1755) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringFormatBlockPattern1755) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringFormatBlockPattern1755(val *BTStringFormatBlockPattern1755) *NullableBTStringFormatBlockPattern1755 {
	return &NullableBTStringFormatBlockPattern1755{value: val, isSet: true}
}

func (v NullableBTStringFormatBlockPattern1755) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringFormatBlockPattern1755) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
