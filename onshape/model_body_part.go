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

// BodyPart struct for BodyPart
type BodyPart struct {
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	Entity *map[string]interface{} `json:"entity,omitempty"`
	Headers *map[string][]string `json:"headers,omitempty"`
	MediaType *BodyPartMediaType `json:"mediaType,omitempty"`
	MessageBodyWorkers *map[string]interface{} `json:"messageBodyWorkers,omitempty"`
	ParameterizedHeaders *map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
	Parent *MultiPart `json:"parent,omitempty"`
	Providers *map[string]interface{} `json:"providers,omitempty"`
}

// NewBodyPart instantiates a new BodyPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBodyPart() *BodyPart {
	this := BodyPart{}
	return &this
}

// NewBodyPartWithDefaults instantiates a new BodyPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodyPartWithDefaults() *BodyPart {
	this := BodyPart{}
	return &this
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *BodyPart) GetContentDisposition() ContentDisposition {
	if o == nil || o.ContentDisposition == nil {
		var ret ContentDisposition
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetContentDispositionOk() (*ContentDisposition, bool) {
	if o == nil || o.ContentDisposition == nil {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *BodyPart) HasContentDisposition() bool {
	if o != nil && o.ContentDisposition != nil {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given ContentDisposition and assigns it to the ContentDisposition field.
func (o *BodyPart) SetContentDisposition(v ContentDisposition) {
	o.ContentDisposition = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *BodyPart) GetEntity() map[string]interface{} {
	if o == nil || o.Entity == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetEntityOk() (*map[string]interface{}, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *BodyPart) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *BodyPart) SetEntity(v map[string]interface{}) {
	o.Entity = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *BodyPart) GetHeaders() map[string][]string {
	if o == nil || o.Headers == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetHeadersOk() (*map[string][]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BodyPart) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string][]string and assigns it to the Headers field.
func (o *BodyPart) SetHeaders(v map[string][]string) {
	o.Headers = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *BodyPart) GetMediaType() BodyPartMediaType {
	if o == nil || o.MediaType == nil {
		var ret BodyPartMediaType
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetMediaTypeOk() (*BodyPartMediaType, bool) {
	if o == nil || o.MediaType == nil {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *BodyPart) HasMediaType() bool {
	if o != nil && o.MediaType != nil {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given BodyPartMediaType and assigns it to the MediaType field.
func (o *BodyPart) SetMediaType(v BodyPartMediaType) {
	o.MediaType = &v
}

// GetMessageBodyWorkers returns the MessageBodyWorkers field value if set, zero value otherwise.
func (o *BodyPart) GetMessageBodyWorkers() map[string]interface{} {
	if o == nil || o.MessageBodyWorkers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.MessageBodyWorkers
}

// GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetMessageBodyWorkersOk() (*map[string]interface{}, bool) {
	if o == nil || o.MessageBodyWorkers == nil {
		return nil, false
	}
	return o.MessageBodyWorkers, true
}

// HasMessageBodyWorkers returns a boolean if a field has been set.
func (o *BodyPart) HasMessageBodyWorkers() bool {
	if o != nil && o.MessageBodyWorkers != nil {
		return true
	}

	return false
}

// SetMessageBodyWorkers gets a reference to the given map[string]interface{} and assigns it to the MessageBodyWorkers field.
func (o *BodyPart) SetMessageBodyWorkers(v map[string]interface{}) {
	o.MessageBodyWorkers = &v
}

// GetParameterizedHeaders returns the ParameterizedHeaders field value if set, zero value otherwise.
func (o *BodyPart) GetParameterizedHeaders() map[string][]ParameterizedHeader {
	if o == nil || o.ParameterizedHeaders == nil {
		var ret map[string][]ParameterizedHeader
		return ret
	}
	return *o.ParameterizedHeaders
}

// GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool) {
	if o == nil || o.ParameterizedHeaders == nil {
		return nil, false
	}
	return o.ParameterizedHeaders, true
}

// HasParameterizedHeaders returns a boolean if a field has been set.
func (o *BodyPart) HasParameterizedHeaders() bool {
	if o != nil && o.ParameterizedHeaders != nil {
		return true
	}

	return false
}

// SetParameterizedHeaders gets a reference to the given map[string][]ParameterizedHeader and assigns it to the ParameterizedHeaders field.
func (o *BodyPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader) {
	o.ParameterizedHeaders = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *BodyPart) GetParent() MultiPart {
	if o == nil || o.Parent == nil {
		var ret MultiPart
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetParentOk() (*MultiPart, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *BodyPart) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given MultiPart and assigns it to the Parent field.
func (o *BodyPart) SetParent(v MultiPart) {
	o.Parent = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *BodyPart) GetProviders() map[string]interface{} {
	if o == nil || o.Providers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetProvidersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *BodyPart) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given map[string]interface{} and assigns it to the Providers field.
func (o *BodyPart) SetProviders(v map[string]interface{}) {
	o.Providers = &v
}

func (o BodyPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentDisposition != nil {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.MediaType != nil {
		toSerialize["mediaType"] = o.MediaType
	}
	if o.MessageBodyWorkers != nil {
		toSerialize["messageBodyWorkers"] = o.MessageBodyWorkers
	}
	if o.ParameterizedHeaders != nil {
		toSerialize["parameterizedHeaders"] = o.ParameterizedHeaders
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	return json.Marshal(toSerialize)
}

type NullableBodyPart struct {
	value *BodyPart
	isSet bool
}

func (v NullableBodyPart) Get() *BodyPart {
	return v.value
}

func (v *NullableBodyPart) Set(val *BodyPart) {
	v.value = val
	v.isSet = true
}

func (v NullableBodyPart) IsSet() bool {
	return v.isSet
}

func (v *NullableBodyPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodyPart(val *BodyPart) *NullableBodyPart {
	return &NullableBodyPart{value: val, isSet: true}
}

func (v NullableBodyPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodyPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
