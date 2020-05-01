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

// BTPModuleId235AllOf struct for BTPModuleId235AllOf
type BTPModuleId235AllOf struct {
	BtType *string `json:"btType,omitempty"`
	DbimportString *string `json:"dbimportString,omitempty"`
	ElementImport *bool `json:"elementImport,omitempty"`
	ExternalDocumentWithVersion *BTDocumentWithVersionId `json:"externalDocumentWithVersion,omitempty"`
	ExternalDocumentWithVersionAndElementId *BTDocumentWithVersionAndElementId `json:"externalDocumentWithVersionAndElementId,omitempty"`
	ExternalImport *bool `json:"externalImport,omitempty"`
	ImportedDocumentId *string `json:"importedDocumentId,omitempty"`
	ImportedElementId *string `json:"importedElementId,omitempty"`
	ImportedVersionId *string `json:"importedVersionId,omitempty"`
	Legacy *bool `json:"legacy,omitempty"`
	Microversion *string `json:"microversion,omitempty"`
	Path *BTPLiteralString259 `json:"path,omitempty"`
	PathPotentiallyValid *bool `json:"pathPotentiallyValid,omitempty"`
	PathVersion *string `json:"pathVersion,omitempty"`
	PotentiallyValid *bool `json:"potentiallyValid,omitempty"`
	SpaceAfterPath *BTPSpace10 `json:"spaceAfterPath,omitempty"`
	SpaceAfterVersion *BTPSpace10 `json:"spaceAfterVersion,omitempty"`
	SpaceBeforePath *BTPSpace10 `json:"spaceBeforePath,omitempty"`
	SpaceBeforeVersion *BTPSpace10 `json:"spaceBeforeVersion,omitempty"`
	StandardLibrary *bool `json:"standardLibrary,omitempty"`
	StandardLibraryCommon *bool `json:"standardLibraryCommon,omitempty"`
	ValidLegacyVersion *bool `json:"validLegacyVersion,omitempty"`
	Version *BTPLiteralString259 `json:"version,omitempty"`
	VersionAndMicroversion *string `json:"versionAndMicroversion,omitempty"`
	VersionPotentiallyValid *bool `json:"versionPotentiallyValid,omitempty"`
}

// NewBTPModuleId235AllOf instantiates a new BTPModuleId235AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPModuleId235AllOf() *BTPModuleId235AllOf {
	this := BTPModuleId235AllOf{}
	return &this
}

// NewBTPModuleId235AllOfWithDefaults instantiates a new BTPModuleId235AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPModuleId235AllOfWithDefaults() *BTPModuleId235AllOf {
	this := BTPModuleId235AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPModuleId235AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetDbimportString returns the DbimportString field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetDbimportString() string {
	if o == nil || o.DbimportString == nil {
		var ret string
		return ret
	}
	return *o.DbimportString
}

// GetDbimportStringOk returns a tuple with the DbimportString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetDbimportStringOk() (*string, bool) {
	if o == nil || o.DbimportString == nil {
		return nil, false
	}
	return o.DbimportString, true
}

// HasDbimportString returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasDbimportString() bool {
	if o != nil && o.DbimportString != nil {
		return true
	}

	return false
}

// SetDbimportString gets a reference to the given string and assigns it to the DbimportString field.
func (o *BTPModuleId235AllOf) SetDbimportString(v string) {
	o.DbimportString = &v
}

// GetElementImport returns the ElementImport field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetElementImport() bool {
	if o == nil || o.ElementImport == nil {
		var ret bool
		return ret
	}
	return *o.ElementImport
}

// GetElementImportOk returns a tuple with the ElementImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetElementImportOk() (*bool, bool) {
	if o == nil || o.ElementImport == nil {
		return nil, false
	}
	return o.ElementImport, true
}

// HasElementImport returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasElementImport() bool {
	if o != nil && o.ElementImport != nil {
		return true
	}

	return false
}

// SetElementImport gets a reference to the given bool and assigns it to the ElementImport field.
func (o *BTPModuleId235AllOf) SetElementImport(v bool) {
	o.ElementImport = &v
}

// GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetExternalDocumentWithVersion() BTDocumentWithVersionId {
	if o == nil || o.ExternalDocumentWithVersion == nil {
		var ret BTDocumentWithVersionId
		return ret
	}
	return *o.ExternalDocumentWithVersion
}

// GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool) {
	if o == nil || o.ExternalDocumentWithVersion == nil {
		return nil, false
	}
	return o.ExternalDocumentWithVersion, true
}

// HasExternalDocumentWithVersion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasExternalDocumentWithVersion() bool {
	if o != nil && o.ExternalDocumentWithVersion != nil {
		return true
	}

	return false
}

// SetExternalDocumentWithVersion gets a reference to the given BTDocumentWithVersionId and assigns it to the ExternalDocumentWithVersion field.
func (o *BTPModuleId235AllOf) SetExternalDocumentWithVersion(v BTDocumentWithVersionId) {
	o.ExternalDocumentWithVersion = &v
}

// GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId {
	if o == nil || o.ExternalDocumentWithVersionAndElementId == nil {
		var ret BTDocumentWithVersionAndElementId
		return ret
	}
	return *o.ExternalDocumentWithVersionAndElementId
}

// GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool) {
	if o == nil || o.ExternalDocumentWithVersionAndElementId == nil {
		return nil, false
	}
	return o.ExternalDocumentWithVersionAndElementId, true
}

// HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasExternalDocumentWithVersionAndElementId() bool {
	if o != nil && o.ExternalDocumentWithVersionAndElementId != nil {
		return true
	}

	return false
}

// SetExternalDocumentWithVersionAndElementId gets a reference to the given BTDocumentWithVersionAndElementId and assigns it to the ExternalDocumentWithVersionAndElementId field.
func (o *BTPModuleId235AllOf) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId) {
	o.ExternalDocumentWithVersionAndElementId = &v
}

// GetExternalImport returns the ExternalImport field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetExternalImport() bool {
	if o == nil || o.ExternalImport == nil {
		var ret bool
		return ret
	}
	return *o.ExternalImport
}

// GetExternalImportOk returns a tuple with the ExternalImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetExternalImportOk() (*bool, bool) {
	if o == nil || o.ExternalImport == nil {
		return nil, false
	}
	return o.ExternalImport, true
}

// HasExternalImport returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasExternalImport() bool {
	if o != nil && o.ExternalImport != nil {
		return true
	}

	return false
}

// SetExternalImport gets a reference to the given bool and assigns it to the ExternalImport field.
func (o *BTPModuleId235AllOf) SetExternalImport(v bool) {
	o.ExternalImport = &v
}

// GetImportedDocumentId returns the ImportedDocumentId field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetImportedDocumentId() string {
	if o == nil || o.ImportedDocumentId == nil {
		var ret string
		return ret
	}
	return *o.ImportedDocumentId
}

// GetImportedDocumentIdOk returns a tuple with the ImportedDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetImportedDocumentIdOk() (*string, bool) {
	if o == nil || o.ImportedDocumentId == nil {
		return nil, false
	}
	return o.ImportedDocumentId, true
}

// HasImportedDocumentId returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasImportedDocumentId() bool {
	if o != nil && o.ImportedDocumentId != nil {
		return true
	}

	return false
}

// SetImportedDocumentId gets a reference to the given string and assigns it to the ImportedDocumentId field.
func (o *BTPModuleId235AllOf) SetImportedDocumentId(v string) {
	o.ImportedDocumentId = &v
}

// GetImportedElementId returns the ImportedElementId field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetImportedElementId() string {
	if o == nil || o.ImportedElementId == nil {
		var ret string
		return ret
	}
	return *o.ImportedElementId
}

// GetImportedElementIdOk returns a tuple with the ImportedElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetImportedElementIdOk() (*string, bool) {
	if o == nil || o.ImportedElementId == nil {
		return nil, false
	}
	return o.ImportedElementId, true
}

// HasImportedElementId returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasImportedElementId() bool {
	if o != nil && o.ImportedElementId != nil {
		return true
	}

	return false
}

// SetImportedElementId gets a reference to the given string and assigns it to the ImportedElementId field.
func (o *BTPModuleId235AllOf) SetImportedElementId(v string) {
	o.ImportedElementId = &v
}

