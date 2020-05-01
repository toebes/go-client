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

// BTWorkflowableTestObjectInfo struct for BTWorkflowableTestObjectInfo
type BTWorkflowableTestObjectInfo struct {
	CompanyId *string `json:"companyId,omitempty"`
	Description *string `json:"description,omitempty"`
	DescriptionAsProperty *string `json:"descriptionAsProperty,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Info *map[string]string `json:"info,omitempty"`
	IsObsoletion *bool `json:"isObsoletion,omitempty"`
	Name *string `json:"name,omitempty"`
	NameAsProperty *string `json:"nameAsProperty,omitempty"`
	Properties *[]BTWorkflowPropertyInfo `json:"properties,omitempty"`
	ViewRef *string `json:"viewRef,omitempty"`
	Workflow *BTWorkflowSnapshotInfo `json:"workflow,omitempty"`
	WorkflowId *BTPublishedWorkflowId `json:"workflowId,omitempty"`
}

// NewBTWorkflowableTestObjectInfo instantiates a new BTWorkflowableTestObjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowableTestObjectInfo() *BTWorkflowableTestObjectInfo {
	this := BTWorkflowableTestObjectInfo{}
	return &this
}

// NewBTWorkflowableTestObjectInfoWithDefaults instantiates a new BTWorkflowableTestObjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowableTestObjectInfoWithDefaults() *BTWorkflowableTestObjectInfo {
	this := BTWorkflowableTestObjectInfo{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTWorkflowableTestObjectInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTWorkflowableTestObjectInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionAsProperty returns the DescriptionAsProperty field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetDescriptionAsProperty() string {
	if o == nil || o.DescriptionAsProperty == nil {
		var ret string
		return ret
	}
	return *o.DescriptionAsProperty
}

// GetDescriptionAsPropertyOk returns a tuple with the DescriptionAsProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetDescriptionAsPropertyOk() (*string, bool) {
	if o == nil || o.DescriptionAsProperty == nil {
		return nil, false
	}
	return o.DescriptionAsProperty, true
}

// HasDescriptionAsProperty returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasDescriptionAsProperty() bool {
	if o != nil && o.DescriptionAsProperty != nil {
		return true
	}

	return false
}

// SetDescriptionAsProperty gets a reference to the given string and assigns it to the DescriptionAsProperty field.
func (o *BTWorkflowableTestObjectInfo) SetDescriptionAsProperty(v string) {
	o.DescriptionAsProperty = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTWorkflowableTestObjectInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTWorkflowableTestObjectInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWorkflowableTestObjectInfo) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetInfo() map[string]string {
	if o == nil || o.Info == nil {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *BTWorkflowableTestObjectInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

// GetIsObsoletion returns the IsObsoletion field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetIsObsoletion() bool {
	if o == nil || o.IsObsoletion == nil {
		var ret bool
		return ret
	}
	return *o.IsObsoletion
}

// GetIsObsoletionOk returns a tuple with the IsObsoletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetIsObsoletionOk() (*bool, bool) {
	if o == nil || o.IsObsoletion == nil {
		return nil, false
	}
	return o.IsObsoletion, true
}

// HasIsObsoletion returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasIsObsoletion() bool {
	if o != nil && o.IsObsoletion != nil {
		return true
	}

	return false
}

// SetIsObsoletion gets a reference to the given bool and assigns it to the IsObsoletion field.
func (o *BTWorkflowableTestObjectInfo) SetIsObsoletion(v bool) {
	o.IsObsoletion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWorkflowableTestObjectInfo) SetName(v string) {
	o.Name = &v
}

// GetNameAsProperty returns the NameAsProperty field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetNameAsProperty() string {
	if o == nil || o.NameAsProperty == nil {
		var ret string
		return ret
	}
	return *o.NameAsProperty
}

// GetNameAsPropertyOk returns a tuple with the NameAsProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetNameAsPropertyOk() (*string, bool) {
	if o == nil || o.NameAsProperty == nil {
		return nil, false
	}
	return o.NameAsProperty, true
}

// HasNameAsProperty returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasNameAsProperty() bool {
	if o != nil && o.NameAsProperty != nil {
		return true
	}

	return false
}

// SetNameAsProperty gets a reference to the given string and assigns it to the NameAsProperty field.
func (o *BTWorkflowableTestObjectInfo) SetNameAsProperty(v string) {
	o.NameAsProperty = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetProperties() []BTWorkflowPropertyInfo {
	if o == nil || o.Properties == nil {
		var ret []BTWorkflowPropertyInfo
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetPropertiesOk() (*[]BTWorkflowPropertyInfo, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTWorkflowPropertyInfo and assigns it to the Properties field.
func (o *BTWorkflowableTestObjectInfo) SetProperties(v []BTWorkflowPropertyInfo) {
	o.Properties = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTWorkflowableTestObjectInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetWorkflow() BTWorkflowSnapshotInfo {
	if o == nil || o.Workflow == nil {
		var ret BTWorkflowSnapshotInfo
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowOk() (*BTWorkflowSnapshotInfo, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given BTWorkflowSnapshotInfo and assigns it to the Workflow field.
func (o *BTWorkflowableTestObjectInfo) SetWorkflow(v BTWorkflowSnapshotInfo) {
	o.Workflow = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowId() BTPublishedWorkflowId {
	if o == nil || o.WorkflowId == nil {
		var ret BTPublishedWorkflowId
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowableTestObjectInfo) GetWorkflowIdOk() (*BTPublishedWorkflowId, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *BTWorkflowableTestObjectInfo) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given BTPublishedWorkflowId and assigns it to the WorkflowId field.
func (o *BTWorkflowableTestObjectInfo) SetWorkflowId(v BTPublishedWorkflowId) {
	o.WorkflowId = &v
}

func (o BTWorkflowableTestObjectInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DescriptionAsProperty != nil {
		toSerialize["descriptionAsProperty"] = o.DescriptionAsProperty
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.IsObsoletion != nil {
		toSerialize["isObsoletion"] = o.IsObsoletion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameAsProperty != nil {
		toSerialize["nameAsProperty"] = o.NameAsProperty
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Workflow != nil {
		toSerialize["workflow"] = o.Workflow
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowableTestObjectInfo struct {
	value *BTWorkflowableTestObjectInfo
	isSet bool
}

func (v NullableBTWorkflowableTestObjectInfo) Get() *BTWorkflowableTestObjectInfo {
	return v.value
}

func (v *NullableBTWorkflowableTestObjectInfo) Set(val *BTWorkflowableTestObjectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowableTestObjectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowableTestObjectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowableTestObjectInfo(val *BTWorkflowableTestObjectInfo) *NullableBTWorkflowableTestObjectInfo {
	return &NullableBTWorkflowableTestObjectInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowableTestObjectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowableTestObjectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
