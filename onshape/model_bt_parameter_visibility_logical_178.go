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

// BTParameterVisibilityLogical178 struct for BTParameterVisibilityLogical178
type BTParameterVisibilityLogical178 struct {
	BTParameterVisibilityCondition177
	BtType *string `json:"btType,omitempty"`
	Children *[]BTParameterVisibilityCondition177 `json:"children,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// NewBTParameterVisibilityLogical178 instantiates a new BTParameterVisibilityLogical178 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterVisibilityLogical178() *BTParameterVisibilityLogical178 {
	this := BTParameterVisibilityLogical178{}
	return &this
}

// NewBTParameterVisibilityLogical178WithDefaults instantiates a new BTParameterVisibilityLogical178 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterVisibilityLogical178WithDefaults() *BTParameterVisibilityLogical178 {
	this := BTParameterVisibilityLogical178{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterVisibilityLogical178) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityLogical178) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterVisibilityLogical178) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterVisibilityLogical178) SetBtType(v string) {
	o.BtType = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *BTParameterVisibilityLogical178) GetChildren() []BTParameterVisibilityCondition177 {
	if o == nil || o.Children == nil {
		var ret []BTParameterVisibilityCondition177
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityLogical178) GetChildrenOk() (*[]BTParameterVisibilityCondition177, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *BTParameterVisibilityLogical178) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []BTParameterVisibilityCondition177 and assigns it to the Children field.
func (o *BTParameterVisibilityLogical178) SetChildren(v []BTParameterVisibilityCondition177) {
	o.Children = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *BTParameterVisibilityLogical178) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityLogical178) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *BTParameterVisibilityLogical178) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *BTParameterVisibilityLogical178) SetOperation(v string) {
	o.Operation = &v
}

func (o BTParameterVisibilityLogical178) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterVisibilityCondition177, errBTParameterVisibilityCondition177 := json.Marshal(o.BTParameterVisibilityCondition177)
	if errBTParameterVisibilityCondition177 != nil {
		return []byte{}, errBTParameterVisibilityCondition177
	}
	errBTParameterVisibilityCondition177 = json.Unmarshal([]byte(serializedBTParameterVisibilityCondition177), &toSerialize)
	if errBTParameterVisibilityCondition177 != nil {
		return []byte{}, errBTParameterVisibilityCondition177
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterVisibilityLogical178 struct {
	value *BTParameterVisibilityLogical178
	isSet bool
}

func (v NullableBTParameterVisibilityLogical178) Get() *BTParameterVisibilityLogical178 {
	return v.value
}

func (v *NullableBTParameterVisibilityLogical178) Set(val *BTParameterVisibilityLogical178) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterVisibilityLogical178) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterVisibilityLogical178) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterVisibilityLogical178(val *BTParameterVisibilityLogical178) *NullableBTParameterVisibilityLogical178 {
	return &NullableBTParameterVisibilityLogical178{value: val, isSet: true}
}

func (v NullableBTParameterVisibilityLogical178) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterVisibilityLogical178) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
