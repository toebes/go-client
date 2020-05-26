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

// BTCurveGeometryInterpolatedSpline116AllOf struct for BTCurveGeometryInterpolatedSpline116AllOf
type BTCurveGeometryInterpolatedSpline116AllOf struct {
	BtType *string `json:"btType,omitempty"`
	EndDerivativeX *float64 `json:"endDerivativeX,omitempty"`
	EndDerivativeY *float64 `json:"endDerivativeY,omitempty"`
	EndHandleX *float64 `json:"endHandleX,omitempty"`
	EndHandleY *float64 `json:"endHandleY,omitempty"`
	InterpolationPoints *[]float64 `json:"interpolationPoints,omitempty"`
	IsPeriodic *bool `json:"isPeriodic,omitempty"`
	StartDerivativeX *float64 `json:"startDerivativeX,omitempty"`
	StartDerivativeY *float64 `json:"startDerivativeY,omitempty"`
	StartHandleX *float64 `json:"startHandleX,omitempty"`
	StartHandleY *float64 `json:"startHandleY,omitempty"`
}

// NewBTCurveGeometryInterpolatedSpline116AllOf instantiates a new BTCurveGeometryInterpolatedSpline116AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryInterpolatedSpline116AllOf() *BTCurveGeometryInterpolatedSpline116AllOf {
	this := BTCurveGeometryInterpolatedSpline116AllOf{}
	return &this
}

// NewBTCurveGeometryInterpolatedSpline116AllOfWithDefaults instantiates a new BTCurveGeometryInterpolatedSpline116AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryInterpolatedSpline116AllOfWithDefaults() *BTCurveGeometryInterpolatedSpline116AllOf {
	this := BTCurveGeometryInterpolatedSpline116AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetEndDerivativeX returns the EndDerivativeX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndDerivativeX() float64 {
	if o == nil || o.EndDerivativeX == nil {
		var ret float64
		return ret
	}
	return *o.EndDerivativeX
}

// GetEndDerivativeXOk returns a tuple with the EndDerivativeX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndDerivativeXOk() (*float64, bool) {
	if o == nil || o.EndDerivativeX == nil {
		return nil, false
	}
	return o.EndDerivativeX, true
}

// HasEndDerivativeX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasEndDerivativeX() bool {
	if o != nil && o.EndDerivativeX != nil {
		return true
	}

	return false
}

// SetEndDerivativeX gets a reference to the given float64 and assigns it to the EndDerivativeX field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetEndDerivativeX(v float64) {
	o.EndDerivativeX = &v
}

// GetEndDerivativeY returns the EndDerivativeY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndDerivativeY() float64 {
	if o == nil || o.EndDerivativeY == nil {
		var ret float64
		return ret
	}
	return *o.EndDerivativeY
}

// GetEndDerivativeYOk returns a tuple with the EndDerivativeY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndDerivativeYOk() (*float64, bool) {
	if o == nil || o.EndDerivativeY == nil {
		return nil, false
	}
	return o.EndDerivativeY, true
}

// HasEndDerivativeY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasEndDerivativeY() bool {
	if o != nil && o.EndDerivativeY != nil {
		return true
	}

	return false
}

// SetEndDerivativeY gets a reference to the given float64 and assigns it to the EndDerivativeY field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetEndDerivativeY(v float64) {
	o.EndDerivativeY = &v
}

// GetEndHandleX returns the EndHandleX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndHandleX() float64 {
	if o == nil || o.EndHandleX == nil {
		var ret float64
		return ret
	}
	return *o.EndHandleX
}

// GetEndHandleXOk returns a tuple with the EndHandleX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndHandleXOk() (*float64, bool) {
	if o == nil || o.EndHandleX == nil {
		return nil, false
	}
	return o.EndHandleX, true
}

// HasEndHandleX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasEndHandleX() bool {
	if o != nil && o.EndHandleX != nil {
		return true
	}

	return false
}

