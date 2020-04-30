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

// BTParameterSpecReferencePartStudio1256 struct for BTParameterSpecReferencePartStudio1256
type BTParameterSpecReferencePartStudio1256 struct {
	BTParameterSpecReference2789
	AllowedInsertableTypes *[]string `json:"allowedInsertableTypes,omitempty"`
	BtType *string `json:"btType,omitempty"`
	ComputedConfigurationInputs *[]BTComputedConfigurationInputSpec2525 `json:"computedConfigurationInputs,omitempty"`
	MaxNumberOfPicks *int32 `json:"maxNumberOfPicks,omitempty"`
}

// NewBTParameterSpecReferencePartStudio1256 instantiates a new BTParameterSpecReferencePartStudio1256 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecReferencePartStudio1256() *BTParameterSpecReferencePartStudio1256 {
	this := BTParameterSpecReferencePartStudio1256{}
	return &this
}

// NewBTParameterSpecReferencePartStudio1256WithDefaults instantiates a new BTParameterSpecReferencePartStudio1256 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecReferencePartStudio1256WithDefaults() *BTParameterSpecReferencePartStudio1256 {
	this := BTParameterSpecReferencePartStudio1256{}
	return &this
}

// GetAllowedInsertableTypes returns the AllowedInsertableTypes field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetAllowedInsertableTypes() []string {
	if o == nil || o.AllowedInsertableTypes == nil {
		var ret []string
		return ret
	}
	return *o.AllowedInsertableTypes
}

// GetAllowedInsertableTypesOk returns a tuple with the AllowedInsertableTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetAllowedInsertableTypesOk() (*[]string, bool) {
	if o == nil || o.AllowedInsertableTypes == nil {
		return nil, false
	}
	return o.AllowedInsertableTypes, true
}

// HasAllowedInsertableTypes returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasAllowedInsertableTypes() bool {
	if o != nil && o.AllowedInsertableTypes != nil {
		return true
	}

	return false
}

// SetAllowedInsertableTypes gets a reference to the given []string and assigns it to the AllowedInsertableTypes field.
func (o *BTParameterSpecReferencePartStudio1256) SetAllowedInsertableTypes(v []string) {
	o.AllowedInsertableTypes = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecReferencePartStudio1256) SetBtType(v string) {
	o.BtType = &v
}

// GetComputedConfigurationInputs returns the ComputedConfigurationInputs field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetComputedConfigurationInputs() []BTComputedConfigurationInputSpec2525 {
	if o == nil || o.ComputedConfigurationInputs == nil {
		var ret []BTComputedConfigurationInputSpec2525
		return ret
	}
	return *o.ComputedConfigurationInputs
}

// GetComputedConfigurationInputsOk returns a tuple with the ComputedConfigurationInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetComputedConfigurationInputsOk() (*[]BTComputedConfigurationInputSpec2525, bool) {
	if o == nil || o.ComputedConfigurationInputs == nil {
		return nil, false
	}
	return o.ComputedConfigurationInputs, true
}

// HasComputedConfigurationInputs returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasComputedConfigurationInputs() bool {
	if o != nil && o.ComputedConfigurationInputs != nil {
		return true
	}

	return false
}

// SetComputedConfigurationInputs gets a reference to the given []BTComputedConfigurationInputSpec2525 and assigns it to the ComputedConfigurationInputs field.
func (o *BTParameterSpecReferencePartStudio1256) SetComputedConfigurationInputs(v []BTComputedConfigurationInputSpec2525) {
	o.ComputedConfigurationInputs = &v
}

// GetMaxNumberOfPicks returns the MaxNumberOfPicks field value if set, zero value otherwise.
func (o *BTParameterSpecReferencePartStudio1256) GetMaxNumberOfPicks() int32 {
	if o == nil || o.MaxNumberOfPicks == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfPicks
}

// GetMaxNumberOfPicksOk returns a tuple with the MaxNumberOfPicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferencePartStudio1256) GetMaxNumberOfPicksOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfPicks == nil {
		return nil, false
	}
	return o.MaxNumberOfPicks, true
}

// HasMaxNumberOfPicks returns a boolean if a field has been set.
func (o *BTParameterSpecReferencePartStudio1256) HasMaxNumberOfPicks() bool {
	if o != nil && o.MaxNumberOfPicks != nil {
		return true
	}

	return false
}

// SetMaxNumberOfPicks gets a reference to the given int32 and assigns it to the MaxNumberOfPicks field.
func (o *BTParameterSpecReferencePartStudio1256) SetMaxNumberOfPicks(v int32) {
	o.MaxNumberOfPicks = &v
}

func (o BTParameterSpecReferencePartStudio1256) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterSpecReference2789, errBTParameterSpecReference2789 := json.Marshal(o.BTParameterSpecReference2789)
	if errBTParameterSpecReference2789 != nil {
		return []byte{}, errBTParameterSpecReference2789
	}
	errBTParameterSpecReference2789 = json.Unmarshal([]byte(serializedBTParameterSpecReference2789), &toSerialize)
	if errBTParameterSpecReference2789 != nil {
		return []byte{}, errBTParameterSpecReference2789
	}
	if o.AllowedInsertableTypes != nil {
		toSerialize["allowedInsertableTypes"] = o.AllowedInsertableTypes
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ComputedConfigurationInputs != nil {
		toSerialize["computedConfigurationInputs"] = o.ComputedConfigurationInputs
	}
	if o.MaxNumberOfPicks != nil {
		toSerialize["maxNumberOfPicks"] = o.MaxNumberOfPicks
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecReferencePartStudio1256 struct {
	value *BTParameterSpecReferencePartStudio1256
	isSet bool
}

func (v NullableBTParameterSpecReferencePartStudio1256) Get() *BTParameterSpecReferencePartStudio1256 {
	return v.value
}

func (v *NullableBTParameterSpecReferencePartStudio1256) Set(val *BTParameterSpecReferencePartStudio1256) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecReferencePartStudio1256) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecReferencePartStudio1256) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecReferencePartStudio1256(val *BTParameterSpecReferencePartStudio1256) *NullableBTParameterSpecReferencePartStudio1256 {
	return &NullableBTParameterSpecReferencePartStudio1256{value: val, isSet: true}
}

func (v NullableBTParameterSpecReferencePartStudio1256) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecReferencePartStudio1256) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
