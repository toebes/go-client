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

// BTMetadataCommonUser struct for BTMetadataCommonUser
type BTMetadataCommonUser struct {
	DefaultValue *string `json:"defaultValue,omitempty"`
	Dirty *bool `json:"dirty,omitempty"`
	Editable *bool `json:"editable,omitempty"`
	EditableInUi *bool `json:"editableInUi,omitempty"`
	InitialValue *map[string]interface{} `json:"initialValue,omitempty"`
	IsApproverProperty *bool `json:"isApproverProperty,omitempty"`
	IsNotifierProperty *bool `json:"isNotifierProperty,omitempty"`
	Multivalued *bool `json:"multivalued,omitempty"`
	Name *string `json:"name,omitempty"`
	PropertyId *string `json:"propertyId,omitempty"`
	PropertySource *int32 `json:"propertySource,omitempty"`
	Required *bool `json:"required,omitempty"`
	SchemaId *string `json:"schemaId,omitempty"`
	UiHints *BTMetadataPropertyUiHintsInfo `json:"uiHints,omitempty"`
	Validator *BTMetadataPropertyValidatorInfo `json:"validator,omitempty"`
	Value *[]BTUserSummaryInfo `json:"value,omitempty"`
	ValueType string `json:"valueType"`
}

// NewBTMetadataCommonUser instantiates a new BTMetadataCommonUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataCommonUser(valueType string, ) *BTMetadataCommonUser {
	this := BTMetadataCommonUser{}
	this.ValueType = valueType
	return &this
}

// NewBTMetadataCommonUserWithDefaults instantiates a new BTMetadataCommonUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataCommonUserWithDefaults() *BTMetadataCommonUser {
	this := BTMetadataCommonUser{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *BTMetadataCommonUser) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDirty returns the Dirty field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetDirty() bool {
	if o == nil || o.Dirty == nil {
		var ret bool
		return ret
	}
	return *o.Dirty
}

// GetDirtyOk returns a tuple with the Dirty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetDirtyOk() (*bool, bool) {
	if o == nil || o.Dirty == nil {
		return nil, false
	}
	return o.Dirty, true
}

// HasDirty returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasDirty() bool {
	if o != nil && o.Dirty != nil {
		return true
	}

	return false
}

// SetDirty gets a reference to the given bool and assigns it to the Dirty field.
func (o *BTMetadataCommonUser) SetDirty(v bool) {
	o.Dirty = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *BTMetadataCommonUser) SetEditable(v bool) {
	o.Editable = &v
}

// GetEditableInUi returns the EditableInUi field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetEditableInUi() bool {
	if o == nil || o.EditableInUi == nil {
		var ret bool
		return ret
	}
	return *o.EditableInUi
}

// GetEditableInUiOk returns a tuple with the EditableInUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetEditableInUiOk() (*bool, bool) {
	if o == nil || o.EditableInUi == nil {
		return nil, false
	}
	return o.EditableInUi, true
}

// HasEditableInUi returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasEditableInUi() bool {
	if o != nil && o.EditableInUi != nil {
		return true
	}

	return false
}

// SetEditableInUi gets a reference to the given bool and assigns it to the EditableInUi field.
func (o *BTMetadataCommonUser) SetEditableInUi(v bool) {
	o.EditableInUi = &v
}

// GetInitialValue returns the InitialValue field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetInitialValue() map[string]interface{} {
	if o == nil || o.InitialValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.InitialValue
}

// GetInitialValueOk returns a tuple with the InitialValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetInitialValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.InitialValue == nil {
		return nil, false
	}
	return o.InitialValue, true
}

// HasInitialValue returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasInitialValue() bool {
	if o != nil && o.InitialValue != nil {
		return true
	}

	return false
}

// SetInitialValue gets a reference to the given map[string]interface{} and assigns it to the InitialValue field.
func (o *BTMetadataCommonUser) SetInitialValue(v map[string]interface{}) {
	o.InitialValue = &v
}

// GetIsApproverProperty returns the IsApproverProperty field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetIsApproverProperty() bool {
	if o == nil || o.IsApproverProperty == nil {
		var ret bool
		return ret
	}
	return *o.IsApproverProperty
}

// GetIsApproverPropertyOk returns a tuple with the IsApproverProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetIsApproverPropertyOk() (*bool, bool) {
	if o == nil || o.IsApproverProperty == nil {
		return nil, false
	}
	return o.IsApproverProperty, true
}

// HasIsApproverProperty returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasIsApproverProperty() bool {
	if o != nil && o.IsApproverProperty != nil {
		return true
	}

	return false
}

// SetIsApproverProperty gets a reference to the given bool and assigns it to the IsApproverProperty field.
func (o *BTMetadataCommonUser) SetIsApproverProperty(v bool) {
	o.IsApproverProperty = &v
}

// GetIsNotifierProperty returns the IsNotifierProperty field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetIsNotifierProperty() bool {
	if o == nil || o.IsNotifierProperty == nil {
		var ret bool
		return ret
	}
	return *o.IsNotifierProperty
}

// GetIsNotifierPropertyOk returns a tuple with the IsNotifierProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetIsNotifierPropertyOk() (*bool, bool) {
	if o == nil || o.IsNotifierProperty == nil {
		return nil, false
	}
	return o.IsNotifierProperty, true
}

// HasIsNotifierProperty returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasIsNotifierProperty() bool {
	if o != nil && o.IsNotifierProperty != nil {
		return true
	}

	return false
}

