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

// BTBillOfMaterialsTableRowMetadata1300 struct for BTBillOfMaterialsTableRowMetadata1300
type BTBillOfMaterialsTableRowMetadata1300 struct {
	BTTableBaseRowMetadata3181
	BtType *string `json:"btType,omitempty"`
	CrossHighlightData *BTTableAssemblyCrossHighlightData2675 `json:"crossHighlightData,omitempty"`
	CrossHighlightDataIfAny *BTTableAssemblyCrossHighlightData2675 `json:"crossHighlightDataIfAny,omitempty"`
}

// NewBTBillOfMaterialsTableRowMetadata1300 instantiates a new BTBillOfMaterialsTableRowMetadata1300 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsTableRowMetadata1300() *BTBillOfMaterialsTableRowMetadata1300 {
	this := BTBillOfMaterialsTableRowMetadata1300{}
	return &this
}

// NewBTBillOfMaterialsTableRowMetadata1300WithDefaults instantiates a new BTBillOfMaterialsTableRowMetadata1300 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsTableRowMetadata1300WithDefaults() *BTBillOfMaterialsTableRowMetadata1300 {
	this := BTBillOfMaterialsTableRowMetadata1300{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightData() BTTableAssemblyCrossHighlightData2675 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableAssemblyCrossHighlightData2675
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataOk() (*BTTableAssemblyCrossHighlightData2675, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableAssemblyCrossHighlightData2675 and assigns it to the CrossHighlightData field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetCrossHighlightData(v BTTableAssemblyCrossHighlightData2675) {
	o.CrossHighlightData = &v
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataIfAny() BTTableAssemblyCrossHighlightData2675 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableAssemblyCrossHighlightData2675
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataIfAnyOk() (*BTTableAssemblyCrossHighlightData2675, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableAssemblyCrossHighlightData2675 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetCrossHighlightDataIfAny(v BTTableAssemblyCrossHighlightData2675) {
	o.CrossHighlightDataIfAny = &v
}

func (o BTBillOfMaterialsTableRowMetadata1300) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTTableBaseRowMetadata3181, errBTTableBaseRowMetadata3181 := json.Marshal(o.BTTableBaseRowMetadata3181)
	if errBTTableBaseRowMetadata3181 != nil {
		return []byte{}, errBTTableBaseRowMetadata3181
	}
	errBTTableBaseRowMetadata3181 = json.Unmarshal([]byte(serializedBTTableBaseRowMetadata3181), &toSerialize)
	if errBTTableBaseRowMetadata3181 != nil {
		return []byte{}, errBTTableBaseRowMetadata3181
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsTableRowMetadata1300 struct {
	value *BTBillOfMaterialsTableRowMetadata1300
	isSet bool
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) Get() *BTBillOfMaterialsTableRowMetadata1300 {
	return v.value
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) Set(val *BTBillOfMaterialsTableRowMetadata1300) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsTableRowMetadata1300(val *BTBillOfMaterialsTableRowMetadata1300) *NullableBTBillOfMaterialsTableRowMetadata1300 {
	return &NullableBTBillOfMaterialsTableRowMetadata1300{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
