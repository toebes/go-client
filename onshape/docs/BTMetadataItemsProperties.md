# BTMetadataItemsProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**EditableInUi** | Pointer to **bool** |  | [optional] 
**EnumValues** | Pointer to [**[]BTMetadataEnumValueInfo**](BTMetadataEnumValueInfo.md) |  | [optional] 
**Multivalued** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PropertyId** | Pointer to **string** |  | [optional] 
**PropertySource** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**UiHints** | Pointer to [**BTMetadataPropertyUiHintsInfo**](BTMetadataPropertyUiHintsInfo.md) |  | [optional] 
**Validator** | Pointer to [**BTMetadataPropertyValidatorInfo**](BTMetadataPropertyValidatorInfo.md) |  | [optional] 
**Value** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewBTMetadataItemsProperties

`func NewBTMetadataItemsProperties() *BTMetadataItemsProperties`

NewBTMetadataItemsProperties instantiates a new BTMetadataItemsProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataItemsPropertiesWithDefaults

`func NewBTMetadataItemsPropertiesWithDefaults() *BTMetadataItemsProperties`

NewBTMetadataItemsPropertiesWithDefaults instantiates a new BTMetadataItemsProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *BTMetadataItemsProperties) GetDefaultValue() map[string]interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BTMetadataItemsProperties) GetDefaultValueOk() (*map[string]interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BTMetadataItemsProperties) SetDefaultValue(v map[string]interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BTMetadataItemsProperties) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetEditable

`func (o *BTMetadataItemsProperties) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *BTMetadataItemsProperties) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *BTMetadataItemsProperties) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *BTMetadataItemsProperties) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEditableInUi

`func (o *BTMetadataItemsProperties) GetEditableInUi() bool`

GetEditableInUi returns the EditableInUi field if non-nil, zero value otherwise.

### GetEditableInUiOk

`func (o *BTMetadataItemsProperties) GetEditableInUiOk() (*bool, bool)`

GetEditableInUiOk returns a tuple with the EditableInUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableInUi

`func (o *BTMetadataItemsProperties) SetEditableInUi(v bool)`

SetEditableInUi sets EditableInUi field to given value.

### HasEditableInUi

`func (o *BTMetadataItemsProperties) HasEditableInUi() bool`

HasEditableInUi returns a boolean if a field has been set.

### GetEnumValues

`func (o *BTMetadataItemsProperties) GetEnumValues() []BTMetadataEnumValueInfo`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *BTMetadataItemsProperties) GetEnumValuesOk() (*[]BTMetadataEnumValueInfo, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *BTMetadataItemsProperties) SetEnumValues(v []BTMetadataEnumValueInfo)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *BTMetadataItemsProperties) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetMultivalued

`func (o *BTMetadataItemsProperties) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *BTMetadataItemsProperties) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *BTMetadataItemsProperties) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *BTMetadataItemsProperties) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.

### GetName

`func (o *BTMetadataItemsProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMetadataItemsProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMetadataItemsProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMetadataItemsProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyId

`func (o *BTMetadataItemsProperties) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTMetadataItemsProperties) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTMetadataItemsProperties) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTMetadataItemsProperties) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPropertySource

`func (o *BTMetadataItemsProperties) GetPropertySource() int32`

GetPropertySource returns the PropertySource field if non-nil, zero value otherwise.

### GetPropertySourceOk

`func (o *BTMetadataItemsProperties) GetPropertySourceOk() (*int32, bool)`

GetPropertySourceOk returns a tuple with the PropertySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySource

`func (o *BTMetadataItemsProperties) SetPropertySource(v int32)`

SetPropertySource sets PropertySource field to given value.

### HasPropertySource

`func (o *BTMetadataItemsProperties) HasPropertySource() bool`

HasPropertySource returns a boolean if a field has been set.

### GetRequired

`func (o *BTMetadataItemsProperties) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BTMetadataItemsProperties) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BTMetadataItemsProperties) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BTMetadataItemsProperties) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchemaId

`func (o *BTMetadataItemsProperties) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *BTMetadataItemsProperties) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *BTMetadataItemsProperties) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *BTMetadataItemsProperties) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetUiHints

`func (o *BTMetadataItemsProperties) GetUiHints() BTMetadataPropertyUiHintsInfo`

GetUiHints returns the UiHints field if non-nil, zero value otherwise.

### GetUiHintsOk

`func (o *BTMetadataItemsProperties) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool)`

GetUiHintsOk returns a tuple with the UiHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHints

`func (o *BTMetadataItemsProperties) SetUiHints(v BTMetadataPropertyUiHintsInfo)`

SetUiHints sets UiHints field to given value.

### HasUiHints

`func (o *BTMetadataItemsProperties) HasUiHints() bool`

HasUiHints returns a boolean if a field has been set.

### GetValidator

`func (o *BTMetadataItemsProperties) GetValidator() BTMetadataPropertyValidatorInfo`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BTMetadataItemsProperties) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BTMetadataItemsProperties) SetValidator(v BTMetadataPropertyValidatorInfo)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *BTMetadataItemsProperties) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetValue

`func (o *BTMetadataItemsProperties) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTMetadataItemsProperties) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTMetadataItemsProperties) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BTMetadataItemsProperties) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *BTMetadataItemsProperties) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BTMetadataItemsProperties) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BTMetadataItemsProperties) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *BTMetadataItemsProperties) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