// GetImportedVersionId returns the ImportedVersionId field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetImportedVersionId() string {
	if o == nil || o.ImportedVersionId == nil {
		var ret string
		return ret
	}
	return *o.ImportedVersionId
}

// GetImportedVersionIdOk returns a tuple with the ImportedVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetImportedVersionIdOk() (*string, bool) {
	if o == nil || o.ImportedVersionId == nil {
		return nil, false
	}
	return o.ImportedVersionId, true
}

// HasImportedVersionId returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasImportedVersionId() bool {
	if o != nil && o.ImportedVersionId != nil {
		return true
	}

	return false
}

// SetImportedVersionId gets a reference to the given string and assigns it to the ImportedVersionId field.
func (o *BTPModuleId235AllOf) SetImportedVersionId(v string) {
	o.ImportedVersionId = &v
}

// GetLegacy returns the Legacy field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetLegacy() bool {
	if o == nil || o.Legacy == nil {
		var ret bool
		return ret
	}
	return *o.Legacy
}

// GetLegacyOk returns a tuple with the Legacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetLegacyOk() (*bool, bool) {
	if o == nil || o.Legacy == nil {
		return nil, false
	}
	return o.Legacy, true
}

// HasLegacy returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasLegacy() bool {
	if o != nil && o.Legacy != nil {
		return true
	}

	return false
}

// SetLegacy gets a reference to the given bool and assigns it to the Legacy field.
func (o *BTPModuleId235AllOf) SetLegacy(v bool) {
	o.Legacy = &v
}

// GetMicroversion returns the Microversion field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetMicroversion() string {
	if o == nil || o.Microversion == nil {
		var ret string
		return ret
	}
	return *o.Microversion
}

// GetMicroversionOk returns a tuple with the Microversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetMicroversionOk() (*string, bool) {
	if o == nil || o.Microversion == nil {
		return nil, false
	}
	return o.Microversion, true
}

// HasMicroversion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasMicroversion() bool {
	if o != nil && o.Microversion != nil {
		return true
	}

	return false
}

// SetMicroversion gets a reference to the given string and assigns it to the Microversion field.
func (o *BTPModuleId235AllOf) SetMicroversion(v string) {
	o.Microversion = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetPath() BTPLiteralString259 {
	if o == nil || o.Path == nil {
		var ret BTPLiteralString259
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetPathOk() (*BTPLiteralString259, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given BTPLiteralString259 and assigns it to the Path field.
func (o *BTPModuleId235AllOf) SetPath(v BTPLiteralString259) {
	o.Path = &v
}

// GetPathPotentiallyValid returns the PathPotentiallyValid field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetPathPotentiallyValid() bool {
	if o == nil || o.PathPotentiallyValid == nil {
		var ret bool
		return ret
	}
	return *o.PathPotentiallyValid
}

// GetPathPotentiallyValidOk returns a tuple with the PathPotentiallyValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetPathPotentiallyValidOk() (*bool, bool) {
	if o == nil || o.PathPotentiallyValid == nil {
		return nil, false
	}
	return o.PathPotentiallyValid, true
}

// HasPathPotentiallyValid returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasPathPotentiallyValid() bool {
	if o != nil && o.PathPotentiallyValid != nil {
		return true
	}

	return false
}

// SetPathPotentiallyValid gets a reference to the given bool and assigns it to the PathPotentiallyValid field.
func (o *BTPModuleId235AllOf) SetPathPotentiallyValid(v bool) {
	o.PathPotentiallyValid = &v
}

// GetPathVersion returns the PathVersion field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetPathVersion() string {
	if o == nil || o.PathVersion == nil {
		var ret string
		return ret
	}
	return *o.PathVersion
}

// GetPathVersionOk returns a tuple with the PathVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetPathVersionOk() (*string, bool) {
	if o == nil || o.PathVersion == nil {
		return nil, false
	}
	return o.PathVersion, true
}

// HasPathVersion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasPathVersion() bool {
	if o != nil && o.PathVersion != nil {
		return true
	}

	return false
}

// SetPathVersion gets a reference to the given string and assigns it to the PathVersion field.
func (o *BTPModuleId235AllOf) SetPathVersion(v string) {
	o.PathVersion = &v
}

