# BTMetadataInfoItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementId** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **int32** | Possible values: 0&#x3D;PARTSTUDIO,1&#x3D;ASSEMBLY,2&#x3D;DRAWING,3&#x3D;FEATURESTUDIO,4&#x3D;BLOB,5&#x3D;APPLICATION,6&#x3D;TABLE,7&#x3D;BILLOFMATERIALS | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]BTMetadataItemsProperties**](BTMetadataItemsProperties.md) |  | [optional] 
**Parts** | Pointer to [**[]BTMetadataParts**](BTMetadataParts.md) |  | [optional] 

## Methods

### NewBTMetadataInfoItems

`func NewBTMetadataInfoItems() *BTMetadataInfoItems`

NewBTMetadataInfoItems instantiates a new BTMetadataInfoItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataInfoItemsWithDefaults

`func NewBTMetadataInfoItemsWithDefaults() *BTMetadataInfoItems`

NewBTMetadataInfoItemsWithDefaults instantiates a new BTMetadataInfoItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElementId

`func (o *BTMetadataInfoItems) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTMetadataInfoItems) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTMetadataInfoItems) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTMetadataInfoItems) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTMetadataInfoItems) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTMetadataInfoItems) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTMetadataInfoItems) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTMetadataInfoItems) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetMimeType

`func (o *BTMetadataInfoItems) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTMetadataInfoItems) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTMetadataInfoItems) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTMetadataInfoItems) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetHref

`func (o *BTMetadataInfoItems) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTMetadataInfoItems) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTMetadataInfoItems) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTMetadataInfoItems) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *BTMetadataInfoItems) GetProperties() []BTMetadataItemsProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTMetadataInfoItems) GetPropertiesOk() (*[]BTMetadataItemsProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTMetadataInfoItems) SetProperties(v []BTMetadataItemsProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTMetadataInfoItems) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetParts

`func (o *BTMetadataInfoItems) GetParts() []BTMetadataParts`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *BTMetadataInfoItems) GetPartsOk() (*[]BTMetadataParts, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *BTMetadataInfoItems) SetParts(v []BTMetadataParts)`

SetParts sets Parts field to given value.

### HasParts

`func (o *BTMetadataInfoItems) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