// SetIsNotifierProperty gets a reference to the given bool and assigns it to the IsNotifierProperty field.
func (o *BTMetadataCommonUser) SetIsNotifierProperty(v bool) {
	o.IsNotifierProperty = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetMultivalued() bool {
	if o == nil || o.Multivalued == nil {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetMultivaluedOk() (*bool, bool) {
	if o == nil || o.Multivalued == nil {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasMultivalued() bool {
	if o != nil && o.Multivalued != nil {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *BTMetadataCommonUser) SetMultivalued(v bool) {
	o.Multivalued = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMetadataCommonUser) SetName(v string) {
	o.Name = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTMetadataCommonUser) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPropertySource returns the PropertySource field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetPropertySource() int32 {
	if o == nil || o.PropertySource == nil {
		var ret int32
		return ret
	}
	return *o.PropertySource
}

// GetPropertySourceOk returns a tuple with the PropertySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetPropertySourceOk() (*int32, bool) {
	if o == nil || o.PropertySource == nil {
		return nil, false
	}
	return o.PropertySource, true
}

// HasPropertySource returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasPropertySource() bool {
	if o != nil && o.PropertySource != nil {
		return true
	}

	return false
}

// SetPropertySource gets a reference to the given int32 and assigns it to the PropertySource field.
func (o *BTMetadataCommonUser) SetPropertySource(v int32) {
	o.PropertySource = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *BTMetadataCommonUser) SetRequired(v bool) {
	o.Required = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *BTMetadataCommonUser) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetUiHints() BTMetadataPropertyUiHintsInfo {
	if o == nil || o.UiHints == nil {
		var ret BTMetadataPropertyUiHintsInfo
		return ret
	}
	return *o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given BTMetadataPropertyUiHintsInfo and assigns it to the UiHints field.
func (o *BTMetadataCommonUser) SetUiHints(v BTMetadataPropertyUiHintsInfo) {
	o.UiHints = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetValidator() BTMetadataPropertyValidatorInfo {
	if o == nil || o.Validator == nil {
		var ret BTMetadataPropertyValidatorInfo
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasValidator() bool {
	if o != nil && o.Validator != nil {
		return true
	}

	return false
}

// SetValidator gets a reference to the given BTMetadataPropertyValidatorInfo and assigns it to the Validator field.
func (o *BTMetadataCommonUser) SetValidator(v BTMetadataPropertyValidatorInfo) {
	o.Validator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMetadataCommonUser) GetValue() []BTUserSummaryInfo {
	if o == nil || o.Value == nil {
		var ret []BTUserSummaryInfo
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetValueOk() (*[]BTUserSummaryInfo, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMetadataCommonUser) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []BTUserSummaryInfo and assigns it to the Value field.
func (o *BTMetadataCommonUser) SetValue(v []BTUserSummaryInfo) {
	o.Value = &v
}

// GetValueType returns the ValueType field value
func (o *BTMetadataCommonUser) GetValueType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *BTMetadataCommonUser) GetValueTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *BTMetadataCommonUser) SetValueType(v string) {
	o.ValueType = v
}

func (o BTMetadataCommonUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Dirty != nil {
		toSerialize["dirty"] = o.Dirty
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.EditableInUi != nil {
		toSerialize["editableInUi"] = o.EditableInUi
	}
	if o.InitialValue != nil {
		toSerialize["initialValue"] = o.InitialValue
	}
	if o.IsApproverProperty != nil {
		toSerialize["isApproverProperty"] = o.IsApproverProperty
	}
	if o.IsNotifierProperty != nil {
		toSerialize["isNotifierProperty"] = o.IsNotifierProperty
	}
	if o.Multivalued != nil {
		toSerialize["multivalued"] = o.Multivalued
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	if o.PropertySource != nil {
		toSerialize["propertySource"] = o.PropertySource
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.SchemaId != nil {
		toSerialize["schemaId"] = o.SchemaId
	}
	if o.UiHints != nil {
		toSerialize["uiHints"] = o.UiHints
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

// AsBTMetadataItemsProperties wraps this instance of BTMetadataCommonUser in BTMetadataItemsProperties
func (s *BTMetadataCommonUser) AsBTMetadataItemsProperties() BTMetadataItemsProperties {
	return BTMetadataItemsProperties{ BTMetadataItemsPropertiesInterface: s }
}
// AsBTMetadataPropertyInfo wraps this instance of BTMetadataCommonUser in BTMetadataPropertyInfo
func (s *BTMetadataCommonUser) AsBTMetadataPropertyInfo() BTMetadataPropertyInfo {
	return BTMetadataPropertyInfo{ BTMetadataPropertyInfoInterface: s }
}
// AsBTWorkflowPropertyInfo wraps this instance of BTMetadataCommonUser in BTWorkflowPropertyInfo
func (s *BTMetadataCommonUser) AsBTWorkflowPropertyInfo() BTWorkflowPropertyInfo {
	return BTWorkflowPropertyInfo{ BTWorkflowPropertyInfoInterface: s }
}
type NullableBTMetadataCommonUser struct {
	value *BTMetadataCommonUser
	isSet bool
}

func (v NullableBTMetadataCommonUser) Get() *BTMetadataCommonUser {
	return v.value
}

func (v *NullableBTMetadataCommonUser) Set(val *BTMetadataCommonUser) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataCommonUser) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataCommonUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataCommonUser(val *BTMetadataCommonUser) *NullableBTMetadataCommonUser {
	return &NullableBTMetadataCommonUser{value: val, isSet: true}
}

func (v NullableBTMetadataCommonUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataCommonUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