// GetPotentiallyValid returns the PotentiallyValid field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetPotentiallyValid() bool {
	if o == nil || o.PotentiallyValid == nil {
		var ret bool
		return ret
	}
	return *o.PotentiallyValid
}

// GetPotentiallyValidOk returns a tuple with the PotentiallyValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetPotentiallyValidOk() (*bool, bool) {
	if o == nil || o.PotentiallyValid == nil {
		return nil, false
	}
	return o.PotentiallyValid, true
}

// HasPotentiallyValid returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasPotentiallyValid() bool {
	if o != nil && o.PotentiallyValid != nil {
		return true
	}

	return false
}

// SetPotentiallyValid gets a reference to the given bool and assigns it to the PotentiallyValid field.
func (o *BTPModuleId235AllOf) SetPotentiallyValid(v bool) {
	o.PotentiallyValid = &v
}

// GetSpaceAfterPath returns the SpaceAfterPath field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetSpaceAfterPath() BTPSpace10 {
	if o == nil || o.SpaceAfterPath == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterPath
}

// GetSpaceAfterPathOk returns a tuple with the SpaceAfterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetSpaceAfterPathOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterPath == nil {
		return nil, false
	}
	return o.SpaceAfterPath, true
}

// HasSpaceAfterPath returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasSpaceAfterPath() bool {
	if o != nil && o.SpaceAfterPath != nil {
		return true
	}

	return false
}

// SetSpaceAfterPath gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterPath field.
func (o *BTPModuleId235AllOf) SetSpaceAfterPath(v BTPSpace10) {
	o.SpaceAfterPath = &v
}

// GetSpaceAfterVersion returns the SpaceAfterVersion field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetSpaceAfterVersion() BTPSpace10 {
	if o == nil || o.SpaceAfterVersion == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterVersion
}

// GetSpaceAfterVersionOk returns a tuple with the SpaceAfterVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetSpaceAfterVersionOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterVersion == nil {
		return nil, false
	}
	return o.SpaceAfterVersion, true
}

// HasSpaceAfterVersion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasSpaceAfterVersion() bool {
	if o != nil && o.SpaceAfterVersion != nil {
		return true
	}

	return false
}

// SetSpaceAfterVersion gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterVersion field.
func (o *BTPModuleId235AllOf) SetSpaceAfterVersion(v BTPSpace10) {
	o.SpaceAfterVersion = &v
}

// GetSpaceBeforePath returns the SpaceBeforePath field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetSpaceBeforePath() BTPSpace10 {
	if o == nil || o.SpaceBeforePath == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforePath
}

// GetSpaceBeforePathOk returns a tuple with the SpaceBeforePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetSpaceBeforePathOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforePath == nil {
		return nil, false
	}
	return o.SpaceBeforePath, true
}

// HasSpaceBeforePath returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasSpaceBeforePath() bool {
	if o != nil && o.SpaceBeforePath != nil {
		return true
	}

	return false
}

// SetSpaceBeforePath gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforePath field.
func (o *BTPModuleId235AllOf) SetSpaceBeforePath(v BTPSpace10) {
	o.SpaceBeforePath = &v
}

// GetSpaceBeforeVersion returns the SpaceBeforeVersion field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetSpaceBeforeVersion() BTPSpace10 {
	if o == nil || o.SpaceBeforeVersion == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeVersion
}

// GetSpaceBeforeVersionOk returns a tuple with the SpaceBeforeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetSpaceBeforeVersionOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeVersion == nil {
		return nil, false
	}
	return o.SpaceBeforeVersion, true
}

// HasSpaceBeforeVersion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasSpaceBeforeVersion() bool {
	if o != nil && o.SpaceBeforeVersion != nil {
		return true
	}

	return false
}

// SetSpaceBeforeVersion gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeVersion field.
func (o *BTPModuleId235AllOf) SetSpaceBeforeVersion(v BTPSpace10) {
	o.SpaceBeforeVersion = &v
}

// GetStandardLibrary returns the StandardLibrary field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetStandardLibrary() bool {
	if o == nil || o.StandardLibrary == nil {
		var ret bool
		return ret
	}
	return *o.StandardLibrary
}

