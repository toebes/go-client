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

// BTParameterSpecQuery174 struct for BTParameterSpecQuery174
type BTParameterSpecQuery174 struct {
	BTParameterSpec6
	AdditionalBoxSelectFilter *BTQueryFilter183 `json:"additionalBoxSelectFilter,omitempty"`
	BtType *string `json:"btType,omitempty"`
	Filter *BTQueryFilter183 `json:"filter,omitempty"`
	MaxNumberOfPicks *int32 `json:"maxNumberOfPicks,omitempty"`
}

// NewBTParameterSpecQuery174 instantiates a new BTParameterSpecQuery174 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecQuery174() *BTParameterSpecQuery174 {
	this := BTParameterSpecQuery174{}
	return &this
}

// NewBTParameterSpecQuery174WithDefaults instantiates a new BTParameterSpecQuery174 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecQuery174WithDefaults() *BTParameterSpecQuery174 {
	this := BTParameterSpecQuery174{}
	return &this
}

// GetAdditionalBoxSelectFilter returns the AdditionalBoxSelectFilter field value if set, zero value otherwise.
func (o *BTParameterSpecQuery174) GetAdditionalBoxSelectFilter() BTQueryFilter183 {
	if o == nil || o.AdditionalBoxSelectFilter == nil {
		var ret BTQueryFilter183
		return ret
	}
	return *o.AdditionalBoxSelectFilter
}

// GetAdditionalBoxSelectFilterOk returns a tuple with the AdditionalBoxSelectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecQuery174) GetAdditionalBoxSelectFilterOk() (*BTQueryFilter183, bool) {
	if o == nil || o.AdditionalBoxSelectFilter == nil {
		return nil, false
	}
	return o.AdditionalBoxSelectFilter, true
}

// HasAdditionalBoxSelectFilter returns a boolean if a field has been set.
func (o *BTParameterSpecQuery174) HasAdditionalBoxSelectFilter() bool {
	if o != nil && o.AdditionalBoxSelectFilter != nil {
		return true
	}

	return false
}

// SetAdditionalBoxSelectFilter gets a reference to the given BTQueryFilter183 and assigns it to the AdditionalBoxSelectFilter field.
func (o *BTParameterSpecQuery174) SetAdditionalBoxSelectFilter(v BTQueryFilter183) {
	o.AdditionalBoxSelectFilter = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecQuery174) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecQuery174) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecQuery174) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecQuery174) SetBtType(v string) {
	o.BtType = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *BTParameterSpecQuery174) GetFilter() BTQueryFilter183 {
	if o == nil || o.Filter == nil {
		var ret BTQueryFilter183
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecQuery174) GetFilterOk() (*BTQueryFilter183, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *BTParameterSpecQuery174) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given BTQueryFilter183 and assigns it to the Filter field.
func (o *BTParameterSpecQuery174) SetFilter(v BTQueryFilter183) {
	o.Filter = &v
}

// GetMaxNumberOfPicks returns the MaxNumberOfPicks field value if set, zero value otherwise.
func (o *BTParameterSpecQuery174) GetMaxNumberOfPicks() int32 {
	if o == nil || o.MaxNumberOfPicks == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfPicks
}

// GetMaxNumberOfPicksOk returns a tuple with the MaxNumberOfPicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecQuery174) GetMaxNumberOfPicksOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfPicks == nil {
		return nil, false
	}
	return o.MaxNumberOfPicks, true
}

// HasMaxNumberOfPicks returns a boolean if a field has been set.
func (o *BTParameterSpecQuery174) HasMaxNumberOfPicks() bool {
	if o != nil && o.MaxNumberOfPicks != nil {
		return true
	}

	return false
}

// SetMaxNumberOfPicks gets a reference to the given int32 and assigns it to the MaxNumberOfPicks field.
func (o *BTParameterSpecQuery174) SetMaxNumberOfPicks(v int32) {
	o.MaxNumberOfPicks = &v
}

func (o BTParameterSpecQuery174) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterSpec6, errBTParameterSpec6 := json.Marshal(o.BTParameterSpec6)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	errBTParameterSpec6 = json.Unmarshal([]byte(serializedBTParameterSpec6), &toSerialize)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	if o.AdditionalBoxSelectFilter != nil {
		toSerialize["additionalBoxSelectFilter"] = o.AdditionalBoxSelectFilter
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.MaxNumberOfPicks != nil {
		toSerialize["maxNumberOfPicks"] = o.MaxNumberOfPicks
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecQuery174 struct {
	value *BTParameterSpecQuery174
	isSet bool
}

func (v NullableBTParameterSpecQuery174) Get() *BTParameterSpecQuery174 {
	return v.value
}

func (v *NullableBTParameterSpecQuery174) Set(val *BTParameterSpecQuery174) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecQuery174) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecQuery174) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecQuery174(val *BTParameterSpecQuery174) *NullableBTParameterSpecQuery174 {
	return &NullableBTParameterSpecQuery174{value: val, isSet: true}
}

func (v NullableBTParameterSpecQuery174) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecQuery174) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
