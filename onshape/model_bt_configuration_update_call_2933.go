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

// BTConfigurationUpdateCall2933 struct for BTConfigurationUpdateCall2933
type BTConfigurationUpdateCall2933 struct {
	BTFeatureApiBase1430
	BtType *string `json:"btType,omitempty"`
	ConfigurationParameters *[]BTMConfigurationParameter819 `json:"configurationParameters,omitempty"`
	CurrentConfiguration *[]BTMParameter1 `json:"currentConfiguration,omitempty"`
}

// NewBTConfigurationUpdateCall2933 instantiates a new BTConfigurationUpdateCall2933 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfigurationUpdateCall2933() *BTConfigurationUpdateCall2933 {
	this := BTConfigurationUpdateCall2933{}
	return &this
}

// NewBTConfigurationUpdateCall2933WithDefaults instantiates a new BTConfigurationUpdateCall2933 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfigurationUpdateCall2933WithDefaults() *BTConfigurationUpdateCall2933 {
	this := BTConfigurationUpdateCall2933{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfigurationUpdateCall2933) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationParameters returns the ConfigurationParameters field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetConfigurationParameters() []BTMConfigurationParameter819 {
	if o == nil || o.ConfigurationParameters == nil {
		var ret []BTMConfigurationParameter819
		return ret
	}
	return *o.ConfigurationParameters
}

// GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetConfigurationParametersOk() (*[]BTMConfigurationParameter819, bool) {
	if o == nil || o.ConfigurationParameters == nil {
		return nil, false
	}
	return o.ConfigurationParameters, true
}

// HasConfigurationParameters returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasConfigurationParameters() bool {
	if o != nil && o.ConfigurationParameters != nil {
		return true
	}

	return false
}

// SetConfigurationParameters gets a reference to the given []BTMConfigurationParameter819 and assigns it to the ConfigurationParameters field.
func (o *BTConfigurationUpdateCall2933) SetConfigurationParameters(v []BTMConfigurationParameter819) {
	o.ConfigurationParameters = &v
}

// GetCurrentConfiguration returns the CurrentConfiguration field value if set, zero value otherwise.
func (o *BTConfigurationUpdateCall2933) GetCurrentConfiguration() []BTMParameter1 {
	if o == nil || o.CurrentConfiguration == nil {
		var ret []BTMParameter1
		return ret
	}
	return *o.CurrentConfiguration
}

// GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationUpdateCall2933) GetCurrentConfigurationOk() (*[]BTMParameter1, bool) {
	if o == nil || o.CurrentConfiguration == nil {
		return nil, false
	}
	return o.CurrentConfiguration, true
}

// HasCurrentConfiguration returns a boolean if a field has been set.
func (o *BTConfigurationUpdateCall2933) HasCurrentConfiguration() bool {
	if o != nil && o.CurrentConfiguration != nil {
		return true
	}

	return false
}

// SetCurrentConfiguration gets a reference to the given []BTMParameter1 and assigns it to the CurrentConfiguration field.
func (o *BTConfigurationUpdateCall2933) SetCurrentConfiguration(v []BTMParameter1) {
	o.CurrentConfiguration = &v
}

func (o BTConfigurationUpdateCall2933) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTFeatureApiBase1430, errBTFeatureApiBase1430 := json.Marshal(o.BTFeatureApiBase1430)
	if errBTFeatureApiBase1430 != nil {
		return []byte{}, errBTFeatureApiBase1430
	}
	errBTFeatureApiBase1430 = json.Unmarshal([]byte(serializedBTFeatureApiBase1430), &toSerialize)
	if errBTFeatureApiBase1430 != nil {
		return []byte{}, errBTFeatureApiBase1430
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationParameters != nil {
		toSerialize["configurationParameters"] = o.ConfigurationParameters
	}
	if o.CurrentConfiguration != nil {
		toSerialize["currentConfiguration"] = o.CurrentConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfigurationUpdateCall2933 struct {
	value *BTConfigurationUpdateCall2933
	isSet bool
}

func (v NullableBTConfigurationUpdateCall2933) Get() *BTConfigurationUpdateCall2933 {
	return v.value
}

func (v *NullableBTConfigurationUpdateCall2933) Set(val *BTConfigurationUpdateCall2933) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfigurationUpdateCall2933) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfigurationUpdateCall2933) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfigurationUpdateCall2933(val *BTConfigurationUpdateCall2933) *NullableBTConfigurationUpdateCall2933 {
	return &NullableBTConfigurationUpdateCall2933{value: val, isSet: true}
}

func (v NullableBTConfigurationUpdateCall2933) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfigurationUpdateCall2933) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}