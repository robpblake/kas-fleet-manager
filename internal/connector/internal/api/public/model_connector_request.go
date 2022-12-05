/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// ConnectorRequest struct for ConnectorRequest
type ConnectorRequest struct {
	Name            string                `json:"name"`
	ConnectorTypeId string                `json:"connector_type_id"`
	NamespaceId     string                `json:"namespace_id"`
	Channel         Channel               `json:"channel,omitempty"`
	DesiredState    ConnectorDesiredState `json:"desired_state"`
	// Name-value string annotations for resource
	Annotations    map[string]string                `json:"annotations,omitempty"`
	Kafka          KafkaConnectionSettings          `json:"kafka"`
	ServiceAccount ServiceAccount                   `json:"service_account"`
	SchemaRegistry SchemaRegistryConnectionSettings `json:"schema_registry,omitempty"`
	Connector      map[string]interface{}           `json:"connector"`
}
