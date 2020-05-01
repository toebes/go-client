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

// BTMParameterConfigured2222AllOf struct for BTMParameterConfigured2222AllOf
type BTMParameterConfigured2222AllOf struct {
	BtType *string `json:"btType,omitempty"`
	ConfigurationParameterId *string `json:"configurationParameterId,omitempty"`
	ConfigurationParameterIdFieldIndex *int32 `json:"configurationParameterIdFieldIndex,omitempty"`
	Values *[]BTMConfiguredValue1341 `json:"values,omitempty"`
	ValuesFieldIndex *int32 `json:"valuesFieldIndex,omitempty"`
}

// NewBTMParameterConfigured2222AllOf instantiates a new BTMParameterConfigured2222AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterConfigured2222AllOf() *BTMParameterConfigured2222AllOf {
	this := BTMParameterConfigured2222AllOf{}
	return &this
}

// NewBTMParameterConfigured2222AllOfWithDefaults instantiates a new BTMParameterConfigured2222AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterConfigured2222AllOfWithDefaults() *BTMParameterConfigured2222AllOf {
	this := BTMParameterConfigured2222AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterConfigured2222AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterConfigured2222AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterConfigured2222AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterConfigured2222AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationParameterId returns the ConfigurationParameterId field value if set, zero value otherwise.
func (o *BTMParameterConfigured2222AllOf) GetConfigurationParameterId() string {
	if o == nil || o.ConfigurationParameterId == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationParameterId
}

// GetConfigurationParameterIdOk returns a tuple with the ConfigurationParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterConfigured2222AllOf) GetConfigurationParameterIdOk() (*string, bool) {
	if o == nil || o.ConfigurationParameterId == nil {
		return nil, false
	}
	return o.ConfigurationParameterId, true
}

// HasConfigurationParameterId returns a boolean if a field has been set.
func (o *BTMParameterConfigured2222AllOf) HasConfigurationParameterId() bool {
	if o != nil && o.ConfigurationParameterId != nil {
		return true
	}

	return false
}

// SetConfigurationParameterId gets a reference to the given string and assigns it to the ConfigurationParameterId field.
func (o *BTMParameterConfigured2222AllOf) SetConfigurationParameterId(v string) {
	o.ConfigurationParameterId = &v
}

// GetConfigurationParameterIdFieldIndex returns the ConfigurationParameterIdFieldIndex field value if set, zero value otherwise.
func (o *BTMParameterConfigured2222AllOf) GetConfigurationParameterIdFieldIndex() int32 {
	if o == nil || o.ConfigurationParameterIdFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.ConfigurationParameterIdFieldIndex
}

// GetConfigurationParameterIdFieldIndexOk returns a tuple with the ConfigurationParameterIdFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterConfigured2222AllOf) GetConfigurationParameterIdFieldIndexOk() (*int32, bool) {
	if o == nil || o.ConfigurationParameterIdFieldIndex == nil {
		return nil, false
	}
	return o.ConfigurationParameterIdFieldIndex, true
}

// HasConfigurationParameterIdFieldIndex returns a boolean if a field has been set.
func (o *BTMParameterConfigured2222AllOf) HasConfigurationParameterIdFieldIndex() bool {
	if o != nil && o.ConfigurationParameterIdFieldIndex != nil {
		return true
	}

	return false
}

// SetConfigurationParameterIdFieldIndex gets a reference to the given int32 and assigns it to the ConfigurationParameterIdFieldIndex field.
func (o *BTMParameterConfigured2222AllOf) SetConfigurationParameterIdFieldIndex(v int32) {
	o.ConfigurationParameterIdFieldIndex = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *BTMParameterConfigured2222AllOf) GetValues() []BTMConfiguredValue1341 {
	if o == nil || o.Values == nil {
		var ret []BTMConfiguredValue1341
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterConfigured2222AllOf) GetValuesOk() (*[]BTMConfiguredValue1341, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *BTMParameterConfigured2222AllOf) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []BTMConfiguredValue1341 and assigns it to the Values field.
func (o *BTMParameterConfigured2222AllOf) SetValues(v []BTMConfiguredValue1341) {
	o.Values = &v
}

// GetValuesFieldIndex returns the ValuesFieldIndex field value if set, zero value otherwise.
func (o *BTMParameterConfigured2222AllOf) GetValuesFieldIndex() int32 {
	if o == nil || o.ValuesFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.ValuesFieldIndex
}

// GetValuesFieldIndexOk returns a tuple with the ValuesFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterConfigured2222AllOf) GetValuesFieldIndexOk() (*int32, bool) {
	if o == nil || o.ValuesFieldIndex == nil {
		return nil, false
	}
	return o.ValuesFieldIndex, true
}

// HasValuesFieldIndex returns a boolean if a field has been set.
func (o *BTMParameterConfigured2222AllOf) HasValuesFieldIndex() bool {
	if o != nil && o.ValuesFieldIndex != nil {
		return true
	}

	return false
}

// SetValuesFieldIndex gets a reference to the given int32 and assigns it to the ValuesFieldIndex field.
func (o *BTMParameterConfigured2222AllOf) SetValuesFieldIndex(v int32) {
	o.ValuesFieldIndex = &v
}

func (o BTMParameterConfigured2222AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationParameterId != nil {
		toSerialize["configurationParameterId"] = o.ConfigurationParameterId
	}
	if o.ConfigurationParameterIdFieldIndex != nil {
		toSerialize["configurationParameterIdFieldIndex"] = o.ConfigurationParameterIdFieldIndex
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.ValuesFieldIndex != nil {
		toSerialize["valuesFieldIndex"] = o.ValuesFieldIndex
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterConfigured2222AllOf struct {
	value *BTMParameterConfigured2222AllOf
	isSet bool
}

func (v NullableBTMParameterConfigured2222AllOf) Get() *BTMParameterConfigured2222AllOf {
	return v.value
}

func (v *NullableBTMParameterConfigured2222AllOf) Set(val *BTMParameterConfigured2222AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterConfigured2222AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterConfigured2222AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterConfigured2222AllOf(val *BTMParameterConfigured2222AllOf) *NullableBTMParameterConfigured2222AllOf {
	return &NullableBTMParameterConfigured2222AllOf{value: val, isSet: true}
}

func (v NullableBTMParameterConfigured2222AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterConfigured2222AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
