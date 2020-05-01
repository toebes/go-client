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

// BTMParameterQueryWithOccurrenceList67 struct for BTMParameterQueryWithOccurrenceList67
type BTMParameterQueryWithOccurrenceList67 struct {
	BTMParameter1
	BtType *string `json:"btType,omitempty"`
	Occurrences *[]BTOccurrence74 `json:"occurrences,omitempty"`
	Queries *[]BTMIndividualQueryWithOccurrenceBase904 `json:"queries,omitempty"`
}

// NewBTMParameterQueryWithOccurrenceList67 instantiates a new BTMParameterQueryWithOccurrenceList67 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterQueryWithOccurrenceList67() *BTMParameterQueryWithOccurrenceList67 {
	this := BTMParameterQueryWithOccurrenceList67{}
	return &this
}

// NewBTMParameterQueryWithOccurrenceList67WithDefaults instantiates a new BTMParameterQueryWithOccurrenceList67 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterQueryWithOccurrenceList67WithDefaults() *BTMParameterQueryWithOccurrenceList67 {
	this := BTMParameterQueryWithOccurrenceList67{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterQueryWithOccurrenceList67) SetBtType(v string) {
	o.BtType = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetOccurrences() []BTOccurrence74 {
	if o == nil || o.Occurrences == nil {
		var ret []BTOccurrence74
		return ret
	}
	return *o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetOccurrencesOk() (*[]BTOccurrence74, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTOccurrence74 and assigns it to the Occurrences field.
func (o *BTMParameterQueryWithOccurrenceList67) SetOccurrences(v []BTOccurrence74) {
	o.Occurrences = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *BTMParameterQueryWithOccurrenceList67) GetQueries() []BTMIndividualQueryWithOccurrenceBase904 {
	if o == nil || o.Queries == nil {
		var ret []BTMIndividualQueryWithOccurrenceBase904
		return ret
	}
	return *o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryWithOccurrenceList67) GetQueriesOk() (*[]BTMIndividualQueryWithOccurrenceBase904, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *BTMParameterQueryWithOccurrenceList67) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []BTMIndividualQueryWithOccurrenceBase904 and assigns it to the Queries field.
func (o *BTMParameterQueryWithOccurrenceList67) SetQueries(v []BTMIndividualQueryWithOccurrenceBase904) {
	o.Queries = &v
}

func (o BTMParameterQueryWithOccurrenceList67) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMParameter1, errBTMParameter1 := json.Marshal(o.BTMParameter1)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	errBTMParameter1 = json.Unmarshal([]byte(serializedBTMParameter1), &toSerialize)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterQueryWithOccurrenceList67 struct {
	value *BTMParameterQueryWithOccurrenceList67
	isSet bool
}

func (v NullableBTMParameterQueryWithOccurrenceList67) Get() *BTMParameterQueryWithOccurrenceList67 {
	return v.value
}

func (v *NullableBTMParameterQueryWithOccurrenceList67) Set(val *BTMParameterQueryWithOccurrenceList67) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterQueryWithOccurrenceList67) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterQueryWithOccurrenceList67) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterQueryWithOccurrenceList67(val *BTMParameterQueryWithOccurrenceList67) *NullableBTMParameterQueryWithOccurrenceList67 {
	return &NullableBTMParameterQueryWithOccurrenceList67{value: val, isSet: true}
}

func (v NullableBTMParameterQueryWithOccurrenceList67) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterQueryWithOccurrenceList67) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
