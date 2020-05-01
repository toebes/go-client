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

// BTAppElementHistoryInfo struct for BTAppElementHistoryInfo
type BTAppElementHistoryInfo struct {
	Changes *[]BTAppElementHistoryEntryInfo `json:"changes,omitempty"`
	ErrorCode *int32 `json:"errorCode,omitempty"`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	ErrorValue *string `json:"errorValue,omitempty"`
}

// NewBTAppElementHistoryInfo instantiates a new BTAppElementHistoryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementHistoryInfo() *BTAppElementHistoryInfo {
	this := BTAppElementHistoryInfo{}
	return &this
}

// NewBTAppElementHistoryInfoWithDefaults instantiates a new BTAppElementHistoryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementHistoryInfoWithDefaults() *BTAppElementHistoryInfo {
	this := BTAppElementHistoryInfo{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *BTAppElementHistoryInfo) GetChanges() []BTAppElementHistoryEntryInfo {
	if o == nil || o.Changes == nil {
		var ret []BTAppElementHistoryEntryInfo
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementHistoryInfo) GetChangesOk() (*[]BTAppElementHistoryEntryInfo, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *BTAppElementHistoryInfo) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []BTAppElementHistoryEntryInfo and assigns it to the Changes field.
func (o *BTAppElementHistoryInfo) SetChanges(v []BTAppElementHistoryEntryInfo) {
	o.Changes = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTAppElementHistoryInfo) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementHistoryInfo) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTAppElementHistoryInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTAppElementHistoryInfo) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BTAppElementHistoryInfo) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementHistoryInfo) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BTAppElementHistoryInfo) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BTAppElementHistoryInfo) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *BTAppElementHistoryInfo) GetErrorValue() string {
	if o == nil || o.ErrorValue == nil {
		var ret string
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementHistoryInfo) GetErrorValueOk() (*string, bool) {
	if o == nil || o.ErrorValue == nil {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *BTAppElementHistoryInfo) HasErrorValue() bool {
	if o != nil && o.ErrorValue != nil {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given string and assigns it to the ErrorValue field.
func (o *BTAppElementHistoryInfo) SetErrorValue(v string) {
	o.ErrorValue = &v
}

func (o BTAppElementHistoryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorValue != nil {
		toSerialize["errorValue"] = o.ErrorValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementHistoryInfo struct {
	value *BTAppElementHistoryInfo
	isSet bool
}

func (v NullableBTAppElementHistoryInfo) Get() *BTAppElementHistoryInfo {
	return v.value
}

func (v *NullableBTAppElementHistoryInfo) Set(val *BTAppElementHistoryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementHistoryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementHistoryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementHistoryInfo(val *BTAppElementHistoryInfo) *NullableBTAppElementHistoryInfo {
	return &NullableBTAppElementHistoryInfo{value: val, isSet: true}
}

func (v NullableBTAppElementHistoryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementHistoryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