// GetStandardLibraryOk returns a tuple with the StandardLibrary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetStandardLibraryOk() (*bool, bool) {
	if o == nil || o.StandardLibrary == nil {
		return nil, false
	}
	return o.StandardLibrary, true
}

// HasStandardLibrary returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasStandardLibrary() bool {
	if o != nil && o.StandardLibrary != nil {
		return true
	}

	return false
}

// SetStandardLibrary gets a reference to the given bool and assigns it to the StandardLibrary field.
func (o *BTPModuleId235AllOf) SetStandardLibrary(v bool) {
	o.StandardLibrary = &v
}

// GetStandardLibraryCommon returns the StandardLibraryCommon field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetStandardLibraryCommon() bool {
	if o == nil || o.StandardLibraryCommon == nil {
		var ret bool
		return ret
	}
	return *o.StandardLibraryCommon
}

// GetStandardLibraryCommonOk returns a tuple with the StandardLibraryCommon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetStandardLibraryCommonOk() (*bool, bool) {
	if o == nil || o.StandardLibraryCommon == nil {
		return nil, false
	}
	return o.StandardLibraryCommon, true
}

// HasStandardLibraryCommon returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasStandardLibraryCommon() bool {
	if o != nil && o.StandardLibraryCommon != nil {
		return true
	}

	return false
}

// SetStandardLibraryCommon gets a reference to the given bool and assigns it to the StandardLibraryCommon field.
func (o *BTPModuleId235AllOf) SetStandardLibraryCommon(v bool) {
	o.StandardLibraryCommon = &v
}

// GetValidLegacyVersion returns the ValidLegacyVersion field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetValidLegacyVersion() bool {
	if o == nil || o.ValidLegacyVersion == nil {
		var ret bool
		return ret
	}
	return *o.ValidLegacyVersion
}

// GetValidLegacyVersionOk returns a tuple with the ValidLegacyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetValidLegacyVersionOk() (*bool, bool) {
	if o == nil || o.ValidLegacyVersion == nil {
		return nil, false
	}
	return o.ValidLegacyVersion, true
}

// HasValidLegacyVersion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasValidLegacyVersion() bool {
	if o != nil && o.ValidLegacyVersion != nil {
		return true
	}

	return false
}

// SetValidLegacyVersion gets a reference to the given bool and assigns it to the ValidLegacyVersion field.
func (o *BTPModuleId235AllOf) SetValidLegacyVersion(v bool) {
	o.ValidLegacyVersion = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetVersion() BTPLiteralString259 {
	if o == nil || o.Version == nil {
		var ret BTPLiteralString259
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetVersionOk() (*BTPLiteralString259, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given BTPLiteralString259 and assigns it to the Version field.
func (o *BTPModuleId235AllOf) SetVersion(v BTPLiteralString259) {
	o.Version = &v
}

// GetVersionAndMicroversion returns the VersionAndMicroversion field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetVersionAndMicroversion() string {
	if o == nil || o.VersionAndMicroversion == nil {
		var ret string
		return ret
	}
	return *o.VersionAndMicroversion
}

// GetVersionAndMicroversionOk returns a tuple with the VersionAndMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetVersionAndMicroversionOk() (*string, bool) {
	if o == nil || o.VersionAndMicroversion == nil {
		return nil, false
	}
	return o.VersionAndMicroversion, true
}

// HasVersionAndMicroversion returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasVersionAndMicroversion() bool {
	if o != nil && o.VersionAndMicroversion != nil {
		return true
	}

	return false
}

// SetVersionAndMicroversion gets a reference to the given string and assigns it to the VersionAndMicroversion field.
func (o *BTPModuleId235AllOf) SetVersionAndMicroversion(v string) {
	o.VersionAndMicroversion = &v
}

// GetVersionPotentiallyValid returns the VersionPotentiallyValid field value if set, zero value otherwise.
func (o *BTPModuleId235AllOf) GetVersionPotentiallyValid() bool {
	if o == nil || o.VersionPotentiallyValid == nil {
		var ret bool
		return ret
	}
	return *o.VersionPotentiallyValid
}

// GetVersionPotentiallyValidOk returns a tuple with the VersionPotentiallyValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPModuleId235AllOf) GetVersionPotentiallyValidOk() (*bool, bool) {
	if o == nil || o.VersionPotentiallyValid == nil {
		return nil, false
	}
	return o.VersionPotentiallyValid, true
}