// SetEndHandleX gets a reference to the given float64 and assigns it to the EndHandleX field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetEndHandleX(v float64) {
	o.EndHandleX = &v
}

// GetEndHandleY returns the EndHandleY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndHandleY() float64 {
	if o == nil || o.EndHandleY == nil {
		var ret float64
		return ret
	}
	return *o.EndHandleY
}

// GetEndHandleYOk returns a tuple with the EndHandleY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetEndHandleYOk() (*float64, bool) {
	if o == nil || o.EndHandleY == nil {
		return nil, false
	}
	return o.EndHandleY, true
}

// HasEndHandleY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasEndHandleY() bool {
	if o != nil && o.EndHandleY != nil {
		return true
	}

	return false
}

// SetEndHandleY gets a reference to the given float64 and assigns it to the EndHandleY field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetEndHandleY(v float64) {
	o.EndHandleY = &v
}

// GetInterpolationPoints returns the InterpolationPoints field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetInterpolationPoints() []float64 {
	if o == nil || o.InterpolationPoints == nil {
		var ret []float64
		return ret
	}
	return *o.InterpolationPoints
}

// GetInterpolationPointsOk returns a tuple with the InterpolationPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetInterpolationPointsOk() (*[]float64, bool) {
	if o == nil || o.InterpolationPoints == nil {
		return nil, false
	}
	return o.InterpolationPoints, true
}

// HasInterpolationPoints returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasInterpolationPoints() bool {
	if o != nil && o.InterpolationPoints != nil {
		return true
	}

	return false
}

// SetInterpolationPoints gets a reference to the given []float64 and assigns it to the InterpolationPoints field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetInterpolationPoints(v []float64) {
	o.InterpolationPoints = &v
}

// GetIsPeriodic returns the IsPeriodic field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetIsPeriodic() bool {
	if o == nil || o.IsPeriodic == nil {
		var ret bool
		return ret
	}
	return *o.IsPeriodic
}

// GetIsPeriodicOk returns a tuple with the IsPeriodic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetIsPeriodicOk() (*bool, bool) {
	if o == nil || o.IsPeriodic == nil {
		return nil, false
	}
	return o.IsPeriodic, true
}

// HasIsPeriodic returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasIsPeriodic() bool {
	if o != nil && o.IsPeriodic != nil {
		return true
	}

	return false
}

// SetIsPeriodic gets a reference to the given bool and assigns it to the IsPeriodic field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetIsPeriodic(v bool) {
	o.IsPeriodic = &v
}

// GetStartDerivativeX returns the StartDerivativeX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartDerivativeX() float64 {
	if o == nil || o.StartDerivativeX == nil {
		var ret float64
		return ret
	}
	return *o.StartDerivativeX
}

// GetStartDerivativeXOk returns a tuple with the StartDerivativeX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartDerivativeXOk() (*float64, bool) {
	if o == nil || o.StartDerivativeX == nil {
		return nil, false
	}
	return o.StartDerivativeX, true
}

// HasStartDerivativeX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasStartDerivativeX() bool {
	if o != nil && o.StartDerivativeX != nil {
		return true
	}

	return false
}

// SetStartDerivativeX gets a reference to the given float64 and assigns it to the StartDerivativeX field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetStartDerivativeX(v float64) {
	o.StartDerivativeX = &v
}

// GetStartDerivativeY returns the StartDerivativeY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartDerivativeY() float64 {
	if o == nil || o.StartDerivativeY == nil {
		var ret float64
		return ret
	}
	return *o.StartDerivativeY
}

// GetStartDerivativeYOk returns a tuple with the StartDerivativeY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartDerivativeYOk() (*float64, bool) {
	if o == nil || o.StartDerivativeY == nil {
		return nil, false
	}
	return o.StartDerivativeY, true
}

// HasStartDerivativeY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasStartDerivativeY() bool {
	if o != nil && o.StartDerivativeY != nil {
		return true
	}

	return false
}

