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

// BTFullElementIdWithDocument1729 struct for BTFullElementIdWithDocument1729
type BTFullElementIdWithDocument1729 struct {
	BTFullElementId756
	BtType *string `json:"btType,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
}

// NewBTFullElementIdWithDocument1729 instantiates a new BTFullElementIdWithDocument1729 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFullElementIdWithDocument1729() *BTFullElementIdWithDocument1729 {
	this := BTFullElementIdWithDocument1729{}
	return &this
}

// NewBTFullElementIdWithDocument1729WithDefaults instantiates a new BTFullElementIdWithDocument1729 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFullElementIdWithDocument1729WithDefaults() *BTFullElementIdWithDocument1729 {
	this := BTFullElementIdWithDocument1729{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFullElementIdWithDocument1729) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTFullElementIdWithDocument1729) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdWithDocument1729) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTFullElementIdWithDocument1729) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTFullElementIdWithDocument1729) SetDocumentId(v string) {
	o.DocumentId = &v
}

func (o BTFullElementIdWithDocument1729) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTFullElementId756, errBTFullElementId756 := json.Marshal(o.BTFullElementId756)
	if errBTFullElementId756 != nil {
		return []byte{}, errBTFullElementId756
	}
	errBTFullElementId756 = json.Unmarshal([]byte(serializedBTFullElementId756), &toSerialize)
	if errBTFullElementId756 != nil {
		return []byte{}, errBTFullElementId756
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	return json.Marshal(toSerialize)
}

type NullableBTFullElementIdWithDocument1729 struct {
	value *BTFullElementIdWithDocument1729
	isSet bool
}

func (v NullableBTFullElementIdWithDocument1729) Get() *BTFullElementIdWithDocument1729 {
	return v.value
}

func (v *NullableBTFullElementIdWithDocument1729) Set(val *BTFullElementIdWithDocument1729) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFullElementIdWithDocument1729) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFullElementIdWithDocument1729) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFullElementIdWithDocument1729(val *BTFullElementIdWithDocument1729) *NullableBTFullElementIdWithDocument1729 {
	return &NullableBTFullElementIdWithDocument1729{value: val, isSet: true}
}

func (v NullableBTFullElementIdWithDocument1729) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFullElementIdWithDocument1729) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}