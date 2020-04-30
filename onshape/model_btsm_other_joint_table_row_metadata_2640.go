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

// BTSMOtherJointTableRowMetadata2640 struct for BTSMOtherJointTableRowMetadata2640
type BTSMOtherJointTableRowMetadata2640 struct {
	BTBaseSMJointTableRowMetadata2232
	BtType *string `json:"btType,omitempty"`
}

// NewBTSMOtherJointTableRowMetadata2640 instantiates a new BTSMOtherJointTableRowMetadata2640 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSMOtherJointTableRowMetadata2640() *BTSMOtherJointTableRowMetadata2640 {
	this := BTSMOtherJointTableRowMetadata2640{}
	return &this
}

// NewBTSMOtherJointTableRowMetadata2640WithDefaults instantiates a new BTSMOtherJointTableRowMetadata2640 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSMOtherJointTableRowMetadata2640WithDefaults() *BTSMOtherJointTableRowMetadata2640 {
	this := BTSMOtherJointTableRowMetadata2640{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSMOtherJointTableRowMetadata2640) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMOtherJointTableRowMetadata2640) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSMOtherJointTableRowMetadata2640) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSMOtherJointTableRowMetadata2640) SetBtType(v string) {
	o.BtType = &v
}

func (o BTSMOtherJointTableRowMetadata2640) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTBaseSMJointTableRowMetadata2232, errBTBaseSMJointTableRowMetadata2232 := json.Marshal(o.BTBaseSMJointTableRowMetadata2232)
	if errBTBaseSMJointTableRowMetadata2232 != nil {
		return []byte{}, errBTBaseSMJointTableRowMetadata2232
	}
	errBTBaseSMJointTableRowMetadata2232 = json.Unmarshal([]byte(serializedBTBaseSMJointTableRowMetadata2232), &toSerialize)
	if errBTBaseSMJointTableRowMetadata2232 != nil {
		return []byte{}, errBTBaseSMJointTableRowMetadata2232
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSMOtherJointTableRowMetadata2640 struct {
	value *BTSMOtherJointTableRowMetadata2640
	isSet bool
}

func (v NullableBTSMOtherJointTableRowMetadata2640) Get() *BTSMOtherJointTableRowMetadata2640 {
	return v.value
}

func (v *NullableBTSMOtherJointTableRowMetadata2640) Set(val *BTSMOtherJointTableRowMetadata2640) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSMOtherJointTableRowMetadata2640) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSMOtherJointTableRowMetadata2640) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSMOtherJointTableRowMetadata2640(val *BTSMOtherJointTableRowMetadata2640) *NullableBTSMOtherJointTableRowMetadata2640 {
	return &NullableBTSMOtherJointTableRowMetadata2640{value: val, isSet: true}
}

func (v NullableBTSMOtherJointTableRowMetadata2640) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSMOtherJointTableRowMetadata2640) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
