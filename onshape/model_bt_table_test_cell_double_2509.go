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

// BTTableTestCellDouble2509 struct for BTTableTestCellDouble2509
type BTTableTestCellDouble2509 struct {
	BTTableCell1114
	BtType *string `json:"btType,omitempty"`
	CellValue *float64 `json:"cellValue,omitempty"`
}

// NewBTTableTestCellDouble2509 instantiates a new BTTableTestCellDouble2509 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableTestCellDouble2509() *BTTableTestCellDouble2509 {
	this := BTTableTestCellDouble2509{}
	return &this
}

// NewBTTableTestCellDouble2509WithDefaults instantiates a new BTTableTestCellDouble2509 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableTestCellDouble2509WithDefaults() *BTTableTestCellDouble2509 {
	this := BTTableTestCellDouble2509{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableTestCellDouble2509) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellDouble2509) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableTestCellDouble2509) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableTestCellDouble2509) SetBtType(v string) {
	o.BtType = &v
}

// GetCellValue returns the CellValue field value if set, zero value otherwise.
func (o *BTTableTestCellDouble2509) GetCellValue() float64 {
	if o == nil || o.CellValue == nil {
		var ret float64
		return ret
	}
	return *o.CellValue
}

// GetCellValueOk returns a tuple with the CellValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableTestCellDouble2509) GetCellValueOk() (*float64, bool) {
	if o == nil || o.CellValue == nil {
		return nil, false
	}
	return o.CellValue, true
}

// HasCellValue returns a boolean if a field has been set.
func (o *BTTableTestCellDouble2509) HasCellValue() bool {
	if o != nil && o.CellValue != nil {
		return true
	}

	return false
}

// SetCellValue gets a reference to the given float64 and assigns it to the CellValue field.
func (o *BTTableTestCellDouble2509) SetCellValue(v float64) {
	o.CellValue = &v
}

func (o BTTableTestCellDouble2509) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTTableCell1114, errBTTableCell1114 := json.Marshal(o.BTTableCell1114)
	if errBTTableCell1114 != nil {
		return []byte{}, errBTTableCell1114
	}
	errBTTableCell1114 = json.Unmarshal([]byte(serializedBTTableCell1114), &toSerialize)
	if errBTTableCell1114 != nil {
		return []byte{}, errBTTableCell1114
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CellValue != nil {
		toSerialize["cellValue"] = o.CellValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableTestCellDouble2509 struct {
	value *BTTableTestCellDouble2509
	isSet bool
}

func (v NullableBTTableTestCellDouble2509) Get() *BTTableTestCellDouble2509 {
	return v.value
}

func (v *NullableBTTableTestCellDouble2509) Set(val *BTTableTestCellDouble2509) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableTestCellDouble2509) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableTestCellDouble2509) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableTestCellDouble2509(val *BTTableTestCellDouble2509) *NullableBTTableTestCellDouble2509 {
	return &NullableBTTableTestCellDouble2509{value: val, isSet: true}
}

func (v NullableBTTableTestCellDouble2509) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableTestCellDouble2509) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
