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

// BTMetadataItemsProperties This is the same as BTWorkflowPropertyInfo with the addition of the multivalued field and dropping the dirty, initialValue, isApproverProperty and is NotifierProperty fields
type BTMetadataItemsProperties struct {
	DefaultValue *map[string]interface{} `json:"defaultValue,omitempty"`
	Editable *bool `json:"editable,omitempty"`
	EditableInUi *bool `json:"editableInUi,omitempty"`
	EnumValues *[]BTMetadataEnumValueInfo `json:"enumValues,omitempty"`
	Multivalued *bool `json:"multivalued,omitempty"`
	Name *string `json:"name,omitempty"`
	PropertyId *string `json:"propertyId,omitempty"`
	PropertySource *int32 `json:"propertySource,omitempty"`
	Required *bool `json:"required,omitempty"`
	SchemaId *string `json:"schemaId,omitempty"`
	UiHints *BTMetadataPropertyUiHintsInfo `json:"uiHints,omitempty"`
	Validator *BTMetadataPropertyValidatorInfo `json:"validator,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
}

// NewBTMetadataItemsProperties instantiates a new BTMetadataItemsProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataItemsProperties() *BTMetadataItemsProperties {
	this := BTMetadataItemsProperties{}
	return &this
}

// NewBTMetadataItemsPropertiesWithDefaults instantiates a new BTMetadataItemsProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataItemsPropertiesWithDefaults() *BTMetadataItemsProperties {
	this := BTMetadataItemsProperties{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetDefaultValue() map[string]interface{} {
	if o == nil || o.DefaultValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetDefaultValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given map[string]interface{} and assigns it to the DefaultValue field.
func (o *BTMetadataItemsProperties) SetDefaultValue(v map[string]interface{}) {
	o.DefaultValue = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *BTMetadataItemsProperties) SetEditable(v bool) {
	o.Editable = &v
}

// GetEditableInUi returns the EditableInUi field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetEditableInUi() bool {
	if o == nil || o.EditableInUi == nil {
		var ret bool
		return ret
	}
	return *o.EditableInUi
}

// GetEditableInUiOk returns a tuple with the EditableInUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetEditableInUiOk() (*bool, bool) {
	if o == nil || o.EditableInUi == nil {
		return nil, false
	}
	return o.EditableInUi, true
}

// HasEditableInUi returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasEditableInUi() bool {
	if o != nil && o.EditableInUi != nil {
		return true
	}

	return false
}

// SetEditableInUi gets a reference to the given bool and assigns it to the EditableInUi field.
func (o *BTMetadataItemsProperties) SetEditableInUi(v bool) {
	o.EditableInUi = &v
}

// GetEnumValues returns the EnumValues field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetEnumValues() []BTMetadataEnumValueInfo {
	if o == nil || o.EnumValues == nil {
		var ret []BTMetadataEnumValueInfo
		return ret
	}
	return *o.EnumValues
}

// GetEnumValuesOk returns a tuple with the EnumValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetEnumValuesOk() (*[]BTMetadataEnumValueInfo, bool) {
	if o == nil || o.EnumValues == nil {
		return nil, false
	}
	return o.EnumValues, true
}

// HasEnumValues returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasEnumValues() bool {
	if o != nil && o.EnumValues != nil {
		return true
	}

	return false
}

// SetEnumValues gets a reference to the given []BTMetadataEnumValueInfo and assigns it to the EnumValues field.
func (o *BTMetadataItemsProperties) SetEnumValues(v []BTMetadataEnumValueInfo) {
	o.EnumValues = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetMultivalued() bool {
	if o == nil || o.Multivalued == nil {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetMultivaluedOk() (*bool, bool) {
	if o == nil || o.Multivalued == nil {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasMultivalued() bool {
	if o != nil && o.Multivalued != nil {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *BTMetadataItemsProperties) SetMultivalued(v bool) {
	o.Multivalued = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMetadataItemsProperties) SetName(v string) {
	o.Name = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTMetadataItemsProperties) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPropertySource returns the PropertySource field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetPropertySource() int32 {
	if o == nil || o.PropertySource == nil {
		var ret int32
		return ret
	}
	return *o.PropertySource
}

// GetPropertySourceOk returns a tuple with the PropertySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetPropertySourceOk() (*int32, bool) {
	if o == nil || o.PropertySource == nil {
		return nil, false
	}
	return o.PropertySource, true
}

// HasPropertySource returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasPropertySource() bool {
	if o != nil && o.PropertySource != nil {
		return true
	}

	return false
}

// SetPropertySource gets a reference to the given int32 and assigns it to the PropertySource field.
func (o *BTMetadataItemsProperties) SetPropertySource(v int32) {
	o.PropertySource = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *BTMetadataItemsProperties) SetRequired(v bool) {
	o.Required = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *BTMetadataItemsProperties) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetUiHints() BTMetadataPropertyUiHintsInfo {
	if o == nil || o.UiHints == nil {
		var ret BTMetadataPropertyUiHintsInfo
		return ret
	}
	return *o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given BTMetadataPropertyUiHintsInfo and assigns it to the UiHints field.
func (o *BTMetadataItemsProperties) SetUiHints(v BTMetadataPropertyUiHintsInfo) {
	o.UiHints = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetValidator() BTMetadataPropertyValidatorInfo {
	if o == nil || o.Validator == nil {
		var ret BTMetadataPropertyValidatorInfo
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasValidator() bool {
	if o != nil && o.Validator != nil {
		return true
	}

	return false
}

// SetValidator gets a reference to the given BTMetadataPropertyValidatorInfo and assigns it to the Validator field.
func (o *BTMetadataItemsProperties) SetValidator(v BTMetadataPropertyValidatorInfo) {
	o.Validator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *BTMetadataItemsProperties) SetValue(v map[string]interface{}) {
	o.Value = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *BTMetadataItemsProperties) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataItemsProperties) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *BTMetadataItemsProperties) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *BTMetadataItemsProperties) SetValueType(v string) {
	o.ValueType = &v
}

func (o BTMetadataItemsProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.EditableInUi != nil {
		toSerialize["editableInUi"] = o.EditableInUi
	}
	if o.EnumValues != nil {
		toSerialize["enumValues"] = o.EnumValues
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
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataItemsProperties struct {
	value *BTMetadataItemsProperties
	isSet bool
}

func (v NullableBTMetadataItemsProperties) Get() *BTMetadataItemsProperties {
	return v.value
}

func (v *NullableBTMetadataItemsProperties) Set(val *BTMetadataItemsProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataItemsProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataItemsProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataItemsProperties(val *BTMetadataItemsProperties) *NullableBTMetadataItemsProperties {
	return &NullableBTMetadataItemsProperties{value: val, isSet: true}
}

func (v NullableBTMetadataItemsProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataItemsProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
