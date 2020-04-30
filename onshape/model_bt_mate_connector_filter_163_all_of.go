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

// BTMateConnectorFilter163AllOf struct for BTMateConnectorFilter163AllOf
type BTMateConnectorFilter163AllOf struct {
	BtType *string `json:"btType,omitempty"`
	RequiresOccurrence *bool `json:"requiresOccurrence,omitempty"`
}

// NewBTMateConnectorFilter163AllOf instantiates a new BTMateConnectorFilter163AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMateConnectorFilter163AllOf() *BTMateConnectorFilter163AllOf {
	this := BTMateConnectorFilter163AllOf{}
	return &this
}

// NewBTMateConnectorFilter163AllOfWithDefaults instantiates a new BTMateConnectorFilter163AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMateConnectorFilter163AllOfWithDefaults() *BTMateConnectorFilter163AllOf {
	this := BTMateConnectorFilter163AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMateConnectorFilter163AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorFilter163AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMateConnectorFilter163AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMateConnectorFilter163AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetRequiresOccurrence returns the RequiresOccurrence field value if set, zero value otherwise.
func (o *BTMateConnectorFilter163AllOf) GetRequiresOccurrence() bool {
	if o == nil || o.RequiresOccurrence == nil {
		var ret bool
		return ret
	}
	return *o.RequiresOccurrence
}

// GetRequiresOccurrenceOk returns a tuple with the RequiresOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorFilter163AllOf) GetRequiresOccurrenceOk() (*bool, bool) {
	if o == nil || o.RequiresOccurrence == nil {
		return nil, false
	}
	return o.RequiresOccurrence, true
}

// HasRequiresOccurrence returns a boolean if a field has been set.
func (o *BTMateConnectorFilter163AllOf) HasRequiresOccurrence() bool {
	if o != nil && o.RequiresOccurrence != nil {
		return true
	}

	return false
}

// SetRequiresOccurrence gets a reference to the given bool and assigns it to the RequiresOccurrence field.
func (o *BTMateConnectorFilter163AllOf) SetRequiresOccurrence(v bool) {
	o.RequiresOccurrence = &v
}

func (o BTMateConnectorFilter163AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.RequiresOccurrence != nil {
		toSerialize["requiresOccurrence"] = o.RequiresOccurrence
	}
	return json.Marshal(toSerialize)
}

type NullableBTMateConnectorFilter163AllOf struct {
	value *BTMateConnectorFilter163AllOf
	isSet bool
}

func (v NullableBTMateConnectorFilter163AllOf) Get() *BTMateConnectorFilter163AllOf {
	return v.value
}

func (v *NullableBTMateConnectorFilter163AllOf) Set(val *BTMateConnectorFilter163AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMateConnectorFilter163AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMateConnectorFilter163AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMateConnectorFilter163AllOf(val *BTMateConnectorFilter163AllOf) *NullableBTMateConnectorFilter163AllOf {
	return &NullableBTMateConnectorFilter163AllOf{value: val, isSet: true}
}

func (v NullableBTMateConnectorFilter163AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMateConnectorFilter163AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
