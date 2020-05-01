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

// BTAssemblyMatedEntity struct for BTAssemblyMatedEntity
type BTAssemblyMatedEntity struct {
	MateCS *BTMateConnectorCSInfo `json:"mateCS,omitempty"`
	MatedOccurrence *[]string `json:"matedOccurrence,omitempty"`
}

// NewBTAssemblyMatedEntity instantiates a new BTAssemblyMatedEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyMatedEntity() *BTAssemblyMatedEntity {
	this := BTAssemblyMatedEntity{}
	return &this
}

// NewBTAssemblyMatedEntityWithDefaults instantiates a new BTAssemblyMatedEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyMatedEntityWithDefaults() *BTAssemblyMatedEntity {
	this := BTAssemblyMatedEntity{}
	return &this
}

// GetMateCS returns the MateCS field value if set, zero value otherwise.
func (o *BTAssemblyMatedEntity) GetMateCS() BTMateConnectorCSInfo {
	if o == nil || o.MateCS == nil {
		var ret BTMateConnectorCSInfo
		return ret
	}
	return *o.MateCS
}

// GetMateCSOk returns a tuple with the MateCS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyMatedEntity) GetMateCSOk() (*BTMateConnectorCSInfo, bool) {
	if o == nil || o.MateCS == nil {
		return nil, false
	}
	return o.MateCS, true
}

// HasMateCS returns a boolean if a field has been set.
func (o *BTAssemblyMatedEntity) HasMateCS() bool {
	if o != nil && o.MateCS != nil {
		return true
	}

	return false
}

// SetMateCS gets a reference to the given BTMateConnectorCSInfo and assigns it to the MateCS field.
func (o *BTAssemblyMatedEntity) SetMateCS(v BTMateConnectorCSInfo) {
	o.MateCS = &v
}

// GetMatedOccurrence returns the MatedOccurrence field value if set, zero value otherwise.
func (o *BTAssemblyMatedEntity) GetMatedOccurrence() []string {
	if o == nil || o.MatedOccurrence == nil {
		var ret []string
		return ret
	}
	return *o.MatedOccurrence
}

// GetMatedOccurrenceOk returns a tuple with the MatedOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyMatedEntity) GetMatedOccurrenceOk() (*[]string, bool) {
	if o == nil || o.MatedOccurrence == nil {
		return nil, false
	}
	return o.MatedOccurrence, true
}

// HasMatedOccurrence returns a boolean if a field has been set.
func (o *BTAssemblyMatedEntity) HasMatedOccurrence() bool {
	if o != nil && o.MatedOccurrence != nil {
		return true
	}

	return false
}

// SetMatedOccurrence gets a reference to the given []string and assigns it to the MatedOccurrence field.
func (o *BTAssemblyMatedEntity) SetMatedOccurrence(v []string) {
	o.MatedOccurrence = &v
}

func (o BTAssemblyMatedEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MateCS != nil {
		toSerialize["mateCS"] = o.MateCS
	}
	if o.MatedOccurrence != nil {
		toSerialize["matedOccurrence"] = o.MatedOccurrence
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyMatedEntity struct {
	value *BTAssemblyMatedEntity
	isSet bool
}

func (v NullableBTAssemblyMatedEntity) Get() *BTAssemblyMatedEntity {
	return v.value
}

func (v *NullableBTAssemblyMatedEntity) Set(val *BTAssemblyMatedEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyMatedEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyMatedEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyMatedEntity(val *BTAssemblyMatedEntity) *NullableBTAssemblyMatedEntity {
	return &NullableBTAssemblyMatedEntity{value: val, isSet: true}
}

func (v NullableBTAssemblyMatedEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyMatedEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
