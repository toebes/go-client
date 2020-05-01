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

// BTFolderInfoAllOf struct for BTFolderInfoAllOf
type BTFolderInfoAllOf struct {
	Active *bool `json:"active,omitempty"`
	CanUnshare *bool `json:"canUnshare,omitempty"`
	IsOrphaned *bool `json:"isOrphaned,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	PermissionSet *[]string `json:"permissionSet,omitempty"`
	Trash *bool `json:"trash,omitempty"`
	TrashedAt *JSONTime `json:"trashedAt,omitempty"`
}

// NewBTFolderInfoAllOf instantiates a new BTFolderInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFolderInfoAllOf() *BTFolderInfoAllOf {
	this := BTFolderInfoAllOf{}
	return &this
}

// NewBTFolderInfoAllOfWithDefaults instantiates a new BTFolderInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFolderInfoAllOfWithDefaults() *BTFolderInfoAllOf {
	this := BTFolderInfoAllOf{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BTFolderInfoAllOf) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfoAllOf) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BTFolderInfoAllOf) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *BTFolderInfoAllOf) SetActive(v bool) {
	o.Active = &v
}

// GetCanUnshare returns the CanUnshare field value if set, zero value otherwise.
func (o *BTFolderInfoAllOf) GetCanUnshare() bool {
	if o == nil || o.CanUnshare == nil {
		var ret bool
		return ret
	}
	return *o.CanUnshare
}

// GetCanUnshareOk returns a tuple with the CanUnshare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfoAllOf) GetCanUnshareOk() (*bool, bool) {
	if o == nil || o.CanUnshare == nil {
		return nil, false
	}
	return o.CanUnshare, true
}

// HasCanUnshare returns a boolean if a field has been set.
func (o *BTFolderInfoAllOf) HasCanUnshare() bool {
	if o != nil && o.CanUnshare != nil {
		return true
	}

	return false
}

// SetCanUnshare gets a reference to the given bool and assigns it to the CanUnshare field.
func (o *BTFolderInfoAllOf) SetCanUnshare(v bool) {
	o.CanUnshare = &v
}

// GetIsOrphaned returns the IsOrphaned field value if set, zero value otherwise.
func (o *BTFolderInfoAllOf) GetIsOrphaned() bool {
	if o == nil || o.IsOrphaned == nil {
		var ret bool
		return ret
	}
	return *o.IsOrphaned
}

// GetIsOrphanedOk returns a tuple with the IsOrphaned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfoAllOf) GetIsOrphanedOk() (*bool, bool) {
	if o == nil || o.IsOrphaned == nil {
		return nil, false
	}
	return o.IsOrphaned, true
}

// HasIsOrphaned returns a boolean if a field has been set.
func (o *BTFolderInfoAllOf) HasIsOrphaned() bool {
	if o != nil && o.IsOrphaned != nil {
		return true
	}

	return false
}

// SetIsOrphaned gets a reference to the given bool and assigns it to the IsOrphaned field.
func (o *BTFolderInfoAllOf) SetIsOrphaned(v bool) {
	o.IsOrphaned = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTFolderInfoAllOf) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfoAllOf) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTFolderInfoAllOf) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTFolderInfoAllOf) SetParentId(v string) {
	o.ParentId = &v
}

// GetPermissionSet returns the PermissionSet field value if set, zero value otherwise.
func (o *BTFolderInfoAllOf) GetPermissionSet() []string {
	if o == nil || o.PermissionSet == nil {
		var ret []string
		return ret
	}
	return *o.PermissionSet
}

// GetPermissionSetOk returns a tuple with the PermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfoAllOf) GetPermissionSetOk() (*[]string, bool) {
	if o == nil || o.PermissionSet == nil {
		return nil, false
	}
	return o.PermissionSet, true
}

// HasPermissionSet returns a boolean if a field has been set.
func (o *BTFolderInfoAllOf) HasPermissionSet() bool {
	if o != nil && o.PermissionSet != nil {
		return true
	}

	return false
}

// SetPermissionSet gets a reference to the given []string and assigns it to the PermissionSet field.
func (o *BTFolderInfoAllOf) SetPermissionSet(v []string) {
	o.PermissionSet = &v
}

// GetTrash returns the Trash field value if set, zero value otherwise.
func (o *BTFolderInfoAllOf) GetTrash() bool {
	if o == nil || o.Trash == nil {
		var ret bool
		return ret
	}
	return *o.Trash
}

// GetTrashOk returns a tuple with the Trash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfoAllOf) GetTrashOk() (*bool, bool) {
	if o == nil || o.Trash == nil {
		return nil, false
	}
	return o.Trash, true
}

// HasTrash returns a boolean if a field has been set.
func (o *BTFolderInfoAllOf) HasTrash() bool {
	if o != nil && o.Trash != nil {
		return true
	}

	return false
}

// SetTrash gets a reference to the given bool and assigns it to the Trash field.
func (o *BTFolderInfoAllOf) SetTrash(v bool) {
	o.Trash = &v
}

// GetTrashedAt returns the TrashedAt field value if set, zero value otherwise.
func (o *BTFolderInfoAllOf) GetTrashedAt() JSONTime {
	if o == nil || o.TrashedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.TrashedAt
}

// GetTrashedAtOk returns a tuple with the TrashedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFolderInfoAllOf) GetTrashedAtOk() (*JSONTime, bool) {
	if o == nil || o.TrashedAt == nil {
		return nil, false
	}
	return o.TrashedAt, true
}

// HasTrashedAt returns a boolean if a field has been set.
func (o *BTFolderInfoAllOf) HasTrashedAt() bool {
	if o != nil && o.TrashedAt != nil {
		return true
	}

	return false
}

// SetTrashedAt gets a reference to the given JSONTime and assigns it to the TrashedAt field.
func (o *BTFolderInfoAllOf) SetTrashedAt(v JSONTime) {
	o.TrashedAt = &v
}

func (o BTFolderInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.CanUnshare != nil {
		toSerialize["canUnshare"] = o.CanUnshare
	}
	if o.IsOrphaned != nil {
		toSerialize["isOrphaned"] = o.IsOrphaned
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.PermissionSet != nil {
		toSerialize["permissionSet"] = o.PermissionSet
	}
	if o.Trash != nil {
		toSerialize["trash"] = o.Trash
	}
	if o.TrashedAt != nil {
		toSerialize["trashedAt"] = o.TrashedAt
	}
	return json.Marshal(toSerialize)
}

type NullableBTFolderInfoAllOf struct {
	value *BTFolderInfoAllOf
	isSet bool
}

func (v NullableBTFolderInfoAllOf) Get() *BTFolderInfoAllOf {
	return v.value
}

func (v *NullableBTFolderInfoAllOf) Set(val *BTFolderInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFolderInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFolderInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFolderInfoAllOf(val *BTFolderInfoAllOf) *NullableBTFolderInfoAllOf {
	return &NullableBTFolderInfoAllOf{value: val, isSet: true}
}

func (v NullableBTFolderInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFolderInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
