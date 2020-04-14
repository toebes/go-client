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
// BtpExpressionSwitch2632AllOf struct for BtpExpressionSwitch2632AllOf
type BtpExpressionSwitch2632AllOf struct {
	Selector BtpExpression9 `json:"selector,omitempty"`
	Choices BtpLiteralMap256 `json:"choices,omitempty"`
	SpaceAfterSwitch BtpSpace10 `json:"spaceAfterSwitch,omitempty"`
	BtType string `json:"btType,omitempty"`
}