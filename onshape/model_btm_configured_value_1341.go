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

// BTMConfiguredValue1341 struct for BTMConfiguredValue1341
type BTMConfiguredValue1341 struct {
	BtType *string `json:"btType,omitempty"`
	ConfigurationValueString *string `json:"configurationValueString,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
	Value *BTMParameter1 `json:"value,omitempty"`
}

// NewBTMConfiguredValue1341 instantiates a new BTMConfiguredValue1341 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfiguredValue1341() *BTMConfiguredValue1341 {
	this := BTMConfiguredValue1341{}
	return &this
}

// NewBTMConfiguredValue1341WithDefaults instantiates a new BTMConfiguredValue1341 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfiguredValue1341WithDefaults() *BTMConfiguredValue1341 {
	this := BTMConfiguredValue1341{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfiguredValue1341) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationValueString returns the ConfigurationValueString field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetConfigurationValueString() string {
	if o == nil || o.ConfigurationValueString == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationValueString
}

// GetConfigurationValueStringOk returns a tuple with the ConfigurationValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetConfigurationValueStringOk() (*string, bool) {
	if o == nil || o.ConfigurationValueString == nil {
		return nil, false
	}
	return o.ConfigurationValueString, true
}

// HasConfigurationValueString returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasConfigurationValueString() bool {
	if o != nil && o.ConfigurationValueString != nil {
		return true
	}

	return false
}

// SetConfigurationValueString gets a reference to the given string and assigns it to the ConfigurationValueString field.
func (o *BTMConfiguredValue1341) SetConfigurationValueString(v string) {
	o.ConfigurationValueString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfiguredValue1341) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfiguredValue1341) SetNodeId(v string) {
	o.NodeId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetValue() BTMParameter1 {
	if o == nil || o.Value == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetValueOk() (*BTMParameter1, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTMParameter1 and assigns it to the Value field.
func (o *BTMConfiguredValue1341) SetValue(v BTMParameter1) {
	o.Value = &v
}

func (o BTMConfiguredValue1341) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationValueString != nil {
		toSerialize["configurationValueString"] = o.ConfigurationValueString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfiguredValue1341 struct {
	value *BTMConfiguredValue1341
	isSet bool
}

func (v NullableBTMConfiguredValue1341) Get() *BTMConfiguredValue1341 {
	return v.value
}

func (v *NullableBTMConfiguredValue1341) Set(val *BTMConfiguredValue1341) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfiguredValue1341) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfiguredValue1341) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfiguredValue1341(val *BTMConfiguredValue1341) *NullableBTMConfiguredValue1341 {
	return &NullableBTMConfiguredValue1341{value: val, isSet: true}
}

func (v NullableBTMConfiguredValue1341) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfiguredValue1341) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
