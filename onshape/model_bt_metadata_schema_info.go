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

// BTMetadataSchemaInfo struct for BTMetadataSchemaInfo
type BTMetadataSchemaInfo struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ObjectType *int32 `json:"objectType,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerType *int32 `json:"ownerType,omitempty"`
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTMetadataSchemaInfo instantiates a new BTMetadataSchemaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataSchemaInfo() *BTMetadataSchemaInfo {
	this := BTMetadataSchemaInfo{}
	return &this
}

// NewBTMetadataSchemaInfoWithDefaults instantiates a new BTMetadataSchemaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataSchemaInfoWithDefaults() *BTMetadataSchemaInfo {
	this := BTMetadataSchemaInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTMetadataSchemaInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataSchemaInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTMetadataSchemaInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTMetadataSchemaInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTMetadataSchemaInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataSchemaInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTMetadataSchemaInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTMetadataSchemaInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMetadataSchemaInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataSchemaInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMetadataSchemaInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMetadataSchemaInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTMetadataSchemaInfo) GetObjectType() int32 {
	if o == nil || o.ObjectType == nil {
		var ret int32
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataSchemaInfo) GetObjectTypeOk() (*int32, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTMetadataSchemaInfo) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given int32 and assigns it to the ObjectType field.
func (o *BTMetadataSchemaInfo) SetObjectType(v int32) {
	o.ObjectType = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTMetadataSchemaInfo) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataSchemaInfo) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTMetadataSchemaInfo) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTMetadataSchemaInfo) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTMetadataSchemaInfo) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataSchemaInfo) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTMetadataSchemaInfo) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTMetadataSchemaInfo) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTMetadataSchemaInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataSchemaInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTMetadataSchemaInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTMetadataSchemaInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTMetadataSchemaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataSchemaInfo struct {
	value *BTMetadataSchemaInfo
	isSet bool
}

func (v NullableBTMetadataSchemaInfo) Get() *BTMetadataSchemaInfo {
	return v.value
}

func (v *NullableBTMetadataSchemaInfo) Set(val *BTMetadataSchemaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataSchemaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataSchemaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataSchemaInfo(val *BTMetadataSchemaInfo) *NullableBTMetadataSchemaInfo {
	return &NullableBTMetadataSchemaInfo{value: val, isSet: true}
}

func (v NullableBTMetadataSchemaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataSchemaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
