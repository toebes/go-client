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

// BTPModule234 struct for BTPModule234
type BTPModule234 struct {
	BTPNode7
	BtType *string `json:"btType,omitempty"`
	DeepImports *map[string][]BTImport `json:"deepImports,omitempty"`
	Imports *[]BTPTopLevelImport285 `json:"imports,omitempty"`
	IsBlob *bool `json:"isBlob,omitempty"`
	IsInternalModule *bool `json:"isInternalModule,omitempty"`
	MayHaveImplicitImports *bool `json:"mayHaveImplicitImports,omitempty"`
	PathMap *map[string]BTMicroversionId366 `json:"pathMap,omitempty"`
	PathToCache *BTCacheDataPath191 `json:"pathToCache,omitempty"`
	TopLevel *[]BTPTopLevelNode286 `json:"topLevel,omitempty"`
	Version *BTPLiteralNumber258 `json:"version,omitempty"`
	VersionNumber *int32 `json:"versionNumber,omitempty"`
}

// NewBTPModule234 instantiates a new BTPModule234 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPModule234() *BTPModule234 {
	this := BTPModule234{}
	return &this
}

// NewBTPModule234WithDefaults instantiates a new BTPModule234 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPModule234WithDefaults() *BTPModule234 {
	this := BTPModule234{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPModule234) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPModule234) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPModule234) SetBtType(v string) {
	o.BtType = &v
}

// GetDeepImports returns the DeepImports field value if set, zero value otherwise.
func (o *BTPModule234) GetDeepImports() map[string][]BTImport {
	if o == nil || o.DeepImports == nil {
		var ret map[string][]BTImport
		return ret
	}
	return *o.DeepImports
}

// GetDeepImportsOk returns a tuple with the DeepImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetDeepImportsOk() (*map[string][]BTImport, bool) {
	if o == nil || o.DeepImports == nil {
		return nil, false
	}
	return o.DeepImports, true
}

// HasDeepImports returns a boolean if a field has been set.
func (o *BTPModule234) HasDeepImports() bool {
	if o != nil && o.DeepImports != nil {
		return true
	}

	return false
}

// SetDeepImports gets a reference to the given map[string][]BTImport and assigns it to the DeepImports field.
func (o *BTPModule234) SetDeepImports(v map[string][]BTImport) {
	o.DeepImports = &v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *BTPModule234) GetImports() []BTPTopLevelImport285 {
	if o == nil || o.Imports == nil {
		var ret []BTPTopLevelImport285
		return ret
	}
	return *o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetImportsOk() (*[]BTPTopLevelImport285, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *BTPModule234) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []BTPTopLevelImport285 and assigns it to the Imports field.
func (o *BTPModule234) SetImports(v []BTPTopLevelImport285) {
	o.Imports = &v
}

// GetIsBlob returns the IsBlob field value if set, zero value otherwise.
func (o *BTPModule234) GetIsBlob() bool {
	if o == nil || o.IsBlob == nil {
		var ret bool
		return ret
	}
	return *o.IsBlob
}

// GetIsBlobOk returns a tuple with the IsBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetIsBlobOk() (*bool, bool) {
	if o == nil || o.IsBlob == nil {
		return nil, false
	}
	return o.IsBlob, true
}

// HasIsBlob returns a boolean if a field has been set.
func (o *BTPModule234) HasIsBlob() bool {
	if o != nil && o.IsBlob != nil {
		return true
	}

	return false
}

// SetIsBlob gets a reference to the given bool and assigns it to the IsBlob field.
func (o *BTPModule234) SetIsBlob(v bool) {
	o.IsBlob = &v
}

// GetIsInternalModule returns the IsInternalModule field value if set, zero value otherwise.
func (o *BTPModule234) GetIsInternalModule() bool {
	if o == nil || o.IsInternalModule == nil {
		var ret bool
		return ret
	}
	return *o.IsInternalModule
}

// GetIsInternalModuleOk returns a tuple with the IsInternalModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetIsInternalModuleOk() (*bool, bool) {
	if o == nil || o.IsInternalModule == nil {
		return nil, false
	}
	return o.IsInternalModule, true
}

// HasIsInternalModule returns a boolean if a field has been set.
func (o *BTPModule234) HasIsInternalModule() bool {
	if o != nil && o.IsInternalModule != nil {
		return true
	}

	return false
}

// SetIsInternalModule gets a reference to the given bool and assigns it to the IsInternalModule field.
func (o *BTPModule234) SetIsInternalModule(v bool) {
	o.IsInternalModule = &v
}

// GetMayHaveImplicitImports returns the MayHaveImplicitImports field value if set, zero value otherwise.
func (o *BTPModule234) GetMayHaveImplicitImports() bool {
	if o == nil || o.MayHaveImplicitImports == nil {
		var ret bool
		return ret
	}
	return *o.MayHaveImplicitImports
}

// GetMayHaveImplicitImportsOk returns a tuple with the MayHaveImplicitImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetMayHaveImplicitImportsOk() (*bool, bool) {
	if o == nil || o.MayHaveImplicitImports == nil {
		return nil, false
	}
	return o.MayHaveImplicitImports, true
}

// HasMayHaveImplicitImports returns a boolean if a field has been set.
func (o *BTPModule234) HasMayHaveImplicitImports() bool {
	if o != nil && o.MayHaveImplicitImports != nil {
		return true
	}

	return false
}

