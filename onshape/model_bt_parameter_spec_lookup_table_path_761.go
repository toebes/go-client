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

// BTParameterSpecLookupTablePath761 struct for BTParameterSpecLookupTablePath761
type BTParameterSpecLookupTablePath761 struct {
	BTParameterSpec6
	BtType *string `json:"btType,omitempty"`
	LookupTable *BTParameterLookupTableListEntry1916 `json:"lookupTable,omitempty"`
}

// NewBTParameterSpecLookupTablePath761 instantiates a new BTParameterSpecLookupTablePath761 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecLookupTablePath761() *BTParameterSpecLookupTablePath761 {
	this := BTParameterSpecLookupTablePath761{}
	return &this
}

// NewBTParameterSpecLookupTablePath761WithDefaults instantiates a new BTParameterSpecLookupTablePath761 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecLookupTablePath761WithDefaults() *BTParameterSpecLookupTablePath761 {
	this := BTParameterSpecLookupTablePath761{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecLookupTablePath761) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecLookupTablePath761) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecLookupTablePath761) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecLookupTablePath761) SetBtType(v string) {
	o.BtType = &v
}

// GetLookupTable returns the LookupTable field value if set, zero value otherwise.
func (o *BTParameterSpecLookupTablePath761) GetLookupTable() BTParameterLookupTableListEntry1916 {
	if o == nil || o.LookupTable == nil {
		var ret BTParameterLookupTableListEntry1916
		return ret
	}
	return *o.LookupTable
}

// GetLookupTableOk returns a tuple with the LookupTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecLookupTablePath761) GetLookupTableOk() (*BTParameterLookupTableListEntry1916, bool) {
	if o == nil || o.LookupTable == nil {
		return nil, false
	}
	return o.LookupTable, true
}

// HasLookupTable returns a boolean if a field has been set.
func (o *BTParameterSpecLookupTablePath761) HasLookupTable() bool {
	if o != nil && o.LookupTable != nil {
		return true
	}

	return false
}

// SetLookupTable gets a reference to the given BTParameterLookupTableListEntry1916 and assigns it to the LookupTable field.
func (o *BTParameterSpecLookupTablePath761) SetLookupTable(v BTParameterLookupTableListEntry1916) {
	o.LookupTable = &v
}

func (o BTParameterSpecLookupTablePath761) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterSpec6, errBTParameterSpec6 := json.Marshal(o.BTParameterSpec6)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	errBTParameterSpec6 = json.Unmarshal([]byte(serializedBTParameterSpec6), &toSerialize)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.LookupTable != nil {
		toSerialize["lookupTable"] = o.LookupTable
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecLookupTablePath761 struct {
	value *BTParameterSpecLookupTablePath761
	isSet bool
}

func (v NullableBTParameterSpecLookupTablePath761) Get() *BTParameterSpecLookupTablePath761 {
	return v.value
}

func (v *NullableBTParameterSpecLookupTablePath761) Set(val *BTParameterSpecLookupTablePath761) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecLookupTablePath761) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecLookupTablePath761) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecLookupTablePath761(val *BTParameterSpecLookupTablePath761) *NullableBTParameterSpecLookupTablePath761 {
	return &NullableBTParameterSpecLookupTablePath761{value: val, isSet: true}
}

func (v NullableBTParameterSpecLookupTablePath761) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecLookupTablePath761) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}