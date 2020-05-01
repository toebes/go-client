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

// BTUserAppSettingsParams struct for BTUserAppSettingsParams
type BTUserAppSettingsParams struct {
	Settings *[]BTSettingParam `json:"settings,omitempty"`
}

// NewBTUserAppSettingsParams instantiates a new BTUserAppSettingsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserAppSettingsParams() *BTUserAppSettingsParams {
	this := BTUserAppSettingsParams{}
	return &this
}

// NewBTUserAppSettingsParamsWithDefaults instantiates a new BTUserAppSettingsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserAppSettingsParamsWithDefaults() *BTUserAppSettingsParams {
	this := BTUserAppSettingsParams{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BTUserAppSettingsParams) GetSettings() []BTSettingParam {
	if o == nil || o.Settings == nil {
		var ret []BTSettingParam
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppSettingsParams) GetSettingsOk() (*[]BTSettingParam, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BTUserAppSettingsParams) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []BTSettingParam and assigns it to the Settings field.
func (o *BTUserAppSettingsParams) SetSettings(v []BTSettingParam) {
	o.Settings = &v
}

func (o BTUserAppSettingsParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserAppSettingsParams struct {
	value *BTUserAppSettingsParams
	isSet bool
}

func (v NullableBTUserAppSettingsParams) Get() *BTUserAppSettingsParams {
	return v.value
}

func (v *NullableBTUserAppSettingsParams) Set(val *BTUserAppSettingsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserAppSettingsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserAppSettingsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserAppSettingsParams(val *BTUserAppSettingsParams) *NullableBTUserAppSettingsParams {
	return &NullableBTUserAppSettingsParams{value: val, isSet: true}
}

func (v NullableBTUserAppSettingsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserAppSettingsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}