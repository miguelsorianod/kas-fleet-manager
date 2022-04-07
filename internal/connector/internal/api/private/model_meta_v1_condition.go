/*
 * Connector Service Fleet Manager Private APIs
 *
 * Connector Service Fleet Manager apis that are used by internal services.
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// MetaV1Condition struct for MetaV1Condition
type MetaV1Condition struct {
	Type               string `json:"type,omitempty"`
	Reason             string `json:"reason,omitempty"`
	Message            string `json:"message,omitempty"`
	Status             string `json:"status,omitempty"`
	LastTransitionTime string `json:"last_transition_time,omitempty"`
}
