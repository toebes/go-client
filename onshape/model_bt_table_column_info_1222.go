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

// BTTableColumnInfo1222 struct for BTTableColumnInfo1222
type BTTableColumnInfo1222 struct {
	BtType *string `json:"btType,omitempty"`
	Id *string `json:"id,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
	Specification *BTTableColumnSpec1967 `json:"specification,omitempty"`
}

// NewBTTableColumnInfo1222 instantiates a new BTTableColumnInfo1222 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	this := BTTableColumnInfo1222{}
	return &this
}

// NewBTTableColumnInfo1222WithDefaults instantiates a new BTTableColumnInfo1222 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableColumnInfo1222WithDefaults() *BTTableColumnInfo1222 {
	this := BTTableColumnInfo1222{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableColumnInfo1222) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnInfo1222) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableColumnInfo1222) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableColumnInfo1222) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTTableColumnInfo1222) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnInfo1222) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTTableColumnInfo1222) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTTableColumnInfo1222) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTTableColumnInfo1222) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnInfo1222) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTTableColumnInfo1222) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTTableColumnInfo1222) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTTableColumnInfo1222) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnInfo1222) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTTableColumnInfo1222) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTTableColumnInfo1222) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

func (o BTTableColumnInfo1222) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Specification != nil {
		toSerialize["specification"] = o.Specification
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableColumnInfo1222 struct {
	value *BTTableColumnInfo1222
	isSet bool
}

func (v NullableBTTableColumnInfo1222) Get() *BTTableColumnInfo1222 {
	return v.value
}

func (v *NullableBTTableColumnInfo1222) Set(val *BTTableColumnInfo1222) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableColumnInfo1222) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableColumnInfo1222) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableColumnInfo1222(val *BTTableColumnInfo1222) *NullableBTTableColumnInfo1222 {
	return &NullableBTTableColumnInfo1222{value: val, isSet: true}
}

func (v NullableBTTableColumnInfo1222) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableColumnInfo1222) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
