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

// BTAppElementContentDeltaInfo struct for BTAppElementContentDeltaInfo
type BTAppElementContentDeltaInfo struct {
	Delta *string `json:"delta,omitempty"`
}

// NewBTAppElementContentDeltaInfo instantiates a new BTAppElementContentDeltaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementContentDeltaInfo() *BTAppElementContentDeltaInfo {
	this := BTAppElementContentDeltaInfo{}
	return &this
}

// NewBTAppElementContentDeltaInfoWithDefaults instantiates a new BTAppElementContentDeltaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementContentDeltaInfoWithDefaults() *BTAppElementContentDeltaInfo {
	this := BTAppElementContentDeltaInfo{}
	return &this
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *BTAppElementContentDeltaInfo) GetDelta() string {
	if o == nil || o.Delta == nil {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentDeltaInfo) GetDeltaOk() (*string, bool) {
	if o == nil || o.Delta == nil {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *BTAppElementContentDeltaInfo) HasDelta() bool {
	if o != nil && o.Delta != nil {
		return true
	}

	return false
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *BTAppElementContentDeltaInfo) SetDelta(v string) {
	o.Delta = &v
}

func (o BTAppElementContentDeltaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delta != nil {
		toSerialize["delta"] = o.Delta
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementContentDeltaInfo struct {
	value *BTAppElementContentDeltaInfo
	isSet bool
}

func (v NullableBTAppElementContentDeltaInfo) Get() *BTAppElementContentDeltaInfo {
	return v.value
}

func (v *NullableBTAppElementContentDeltaInfo) Set(val *BTAppElementContentDeltaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementContentDeltaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementContentDeltaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementContentDeltaInfo(val *BTAppElementContentDeltaInfo) *NullableBTAppElementContentDeltaInfo {
	return &NullableBTAppElementContentDeltaInfo{value: val, isSet: true}
}

func (v NullableBTAppElementContentDeltaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementContentDeltaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}