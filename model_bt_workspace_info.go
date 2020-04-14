/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.111
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// BtWorkspaceInfo struct for BtWorkspaceInfo
type BtWorkspaceInfo struct {
	CanDelete bool `json:"canDelete,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Creator BtUserBasicSummaryInfo `json:"creator,omitempty"`
	Description string `json:"description,omitempty"`
	DocumentId string `json:"documentId,omitempty"`
	Href string `json:"href,omitempty"`
	Id string `json:"id,omitempty"`
	IsReadOnly bool `json:"isReadOnly,omitempty"`
	LastModifier BtUserBasicSummaryInfo `json:"lastModifier,omitempty"`
	Microversion string `json:"microversion,omitempty"`
	ModifiedAt time.Time `json:"modifiedAt,omitempty"`
	Name string `json:"name,omitempty"`
	OverrideDate time.Time `json:"overrideDate,omitempty"`
	Parent string `json:"parent,omitempty"`
	Parents []BtVersionInfo `json:"parents,omitempty"`
	Thumbnail BtThumbnailInfo `json:"thumbnail,omitempty"`
	Type string `json:"type,omitempty"`
	ViewRef string `json:"viewRef,omitempty"`
}