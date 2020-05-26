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

// BTUserSummaryInfo struct for BTUserSummaryInfo
type BTUserSummaryInfo struct {
	BTUserBasicSummaryInfo
	Company *BTCompanySummaryInfo `json:"company,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	GlobalPermissions *GlobalPermissionInfo `json:"globalPermissions,omitempty"`
	IsGuest *bool `json:"isGuest,omitempty"`
	IsLight *bool `json:"isLight,omitempty"`
	LastLoginTime *JSONTime `json:"lastLoginTime,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Source *int32 `json:"source,omitempty"`
}

// NewBTUserSummaryInfo instantiates a new BTUserSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserSummaryInfo() *BTUserSummaryInfo {
	this := BTUserSummaryInfo{}
	return &this
}

// NewBTUserSummaryInfoWithDefaults instantiates a new BTUserSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserSummaryInfoWithDefaults() *BTUserSummaryInfo {
	this := BTUserSummaryInfo{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetCompany() BTCompanySummaryInfo {
	if o == nil || o.Company == nil {
		var ret BTCompanySummaryInfo
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given BTCompanySummaryInfo and assigns it to the Company field.
func (o *BTUserSummaryInfo) SetCompany(v BTCompanySummaryInfo) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTUserSummaryInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BTUserSummaryInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetGlobalPermissions() GlobalPermissionInfo {
	if o == nil || o.GlobalPermissions == nil {
		var ret GlobalPermissionInfo
		return ret
	}
	return *o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool) {
	if o == nil || o.GlobalPermissions == nil {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasGlobalPermissions() bool {
	if o != nil && o.GlobalPermissions != nil {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given GlobalPermissionInfo and assigns it to the GlobalPermissions field.
func (o *BTUserSummaryInfo) SetGlobalPermissions(v GlobalPermissionInfo) {
	o.GlobalPermissions = &v
}

// GetIsGuest returns the IsGuest field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetIsGuest() bool {
	if o == nil || o.IsGuest == nil {
		var ret bool
		return ret
	}
	return *o.IsGuest
}

// GetIsGuestOk returns a tuple with the IsGuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetIsGuestOk() (*bool, bool) {
	if o == nil || o.IsGuest == nil {
		return nil, false
	}
	return o.IsGuest, true
}

// HasIsGuest returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasIsGuest() bool {
	if o != nil && o.IsGuest != nil {
		return true
	}

	return false
}

// SetIsGuest gets a reference to the given bool and assigns it to the IsGuest field.
func (o *BTUserSummaryInfo) SetIsGuest(v bool) {
	o.IsGuest = &v
}

// GetIsLight returns the IsLight field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetIsLight() bool {
	if o == nil || o.IsLight == nil {
		var ret bool
		return ret
	}
	return *o.IsLight
}

// GetIsLightOk returns a tuple with the IsLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetIsLightOk() (*bool, bool) {
	if o == nil || o.IsLight == nil {
		return nil, false
	}
	return o.IsLight, true
}

// HasIsLight returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasIsLight() bool {
	if o != nil && o.IsLight != nil {
		return true
	}

	return false
}

// SetIsLight gets a reference to the given bool and assigns it to the IsLight field.
func (o *BTUserSummaryInfo) SetIsLight(v bool) {
	o.IsLight = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetLastLoginTime() JSONTime {
	if o == nil || o.LastLoginTime == nil {
		var ret JSONTime
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetLastLoginTimeOk() (*JSONTime, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given JSONTime and assigns it to the LastLoginTime field.
func (o *BTUserSummaryInfo) SetLastLoginTime(v JSONTime) {
	o.LastLoginTime = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BTUserSummaryInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTUserSummaryInfo) GetSource() int32 {
	if o == nil || o.Source == nil {
		var ret int32
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserSummaryInfo) GetSourceOk() (*int32, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTUserSummaryInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given int32 and assigns it to the Source field.
func (o *BTUserSummaryInfo) SetSource(v int32) {
	o.Source = &v
}

func (o BTUserSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTUserBasicSummaryInfo, errBTUserBasicSummaryInfo := json.Marshal(o.BTUserBasicSummaryInfo)
	if errBTUserBasicSummaryInfo != nil {
		return []byte{}, errBTUserBasicSummaryInfo
	}
	errBTUserBasicSummaryInfo = json.Unmarshal([]byte(serializedBTUserBasicSummaryInfo), &toSerialize)
	if errBTUserBasicSummaryInfo != nil {
		return []byte{}, errBTUserBasicSummaryInfo
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.GlobalPermissions != nil {
		toSerialize["globalPermissions"] = o.GlobalPermissions
	}
	if o.IsGuest != nil {
		toSerialize["isGuest"] = o.IsGuest
	}
	if o.IsLight != nil {
		toSerialize["isLight"] = o.IsLight
	}
	if o.LastLoginTime != nil {
		toSerialize["lastLoginTime"] = o.LastLoginTime
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserSummaryInfo struct {
	value *BTUserSummaryInfo
	isSet bool
}

func (v NullableBTUserSummaryInfo) Get() *BTUserSummaryInfo {
	return v.value
}

func (v *NullableBTUserSummaryInfo) Set(val *BTUserSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserSummaryInfo(val *BTUserSummaryInfo) *NullableBTUserSummaryInfo {
	return &NullableBTUserSummaryInfo{value: val, isSet: true}
}

func (v NullableBTUserSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}