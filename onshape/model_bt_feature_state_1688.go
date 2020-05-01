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

// BTFeatureState1688 struct for BTFeatureState1688
type BTFeatureState1688 struct {
	BtType *string `json:"btType,omitempty"`
	FeatureStatus *string `json:"featureStatus,omitempty"`
}

// NewBTFeatureState1688 instantiates a new BTFeatureState1688 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureState1688() *BTFeatureState1688 {
	this := BTFeatureState1688{}
	return &this
}

// NewBTFeatureState1688WithDefaults instantiates a new BTFeatureState1688 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureState1688WithDefaults() *BTFeatureState1688 {
	this := BTFeatureState1688{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureState1688) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureState1688) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureState1688) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureState1688) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureStatus returns the FeatureStatus field value if set, zero value otherwise.
func (o *BTFeatureState1688) GetFeatureStatus() string {
	if o == nil || o.FeatureStatus == nil {
		var ret string
		return ret
	}
	return *o.FeatureStatus
}

// GetFeatureStatusOk returns a tuple with the FeatureStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureState1688) GetFeatureStatusOk() (*string, bool) {
	if o == nil || o.FeatureStatus == nil {
		return nil, false
	}
	return o.FeatureStatus, true
}

// HasFeatureStatus returns a boolean if a field has been set.
func (o *BTFeatureState1688) HasFeatureStatus() bool {
	if o != nil && o.FeatureStatus != nil {
		return true
	}

	return false
}

// SetFeatureStatus gets a reference to the given string and assigns it to the FeatureStatus field.
func (o *BTFeatureState1688) SetFeatureStatus(v string) {
	o.FeatureStatus = &v
}

func (o BTFeatureState1688) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureStatus != nil {
		toSerialize["featureStatus"] = o.FeatureStatus
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureState1688 struct {
	value *BTFeatureState1688
	isSet bool
}

func (v NullableBTFeatureState1688) Get() *BTFeatureState1688 {
	return v.value
}

func (v *NullableBTFeatureState1688) Set(val *BTFeatureState1688) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureState1688) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureState1688) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureState1688(val *BTFeatureState1688) *NullableBTFeatureState1688 {
	return &NullableBTFeatureState1688{value: val, isSet: true}
}

func (v NullableBTFeatureState1688) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureState1688) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
