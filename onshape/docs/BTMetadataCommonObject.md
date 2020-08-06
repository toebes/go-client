# BTMetadataCommonObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** |  | [optional] 
**Dirty** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**EditableInUi** | Pointer to **bool** |  | [optional] 
**EnumValues** | Pointer to [**[]BTMetadataEnumValueInfo**](BTMetadataEnumValueInfo.md) |  | [optional] 
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
**Value** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**ValueType** | Pointer to **string** |  | 

## Methods

### NewBTMetadataCommonObject

`func NewBTMetadataCommonObject(valueType string, ) *BTMetadataCommonObject`

NewBTMetadataCommonObject instantiates a new BTMetadataCommonObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataCommonObjectWithDefaults

`func NewBTMetadataCommonObjectWithDefaults() *BTMetadataCommonObject`

NewBTMetadataCommonObjectWithDefaults instantiates a new BTMetadataCommonObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *BTMetadataCommonObject) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BTMetadataCommonObject) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BTMetadataCommonObject) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BTMetadataCommonObject) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDirty

`func (o *BTMetadataCommonObject) GetDirty() bool`

GetDirty returns the Dirty field if non-nil, zero value otherwise.

### GetDirtyOk

`func (o *BTMetadataCommonObject) GetDirtyOk() (*bool, bool)`

GetDirtyOk returns a tuple with the Dirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirty

`func (o *BTMetadataCommonObject) SetDirty(v bool)`

SetDirty sets Dirty field to given value.

### HasDirty

`func (o *BTMetadataCommonObject) HasDirty() bool`

HasDirty returns a boolean if a field has been set.

### GetEditable

`func (o *BTMetadataCommonObject) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *BTMetadataCommonObject) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *BTMetadataCommonObject) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *BTMetadataCommonObject) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEditableInUi

`func (o *BTMetadataCommonObject) GetEditableInUi() bool`

GetEditableInUi returns the EditableInUi field if non-nil, zero value otherwise.

### GetEditableInUiOk

`func (o *BTMetadataCommonObject) GetEditableInUiOk() (*bool, bool)`

GetEditableInUiOk returns a tuple with the EditableInUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInUi

`func (o *BTMetadataCommonObject) SetEditableInUi(v bool)`

SetEditableInUi sets EditableInUi field to given value.

### HasEditableInUi

`func (o *BTMetadataCommonObject) HasEditableInUi() bool`

HasEditableInUi returns a boolean if a field has been set.

### GetEnumValues

`func (o *BTMetadataCommonObject) GetEnumValues() []BTMetadataEnumValueInfo`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *BTMetadataCommonObject) GetEnumValuesOk() (*[]BTMetadataEnumValueInfo, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *BTMetadataCommonObject) SetEnumValues(v []BTMetadataEnumValueInfo)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *BTMetadataCommonObject) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetInitialValue

`func (o *BTMetadataCommonObject) GetInitialValue() map[string]interface{}`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *BTMetadataCommonObject) GetInitialValueOk() (*map[string]interface{}, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *BTMetadataCommonObject) SetInitialValue(v map[string]interface{})`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *BTMetadataCommonObject) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetIsApproverProperty

`func (o *BTMetadataCommonObject) GetIsApproverProperty() bool`

GetIsApproverProperty returns the IsApproverProperty field if non-nil, zero value otherwise.

### GetIsApproverPropertyOk

`func (o *BTMetadataCommonObject) GetIsApproverPropertyOk() (*bool, bool)`

GetIsApproverPropertyOk returns a tuple with the IsApproverProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproverProperty

`func (o *BTMetadataCommonObject) SetIsApproverProperty(v bool)`

SetIsApproverProperty sets IsApproverProperty field to given value.

### HasIsApproverProperty

`func (o *BTMetadataCommonObject) HasIsApproverProperty() bool`

HasIsApproverProperty returns a boolean if a field has been set.

### GetIsNotifierProperty

`func (o *BTMetadataCommonObject) GetIsNotifierProperty() bool`

