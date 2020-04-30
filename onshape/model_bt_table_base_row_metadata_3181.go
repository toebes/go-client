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

// BTTableBaseRowMetadata3181 struct for BTTableBaseRowMetadata3181
type BTTableBaseRowMetadata3181 struct {
	BtType *string `json:"btType,omitempty"`
	CrossHighlightDataIfAny *BTTableBaseCrossHighlightData2609 `json:"crossHighlightDataIfAny,omitempty"`
}

// NewBTTableBaseRowMetadata3181 instantiates a new BTTableBaseRowMetadata3181 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableBaseRowMetadata3181() *BTTableBaseRowMetadata3181 {
	this := BTTableBaseRowMetadata3181{}
	return &this
}

// NewBTTableBaseRowMetadata3181WithDefaults instantiates a new BTTableBaseRowMetadata3181 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableBaseRowMetadata3181WithDefaults() *BTTableBaseRowMetadata3181 {
	this := BTTableBaseRowMetadata3181{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableBaseRowMetadata3181) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableBaseRowMetadata3181) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableBaseRowMetadata3181) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableBaseRowMetadata3181) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTTableBaseRowMetadata3181) GetCrossHighlightDataIfAny() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableBaseRowMetadata3181) GetCrossHighlightDataIfAnyOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTTableBaseRowMetadata3181) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTTableBaseRowMetadata3181) SetCrossHighlightDataIfAny(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightDataIfAny = &v
}

func (o BTTableBaseRowMetadata3181) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableBaseRowMetadata3181 struct {
	value *BTTableBaseRowMetadata3181
	isSet bool
}

func (v NullableBTTableBaseRowMetadata3181) Get() *BTTableBaseRowMetadata3181 {
	return v.value
}

func (v *NullableBTTableBaseRowMetadata3181) Set(val *BTTableBaseRowMetadata3181) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableBaseRowMetadata3181) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableBaseRowMetadata3181) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableBaseRowMetadata3181(val *BTTableBaseRowMetadata3181) *NullableBTTableBaseRowMetadata3181 {
	return &NullableBTTableBaseRowMetadata3181{value: val, isSet: true}
}

func (v NullableBTTableBaseRowMetadata3181) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableBaseRowMetadata3181) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
