/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectorClusterUpdateStatusNodeInfo struct for ConnectorClusterUpdateStatusNodeInfo
type ConnectorClusterUpdateStatusNodeInfo struct {
	Ceiling                int32 `json:"ceiling,omitempty"`
	Floor                  int32 `json:"floor,omitempty"`
	Current                int32 `json:"current,omitempty"`
	CurrentWorkLoadMinimum int32 `json:"currentWorkLoadMinimum,omitempty"`
}