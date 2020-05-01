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

// BTPartMaterialPropertyInfo struct for BTPartMaterialPropertyInfo
type BTPartMaterialPropertyInfo struct {
	Category *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Units *string `json:"units,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewBTPartMaterialPropertyInfo instantiates a new BTPartMaterialPropertyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartMaterialPropertyInfo() *BTPartMaterialPropertyInfo {
	this := BTPartMaterialPropertyInfo{}
	return &this
}

// NewBTPartMaterialPropertyInfoWithDefaults instantiates a new BTPartMaterialPropertyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartMaterialPropertyInfoWithDefaults() *BTPartMaterialPropertyInfo {
	this := BTPartMaterialPropertyInfo{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *BTPartMaterialPropertyInfo) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialPropertyInfo) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *BTPartMaterialPropertyInfo) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *BTPartMaterialPropertyInfo) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTPartMaterialPropertyInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialPropertyInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTPartMaterialPropertyInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTPartMaterialPropertyInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTPartMaterialPropertyInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialPropertyInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTPartMaterialPropertyInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTPartMaterialPropertyInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPartMaterialPropertyInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialPropertyInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPartMaterialPropertyInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPartMaterialPropertyInfo) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTPartMaterialPropertyInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialPropertyInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTPartMaterialPropertyInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTPartMaterialPropertyInfo) SetType(v string) {
	o.Type = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTPartMaterialPropertyInfo) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialPropertyInfo) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTPartMaterialPropertyInfo) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTPartMaterialPropertyInfo) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTPartMaterialPropertyInfo) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialPropertyInfo) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTPartMaterialPropertyInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTPartMaterialPropertyInfo) SetValue(v string) {
	o.Value = &v
}

func (o BTPartMaterialPropertyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartMaterialPropertyInfo struct {
	value *BTPartMaterialPropertyInfo
	isSet bool
}

func (v NullableBTPartMaterialPropertyInfo) Get() *BTPartMaterialPropertyInfo {
	return v.value
}

func (v *NullableBTPartMaterialPropertyInfo) Set(val *BTPartMaterialPropertyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartMaterialPropertyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartMaterialPropertyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartMaterialPropertyInfo(val *BTPartMaterialPropertyInfo) *NullableBTPartMaterialPropertyInfo {
	return &NullableBTPartMaterialPropertyInfo{value: val, isSet: true}
}

func (v NullableBTPartMaterialPropertyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartMaterialPropertyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}