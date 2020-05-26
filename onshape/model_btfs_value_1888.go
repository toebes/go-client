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

// BTFSValue1888 struct for BTFSValue1888
type BTFSValue1888 struct {
	BtType *string `json:"btType,omitempty"`
	ConfigurationValueString *string `json:"configurationValueString,omitempty"`
	StandardTypeName *string `json:"standardTypeName,omitempty"`
	TypeTag *string `json:"typeTag,omitempty"`
	ValueObject *map[string]interface{} `json:"valueObject,omitempty"`
}

// NewBTFSValue1888 instantiates a new BTFSValue1888 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValue1888() *BTFSValue1888 {
	this := BTFSValue1888{}
	return &this
}

// NewBTFSValue1888WithDefaults instantiates a new BTFSValue1888 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValue1888WithDefaults() *BTFSValue1888 {
	this := BTFSValue1888{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSValue1888) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValue1888) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSValue1888) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSValue1888) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationValueString returns the ConfigurationValueString field value if set, zero value otherwise.
func (o *BTFSValue1888) GetConfigurationValueString() string {
	if o == nil || o.ConfigurationValueString == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationValueString
}

// GetConfigurationValueStringOk returns a tuple with the ConfigurationValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValue1888) GetConfigurationValueStringOk() (*string, bool) {
	if o == nil || o.ConfigurationValueString == nil {
		return nil, false
	}
	return o.ConfigurationValueString, true
}

// HasConfigurationValueString returns a boolean if a field has been set.
func (o *BTFSValue1888) HasConfigurationValueString() bool {
	if o != nil && o.ConfigurationValueString != nil {
		return true
	}

	return false
}

// SetConfigurationValueString gets a reference to the given string and assigns it to the ConfigurationValueString field.
func (o *BTFSValue1888) SetConfigurationValueString(v string) {
	o.ConfigurationValueString = &v
}

// GetStandardTypeName returns the StandardTypeName field value if set, zero value otherwise.
func (o *BTFSValue1888) GetStandardTypeName() string {
	if o == nil || o.StandardTypeName == nil {
		var ret string
		return ret
	}
	return *o.StandardTypeName
}

// GetStandardTypeNameOk returns a tuple with the StandardTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValue1888) GetStandardTypeNameOk() (*string, bool) {
	if o == nil || o.StandardTypeName == nil {
		return nil, false
	}
	return o.StandardTypeName, true
}

// HasStandardTypeName returns a boolean if a field has been set.
func (o *BTFSValue1888) HasStandardTypeName() bool {
	if o != nil && o.StandardTypeName != nil {
		return true
	}

	return false
}

// SetStandardTypeName gets a reference to the given string and assigns it to the StandardTypeName field.
func (o *BTFSValue1888) SetStandardTypeName(v string) {
	o.StandardTypeName = &v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValue1888) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValue1888) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValue1888) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValue1888) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetValueObject returns the ValueObject field value if set, zero value otherwise.
func (o *BTFSValue1888) GetValueObject() map[string]interface{} {
	if o == nil || o.ValueObject == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ValueObject
}

// GetValueObjectOk returns a tuple with the ValueObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValue1888) GetValueObjectOk() (*map[string]interface{}, bool) {
	if o == nil || o.ValueObject == nil {
		return nil, false
	}
	return o.ValueObject, true
}

// HasValueObject returns a boolean if a field has been set.
func (o *BTFSValue1888) HasValueObject() bool {
	if o != nil && o.ValueObject != nil {
		return true
	}

	return false
}

// SetValueObject gets a reference to the given map[string]interface{} and assigns it to the ValueObject field.
func (o *BTFSValue1888) SetValueObject(v map[string]interface{}) {
	o.ValueObject = &v
}

func (o BTFSValue1888) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationValueString != nil {
		toSerialize["configurationValueString"] = o.ConfigurationValueString
	}
	if o.StandardTypeName != nil {
		toSerialize["standardTypeName"] = o.StandardTypeName
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	if o.ValueObject != nil {
		toSerialize["valueObject"] = o.ValueObject
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValue1888 struct {
	value *BTFSValue1888
	isSet bool
}

func (v NullableBTFSValue1888) Get() *BTFSValue1888 {
	return v.value
}

func (v *NullableBTFSValue1888) Set(val *BTFSValue1888) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValue1888) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValue1888) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValue1888(val *BTFSValue1888) *NullableBTFSValue1888 {
	return &NullableBTFSValue1888{value: val, isSet: true}
}

func (v NullableBTFSValue1888) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValue1888) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}