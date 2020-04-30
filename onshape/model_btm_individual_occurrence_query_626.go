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

// BTMIndividualOccurrenceQuery626 struct for BTMIndividualOccurrenceQuery626
type BTMIndividualOccurrenceQuery626 struct {
	BTMIndividualQueryWithOccurrenceBase904
	BtType *string `json:"btType,omitempty"`
}

// NewBTMIndividualOccurrenceQuery626 instantiates a new BTMIndividualOccurrenceQuery626 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualOccurrenceQuery626() *BTMIndividualOccurrenceQuery626 {
	this := BTMIndividualOccurrenceQuery626{}
	return &this
}

// NewBTMIndividualOccurrenceQuery626WithDefaults instantiates a new BTMIndividualOccurrenceQuery626 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualOccurrenceQuery626WithDefaults() *BTMIndividualOccurrenceQuery626 {
	this := BTMIndividualOccurrenceQuery626{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualOccurrenceQuery626) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualOccurrenceQuery626) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualOccurrenceQuery626) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualOccurrenceQuery626) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMIndividualOccurrenceQuery626) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMIndividualQueryWithOccurrenceBase904, errBTMIndividualQueryWithOccurrenceBase904 := json.Marshal(o.BTMIndividualQueryWithOccurrenceBase904)
	if errBTMIndividualQueryWithOccurrenceBase904 != nil {
		return []byte{}, errBTMIndividualQueryWithOccurrenceBase904
	}
	errBTMIndividualQueryWithOccurrenceBase904 = json.Unmarshal([]byte(serializedBTMIndividualQueryWithOccurrenceBase904), &toSerialize)
	if errBTMIndividualQueryWithOccurrenceBase904 != nil {
		return []byte{}, errBTMIndividualQueryWithOccurrenceBase904
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMIndividualOccurrenceQuery626 struct {
	value *BTMIndividualOccurrenceQuery626
	isSet bool
}

func (v NullableBTMIndividualOccurrenceQuery626) Get() *BTMIndividualOccurrenceQuery626 {
	return v.value
}

func (v *NullableBTMIndividualOccurrenceQuery626) Set(val *BTMIndividualOccurrenceQuery626) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualOccurrenceQuery626) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualOccurrenceQuery626) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualOccurrenceQuery626(val *BTMIndividualOccurrenceQuery626) *NullableBTMIndividualOccurrenceQuery626 {
	return &NullableBTMIndividualOccurrenceQuery626{value: val, isSet: true}
}

func (v NullableBTMIndividualOccurrenceQuery626) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualOccurrenceQuery626) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
