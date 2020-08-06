# BTMetadataCommonString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** |  | [optional] 
**Dirty** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**EditableInUi** | Pointer to **bool** |  | [optional] 
**InitialValue** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**IsApproverProperty** | Pointer to **bool** |  | [optional] 
**IsNotifierProperty** | Pointer to **bool** |  | [optional] 
**Multivalued** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PropertyId** | Pointer to **string** |  | [optional] 
**PropertySource** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**UiHints** | Pointer to [**BTMetadataPropertyUiHintsInfo**](BTMetadataPropertyUiHintsInfo.md) |  | [optional] 
**Validator** | Pointer to [**BTMetadataPropertyValidatorInfo**](BTMetadataPropertyValidatorInfo.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**ValueType** | Pointer to **string** |  | 

## Methods

### NewBTMetadataCommonString

`func NewBTMetadataCommonString(valueType string, ) *BTMetadataCommonString`

NewBTMetadataCommonString instantiates a new BTMetadataCommonString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataCommonStringWithDefaults

`func NewBTMetadataCommonStringWithDefaults() *BTMetadataCommonString`

NewBTMetadataCommonStringWithDefaults instantiates a new BTMetadataCommonString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *BTMetadataCommonString) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BTMetadataCommonString) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BTMetadataCommonString) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BTMetadataCommonString) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDirty

`func (o *BTMetadataCommonString) GetDirty() bool`

GetDirty returns the Dirty field if non-nil, zero value otherwise.

### GetDirtyOk

`func (o *BTMetadataCommonString) GetDirtyOk() (*bool, bool)`

GetDirtyOk returns a tuple with the Dirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirty

`func (o *BTMetadataCommonString) SetDirty(v bool)`

SetDirty sets Dirty field to given value.

### HasDirty

`func (o *BTMetadataCommonString) HasDirty() bool`

HasDirty returns a boolean if a field has been set.

### GetEditable

`func (o *BTMetadataCommonString) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *BTMetadataCommonString) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *BTMetadataCommonString) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *BTMetadataCommonString) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEditableInUi

`func (o *BTMetadataCommonString) GetEditableInUi() bool`

GetEditableInUi returns the EditableInUi field if non-nil, zero value otherwise.

### GetEditableInUiOk

`func (o *BTMetadataCommonString) GetEditableInUiOk() (*bool, bool)`

GetEditableInUiOk returns a tuple with the EditableInUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInUi

`func (o *BTMetadataCommonString) SetEditableInUi(v bool)`

SetEditableInUi sets EditableInUi field to given value.

### HasEditableInUi

`func (o *BTMetadataCommonString) HasEditableInUi() bool`

HasEditableInUi returns a boolean if a field has been set.

### GetInitialValue

`func (o *BTMetadataCommonString) GetInitialValue() map[string]interface{}`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *BTMetadataCommonString) GetInitialValueOk() (*map[string]interface{}, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *BTMetadataCommonString) SetInitialValue(v map[string]interface{})`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *BTMetadataCommonString) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetIsApproverProperty

`func (o *BTMetadataCommonString) GetIsApproverProperty() bool`

GetIsApproverProperty returns the IsApproverProperty field if non-nil, zero value otherwise.

### GetIsApproverPropertyOk

`func (o *BTMetadataCommonString) GetIsApproverPropertyOk() (*bool, bool)`

GetIsApproverPropertyOk returns a tuple with the IsApproverProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproverProperty

`func (o *BTMetadataCommonString) SetIsApproverProperty(v bool)`

SetIsApproverProperty sets IsApproverProperty field to given value.

### HasIsApproverProperty

`func (o *BTMetadataCommonString) HasIsApproverProperty() bool`

HasIsApproverProperty returns a boolean if a field has been set.

### GetIsNotifierProperty

`func (o *BTMetadataCommonString) GetIsNotifierProperty() bool`

GetIsNotifierProperty returns the IsNotifierProperty field if non-nil, zero value otherwise.

### GetIsNotifierPropertyOk

`func (o *BTMetadataCommonString) GetIsNotifierPropertyOk() (*bool, bool)`

GetIsNotifierPropertyOk returns a tuple with the IsNotifierProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotifierProperty

`func (o *BTMetadataCommonString) SetIsNotifierProperty(v bool)`

SetIsNotifierProperty sets IsNotifierProperty field to given value.

### HasIsNotifierProperty

`func (o *BTMetadataCommonString) HasIsNotifierProperty() bool`

HasIsNotifierProperty returns a boolean if a field has been set.

### GetMultivalued

`func (o *BTMetadataCommonString) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *BTMetadataCommonString) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *BTMetadataCommonString) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *BTMetadataCommonString) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.

### GetName

`func (o *BTMetadataCommonString) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMetadataCommonString) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMetadataCommonString) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMetadataCommonString) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyId

`func (o *BTMetadataCommonString) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTMetadataCommonString) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTMetadataCommonString) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTMetadataCommonString) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPropertySource

`func (o *BTMetadataCommonString) GetPropertySource() int32`

GetPropertySource returns the PropertySource field if non-nil, zero value otherwise.

### GetPropertySourceOk

`func (o *BTMetadataCommonString) GetPropertySourceOk() (*int32, bool)`

GetPropertySourceOk returns a tuple with the PropertySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySource

`func (o *BTMetadataCommonString) SetPropertySource(v int32)`

SetPropertySource sets PropertySource field to given value.

### HasPropertySource

`func (o *BTMetadataCommonString) HasPropertySource() bool`

HasPropertySource returns a boolean if a field has been set.

### GetRequired

`func (o *BTMetadataCommonString) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BTMetadataCommonString) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BTMetadataCommonString) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BTMetadataCommonString) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchemaId

`func (o *BTMetadataCommonString) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *BTMetadataCommonString) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *BTMetadataCommonString) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *BTMetadataCommonString) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetUiHints

`func (o *BTMetadataCommonString) GetUiHints() BTMetadataPropertyUiHintsInfo`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *BTMetadataCommonString) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *BTMetadataCommonString) SetUiHints(v BTMetadataPropertyUiHintsInfo)`

SetUiHints sets UiHints field to given value.

### HasUiHints

`func (o *BTMetadataCommonString) HasUiHints() bool`

HasUiHints returns a boolean if a field has been set.

### GetValidator

`func (o *BTMetadataCommonString) GetValidator() BTMetadataPropertyValidatorInfo`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BTMetadataCommonString) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BTMetadataCommonString) SetValidator(v BTMetadataPropertyValidatorInfo)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *BTMetadataCommonString) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetValue

`func (o *BTMetadataCommonString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTMetadataCommonString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTMetadataCommonString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BTMetadataCommonString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *BTMetadataCommonString) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTMetadataCommonString) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTMetadataCommonString) SetValueType(v string)`

SetValueType sets ValueType field to given value.



### AsBTMetadataItemsProperties

`func (s *BTMetadataCommonString) AsBTMetadataItemsProperties() BTMetadataItemsProperties`

Convenience method to wrap this instance of BTMetadataCommonString in BTMetadataItemsProperties

### AsBTMetadataPropertyInfo

`func (s *BTMetadataCommonString) AsBTMetadataPropertyInfo() BTMetadataPropertyInfo`

Convenience method to wrap this instance of BTMetadataCommonString in BTMetadataPropertyInfo

### AsBTWorkflowPropertyInfo

`func (s *BTMetadataCommonString) AsBTWorkflowPropertyInfo() BTWorkflowPropertyInfo`

Convenience method to wrap this instance of BTMetadataCommonString in BTWorkflowPropertyInfo

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


