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

// BodyPartMediaType struct for BodyPartMediaType
type BodyPartMediaType struct {
	Type *string `json:"type,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
	Parameters *map[string]string `json:"parameters,omitempty"`
	WildcardType *bool `json:"wildcardType,omitempty"`
	WildcardSubtype *bool `json:"wildcardSubtype,omitempty"`
}

// NewBodyPartMediaType instantiates a new BodyPartMediaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBodyPartMediaType() *BodyPartMediaType {
	this := BodyPartMediaType{}
	return &this
}

// NewBodyPartMediaTypeWithDefaults instantiates a new BodyPartMediaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodyPartMediaTypeWithDefaults() *BodyPartMediaType {
	this := BodyPartMediaType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BodyPartMediaType) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BodyPartMediaType) SetSubtype(v string) {
	o.Subtype = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *BodyPartMediaType) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetWildcardType returns the WildcardType field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetWildcardType() bool {
	if o == nil || o.WildcardType == nil {
		var ret bool
		return ret
	}
	return *o.WildcardType
}

// GetWildcardTypeOk returns a tuple with the WildcardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetWildcardTypeOk() (*bool, bool) {
	if o == nil || o.WildcardType == nil {
		return nil, false
	}
	return o.WildcardType, true
}

// HasWildcardType returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasWildcardType() bool {
	if o != nil && o.WildcardType != nil {
		return true
	}

	return false
}

// SetWildcardType gets a reference to the given bool and assigns it to the WildcardType field.
func (o *BodyPartMediaType) SetWildcardType(v bool) {
	o.WildcardType = &v
}

// GetWildcardSubtype returns the WildcardSubtype field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetWildcardSubtype() bool {
	if o == nil || o.WildcardSubtype == nil {
		var ret bool
		return ret
	}
	return *o.WildcardSubtype
}

// GetWildcardSubtypeOk returns a tuple with the WildcardSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetWildcardSubtypeOk() (*bool, bool) {
	if o == nil || o.WildcardSubtype == nil {
		return nil, false
	}
	return o.WildcardSubtype, true
}

// HasWildcardSubtype returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasWildcardSubtype() bool {
	if o != nil && o.WildcardSubtype != nil {
		return true
	}

	return false
}

// SetWildcardSubtype gets a reference to the given bool and assigns it to the WildcardSubtype field.
func (o *BodyPartMediaType) SetWildcardSubtype(v bool) {
	o.WildcardSubtype = &v
}

func (o BodyPartMediaType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.WildcardType != nil {
		toSerialize["wildcardType"] = o.WildcardType
	}
	if o.WildcardSubtype != nil {
		toSerialize["wildcardSubtype"] = o.WildcardSubtype
	}
	return json.Marshal(toSerialize)
}

type NullableBodyPartMediaType struct {
	value *BodyPartMediaType
	isSet bool
}

func (v NullableBodyPartMediaType) Get() *BodyPartMediaType {
	return v.value
}

func (v *NullableBodyPartMediaType) Set(val *BodyPartMediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableBodyPartMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableBodyPartMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodyPartMediaType(val *BodyPartMediaType) *NullableBodyPartMediaType {
	return &NullableBodyPartMediaType{value: val, isSet: true}
}

func (v NullableBodyPartMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodyPartMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
