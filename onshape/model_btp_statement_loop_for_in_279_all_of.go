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

// BTPStatementLoopForIn279AllOf struct for BTPStatementLoopForIn279AllOf
type BTPStatementLoopForIn279AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Container *BTPExpression9 `json:"container,omitempty"`
	IsVarDeclaredHere *bool `json:"isVarDeclaredHere,omitempty"`
	Name *BTPIdentifier8 `json:"name,omitempty"`
	SpaceBeforeVar *BTPSpace10 `json:"spaceBeforeVar,omitempty"`
	StandardType *string `json:"standardType,omitempty"`
	TypeName *string `json:"typeName,omitempty"`
	Var *BTPIdentifier8 `json:"var,omitempty"`
}

// NewBTPStatementLoopForIn279AllOf instantiates a new BTPStatementLoopForIn279AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementLoopForIn279AllOf() *BTPStatementLoopForIn279AllOf {
	this := BTPStatementLoopForIn279AllOf{}
	return &this
}

// NewBTPStatementLoopForIn279AllOfWithDefaults instantiates a new BTPStatementLoopForIn279AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementLoopForIn279AllOfWithDefaults() *BTPStatementLoopForIn279AllOf {
	this := BTPStatementLoopForIn279AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementLoopForIn279AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetContainer() BTPExpression9 {
	if o == nil || o.Container == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetContainerOk() (*BTPExpression9, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given BTPExpression9 and assigns it to the Container field.
func (o *BTPStatementLoopForIn279AllOf) SetContainer(v BTPExpression9) {
	o.Container = &v
}

// GetIsVarDeclaredHere returns the IsVarDeclaredHere field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetIsVarDeclaredHere() bool {
	if o == nil || o.IsVarDeclaredHere == nil {
		var ret bool
		return ret
	}
	return *o.IsVarDeclaredHere
}

// GetIsVarDeclaredHereOk returns a tuple with the IsVarDeclaredHere field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetIsVarDeclaredHereOk() (*bool, bool) {
	if o == nil || o.IsVarDeclaredHere == nil {
		return nil, false
	}
	return o.IsVarDeclaredHere, true
}

// HasIsVarDeclaredHere returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasIsVarDeclaredHere() bool {
	if o != nil && o.IsVarDeclaredHere != nil {
		return true
	}

	return false
}

// SetIsVarDeclaredHere gets a reference to the given bool and assigns it to the IsVarDeclaredHere field.
func (o *BTPStatementLoopForIn279AllOf) SetIsVarDeclaredHere(v bool) {
	o.IsVarDeclaredHere = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetName() BTPIdentifier8 {
	if o == nil || o.Name == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given BTPIdentifier8 and assigns it to the Name field.
func (o *BTPStatementLoopForIn279AllOf) SetName(v BTPIdentifier8) {
	o.Name = &v
}

// GetSpaceBeforeVar returns the SpaceBeforeVar field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetSpaceBeforeVar() BTPSpace10 {
	if o == nil || o.SpaceBeforeVar == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeVar
}

// GetSpaceBeforeVarOk returns a tuple with the SpaceBeforeVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetSpaceBeforeVarOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeVar == nil {
		return nil, false
	}
	return o.SpaceBeforeVar, true
}

// HasSpaceBeforeVar returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasSpaceBeforeVar() bool {
	if o != nil && o.SpaceBeforeVar != nil {
		return true
	}

	return false
}

// SetSpaceBeforeVar gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeVar field.
func (o *BTPStatementLoopForIn279AllOf) SetSpaceBeforeVar(v BTPSpace10) {
	o.SpaceBeforeVar = &v
}

// GetStandardType returns the StandardType field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetStandardType() string {
	if o == nil || o.StandardType == nil {
		var ret string
		return ret
	}
	return *o.StandardType
}

// GetStandardTypeOk returns a tuple with the StandardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetStandardTypeOk() (*string, bool) {
	if o == nil || o.StandardType == nil {
		return nil, false
	}
	return o.StandardType, true
}

// HasStandardType returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasStandardType() bool {
	if o != nil && o.StandardType != nil {
		return true
	}

	return false
}

// SetStandardType gets a reference to the given string and assigns it to the StandardType field.
func (o *BTPStatementLoopForIn279AllOf) SetStandardType(v string) {
	o.StandardType = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetTypeName() string {
	if o == nil || o.TypeName == nil {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetTypeNameOk() (*string, bool) {
	if o == nil || o.TypeName == nil {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasTypeName() bool {
	if o != nil && o.TypeName != nil {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *BTPStatementLoopForIn279AllOf) SetTypeName(v string) {
	o.TypeName = &v
}

// GetVar returns the Var field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279AllOf) GetVar() BTPIdentifier8 {
	if o == nil || o.Var == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Var
}

// GetVarOk returns a tuple with the Var field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279AllOf) GetVarOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Var == nil {
		return nil, false
	}
	return o.Var, true
}

// HasVar returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279AllOf) HasVar() bool {
	if o != nil && o.Var != nil {
		return true
	}

	return false
}

// SetVar gets a reference to the given BTPIdentifier8 and assigns it to the Var field.
func (o *BTPStatementLoopForIn279AllOf) SetVar(v BTPIdentifier8) {
	o.Var = &v
}

func (o BTPStatementLoopForIn279AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.IsVarDeclaredHere != nil {
		toSerialize["isVarDeclaredHere"] = o.IsVarDeclaredHere
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SpaceBeforeVar != nil {
		toSerialize["spaceBeforeVar"] = o.SpaceBeforeVar
	}
	if o.StandardType != nil {
		toSerialize["standardType"] = o.StandardType
	}
	if o.TypeName != nil {
		toSerialize["typeName"] = o.TypeName
	}
	if o.Var != nil {
		toSerialize["var"] = o.Var
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementLoopForIn279AllOf struct {
	value *BTPStatementLoopForIn279AllOf
	isSet bool
}

func (v NullableBTPStatementLoopForIn279AllOf) Get() *BTPStatementLoopForIn279AllOf {
	return v.value
}

func (v *NullableBTPStatementLoopForIn279AllOf) Set(val *BTPStatementLoopForIn279AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementLoopForIn279AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementLoopForIn279AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementLoopForIn279AllOf(val *BTPStatementLoopForIn279AllOf) *NullableBTPStatementLoopForIn279AllOf {
	return &NullableBTPStatementLoopForIn279AllOf{value: val, isSet: true}
}

func (v NullableBTPStatementLoopForIn279AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementLoopForIn279AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
