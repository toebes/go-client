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

// BTMIndividualQueryWithOccurrence811AllOf struct for BTMIndividualQueryWithOccurrence811AllOf
type BTMIndividualQueryWithOccurrence811AllOf struct {
	BtType *string `json:"btType,omitempty"`
	EntityQuery *string `json:"entityQuery,omitempty"`
}

// NewBTMIndividualQueryWithOccurrence811AllOf instantiates a new BTMIndividualQueryWithOccurrence811AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualQueryWithOccurrence811AllOf() *BTMIndividualQueryWithOccurrence811AllOf {
	this := BTMIndividualQueryWithOccurrence811AllOf{}
	return &this
}

// NewBTMIndividualQueryWithOccurrence811AllOfWithDefaults instantiates a new BTMIndividualQueryWithOccurrence811AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualQueryWithOccurrence811AllOfWithDefaults() *BTMIndividualQueryWithOccurrence811AllOf {
	this := BTMIndividualQueryWithOccurrence811AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualQueryWithOccurrence811AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetEntityQuery returns the EntityQuery field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811AllOf) GetEntityQuery() string {
	if o == nil || o.EntityQuery == nil {
		var ret string
		return ret
	}
	return *o.EntityQuery
}

// GetEntityQueryOk returns a tuple with the EntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811AllOf) GetEntityQueryOk() (*string, bool) {
	if o == nil || o.EntityQuery == nil {
		return nil, false
	}
	return o.EntityQuery, true
}

// HasEntityQuery returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811AllOf) HasEntityQuery() bool {
	if o != nil && o.EntityQuery != nil {
		return true
	}

	return false
}

// SetEntityQuery gets a reference to the given string and assigns it to the EntityQuery field.
func (o *BTMIndividualQueryWithOccurrence811AllOf) SetEntityQuery(v string) {
	o.EntityQuery = &v
}

func (o BTMIndividualQueryWithOccurrence811AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EntityQuery != nil {
		toSerialize["entityQuery"] = o.EntityQuery
	}
	return json.Marshal(toSerialize)
}

type NullableBTMIndividualQueryWithOccurrence811AllOf struct {
	value *BTMIndividualQueryWithOccurrence811AllOf
	isSet bool
}

func (v NullableBTMIndividualQueryWithOccurrence811AllOf) Get() *BTMIndividualQueryWithOccurrence811AllOf {
	return v.value
}

func (v *NullableBTMIndividualQueryWithOccurrence811AllOf) Set(val *BTMIndividualQueryWithOccurrence811AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualQueryWithOccurrence811AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualQueryWithOccurrence811AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualQueryWithOccurrence811AllOf(val *BTMIndividualQueryWithOccurrence811AllOf) *NullableBTMIndividualQueryWithOccurrence811AllOf {
	return &NullableBTMIndividualQueryWithOccurrence811AllOf{value: val, isSet: true}
}

func (v NullableBTMIndividualQueryWithOccurrence811AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualQueryWithOccurrence811AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
