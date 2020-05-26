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

// BTPExpressionOperator244 struct for BTPExpressionOperator244
type BTPExpressionOperator244 struct {
	BTPExpression9
	BtType *string `json:"btType,omitempty"`
	ForExport *bool `json:"forExport,omitempty"`
	GlobalNamespace *bool `json:"globalNamespace,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	Namespace *[]BTPIdentifier8 `json:"namespace,omitempty"`
	Operand1 *BTPExpression9 `json:"operand1,omitempty"`
	Operand2 *BTPExpression9 `json:"operand2,omitempty"`
	Operand3 *BTPExpression9 `json:"operand3,omitempty"`
	Operator *string `json:"operator,omitempty"`
	SpaceAfterNamespace *BTPSpace10 `json:"spaceAfterNamespace,omitempty"`
	SpaceAfterOperator *BTPSpace10 `json:"spaceAfterOperator,omitempty"`
	SpaceBeforeOperator *BTPSpace10 `json:"spaceBeforeOperator,omitempty"`
	WrittenAsFunctionCall *bool `json:"writtenAsFunctionCall,omitempty"`
}

// NewBTPExpressionOperator244 instantiates a new BTPExpressionOperator244 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionOperator244() *BTPExpressionOperator244 {
	this := BTPExpressionOperator244{}
	return &this
}

// NewBTPExpressionOperator244WithDefaults instantiates a new BTPExpressionOperator244 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionOperator244WithDefaults() *BTPExpressionOperator244 {
	this := BTPExpressionOperator244{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionOperator244) SetBtType(v string) {
	o.BtType = &v
}

// GetForExport returns the ForExport field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetForExport() bool {
	if o == nil || o.ForExport == nil {
		var ret bool
		return ret
	}
	return *o.ForExport
}

// GetForExportOk returns a tuple with the ForExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetForExportOk() (*bool, bool) {
	if o == nil || o.ForExport == nil {
		return nil, false
	}
	return o.ForExport, true
}

// HasForExport returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasForExport() bool {
	if o != nil && o.ForExport != nil {
		return true
	}

	return false
}

// SetForExport gets a reference to the given bool and assigns it to the ForExport field.
func (o *BTPExpressionOperator244) SetForExport(v bool) {
	o.ForExport = &v
}

// GetGlobalNamespace returns the GlobalNamespace field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetGlobalNamespace() bool {
	if o == nil || o.GlobalNamespace == nil {
		var ret bool
		return ret
	}
	return *o.GlobalNamespace
}

// GetGlobalNamespaceOk returns a tuple with the GlobalNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetGlobalNamespaceOk() (*bool, bool) {
	if o == nil || o.GlobalNamespace == nil {
		return nil, false
	}
	return o.GlobalNamespace, true
}

// HasGlobalNamespace returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasGlobalNamespace() bool {
	if o != nil && o.GlobalNamespace != nil {
		return true
	}

	return false
}

// SetGlobalNamespace gets a reference to the given bool and assigns it to the GlobalNamespace field.
func (o *BTPExpressionOperator244) SetGlobalNamespace(v bool) {
	o.GlobalNamespace = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTPExpressionOperator244) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetNamespace() []BTPIdentifier8 {
	if o == nil || o.Namespace == nil {
		var ret []BTPIdentifier8
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetNamespaceOk() (*[]BTPIdentifier8, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given []BTPIdentifier8 and assigns it to the Namespace field.
func (o *BTPExpressionOperator244) SetNamespace(v []BTPIdentifier8) {
	o.Namespace = &v
}

// GetOperand1 returns the Operand1 field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperand1() BTPExpression9 {
	if o == nil || o.Operand1 == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Operand1
}

// GetOperand1Ok returns a tuple with the Operand1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperand1Ok() (*BTPExpression9, bool) {
	if o == nil || o.Operand1 == nil {
		return nil, false
	}
	return o.Operand1, true
}

// HasOperand1 returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperand1() bool {
	if o != nil && o.Operand1 != nil {
		return true
	}

	return false
}

// SetOperand1 gets a reference to the given BTPExpression9 and assigns it to the Operand1 field.
func (o *BTPExpressionOperator244) SetOperand1(v BTPExpression9) {
	o.Operand1 = &v
}

// GetOperand2 returns the Operand2 field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperand2() BTPExpression9 {
	if o == nil || o.Operand2 == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Operand2
}

// GetOperand2Ok returns a tuple with the Operand2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperand2Ok() (*BTPExpression9, bool) {
	if o == nil || o.Operand2 == nil {
		return nil, false
	}
	return o.Operand2, true
}

// HasOperand2 returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperand2() bool {
	if o != nil && o.Operand2 != nil {
		return true
	}

	return false
}

// SetOperand2 gets a reference to the given BTPExpression9 and assigns it to the Operand2 field.
func (o *BTPExpressionOperator244) SetOperand2(v BTPExpression9) {
	o.Operand2 = &v
}

// GetOperand3 returns the Operand3 field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperand3() BTPExpression9 {
	if o == nil || o.Operand3 == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Operand3
}

// GetOperand3Ok returns a tuple with the Operand3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperand3Ok() (*BTPExpression9, bool) {
	if o == nil || o.Operand3 == nil {
		return nil, false
	}
	return o.Operand3, true
}

// HasOperand3 returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperand3() bool {
	if o != nil && o.Operand3 != nil {
		return true
	}

	return false
}

// SetOperand3 gets a reference to the given BTPExpression9 and assigns it to the Operand3 field.
func (o *BTPExpressionOperator244) SetOperand3(v BTPExpression9) {
	o.Operand3 = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *BTPExpressionOperator244) SetOperator(v string) {
	o.Operator = &v
}

// GetSpaceAfterNamespace returns the SpaceAfterNamespace field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceAfterNamespace() BTPSpace10 {
	if o == nil || o.SpaceAfterNamespace == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterNamespace
}

// GetSpaceAfterNamespaceOk returns a tuple with the SpaceAfterNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceAfterNamespaceOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterNamespace == nil {
		return nil, false
	}
	return o.SpaceAfterNamespace, true
}

// HasSpaceAfterNamespace returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceAfterNamespace() bool {
	if o != nil && o.SpaceAfterNamespace != nil {
		return true
	}

	return false
}

// SetSpaceAfterNamespace gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterNamespace field.
func (o *BTPExpressionOperator244) SetSpaceAfterNamespace(v BTPSpace10) {
	o.SpaceAfterNamespace = &v
}

// GetSpaceAfterOperator returns the SpaceAfterOperator field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceAfterOperator() BTPSpace10 {
	if o == nil || o.SpaceAfterOperator == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterOperator
}

// GetSpaceAfterOperatorOk returns a tuple with the SpaceAfterOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceAfterOperatorOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterOperator == nil {
		return nil, false
	}
	return o.SpaceAfterOperator, true
}

// HasSpaceAfterOperator returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceAfterOperator() bool {
	if o != nil && o.SpaceAfterOperator != nil {
		return true
	}

	return false
}

// SetSpaceAfterOperator gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterOperator field.
func (o *BTPExpressionOperator244) SetSpaceAfterOperator(v BTPSpace10) {
	o.SpaceAfterOperator = &v
}

// GetSpaceBeforeOperator returns the SpaceBeforeOperator field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetSpaceBeforeOperator() BTPSpace10 {
	if o == nil || o.SpaceBeforeOperator == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeOperator
}

// GetSpaceBeforeOperatorOk returns a tuple with the SpaceBeforeOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetSpaceBeforeOperatorOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeOperator == nil {
		return nil, false
	}
	return o.SpaceBeforeOperator, true
}

// HasSpaceBeforeOperator returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasSpaceBeforeOperator() bool {
	if o != nil && o.SpaceBeforeOperator != nil {
		return true
	}

	return false
}

// SetSpaceBeforeOperator gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeOperator field.
func (o *BTPExpressionOperator244) SetSpaceBeforeOperator(v BTPSpace10) {
	o.SpaceBeforeOperator = &v
}

// GetWrittenAsFunctionCall returns the WrittenAsFunctionCall field value if set, zero value otherwise.
func (o *BTPExpressionOperator244) GetWrittenAsFunctionCall() bool {
	if o == nil || o.WrittenAsFunctionCall == nil {
		var ret bool
		return ret
	}
	return *o.WrittenAsFunctionCall
}

// GetWrittenAsFunctionCallOk returns a tuple with the WrittenAsFunctionCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionOperator244) GetWrittenAsFunctionCallOk() (*bool, bool) {
	if o == nil || o.WrittenAsFunctionCall == nil {
		return nil, false
	}
	return o.WrittenAsFunctionCall, true
}

// HasWrittenAsFunctionCall returns a boolean if a field has been set.
func (o *BTPExpressionOperator244) HasWrittenAsFunctionCall() bool {
	if o != nil && o.WrittenAsFunctionCall != nil {
		return true
	}

	return false
}

// SetWrittenAsFunctionCall gets a reference to the given bool and assigns it to the WrittenAsFunctionCall field.
func (o *BTPExpressionOperator244) SetWrittenAsFunctionCall(v bool) {
	o.WrittenAsFunctionCall = &v
}

func (o BTPExpressionOperator244) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPExpression9, errBTPExpression9 := json.Marshal(o.BTPExpression9)
	if errBTPExpression9 != nil {
		return []byte{}, errBTPExpression9
	}
	errBTPExpression9 = json.Unmarshal([]byte(serializedBTPExpression9), &toSerialize)
	if errBTPExpression9 != nil {
		return []byte{}, errBTPExpression9
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ForExport != nil {
		toSerialize["forExport"] = o.ForExport
	}
	if o.GlobalNamespace != nil {
		toSerialize["globalNamespace"] = o.GlobalNamespace
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Operand1 != nil {
		toSerialize["operand1"] = o.Operand1
	}
	if o.Operand2 != nil {
		toSerialize["operand2"] = o.Operand2
	}
	if o.Operand3 != nil {
		toSerialize["operand3"] = o.Operand3
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.SpaceAfterNamespace != nil {
		toSerialize["spaceAfterNamespace"] = o.SpaceAfterNamespace
	}
	if o.SpaceAfterOperator != nil {
		toSerialize["spaceAfterOperator"] = o.SpaceAfterOperator
	}
	if o.SpaceBeforeOperator != nil {
		toSerialize["spaceBeforeOperator"] = o.SpaceBeforeOperator
	}
	if o.WrittenAsFunctionCall != nil {
		toSerialize["writtenAsFunctionCall"] = o.WrittenAsFunctionCall
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionOperator244 struct {
	value *BTPExpressionOperator244
	isSet bool
}

func (v NullableBTPExpressionOperator244) Get() *BTPExpressionOperator244 {
	return v.value
}

func (v *NullableBTPExpressionOperator244) Set(val *BTPExpressionOperator244) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionOperator244) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionOperator244) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionOperator244(val *BTPExpressionOperator244) *NullableBTPExpressionOperator244 {
	return &NullableBTPExpressionOperator244{value: val, isSet: true}
}

func (v NullableBTPExpressionOperator244) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionOperator244) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}