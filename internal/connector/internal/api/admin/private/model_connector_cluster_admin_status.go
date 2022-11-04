/*
 * Connector Service Fleet Manager Admin APIs
 *
 * Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ConnectorClusterAdminStatus struct for ConnectorClusterAdminStatus
type ConnectorClusterAdminStatus struct {
	State      ConnectorClusterState    `json:"state,omitempty"`
	Version    string                   `json:"version,omitempty"`
	Conditions []MetaV1Condition        `json:"conditions,omitempty"`
	Platform   ConnectorClusterPlatform `json:"platform,omitempty"`
	// the list of installed operators
	Operators []ConnectorClusterAdminStatusOperators `json:"operators,omitempty"`
}