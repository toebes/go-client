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

// BTPExpressionBuiltinCall239 struct for BTPExpressionBuiltinCall239
type BTPExpressionBuiltinCall239 struct {
	BTPExpression9
	Arguments *[]BTPExpression9 `json:"arguments,omitempty"`
	BtType *string `json:"btType,omitempty"`
	Name *BTPBuiltinIdentifier233 `json:"name,omitempty"`
	SpaceInEmptyList *BTPSpace10 `json:"spaceInEmptyList,omitempty"`
}

// NewBTPExpressionBuiltinCall239 instantiates a new BTPExpressionBuiltinCall239 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionBuiltinCall239() *BTPExpressionBuiltinCall239 {
	this := BTPExpressionBuiltinCall239{}
	return &this
}

// NewBTPExpressionBuiltinCall239WithDefaults instantiates a new BTPExpressionBuiltinCall239 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionBuiltinCall239WithDefaults() *BTPExpressionBuiltinCall239 {
	this := BTPExpressionBuiltinCall239{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *BTPExpressionBuiltinCall239) GetArguments() []BTPExpression9 {
	if o == nil || o.Arguments == nil {
		var ret []BTPExpression9
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionBuiltinCall239) GetArgumentsOk() (*[]BTPExpression9, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *BTPExpressionBuiltinCall239) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []BTPExpression9 and assigns it to the Arguments field.
func (o *BTPExpressionBuiltinCall239) SetArguments(v []BTPExpression9) {
	o.Arguments = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionBuiltinCall239) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionBuiltinCall239) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionBuiltinCall239) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionBuiltinCall239) SetBtType(v string) {
	o.BtType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPExpressionBuiltinCall239) GetName() BTPBuiltinIdentifier233 {
	if o == nil || o.Name == nil {
		var ret BTPBuiltinIdentifier233
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionBuiltinCall239) GetNameOk() (*BTPBuiltinIdentifier233, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPExpressionBuiltinCall239) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given BTPBuiltinIdentifier233 and assigns it to the Name field.
func (o *BTPExpressionBuiltinCall239) SetName(v BTPBuiltinIdentifier233) {
	o.Name = &v
}

// GetSpaceInEmptyList returns the SpaceInEmptyList field value if set, zero value otherwise.
func (o *BTPExpressionBuiltinCall239) GetSpaceInEmptyList() BTPSpace10 {
	if o == nil || o.SpaceInEmptyList == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInEmptyList
}

// GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionBuiltinCall239) GetSpaceInEmptyListOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInEmptyList == nil {
		return nil, false
	}
	return o.SpaceInEmptyList, true
}

// HasSpaceInEmptyList returns a boolean if a field has been set.
func (o *BTPExpressionBuiltinCall239) HasSpaceInEmptyList() bool {
	if o != nil && o.SpaceInEmptyList != nil {
		return true
	}

	return false
}

// SetSpaceInEmptyList gets a reference to the given BTPSpace10 and assigns it to the SpaceInEmptyList field.
func (o *BTPExpressionBuiltinCall239) SetSpaceInEmptyList(v BTPSpace10) {
	o.SpaceInEmptyList = &v
}

func (o BTPExpressionBuiltinCall239) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPExpression9, errBTPExpression9 := json.Marshal(o.BTPExpression9)
	if errBTPExpression9 != nil {
		return []byte{}, errBTPExpression9
	}
	errBTPExpression9 = json.Unmarshal([]byte(serializedBTPExpression9), &toSerialize)
	if errBTPExpression9 != nil {
		return []byte{}, errBTPExpression9
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SpaceInEmptyList != nil {
		toSerialize["spaceInEmptyList"] = o.SpaceInEmptyList
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionBuiltinCall239 struct {
	value *BTPExpressionBuiltinCall239
	isSet bool
}

func (v NullableBTPExpressionBuiltinCall239) Get() *BTPExpressionBuiltinCall239 {
	return v.value
}

func (v *NullableBTPExpressionBuiltinCall239) Set(val *BTPExpressionBuiltinCall239) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionBuiltinCall239) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionBuiltinCall239) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionBuiltinCall239(val *BTPExpressionBuiltinCall239) *NullableBTPExpressionBuiltinCall239 {
	return &NullableBTPExpressionBuiltinCall239{value: val, isSet: true}
}

func (v NullableBTPExpressionBuiltinCall239) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionBuiltinCall239) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
