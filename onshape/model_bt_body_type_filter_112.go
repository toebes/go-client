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

// BTBodyTypeFilter112 struct for BTBodyTypeFilter112
type BTBodyTypeFilter112 struct {
	BTQueryFilter183
	BodyType *string `json:"bodyType,omitempty"`
	BtType *string `json:"btType,omitempty"`
}

// NewBTBodyTypeFilter112 instantiates a new BTBodyTypeFilter112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBodyTypeFilter112() *BTBodyTypeFilter112 {
	this := BTBodyTypeFilter112{}
	return &this
}

// NewBTBodyTypeFilter112WithDefaults instantiates a new BTBodyTypeFilter112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBodyTypeFilter112WithDefaults() *BTBodyTypeFilter112 {
	this := BTBodyTypeFilter112{}
	return &this
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *BTBodyTypeFilter112) GetBodyType() string {
	if o == nil || o.BodyType == nil {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBodyTypeFilter112) GetBodyTypeOk() (*string, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *BTBodyTypeFilter112) HasBodyType() bool {
	if o != nil && o.BodyType != nil {
		return true
	}

	return false
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *BTBodyTypeFilter112) SetBodyType(v string) {
	o.BodyType = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBodyTypeFilter112) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBodyTypeFilter112) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBodyTypeFilter112) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBodyTypeFilter112) SetBtType(v string) {
	o.BtType = &v
}

func (o BTBodyTypeFilter112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTBodyTypeFilter112 struct {
	value *BTBodyTypeFilter112
	isSet bool
}

func (v NullableBTBodyTypeFilter112) Get() *BTBodyTypeFilter112 {
	return v.value
}

func (v *NullableBTBodyTypeFilter112) Set(val *BTBodyTypeFilter112) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBodyTypeFilter112) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBodyTypeFilter112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBodyTypeFilter112(val *BTBodyTypeFilter112) *NullableBTBodyTypeFilter112 {
	return &NullableBTBodyTypeFilter112{value: val, isSet: true}
}

func (v NullableBTBodyTypeFilter112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBodyTypeFilter112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
