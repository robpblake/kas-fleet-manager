/*
 * Connector Service Fleet Manager Admin APIs
 *
 * Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ConnectorNamespaceWithTenantRequest struct for ConnectorNamespaceWithTenantRequest
type ConnectorNamespaceWithTenantRequest struct {
	// Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated.
	Name string `json:"name"`
	// Name-value string annotations for resource
	Annotations map[string]string        `json:"annotations,omitempty"`
	ClusterId   string                   `json:"cluster_id"`
	Tenant      ConnectorNamespaceTenant `json:"tenant"`
	// Namespace expiration timestamp in RFC 3339 format
	Expiration string `json:"expiration,omitempty"`
}
