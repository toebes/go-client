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

// BTWorkspacePartParams struct for BTWorkspacePartParams
type BTWorkspacePartParams struct {
	Appearance *BTPartAppearanceParams `json:"appearance,omitempty"`
	ApplyUpdateToAllConfigurations *bool `json:"applyUpdateToAllConfigurations,omitempty"`
	Configuration *string `json:"configuration,omitempty"`
	ConnectionId *string `json:"connectionId,omitempty"`
	CustomProperties *[]BTNameValuePair `json:"customProperties,omitempty"`
	CustomPropertyDefinitions *[]BTCustomPropertyDefinitionParams `json:"customPropertyDefinitions,omitempty"`
	Description *string `json:"description,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	Material *BTMaterialParams `json:"material,omitempty"`
	Name *string `json:"name,omitempty"`
	PartId *string `json:"partId,omitempty"`
	PartNumber *string `json:"partNumber,omitempty"`
	ProductLine *string `json:"productLine,omitempty"`
	Project *string `json:"project,omitempty"`
	Revision *string `json:"revision,omitempty"`
	Title1 *string `json:"title1,omitempty"`
	Title2 *string `json:"title2,omitempty"`
	Title3 *string `json:"title3,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewBTWorkspacePartParams instantiates a new BTWorkspacePartParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkspacePartParams() *BTWorkspacePartParams {
	this := BTWorkspacePartParams{}
	return &this
}

// NewBTWorkspacePartParamsWithDefaults instantiates a new BTWorkspacePartParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkspacePartParamsWithDefaults() *BTWorkspacePartParams {
	this := BTWorkspacePartParams{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetAppearance() BTPartAppearanceParams {
	if o == nil || o.Appearance == nil {
		var ret BTPartAppearanceParams
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetAppearanceOk() (*BTPartAppearanceParams, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTPartAppearanceParams and assigns it to the Appearance field.
func (o *BTWorkspacePartParams) SetAppearance(v BTPartAppearanceParams) {
	o.Appearance = &v
}

// GetApplyUpdateToAllConfigurations returns the ApplyUpdateToAllConfigurations field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetApplyUpdateToAllConfigurations() bool {
	if o == nil || o.ApplyUpdateToAllConfigurations == nil {
		var ret bool
		return ret
	}
	return *o.ApplyUpdateToAllConfigurations
}

// GetApplyUpdateToAllConfigurationsOk returns a tuple with the ApplyUpdateToAllConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetApplyUpdateToAllConfigurationsOk() (*bool, bool) {
	if o == nil || o.ApplyUpdateToAllConfigurations == nil {
		return nil, false
	}
	return o.ApplyUpdateToAllConfigurations, true
}

// HasApplyUpdateToAllConfigurations returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasApplyUpdateToAllConfigurations() bool {
	if o != nil && o.ApplyUpdateToAllConfigurations != nil {
		return true
	}

	return false
}

// SetApplyUpdateToAllConfigurations gets a reference to the given bool and assigns it to the ApplyUpdateToAllConfigurations field.
func (o *BTWorkspacePartParams) SetApplyUpdateToAllConfigurations(v bool) {
	o.ApplyUpdateToAllConfigurations = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTWorkspacePartParams) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *BTWorkspacePartParams) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetCustomProperties() []BTNameValuePair {
	if o == nil || o.CustomProperties == nil {
		var ret []BTNameValuePair
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetCustomPropertiesOk() (*[]BTNameValuePair, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []BTNameValuePair and assigns it to the CustomProperties field.
func (o *BTWorkspacePartParams) SetCustomProperties(v []BTNameValuePair) {
	o.CustomProperties = &v
}

// GetCustomPropertyDefinitions returns the CustomPropertyDefinitions field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetCustomPropertyDefinitions() []BTCustomPropertyDefinitionParams {
	if o == nil || o.CustomPropertyDefinitions == nil {
		var ret []BTCustomPropertyDefinitionParams
		return ret
	}
	return *o.CustomPropertyDefinitions
}

// GetCustomPropertyDefinitionsOk returns a tuple with the CustomPropertyDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetCustomPropertyDefinitionsOk() (*[]BTCustomPropertyDefinitionParams, bool) {
	if o == nil || o.CustomPropertyDefinitions == nil {
		return nil, false
	}
	return o.CustomPropertyDefinitions, true
}

// HasCustomPropertyDefinitions returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasCustomPropertyDefinitions() bool {
	if o != nil && o.CustomPropertyDefinitions != nil {
		return true
	}

	return false
}

// SetCustomPropertyDefinitions gets a reference to the given []BTCustomPropertyDefinitionParams and assigns it to the CustomPropertyDefinitions field.
func (o *BTWorkspacePartParams) SetCustomPropertyDefinitions(v []BTCustomPropertyDefinitionParams) {
	o.CustomPropertyDefinitions = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTWorkspacePartParams) SetDescription(v string) {
	o.Description = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTWorkspacePartParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetMaterial returns the Material field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetMaterial() BTMaterialParams {
	if o == nil || o.Material == nil {
		var ret BTMaterialParams
		return ret
	}
	return *o.Material
}

// GetMaterialOk returns a tuple with the Material field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetMaterialOk() (*BTMaterialParams, bool) {
	if o == nil || o.Material == nil {
		return nil, false
	}
	return o.Material, true
}

