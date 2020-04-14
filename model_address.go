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
// Address struct for Address
type Address struct {
	City string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Line1 string `json:"line1,omitempty"`
	Line2 string `json:"line2,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	State string `json:"state,omitempty"`
}