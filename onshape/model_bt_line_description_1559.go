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

// BTLineDescription1559 struct for BTLineDescription1559
type BTLineDescription1559 struct {
	BTCurveDescription1583
	BtType *string `json:"btType,omitempty"`
	Direction *BTVector3d389 `json:"direction,omitempty"`
	Origin *BTVector3d389 `json:"origin,omitempty"`
}

// NewBTLineDescription1559 instantiates a new BTLineDescription1559 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLineDescription1559() *BTLineDescription1559 {
	this := BTLineDescription1559{}
	return &this
}

// NewBTLineDescription1559WithDefaults instantiates a new BTLineDescription1559 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLineDescription1559WithDefaults() *BTLineDescription1559 {
	this := BTLineDescription1559{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTLineDescription1559) SetBtType(v string) {
	o.BtType = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTLineDescription1559) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTLineDescription1559) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLineDescription1559) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTLineDescription1559) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTLineDescription1559) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

func (o BTLineDescription1559) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTCurveDescription1583, errBTCurveDescription1583 := json.Marshal(o.BTCurveDescription1583)
	if errBTCurveDescription1583 != nil {
		return []byte{}, errBTCurveDescription1583
	}
	errBTCurveDescription1583 = json.Unmarshal([]byte(serializedBTCurveDescription1583), &toSerialize)
	if errBTCurveDescription1583 != nil {
		return []byte{}, errBTCurveDescription1583
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableBTLineDescription1559 struct {
	value *BTLineDescription1559
	isSet bool
}

func (v NullableBTLineDescription1559) Get() *BTLineDescription1559 {
	return v.value
}

func (v *NullableBTLineDescription1559) Set(val *BTLineDescription1559) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLineDescription1559) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLineDescription1559) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLineDescription1559(val *BTLineDescription1559) *NullableBTLineDescription1559 {
	return &NullableBTLineDescription1559{value: val, isSet: true}
}

func (v NullableBTLineDescription1559) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLineDescription1559) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
