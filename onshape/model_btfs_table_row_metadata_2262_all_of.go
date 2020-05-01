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

// BTFSTableRowMetadata2262AllOf struct for BTFSTableRowMetadata2262AllOf
type BTFSTableRowMetadata2262AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Callout *string `json:"callout,omitempty"`
	CrossHighlightData *BTTableBaseCrossHighlightData2609 `json:"crossHighlightData,omitempty"`
}

// NewBTFSTableRowMetadata2262AllOf instantiates a new BTFSTableRowMetadata2262AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSTableRowMetadata2262AllOf() *BTFSTableRowMetadata2262AllOf {
	this := BTFSTableRowMetadata2262AllOf{}
	return &this
}

// NewBTFSTableRowMetadata2262AllOfWithDefaults instantiates a new BTFSTableRowMetadata2262AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSTableRowMetadata2262AllOfWithDefaults() *BTFSTableRowMetadata2262AllOf {
	this := BTFSTableRowMetadata2262AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSTableRowMetadata2262AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableRowMetadata2262AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSTableRowMetadata2262AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSTableRowMetadata2262AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetCallout returns the Callout field value if set, zero value otherwise.
func (o *BTFSTableRowMetadata2262AllOf) GetCallout() string {
	if o == nil || o.Callout == nil {
		var ret string
		return ret
	}
	return *o.Callout
}

// GetCalloutOk returns a tuple with the Callout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableRowMetadata2262AllOf) GetCalloutOk() (*string, bool) {
	if o == nil || o.Callout == nil {
		return nil, false
	}
	return o.Callout, true
}

// HasCallout returns a boolean if a field has been set.
func (o *BTFSTableRowMetadata2262AllOf) HasCallout() bool {
	if o != nil && o.Callout != nil {
		return true
	}

	return false
}

// SetCallout gets a reference to the given string and assigns it to the Callout field.
func (o *BTFSTableRowMetadata2262AllOf) SetCallout(v string) {
	o.Callout = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTFSTableRowMetadata2262AllOf) GetCrossHighlightData() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableRowMetadata2262AllOf) GetCrossHighlightDataOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTFSTableRowMetadata2262AllOf) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightData field.
func (o *BTFSTableRowMetadata2262AllOf) SetCrossHighlightData(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightData = &v
}

func (o BTFSTableRowMetadata2262AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Callout != nil {
		toSerialize["callout"] = o.Callout
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSTableRowMetadata2262AllOf struct {
	value *BTFSTableRowMetadata2262AllOf
	isSet bool
}

func (v NullableBTFSTableRowMetadata2262AllOf) Get() *BTFSTableRowMetadata2262AllOf {
	return v.value
}

func (v *NullableBTFSTableRowMetadata2262AllOf) Set(val *BTFSTableRowMetadata2262AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSTableRowMetadata2262AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSTableRowMetadata2262AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSTableRowMetadata2262AllOf(val *BTFSTableRowMetadata2262AllOf) *NullableBTFSTableRowMetadata2262AllOf {
	return &NullableBTFSTableRowMetadata2262AllOf{value: val, isSet: true}
}

func (v NullableBTFSTableRowMetadata2262AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSTableRowMetadata2262AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}