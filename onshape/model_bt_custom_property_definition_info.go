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

// BTCustomPropertyDefinitionInfo struct for BTCustomPropertyDefinitionInfo
type BTCustomPropertyDefinitionInfo struct {
	Description *string `json:"description,omitempty"`
	EnumDefinition *[]string `json:"enumDefinition,omitempty"`
	Name *string `json:"name,omitempty"`
	Template *string `json:"template,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewBTCustomPropertyDefinitionInfo instantiates a new BTCustomPropertyDefinitionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCustomPropertyDefinitionInfo() *BTCustomPropertyDefinitionInfo {
	this := BTCustomPropertyDefinitionInfo{}
	return &this
}

// NewBTCustomPropertyDefinitionInfoWithDefaults instantiates a new BTCustomPropertyDefinitionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCustomPropertyDefinitionInfoWithDefaults() *BTCustomPropertyDefinitionInfo {
	this := BTCustomPropertyDefinitionInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTCustomPropertyDefinitionInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCustomPropertyDefinitionInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTCustomPropertyDefinitionInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTCustomPropertyDefinitionInfo) SetDescription(v string) {
	o.Description = &v
}

// GetEnumDefinition returns the EnumDefinition field value if set, zero value otherwise.
func (o *BTCustomPropertyDefinitionInfo) GetEnumDefinition() []string {
	if o == nil || o.EnumDefinition == nil {
		var ret []string
		return ret
	}
	return *o.EnumDefinition
}

// GetEnumDefinitionOk returns a tuple with the EnumDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCustomPropertyDefinitionInfo) GetEnumDefinitionOk() (*[]string, bool) {
	if o == nil || o.EnumDefinition == nil {
		return nil, false
	}
	return o.EnumDefinition, true
}

// HasEnumDefinition returns a boolean if a field has been set.
func (o *BTCustomPropertyDefinitionInfo) HasEnumDefinition() bool {
	if o != nil && o.EnumDefinition != nil {
		return true
	}

	return false
}

// SetEnumDefinition gets a reference to the given []string and assigns it to the EnumDefinition field.
func (o *BTCustomPropertyDefinitionInfo) SetEnumDefinition(v []string) {
	o.EnumDefinition = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCustomPropertyDefinitionInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCustomPropertyDefinitionInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCustomPropertyDefinitionInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCustomPropertyDefinitionInfo) SetName(v string) {
	o.Name = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *BTCustomPropertyDefinitionInfo) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCustomPropertyDefinitionInfo) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *BTCustomPropertyDefinitionInfo) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *BTCustomPropertyDefinitionInfo) SetTemplate(v string) {
	o.Template = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTCustomPropertyDefinitionInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCustomPropertyDefinitionInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTCustomPropertyDefinitionInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTCustomPropertyDefinitionInfo) SetType(v string) {
	o.Type = &v
}

func (o BTCustomPropertyDefinitionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EnumDefinition != nil {
		toSerialize["enumDefinition"] = o.EnumDefinition
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTCustomPropertyDefinitionInfo struct {
	value *BTCustomPropertyDefinitionInfo
	isSet bool
}

func (v NullableBTCustomPropertyDefinitionInfo) Get() *BTCustomPropertyDefinitionInfo {
	return v.value
}

func (v *NullableBTCustomPropertyDefinitionInfo) Set(val *BTCustomPropertyDefinitionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCustomPropertyDefinitionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCustomPropertyDefinitionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCustomPropertyDefinitionInfo(val *BTCustomPropertyDefinitionInfo) *NullableBTCustomPropertyDefinitionInfo {
	return &NullableBTCustomPropertyDefinitionInfo{value: val, isSet: true}
}

func (v NullableBTCustomPropertyDefinitionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCustomPropertyDefinitionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
