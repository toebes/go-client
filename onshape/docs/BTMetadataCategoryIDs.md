# BTMetadataCategoryIDs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberCategoryIds** | Pointer to [**map[string]interface{}**](.md) | TODO: Not sure the type of this. | [optional] 
**MemberCategories** | Pointer to [**map[string]interface{}**](.md) | TODO: Not sure the type of this. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PublishState** | Pointer to **int32** |  | [optional] 
**ObjectTypes** | Pointer to **[]int32** |  | [optional] 
**DefaultObjectType** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewBTMetadataCategoryIDs

`func NewBTMetadataCategoryIDs() *BTMetadataCategoryIDs`

NewBTMetadataCategoryIDs instantiates a new BTMetadataCategoryIDs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataCategoryIDsWithDefaults

`func NewBTMetadataCategoryIDsWithDefaults() *BTMetadataCategoryIDs`

NewBTMetadataCategoryIDsWithDefaults instantiates a new BTMetadataCategoryIDs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberCategoryIds

`func (o *BTMetadataCategoryIDs) GetMemberCategoryIds() map[string]interface{}`

GetMemberCategoryIds returns the MemberCategoryIds field if non-nil, zero value otherwise.

### GetMemberCategoryIdsOk

`func (o *BTMetadataCategoryIDs) GetMemberCategoryIdsOk() (*map[string]interface{}, bool)`

GetMemberCategoryIdsOk returns a tuple with the MemberCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCategoryIds

`func (o *BTMetadataCategoryIDs) SetMemberCategoryIds(v map[string]interface{})`

SetMemberCategoryIds sets MemberCategoryIds field to given value.

### HasMemberCategoryIds

`func (o *BTMetadataCategoryIDs) HasMemberCategoryIds() bool`

HasMemberCategoryIds returns a boolean if a field has been set.

### GetMemberCategories

`func (o *BTMetadataCategoryIDs) GetMemberCategories() map[string]interface{}`

GetMemberCategories returns the MemberCategories field if non-nil, zero value otherwise.

### GetMemberCategoriesOk

`func (o *BTMetadataCategoryIDs) GetMemberCategoriesOk() (*map[string]interface{}, bool)`

GetMemberCategoriesOk returns a tuple with the MemberCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCategories

`func (o *BTMetadataCategoryIDs) SetMemberCategories(v map[string]interface{})`

SetMemberCategories sets MemberCategories field to given value.

### HasMemberCategories

`func (o *BTMetadataCategoryIDs) HasMemberCategories() bool`

HasMemberCategories returns a boolean if a field has been set.

### GetDescription

`func (o *BTMetadataCategoryIDs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTMetadataCategoryIDs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTMetadataCategoryIDs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTMetadataCategoryIDs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTMetadataCategoryIDs) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTMetadataCategoryIDs) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTMetadataCategoryIDs) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTMetadataCategoryIDs) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTMetadataCategoryIDs) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTMetadataCategoryIDs) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTMetadataCategoryIDs) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTMetadataCategoryIDs) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPublishState

`func (o *BTMetadataCategoryIDs) GetPublishState() int32`

GetPublishState returns the PublishState field if non-nil, zero value otherwise.

### GetPublishStateOk

`func (o *BTMetadataCategoryIDs) GetPublishStateOk() (*int32, bool)`

GetPublishStateOk returns a tuple with the PublishState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishState

`func (o *BTMetadataCategoryIDs) SetPublishState(v int32)`

SetPublishState sets PublishState field to given value.

### HasPublishState

`func (o *BTMetadataCategoryIDs) HasPublishState() bool`

HasPublishState returns a boolean if a field has been set.

### GetObjectTypes

`func (o *BTMetadataCategoryIDs) GetObjectTypes() []int32`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *BTMetadataCategoryIDs) GetObjectTypesOk() (*[]int32, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *BTMetadataCategoryIDs) SetObjectTypes(v []int32)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *BTMetadataCategoryIDs) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetDefaultObjectType

`func (o *BTMetadataCategoryIDs) GetDefaultObjectType() int32`

GetDefaultObjectType returns the DefaultObjectType field if non-nil, zero value otherwise.

### GetDefaultObjectTypeOk

`func (o *BTMetadataCategoryIDs) GetDefaultObjectTypeOk() (*int32, bool)`

GetDefaultObjectTypeOk returns a tuple with the DefaultObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultObjectType

`func (o *BTMetadataCategoryIDs) SetDefaultObjectType(v int32)`

SetDefaultObjectType sets DefaultObjectType field to given value.

### HasDefaultObjectType

`func (o *BTMetadataCategoryIDs) HasDefaultObjectType() bool`

HasDefaultObjectType returns a boolean if a field has been set.

### GetName

`func (o *BTMetadataCategoryIDs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTMetadataCategoryIDs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTMetadataCategoryIDs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTMetadataCategoryIDs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *BTMetadataCategoryIDs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTMetadataCategoryIDs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTMetadataCategoryIDs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTMetadataCategoryIDs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *BTMetadataCategoryIDs) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTMetadataCategoryIDs) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTMetadataCategoryIDs) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTMetadataCategoryIDs) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


