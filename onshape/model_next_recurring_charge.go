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

// NextRecurringCharge struct for NextRecurringCharge
type NextRecurringCharge struct {
	Amount *int64 `json:"amount,omitempty"`
	Date *string `json:"date,omitempty"`
}

// NewNextRecurringCharge instantiates a new NextRecurringCharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextRecurringCharge() *NextRecurringCharge {
	this := NextRecurringCharge{}
	return &this
}

// NewNextRecurringChargeWithDefaults instantiates a new NextRecurringCharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextRecurringChargeWithDefaults() *NextRecurringCharge {
	this := NextRecurringCharge{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NextRecurringCharge) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextRecurringCharge) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NextRecurringCharge) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *NextRecurringCharge) SetAmount(v int64) {
	o.Amount = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *NextRecurringCharge) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextRecurringCharge) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *NextRecurringCharge) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *NextRecurringCharge) SetDate(v string) {
	o.Date = &v
}

func (o NextRecurringCharge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableNextRecurringCharge struct {
	value *NextRecurringCharge
	isSet bool
}

func (v NullableNextRecurringCharge) Get() *NextRecurringCharge {
	return v.value
}

func (v *NullableNextRecurringCharge) Set(val *NextRecurringCharge) {
	v.value = val
	v.isSet = true
}

func (v NullableNextRecurringCharge) IsSet() bool {
	return v.isSet
}

func (v *NullableNextRecurringCharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextRecurringCharge(val *NextRecurringCharge) *NullableNextRecurringCharge {
	return &NullableNextRecurringCharge{value: val, isSet: true}
}

func (v NullableNextRecurringCharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextRecurringCharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}