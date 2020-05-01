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

// BTEntityTypeFilter124AllOf struct for BTEntityTypeFilter124AllOf
type BTEntityTypeFilter124AllOf struct {
	BtType *string `json:"btType,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
}

// NewBTEntityTypeFilter124AllOf instantiates a new BTEntityTypeFilter124AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEntityTypeFilter124AllOf() *BTEntityTypeFilter124AllOf {
	this := BTEntityTypeFilter124AllOf{}
	return &this
}

// NewBTEntityTypeFilter124AllOfWithDefaults instantiates a new BTEntityTypeFilter124AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEntityTypeFilter124AllOfWithDefaults() *BTEntityTypeFilter124AllOf {
	this := BTEntityTypeFilter124AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEntityTypeFilter124AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityTypeFilter124AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEntityTypeFilter124AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEntityTypeFilter124AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BTEntityTypeFilter124AllOf) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityTypeFilter124AllOf) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BTEntityTypeFilter124AllOf) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *BTEntityTypeFilter124AllOf) SetEntityType(v string) {
	o.EntityType = &v
}

func (o BTEntityTypeFilter124AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	return json.Marshal(toSerialize)
}

type NullableBTEntityTypeFilter124AllOf struct {
	value *BTEntityTypeFilter124AllOf
	isSet bool
}

func (v NullableBTEntityTypeFilter124AllOf) Get() *BTEntityTypeFilter124AllOf {
	return v.value
}

func (v *NullableBTEntityTypeFilter124AllOf) Set(val *BTEntityTypeFilter124AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEntityTypeFilter124AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEntityTypeFilter124AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEntityTypeFilter124AllOf(val *BTEntityTypeFilter124AllOf) *NullableBTEntityTypeFilter124AllOf {
	return &NullableBTEntityTypeFilter124AllOf{value: val, isSet: true}
}

func (v NullableBTEntityTypeFilter124AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEntityTypeFilter124AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}