// SetStartDerivativeY gets a reference to the given float64 and assigns it to the StartDerivativeY field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetStartDerivativeY(v float64) {
	o.StartDerivativeY = &v
}

// GetStartHandleX returns the StartHandleX field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartHandleX() float64 {
	if o == nil || o.StartHandleX == nil {
		var ret float64
		return ret
	}
	return *o.StartHandleX
}

// GetStartHandleXOk returns a tuple with the StartHandleX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartHandleXOk() (*float64, bool) {
	if o == nil || o.StartHandleX == nil {
		return nil, false
	}
	return o.StartHandleX, true
}

// HasStartHandleX returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasStartHandleX() bool {
	if o != nil && o.StartHandleX != nil {
		return true
	}

	return false
}

// SetStartHandleX gets a reference to the given float64 and assigns it to the StartHandleX field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetStartHandleX(v float64) {
	o.StartHandleX = &v
}

// GetStartHandleY returns the StartHandleY field value if set, zero value otherwise.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartHandleY() float64 {
	if o == nil || o.StartHandleY == nil {
		var ret float64
		return ret
	}
	return *o.StartHandleY
}

// GetStartHandleYOk returns a tuple with the StartHandleY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) GetStartHandleYOk() (*float64, bool) {
	if o == nil || o.StartHandleY == nil {
		return nil, false
	}
	return o.StartHandleY, true
}

// HasStartHandleY returns a boolean if a field has been set.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) HasStartHandleY() bool {
	if o != nil && o.StartHandleY != nil {
		return true
	}

	return false
}

// SetStartHandleY gets a reference to the given float64 and assigns it to the StartHandleY field.
func (o *BTCurveGeometryInterpolatedSpline116AllOf) SetStartHandleY(v float64) {
	o.StartHandleY = &v
}

func (o BTCurveGeometryInterpolatedSpline116AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EndDerivativeX != nil {
		toSerialize["endDerivativeX"] = o.EndDerivativeX
	}
	if o.EndDerivativeY != nil {
		toSerialize["endDerivativeY"] = o.EndDerivativeY
	}
	if o.EndHandleX != nil {
		toSerialize["endHandleX"] = o.EndHandleX
	}
	if o.EndHandleY != nil {
		toSerialize["endHandleY"] = o.EndHandleY
	}
	if o.InterpolationPoints != nil {
		toSerialize["interpolationPoints"] = o.InterpolationPoints
	}
	if o.IsPeriodic != nil {
		toSerialize["isPeriodic"] = o.IsPeriodic
	}
	if o.StartDerivativeX != nil {
		toSerialize["startDerivativeX"] = o.StartDerivativeX
	}
	if o.StartDerivativeY != nil {
		toSerialize["startDerivativeY"] = o.StartDerivativeY
	}
	if o.StartHandleX != nil {
		toSerialize["startHandleX"] = o.StartHandleX
	}
	if o.StartHandleY != nil {
		toSerialize["startHandleY"] = o.StartHandleY
	}
	return json.Marshal(toSerialize)
}

type NullableBTCurveGeometryInterpolatedSpline116AllOf struct {
	value *BTCurveGeometryInterpolatedSpline116AllOf
	isSet bool
}

func (v NullableBTCurveGeometryInterpolatedSpline116AllOf) Get() *BTCurveGeometryInterpolatedSpline116AllOf {
	return v.value
}

func (v *NullableBTCurveGeometryInterpolatedSpline116AllOf) Set(val *BTCurveGeometryInterpolatedSpline116AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryInterpolatedSpline116AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryInterpolatedSpline116AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryInterpolatedSpline116AllOf(val *BTCurveGeometryInterpolatedSpline116AllOf) *NullableBTCurveGeometryInterpolatedSpline116AllOf {
	return &NullableBTCurveGeometryInterpolatedSpline116AllOf{value: val, isSet: true}
}

func (v NullableBTCurveGeometryInterpolatedSpline116AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryInterpolatedSpline116AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}