GetIsNotifierProperty returns the IsNotifierProperty field if non-nil, zero value otherwise.

### GetIsNotifierPropertyOk

`func (o *BTMetadataCommonObject) GetIsNotifierPropertyOk() (*bool, bool)`

GetIsNotifierPropertyOk returns a tuple with the IsNotifierProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotifierProperty

`func (o *BTMetadataCommonObject) SetIsNotifierProperty(v bool)`

SetIsNotifierProperty sets IsNotifierProperty field to given value.

### HasIsNotifierProperty

`func (o *BTMetadataCommonObject) HasIsNotifierProperty() bool`

HasIsNotifierProperty returns a boolean if a field has been set.

### GetMultivalued

`func (o *BTMetadataCommonObject) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *BTMetadataCommonObject) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *BTMetadataCommonObject) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *BTMetadataCommonObject) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.

### GetName

`func (o *BTMetadataCommonObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMetadataCommonObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMetadataCommonObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMetadataCommonObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyId

`func (o *BTMetadataCommonObject) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTMetadataCommonObject) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTMetadataCommonObject) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTMetadataCommonObject) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPropertySource

`func (o *BTMetadataCommonObject) GetPropertySource() int32`

GetPropertySource returns the PropertySource field if non-nil, zero value otherwise.

### GetPropertySourceOk

`func (o *BTMetadataCommonObject) GetPropertySourceOk() (*int32, bool)`

GetPropertySourceOk returns a tuple with the PropertySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySource

`func (o *BTMetadataCommonObject) SetPropertySource(v int32)`

SetPropertySource sets PropertySource field to given value.

### HasPropertySource

`func (o *BTMetadataCommonObject) HasPropertySource() bool`

HasPropertySource returns a boolean if a field has been set.

### GetRequired

`func (o *BTMetadataCommonObject) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BTMetadataCommonObject) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BTMetadataCommonObject) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BTMetadataCommonObject) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchemaId

`func (o *BTMetadataCommonObject) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *BTMetadataCommonObject) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *BTMetadataCommonObject) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *BTMetadataCommonObject) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetUiHints

`func (o *BTMetadataCommonObject) GetUiHints() BTMetadataPropertyUiHintsInfo`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *BTMetadataCommonObject) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *BTMetadataCommonObject) SetUiHints(v BTMetadataPropertyUiHintsInfo)`

SetUiHints sets UiHints field to given value.

### HasUiHints

`func (o *BTMetadataCommonObject) HasUiHints() bool`

HasUiHints returns a boolean if a field has been set.

### GetValidator

`func (o *BTMetadataCommonObject) GetValidator() BTMetadataPropertyValidatorInfo`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BTMetadataCommonObject) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BTMetadataCommonObject) SetValidator(v BTMetadataPropertyValidatorInfo)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *BTMetadataCommonObject) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetValue

`func (o *BTMetadataCommonObject) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTMetadataCommonObject) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTMetadataCommonObject) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BTMetadataCommonObject) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *BTMetadataCommonObject) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTMetadataCommonObject) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTMetadataCommonObject) SetValueType(v string)`

SetValueType sets ValueType field to given value.



### AsBTMetadataItemsProperties

`func (s *BTMetadataCommonObject) AsBTMetadataItemsProperties() BTMetadataItemsProperties`

Convenience method to wrap this instance of BTMetadataCommonObject in BTMetadataItemsProperties

### AsBTMetadataPropertyInfo

`func (s *BTMetadataCommonObject) AsBTMetadataPropertyInfo() BTMetadataPropertyInfo`

Convenience method to wrap this instance of BTMetadataCommonObject in BTMetadataPropertyInfo

### AsBTWorkflowPropertyInfo

`func (s *BTMetadataCommonObject) AsBTWorkflowPropertyInfo() BTWorkflowPropertyInfo`

Convenience method to wrap this instance of BTMetadataCommonObject in BTWorkflowPropertyInfo

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


