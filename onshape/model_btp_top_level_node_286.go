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

// BTPTopLevelNode286 struct for BTPTopLevelNode286
type BTPTopLevelNode286 struct {
	BTPNode7
	BtType *string `json:"btType,omitempty"`
	Annotation *BTPAnnotation231 `json:"annotation,omitempty"`
	ArgumentsToDocument *[]BTPArgumentDeclaration232 `json:"argumentsToDocument,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	DeprecatedExplanation *string `json:"deprecatedExplanation,omitempty"`
	ForExport *bool `json:"forExport,omitempty"`
	SpaceAfterExport *BTPSpace10 `json:"spaceAfterExport,omitempty"`
	SymbolName *BTPIdentifier8 `json:"symbolName,omitempty"`
}

// NewBTPTopLevelNode286 instantiates a new BTPTopLevelNode286 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPTopLevelNode286() *BTPTopLevelNode286 {
	this := BTPTopLevelNode286{}
	return &this
}

// NewBTPTopLevelNode286WithDefaults instantiates a new BTPTopLevelNode286 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPTopLevelNode286WithDefaults() *BTPTopLevelNode286 {
	this := BTPTopLevelNode286{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPTopLevelNode286) SetBtType(v string) {
	o.BtType = &v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPTopLevelNode286) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetArgumentsToDocument returns the ArgumentsToDocument field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetArgumentsToDocument() []BTPArgumentDeclaration232 {
	if o == nil || o.ArgumentsToDocument == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return *o.ArgumentsToDocument
}

// GetArgumentsToDocumentOk returns a tuple with the ArgumentsToDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetArgumentsToDocumentOk() (*[]BTPArgumentDeclaration232, bool) {
	if o == nil || o.ArgumentsToDocument == nil {
		return nil, false
	}
	return o.ArgumentsToDocument, true
}

// HasArgumentsToDocument returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasArgumentsToDocument() bool {
	if o != nil && o.ArgumentsToDocument != nil {
		return true
	}

	return false
}

// SetArgumentsToDocument gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the ArgumentsToDocument field.
func (o *BTPTopLevelNode286) SetArgumentsToDocument(v []BTPArgumentDeclaration232) {
	o.ArgumentsToDocument = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *BTPTopLevelNode286) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDeprecatedExplanation returns the DeprecatedExplanation field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetDeprecatedExplanation() string {
	if o == nil || o.DeprecatedExplanation == nil {
		var ret string
		return ret
	}
	return *o.DeprecatedExplanation
}

// GetDeprecatedExplanationOk returns a tuple with the DeprecatedExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetDeprecatedExplanationOk() (*string, bool) {
	if o == nil || o.DeprecatedExplanation == nil {
		return nil, false
	}
	return o.DeprecatedExplanation, true
}

// HasDeprecatedExplanation returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasDeprecatedExplanation() bool {
	if o != nil && o.DeprecatedExplanation != nil {
		return true
	}

	return false
}

// SetDeprecatedExplanation gets a reference to the given string and assigns it to the DeprecatedExplanation field.
func (o *BTPTopLevelNode286) SetDeprecatedExplanation(v string) {
	o.DeprecatedExplanation = &v
}

// GetForExport returns the ForExport field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetForExport() bool {
	if o == nil || o.ForExport == nil {
		var ret bool
		return ret
	}
	return *o.ForExport
}

// GetForExportOk returns a tuple with the ForExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetForExportOk() (*bool, bool) {
	if o == nil || o.ForExport == nil {
		return nil, false
	}
	return o.ForExport, true
}

// HasForExport returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasForExport() bool {
	if o != nil && o.ForExport != nil {
		return true
	}

	return false
}

// SetForExport gets a reference to the given bool and assigns it to the ForExport field.
func (o *BTPTopLevelNode286) SetForExport(v bool) {
	o.ForExport = &v
}

// GetSpaceAfterExport returns the SpaceAfterExport field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetSpaceAfterExport() BTPSpace10 {
	if o == nil || o.SpaceAfterExport == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterExport
}

// GetSpaceAfterExportOk returns a tuple with the SpaceAfterExport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetSpaceAfterExportOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterExport == nil {
		return nil, false
	}
	return o.SpaceAfterExport, true
}

// HasSpaceAfterExport returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasSpaceAfterExport() bool {
	if o != nil && o.SpaceAfterExport != nil {
		return true
	}

	return false
}

// SetSpaceAfterExport gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterExport field.
func (o *BTPTopLevelNode286) SetSpaceAfterExport(v BTPSpace10) {
	o.SpaceAfterExport = &v
}

// GetSymbolName returns the SymbolName field value if set, zero value otherwise.
func (o *BTPTopLevelNode286) GetSymbolName() BTPIdentifier8 {
	if o == nil || o.SymbolName == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.SymbolName
}

// GetSymbolNameOk returns a tuple with the SymbolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelNode286) GetSymbolNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.SymbolName == nil {
		return nil, false
	}
	return o.SymbolName, true
}

// HasSymbolName returns a boolean if a field has been set.
func (o *BTPTopLevelNode286) HasSymbolName() bool {
	if o != nil && o.SymbolName != nil {
		return true
	}

	return false
}

// SetSymbolName gets a reference to the given BTPIdentifier8 and assigns it to the SymbolName field.
func (o *BTPTopLevelNode286) SetSymbolName(v BTPIdentifier8) {
	o.SymbolName = &v
}

func (o BTPTopLevelNode286) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPNode7, errBTPNode7 := json.Marshal(o.BTPNode7)
	if errBTPNode7 != nil {
		return []byte{}, errBTPNode7
	}
	errBTPNode7 = json.Unmarshal([]byte(serializedBTPNode7), &toSerialize)
	if errBTPNode7 != nil {
		return []byte{}, errBTPNode7
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Annotation != nil {
		toSerialize["annotation"] = o.Annotation
	}
	if o.ArgumentsToDocument != nil {
		toSerialize["argumentsToDocument"] = o.ArgumentsToDocument
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if o.DeprecatedExplanation != nil {
		toSerialize["deprecatedExplanation"] = o.DeprecatedExplanation
	}
	if o.ForExport != nil {
		toSerialize["forExport"] = o.ForExport
	}
	if o.SpaceAfterExport != nil {
		toSerialize["spaceAfterExport"] = o.SpaceAfterExport
	}
	if o.SymbolName != nil {
		toSerialize["symbolName"] = o.SymbolName
	}
	return json.Marshal(toSerialize)
}

type NullableBTPTopLevelNode286 struct {
	value *BTPTopLevelNode286
	isSet bool
}

func (v NullableBTPTopLevelNode286) Get() *BTPTopLevelNode286 {
	return v.value
}

func (v *NullableBTPTopLevelNode286) Set(val *BTPTopLevelNode286) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPTopLevelNode286) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPTopLevelNode286) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPTopLevelNode286(val *BTPTopLevelNode286) *NullableBTPTopLevelNode286 {
	return &NullableBTPTopLevelNode286{value: val, isSet: true}
}

func (v NullableBTPTopLevelNode286) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPTopLevelNode286) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
