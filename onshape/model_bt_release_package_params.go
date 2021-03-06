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

// BTReleasePackageParams struct for BTReleasePackageParams
type BTReleasePackageParams struct {
	Items *[]BTReleasePackageItemParams `json:"items,omitempty"`
}

// NewBTReleasePackageParams instantiates a new BTReleasePackageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleasePackageParams() *BTReleasePackageParams {
	this := BTReleasePackageParams{}
	return &this
}

// NewBTReleasePackageParamsWithDefaults instantiates a new BTReleasePackageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleasePackageParamsWithDefaults() *BTReleasePackageParams {
	this := BTReleasePackageParams{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTReleasePackageParams) GetItems() []BTReleasePackageItemParams {
	if o == nil || o.Items == nil {
		var ret []BTReleasePackageItemParams
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageParams) GetItemsOk() (*[]BTReleasePackageItemParams, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTReleasePackageParams) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTReleasePackageItemParams and assigns it to the Items field.
func (o *BTReleasePackageParams) SetItems(v []BTReleasePackageItemParams) {
	o.Items = &v
}

func (o BTReleasePackageParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleasePackageParams struct {
	value *BTReleasePackageParams
	isSet bool
}

func (v NullableBTReleasePackageParams) Get() *BTReleasePackageParams {
	return v.value
}

func (v *NullableBTReleasePackageParams) Set(val *BTReleasePackageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleasePackageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleasePackageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleasePackageParams(val *BTReleasePackageParams) *NullableBTReleasePackageParams {
	return &NullableBTReleasePackageParams{value: val, isSet: true}
}

func (v NullableBTReleasePackageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleasePackageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
