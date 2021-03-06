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

// BTParameterVisibilityLogical178AllOf struct for BTParameterVisibilityLogical178AllOf
type BTParameterVisibilityLogical178AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Children *[]BTParameterVisibilityCondition177 `json:"children,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// NewBTParameterVisibilityLogical178AllOf instantiates a new BTParameterVisibilityLogical178AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterVisibilityLogical178AllOf() *BTParameterVisibilityLogical178AllOf {
	this := BTParameterVisibilityLogical178AllOf{}
	return &this
}

// NewBTParameterVisibilityLogical178AllOfWithDefaults instantiates a new BTParameterVisibilityLogical178AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterVisibilityLogical178AllOfWithDefaults() *BTParameterVisibilityLogical178AllOf {
	this := BTParameterVisibilityLogical178AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterVisibilityLogical178AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityLogical178AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterVisibilityLogical178AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterVisibilityLogical178AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *BTParameterVisibilityLogical178AllOf) GetChildren() []BTParameterVisibilityCondition177 {
	if o == nil || o.Children == nil {
		var ret []BTParameterVisibilityCondition177
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityLogical178AllOf) GetChildrenOk() (*[]BTParameterVisibilityCondition177, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *BTParameterVisibilityLogical178AllOf) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []BTParameterVisibilityCondition177 and assigns it to the Children field.
func (o *BTParameterVisibilityLogical178AllOf) SetChildren(v []BTParameterVisibilityCondition177) {
	o.Children = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *BTParameterVisibilityLogical178AllOf) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityLogical178AllOf) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *BTParameterVisibilityLogical178AllOf) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *BTParameterVisibilityLogical178AllOf) SetOperation(v string) {
	o.Operation = &v
}

func (o BTParameterVisibilityLogical178AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableBTParameterVisibilityLogical178AllOf struct {
	value *BTParameterVisibilityLogical178AllOf
	isSet bool
}

func (v NullableBTParameterVisibilityLogical178AllOf) Get() *BTParameterVisibilityLogical178AllOf {
	return v.value
}

func (v *NullableBTParameterVisibilityLogical178AllOf) Set(val *BTParameterVisibilityLogical178AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterVisibilityLogical178AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterVisibilityLogical178AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterVisibilityLogical178AllOf(val *BTParameterVisibilityLogical178AllOf) *NullableBTParameterVisibilityLogical178AllOf {
	return &NullableBTParameterVisibilityLogical178AllOf{value: val, isSet: true}
}

func (v NullableBTParameterVisibilityLogical178AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterVisibilityLogical178AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
