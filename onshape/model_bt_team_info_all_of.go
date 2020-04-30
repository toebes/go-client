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

// BTTeamInfoAllOf struct for BTTeamInfoAllOf
type BTTeamInfoAllOf struct {
	Admin *bool `json:"admin,omitempty"`
	Member *bool `json:"member,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewBTTeamInfoAllOf instantiates a new BTTeamInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTeamInfoAllOf() *BTTeamInfoAllOf {
	this := BTTeamInfoAllOf{}
	return &this
}

// NewBTTeamInfoAllOfWithDefaults instantiates a new BTTeamInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTeamInfoAllOfWithDefaults() *BTTeamInfoAllOf {
	this := BTTeamInfoAllOf{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *BTTeamInfoAllOf) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamInfoAllOf) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *BTTeamInfoAllOf) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *BTTeamInfoAllOf) SetAdmin(v bool) {
	o.Admin = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *BTTeamInfoAllOf) GetMember() bool {
	if o == nil || o.Member == nil {
		var ret bool
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamInfoAllOf) GetMemberOk() (*bool, bool) {
	if o == nil || o.Member == nil {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *BTTeamInfoAllOf) HasMember() bool {
	if o != nil && o.Member != nil {
		return true
	}

	return false
}

// SetMember gets a reference to the given bool and assigns it to the Member field.
func (o *BTTeamInfoAllOf) SetMember(v bool) {
	o.Member = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BTTeamInfoAllOf) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTeamInfoAllOf) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BTTeamInfoAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *BTTeamInfoAllOf) SetSize(v int32) {
	o.Size = &v
}

func (o BTTeamInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.Member != nil {
		toSerialize["member"] = o.Member
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableBTTeamInfoAllOf struct {
	value *BTTeamInfoAllOf
	isSet bool
}

func (v NullableBTTeamInfoAllOf) Get() *BTTeamInfoAllOf {
	return v.value
}

func (v *NullableBTTeamInfoAllOf) Set(val *BTTeamInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTeamInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTeamInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTeamInfoAllOf(val *BTTeamInfoAllOf) *NullableBTTeamInfoAllOf {
	return &NullableBTTeamInfoAllOf{value: val, isSet: true}
}

func (v NullableBTTeamInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTeamInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