// HasVersionPotentiallyValid returns a boolean if a field has been set.
func (o *BTPModuleId235AllOf) HasVersionPotentiallyValid() bool {
	if o != nil && o.VersionPotentiallyValid != nil {
		return true
	}

	return false
}

// SetVersionPotentiallyValid gets a reference to the given bool and assigns it to the VersionPotentiallyValid field.
func (o *BTPModuleId235AllOf) SetVersionPotentiallyValid(v bool) {
	o.VersionPotentiallyValid = &v
}

func (o BTPModuleId235AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DbimportString != nil {
		toSerialize["dbimportString"] = o.DbimportString
	}
	if o.ElementImport != nil {
		toSerialize["elementImport"] = o.ElementImport
	}
	if o.ExternalDocumentWithVersion != nil {
		toSerialize["externalDocumentWithVersion"] = o.ExternalDocumentWithVersion
	}
	if o.ExternalDocumentWithVersionAndElementId != nil {
		toSerialize["externalDocumentWithVersionAndElementId"] = o.ExternalDocumentWithVersionAndElementId
	}
	if o.ExternalImport != nil {
		toSerialize["externalImport"] = o.ExternalImport
	}
	if o.ImportedDocumentId != nil {
		toSerialize["importedDocumentId"] = o.ImportedDocumentId
	}
	if o.ImportedElementId != nil {
		toSerialize["importedElementId"] = o.ImportedElementId
	}
	if o.ImportedVersionId != nil {
		toSerialize["importedVersionId"] = o.ImportedVersionId
	}
	if o.Legacy != nil {
		toSerialize["legacy"] = o.Legacy
	}
	if o.Microversion != nil {
		toSerialize["microversion"] = o.Microversion
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.PathPotentiallyValid != nil {
		toSerialize["pathPotentiallyValid"] = o.PathPotentiallyValid
	}
	if o.PathVersion != nil {
		toSerialize["pathVersion"] = o.PathVersion
	}
	if o.PotentiallyValid != nil {
		toSerialize["potentiallyValid"] = o.PotentiallyValid
	}
	if o.SpaceAfterPath != nil {
		toSerialize["spaceAfterPath"] = o.SpaceAfterPath
	}
	if o.SpaceAfterVersion != nil {
		toSerialize["spaceAfterVersion"] = o.SpaceAfterVersion
	}
	if o.SpaceBeforePath != nil {
		toSerialize["spaceBeforePath"] = o.SpaceBeforePath
	}
	if o.SpaceBeforeVersion != nil {
		toSerialize["spaceBeforeVersion"] = o.SpaceBeforeVersion
	}
	if o.StandardLibrary != nil {
		toSerialize["standardLibrary"] = o.StandardLibrary
	}
	if o.StandardLibraryCommon != nil {
		toSerialize["standardLibraryCommon"] = o.StandardLibraryCommon
	}
	if o.ValidLegacyVersion != nil {
		toSerialize["validLegacyVersion"] = o.ValidLegacyVersion
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VersionAndMicroversion != nil {
		toSerialize["versionAndMicroversion"] = o.VersionAndMicroversion
	}
	if o.VersionPotentiallyValid != nil {
		toSerialize["versionPotentiallyValid"] = o.VersionPotentiallyValid
	}
	return json.Marshal(toSerialize)
}

type NullableBTPModuleId235AllOf struct {
	value *BTPModuleId235AllOf
	isSet bool
}

func (v NullableBTPModuleId235AllOf) Get() *BTPModuleId235AllOf {
	return v.value
}

func (v *NullableBTPModuleId235AllOf) Set(val *BTPModuleId235AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPModuleId235AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPModuleId235AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPModuleId235AllOf(val *BTPModuleId235AllOf) *NullableBTPModuleId235AllOf {
	return &NullableBTPModuleId235AllOf{value: val, isSet: true}
}

func (v NullableBTPModuleId235AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPModuleId235AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}