// SetMayHaveImplicitImports gets a reference to the given bool and assigns it to the MayHaveImplicitImports field.
func (o *BTPModule234) SetMayHaveImplicitImports(v bool) {
	o.MayHaveImplicitImports = &v
}

// GetPathMap returns the PathMap field value if set, zero value otherwise.
func (o *BTPModule234) GetPathMap() map[string]BTMicroversionId366 {
	if o == nil || o.PathMap == nil {
		var ret map[string]BTMicroversionId366
		return ret
	}
	return *o.PathMap
}

// GetPathMapOk returns a tuple with the PathMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetPathMapOk() (*map[string]BTMicroversionId366, bool) {
	if o == nil || o.PathMap == nil {
		return nil, false
	}
	return o.PathMap, true
}

// HasPathMap returns a boolean if a field has been set.
func (o *BTPModule234) HasPathMap() bool {
	if o != nil && o.PathMap != nil {
		return true
	}

	return false
}

// SetPathMap gets a reference to the given map[string]BTMicroversionId366 and assigns it to the PathMap field.
func (o *BTPModule234) SetPathMap(v map[string]BTMicroversionId366) {
	o.PathMap = &v
}

// GetPathToCache returns the PathToCache field value if set, zero value otherwise.
func (o *BTPModule234) GetPathToCache() BTCacheDataPath191 {
	if o == nil || o.PathToCache == nil {
		var ret BTCacheDataPath191
		return ret
	}
	return *o.PathToCache
}

// GetPathToCacheOk returns a tuple with the PathToCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetPathToCacheOk() (*BTCacheDataPath191, bool) {
	if o == nil || o.PathToCache == nil {
		return nil, false
	}
	return o.PathToCache, true
}

// HasPathToCache returns a boolean if a field has been set.
func (o *BTPModule234) HasPathToCache() bool {
	if o != nil && o.PathToCache != nil {
		return true
	}

	return false
}

// SetPathToCache gets a reference to the given BTCacheDataPath191 and assigns it to the PathToCache field.
func (o *BTPModule234) SetPathToCache(v BTCacheDataPath191) {
	o.PathToCache = &v
}

// GetTopLevel returns the TopLevel field value if set, zero value otherwise.
func (o *BTPModule234) GetTopLevel() []BTPTopLevelNode286 {
	if o == nil || o.TopLevel == nil {
		var ret []BTPTopLevelNode286
		return ret
	}
	return *o.TopLevel
}

// GetTopLevelOk returns a tuple with the TopLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetTopLevelOk() (*[]BTPTopLevelNode286, bool) {
	if o == nil || o.TopLevel == nil {
		return nil, false
	}
	return o.TopLevel, true
}

// HasTopLevel returns a boolean if a field has been set.
func (o *BTPModule234) HasTopLevel() bool {
	if o != nil && o.TopLevel != nil {
		return true
	}

	return false
}

// SetTopLevel gets a reference to the given []BTPTopLevelNode286 and assigns it to the TopLevel field.
func (o *BTPModule234) SetTopLevel(v []BTPTopLevelNode286) {
	o.TopLevel = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTPModule234) GetVersion() BTPLiteralNumber258 {
	if o == nil || o.Version == nil {
		var ret BTPLiteralNumber258
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetVersionOk() (*BTPLiteralNumber258, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTPModule234) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given BTPLiteralNumber258 and assigns it to the Version field.
func (o *BTPModule234) SetVersion(v BTPLiteralNumber258) {
	o.Version = &v
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *BTPModule234) GetVersionNumber() int32 {
	if o == nil || o.VersionNumber == nil {
		var ret int32
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModule234) GetVersionNumberOk() (*int32, bool) {
	if o == nil || o.VersionNumber == nil {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *BTPModule234) HasVersionNumber() bool {
	if o != nil && o.VersionNumber != nil {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given int32 and assigns it to the VersionNumber field.
func (o *BTPModule234) SetVersionNumber(v int32) {
	o.VersionNumber = &v
}

func (o BTPModule234) MarshalJSON() ([]byte, error) {
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
	if o.DeepImports != nil {
		toSerialize["deepImports"] = o.DeepImports
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.IsBlob != nil {
		toSerialize["isBlob"] = o.IsBlob
	}
	if o.IsInternalModule != nil {
		toSerialize["isInternalModule"] = o.IsInternalModule
	}
	if o.MayHaveImplicitImports != nil {
		toSerialize["mayHaveImplicitImports"] = o.MayHaveImplicitImports
	}
	if o.PathMap != nil {
		toSerialize["pathMap"] = o.PathMap
	}
	if o.PathToCache != nil {
		toSerialize["pathToCache"] = o.PathToCache
	}
	if o.TopLevel != nil {
		toSerialize["topLevel"] = o.TopLevel
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VersionNumber != nil {
		toSerialize["versionNumber"] = o.VersionNumber
	}
	return json.Marshal(toSerialize)
}

type NullableBTPModule234 struct {
	value *BTPModule234
	isSet bool
}

func (v NullableBTPModule234) Get() *BTPModule234 {
	return v.value
}

func (v *NullableBTPModule234) Set(val *BTPModule234) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPModule234) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPModule234) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPModule234(val *BTPModule234) *NullableBTPModule234 {
	return &NullableBTPModule234{value: val, isSet: true}
}

func (v NullableBTPModule234) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPModule234) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
