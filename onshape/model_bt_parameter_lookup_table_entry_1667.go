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

// BTParameterLookupTableEntry1667 struct for BTParameterLookupTableEntry1667
type BTParameterLookupTableEntry1667 struct {
	AdditionalLocalizedStrings *int32 `json:"additionalLocalizedStrings,omitempty"`
	BtType *string `json:"btType,omitempty"`
	Label *string `json:"label,omitempty"`
	LocalizableName *string `json:"localizableName,omitempty"`
	LocalizedLabel *string `json:"localizedLabel,omitempty"`
	LocalizedName *string `json:"localizedName,omitempty"`
	StringsToLocalize *[]string `json:"stringsToLocalize,omitempty"`
}

// NewBTParameterLookupTableEntry1667 instantiates a new BTParameterLookupTableEntry1667 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterLookupTableEntry1667() *BTParameterLookupTableEntry1667 {
	this := BTParameterLookupTableEntry1667{}
	return &this
}

// NewBTParameterLookupTableEntry1667WithDefaults instantiates a new BTParameterLookupTableEntry1667 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterLookupTableEntry1667WithDefaults() *BTParameterLookupTableEntry1667 {
	this := BTParameterLookupTableEntry1667{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterLookupTableEntry1667) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterLookupTableEntry1667) SetBtType(v string) {
	o.BtType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BTParameterLookupTableEntry1667) SetLabel(v string) {
	o.Label = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterLookupTableEntry1667) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedLabel returns the LocalizedLabel field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLocalizedLabel() string {
	if o == nil || o.LocalizedLabel == nil {
		var ret string
		return ret
	}
	return *o.LocalizedLabel
}

// GetLocalizedLabelOk returns a tuple with the LocalizedLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLocalizedLabelOk() (*string, bool) {
	if o == nil || o.LocalizedLabel == nil {
		return nil, false
	}
	return o.LocalizedLabel, true
}

// HasLocalizedLabel returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLocalizedLabel() bool {
	if o != nil && o.LocalizedLabel != nil {
		return true
	}

	return false
}

// SetLocalizedLabel gets a reference to the given string and assigns it to the LocalizedLabel field.
func (o *BTParameterLookupTableEntry1667) SetLocalizedLabel(v string) {
	o.LocalizedLabel = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterLookupTableEntry1667) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return *o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetStringsToLocalizeOk() (*[]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterLookupTableEntry1667) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = &v
}

func (o BTParameterLookupTableEntry1667) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LocalizableName != nil {
		toSerialize["localizableName"] = o.LocalizableName
	}
	if o.LocalizedLabel != nil {
		toSerialize["localizedLabel"] = o.LocalizedLabel
	}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.StringsToLocalize != nil {
		toSerialize["stringsToLocalize"] = o.StringsToLocalize
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterLookupTableEntry1667 struct {
	value *BTParameterLookupTableEntry1667
	isSet bool
}

func (v NullableBTParameterLookupTableEntry1667) Get() *BTParameterLookupTableEntry1667 {
	return v.value
}

func (v *NullableBTParameterLookupTableEntry1667) Set(val *BTParameterLookupTableEntry1667) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterLookupTableEntry1667) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterLookupTableEntry1667) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterLookupTableEntry1667(val *BTParameterLookupTableEntry1667) *NullableBTParameterLookupTableEntry1667 {
	return &NullableBTParameterLookupTableEntry1667{value: val, isSet: true}
}

func (v NullableBTParameterLookupTableEntry1667) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterLookupTableEntry1667) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}