// HasMaterial returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasMaterial() bool {
	if o != nil && o.Material != nil {
		return true
	}

	return false
}

// SetMaterial gets a reference to the given BTMaterialParams and assigns it to the Material field.
func (o *BTWorkspacePartParams) SetMaterial(v BTMaterialParams) {
	o.Material = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWorkspacePartParams) SetName(v string) {
	o.Name = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTWorkspacePartParams) SetPartId(v string) {
	o.PartId = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTWorkspacePartParams) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProductLine returns the ProductLine field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetProductLine() string {
	if o == nil || o.ProductLine == nil {
		var ret string
		return ret
	}
	return *o.ProductLine
}

// GetProductLineOk returns a tuple with the ProductLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetProductLineOk() (*string, bool) {
	if o == nil || o.ProductLine == nil {
		return nil, false
	}
	return o.ProductLine, true
}

// HasProductLine returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasProductLine() bool {
	if o != nil && o.ProductLine != nil {
		return true
	}

	return false
}

// SetProductLine gets a reference to the given string and assigns it to the ProductLine field.
func (o *BTWorkspacePartParams) SetProductLine(v string) {
	o.ProductLine = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *BTWorkspacePartParams) SetProject(v string) {
	o.Project = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTWorkspacePartParams) SetRevision(v string) {
	o.Revision = &v
}

// GetTitle1 returns the Title1 field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetTitle1() string {
	if o == nil || o.Title1 == nil {
		var ret string
		return ret
	}
	return *o.Title1
}

// GetTitle1Ok returns a tuple with the Title1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetTitle1Ok() (*string, bool) {
	if o == nil || o.Title1 == nil {
		return nil, false
	}
	return o.Title1, true
}

// HasTitle1 returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasTitle1() bool {
	if o != nil && o.Title1 != nil {
		return true
	}

	return false
}

// SetTitle1 gets a reference to the given string and assigns it to the Title1 field.
func (o *BTWorkspacePartParams) SetTitle1(v string) {
	o.Title1 = &v
}

// GetTitle2 returns the Title2 field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetTitle2() string {
	if o == nil || o.Title2 == nil {
		var ret string
		return ret
	}
	return *o.Title2
}

// GetTitle2Ok returns a tuple with the Title2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetTitle2Ok() (*string, bool) {
	if o == nil || o.Title2 == nil {
		return nil, false
	}
	return o.Title2, true
}

// HasTitle2 returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasTitle2() bool {
	if o != nil && o.Title2 != nil {
		return true
	}

	return false
}

// SetTitle2 gets a reference to the given string and assigns it to the Title2 field.
func (o *BTWorkspacePartParams) SetTitle2(v string) {
	o.Title2 = &v
}

// GetTitle3 returns the Title3 field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetTitle3() string {
	if o == nil || o.Title3 == nil {
		var ret string
		return ret
	}
	return *o.Title3
}

// GetTitle3Ok returns a tuple with the Title3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetTitle3Ok() (*string, bool) {
	if o == nil || o.Title3 == nil {
		return nil, false
	}
	return o.Title3, true
}

// HasTitle3 returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasTitle3() bool {
	if o != nil && o.Title3 != nil {
		return true
	}

	return false
}

// SetTitle3 gets a reference to the given string and assigns it to the Title3 field.
func (o *BTWorkspacePartParams) SetTitle3(v string) {
	o.Title3 = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *BTWorkspacePartParams) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspacePartParams) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *BTWorkspacePartParams) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *BTWorkspacePartParams) SetVendor(v string) {
	o.Vendor = &v
}

func (o BTWorkspacePartParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.ApplyUpdateToAllConfigurations != nil {
		toSerialize["applyUpdateToAllConfigurations"] = o.ApplyUpdateToAllConfigurations
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ConnectionId != nil {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if o.CustomPropertyDefinitions != nil {
		toSerialize["customPropertyDefinitions"] = o.CustomPropertyDefinitions
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Material != nil {
		toSerialize["material"] = o.Material
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.ProductLine != nil {
		toSerialize["productLine"] = o.ProductLine
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Title1 != nil {
		toSerialize["title1"] = o.Title1
	}
	if o.Title2 != nil {
		toSerialize["title2"] = o.Title2
	}
	if o.Title3 != nil {
		toSerialize["title3"] = o.Title3
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkspacePartParams struct {
	value *BTWorkspacePartParams
	isSet bool
}

func (v NullableBTWorkspacePartParams) Get() *BTWorkspacePartParams {
	return v.value
}

func (v *NullableBTWorkspacePartParams) Set(val *BTWorkspacePartParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkspacePartParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkspacePartParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkspacePartParams(val *BTWorkspacePartParams) *NullableBTWorkspacePartParams {
	return &NullableBTWorkspacePartParams{value: val, isSet: true}
}

func (v NullableBTWorkspacePartParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkspacePartParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
