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

// BTVersionOrWorkspaceInfo struct for BTVersionOrWorkspaceInfo
type BTVersionOrWorkspaceInfo struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewBTVersionOrWorkspaceInfo instantiates a new BTVersionOrWorkspaceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVersionOrWorkspaceInfo() *BTVersionOrWorkspaceInfo {
	this := BTVersionOrWorkspaceInfo{}
	return &this
}

// NewBTVersionOrWorkspaceInfoWithDefaults instantiates a new BTVersionOrWorkspaceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVersionOrWorkspaceInfoWithDefaults() *BTVersionOrWorkspaceInfo {
	this := BTVersionOrWorkspaceInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTVersionOrWorkspaceInfo) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTVersionOrWorkspaceInfo) SetType(v string) {
	o.Type = &v
}

func (o BTVersionOrWorkspaceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTVersionOrWorkspaceInfo struct {
	value *BTVersionOrWorkspaceInfo
	isSet bool
}

func (v NullableBTVersionOrWorkspaceInfo) Get() *BTVersionOrWorkspaceInfo {
	return v.value
}

func (v *NullableBTVersionOrWorkspaceInfo) Set(val *BTVersionOrWorkspaceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVersionOrWorkspaceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVersionOrWorkspaceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVersionOrWorkspaceInfo(val *BTVersionOrWorkspaceInfo) *NullableBTVersionOrWorkspaceInfo {
	return &NullableBTVersionOrWorkspaceInfo{value: val, isSet: true}
}

func (v NullableBTVersionOrWorkspaceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVersionOrWorkspaceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}