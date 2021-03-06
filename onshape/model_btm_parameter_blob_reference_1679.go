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

// BTMParameterBlobReference1679 struct for BTMParameterBlobReference1679
type BTMParameterBlobReference1679 struct {
	BTMParameter1
	BlobImport *BTMImport136 `json:"blobImport,omitempty"`
	BtType *string `json:"btType,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewBTMParameterBlobReference1679 instantiates a new BTMParameterBlobReference1679 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterBlobReference1679() *BTMParameterBlobReference1679 {
	this := BTMParameterBlobReference1679{}
	return &this
}

// NewBTMParameterBlobReference1679WithDefaults instantiates a new BTMParameterBlobReference1679 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterBlobReference1679WithDefaults() *BTMParameterBlobReference1679 {
	this := BTMParameterBlobReference1679{}
	return &this
}

// GetBlobImport returns the BlobImport field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetBlobImport() BTMImport136 {
	if o == nil || o.BlobImport == nil {
		var ret BTMImport136
		return ret
	}
	return *o.BlobImport
}

// GetBlobImportOk returns a tuple with the BlobImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetBlobImportOk() (*BTMImport136, bool) {
	if o == nil || o.BlobImport == nil {
		return nil, false
	}
	return o.BlobImport, true
}

// HasBlobImport returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasBlobImport() bool {
	if o != nil && o.BlobImport != nil {
		return true
	}

	return false
}

// SetBlobImport gets a reference to the given BTMImport136 and assigns it to the BlobImport field.
func (o *BTMParameterBlobReference1679) SetBlobImport(v BTMImport136) {
	o.BlobImport = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterBlobReference1679) SetBtType(v string) {
	o.BtType = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMParameterBlobReference1679) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterBlobReference1679) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMParameterBlobReference1679) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMParameterBlobReference1679) SetNamespace(v string) {
	o.Namespace = &v
}

func (o BTMParameterBlobReference1679) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMParameter1, errBTMParameter1 := json.Marshal(o.BTMParameter1)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	errBTMParameter1 = json.Unmarshal([]byte(serializedBTMParameter1), &toSerialize)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	if o.BlobImport != nil {
		toSerialize["blobImport"] = o.BlobImport
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterBlobReference1679 struct {
	value *BTMParameterBlobReference1679
	isSet bool
}

func (v NullableBTMParameterBlobReference1679) Get() *BTMParameterBlobReference1679 {
	return v.value
}

func (v *NullableBTMParameterBlobReference1679) Set(val *BTMParameterBlobReference1679) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterBlobReference1679) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterBlobReference1679) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterBlobReference1679(val *BTMParameterBlobReference1679) *NullableBTMParameterBlobReference1679 {
	return &NullableBTMParameterBlobReference1679{value: val, isSet: true}
}

func (v NullableBTMParameterBlobReference1679) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterBlobReference1679) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
