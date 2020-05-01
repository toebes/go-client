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

// BTSphereDescription1263AllOf struct for BTSphereDescription1263AllOf
type BTSphereDescription1263AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Origin *BTVector3d389 `json:"origin,omitempty"`
	Radius *float64 `json:"radius,omitempty"`
}

// NewBTSphereDescription1263AllOf instantiates a new BTSphereDescription1263AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSphereDescription1263AllOf() *BTSphereDescription1263AllOf {
	this := BTSphereDescription1263AllOf{}
	return &this
}

// NewBTSphereDescription1263AllOfWithDefaults instantiates a new BTSphereDescription1263AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSphereDescription1263AllOfWithDefaults() *BTSphereDescription1263AllOf {
	this := BTSphereDescription1263AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSphereDescription1263AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSphereDescription1263AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSphereDescription1263AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSphereDescription1263AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTSphereDescription1263AllOf) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSphereDescription1263AllOf) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTSphereDescription1263AllOf) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTSphereDescription1263AllOf) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTSphereDescription1263AllOf) GetRadius() float64 {
	if o == nil || o.Radius == nil {
		var ret float64
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSphereDescription1263AllOf) GetRadiusOk() (*float64, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BTSphereDescription1263AllOf) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *BTSphereDescription1263AllOf) SetRadius(v float64) {
	o.Radius = &v
}

func (o BTSphereDescription1263AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableBTSphereDescription1263AllOf struct {
	value *BTSphereDescription1263AllOf
	isSet bool
}

func (v NullableBTSphereDescription1263AllOf) Get() *BTSphereDescription1263AllOf {
	return v.value
}

func (v *NullableBTSphereDescription1263AllOf) Set(val *BTSphereDescription1263AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSphereDescription1263AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSphereDescription1263AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSphereDescription1263AllOf(val *BTSphereDescription1263AllOf) *NullableBTSphereDescription1263AllOf {
	return &NullableBTSphereDescription1263AllOf{value: val, isSet: true}
}

func (v NullableBTSphereDescription1263AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSphereDescription1263AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}