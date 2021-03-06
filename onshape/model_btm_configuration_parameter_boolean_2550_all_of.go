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

// BTMConfigurationParameterBoolean2550AllOf struct for BTMConfigurationParameterBoolean2550AllOf
type BTMConfigurationParameterBoolean2550AllOf struct {
	BtType *string `json:"btType,omitempty"`
	DefaultValue *bool `json:"defaultValue,omitempty"`
}

// NewBTMConfigurationParameterBoolean2550AllOf instantiates a new BTMConfigurationParameterBoolean2550AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfigurationParameterBoolean2550AllOf() *BTMConfigurationParameterBoolean2550AllOf {
	this := BTMConfigurationParameterBoolean2550AllOf{}
	return &this
}

// NewBTMConfigurationParameterBoolean2550AllOfWithDefaults instantiates a new BTMConfigurationParameterBoolean2550AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfigurationParameterBoolean2550AllOfWithDefaults() *BTMConfigurationParameterBoolean2550AllOf {
	this := BTMConfigurationParameterBoolean2550AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterBoolean2550AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterBoolean2550AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterBoolean2550AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfigurationParameterBoolean2550AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTMConfigurationParameterBoolean2550AllOf) GetDefaultValue() bool {
	if o == nil || o.DefaultValue == nil {
		var ret bool
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterBoolean2550AllOf) GetDefaultValueOk() (*bool, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTMConfigurationParameterBoolean2550AllOf) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given bool and assigns it to the DefaultValue field.
func (o *BTMConfigurationParameterBoolean2550AllOf) SetDefaultValue(v bool) {
	o.DefaultValue = &v
}

func (o BTMConfigurationParameterBoolean2550AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfigurationParameterBoolean2550AllOf struct {
	value *BTMConfigurationParameterBoolean2550AllOf
	isSet bool
}

func (v NullableBTMConfigurationParameterBoolean2550AllOf) Get() *BTMConfigurationParameterBoolean2550AllOf {
	return v.value
}

func (v *NullableBTMConfigurationParameterBoolean2550AllOf) Set(val *BTMConfigurationParameterBoolean2550AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfigurationParameterBoolean2550AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfigurationParameterBoolean2550AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfigurationParameterBoolean2550AllOf(val *BTMConfigurationParameterBoolean2550AllOf) *NullableBTMConfigurationParameterBoolean2550AllOf {
	return &NullableBTMConfigurationParameterBoolean2550AllOf{value: val, isSet: true}
}

func (v NullableBTMConfigurationParameterBoolean2550AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfigurationParameterBoolean2550AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
