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

// BTMConfigurationParameter819 struct for BTMConfigurationParameter819
type BTMConfigurationParameter819 struct {
	BtType *string `json:"btType,omitempty"`
	GeneratedParameterId *BTTreeNode20 `json:"generatedParameterId,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
	ParameterId *string `json:"parameterId,omitempty"`
	ParameterName *string `json:"parameterName,omitempty"`
	ParameterType *string `json:"parameterType,omitempty"`
	Valid *bool `json:"valid,omitempty"`
}

// NewBTMConfigurationParameter819 instantiates a new BTMConfigurationParameter819 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfigurationParameter819() *BTMConfigurationParameter819 {
	this := BTMConfigurationParameter819{}
	return &this
}

// NewBTMConfigurationParameter819WithDefaults instantiates a new BTMConfigurationParameter819 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfigurationParameter819WithDefaults() *BTMConfigurationParameter819 {
	this := BTMConfigurationParameter819{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfigurationParameter819) SetBtType(v string) {
	o.BtType = &v
}

// GetGeneratedParameterId returns the GeneratedParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetGeneratedParameterId() BTTreeNode20 {
	if o == nil || o.GeneratedParameterId == nil {
		var ret BTTreeNode20
		return ret
	}
	return *o.GeneratedParameterId
}

// GetGeneratedParameterIdOk returns a tuple with the GeneratedParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetGeneratedParameterIdOk() (*BTTreeNode20, bool) {
	if o == nil || o.GeneratedParameterId == nil {
		return nil, false
	}
	return o.GeneratedParameterId, true
}

// HasGeneratedParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasGeneratedParameterId() bool {
	if o != nil && o.GeneratedParameterId != nil {
		return true
	}

	return false
}

// SetGeneratedParameterId gets a reference to the given BTTreeNode20 and assigns it to the GeneratedParameterId field.
func (o *BTMConfigurationParameter819) SetGeneratedParameterId(v BTTreeNode20) {
	o.GeneratedParameterId = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfigurationParameter819) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfigurationParameter819) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMConfigurationParameter819) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTMConfigurationParameter819) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetParameterType() string {
	if o == nil || o.ParameterType == nil {
		var ret string
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetParameterTypeOk() (*string, bool) {
	if o == nil || o.ParameterType == nil {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasParameterType() bool {
	if o != nil && o.ParameterType != nil {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given string and assigns it to the ParameterType field.
func (o *BTMConfigurationParameter819) SetParameterType(v string) {
	o.ParameterType = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTMConfigurationParameter819) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameter819) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *BTMConfigurationParameter819) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTMConfigurationParameter819) SetValid(v bool) {
	o.Valid = &v
}

func (o BTMConfigurationParameter819) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.GeneratedParameterId != nil {
		toSerialize["generatedParameterId"] = o.GeneratedParameterId
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.ParameterType != nil {
		toSerialize["parameterType"] = o.ParameterType
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfigurationParameter819 struct {
	value *BTMConfigurationParameter819
	isSet bool
}

func (v NullableBTMConfigurationParameter819) Get() *BTMConfigurationParameter819 {
	return v.value
}

func (v *NullableBTMConfigurationParameter819) Set(val *BTMConfigurationParameter819) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfigurationParameter819) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfigurationParameter819) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfigurationParameter819(val *BTMConfigurationParameter819) *NullableBTMConfigurationParameter819 {
	return &NullableBTMConfigurationParameter819{value: val, isSet: true}
}

func (v NullableBTMConfigurationParameter819) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfigurationParameter819) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
