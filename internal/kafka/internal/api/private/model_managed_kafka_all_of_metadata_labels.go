/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager APIs that are used by internal services e.g kas-fleetshard operators.
 *
 * API version: 1.5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ManagedKafkaAllOfMetadataLabels struct for ManagedKafkaAllOfMetadataLabels
type ManagedKafkaAllOfMetadataLabels struct {
	Bf2OrgKafkaInstanceProfileType          string `json:"bf2.org/kafkaInstanceProfileType"`
	Bf2OrgKafkaInstanceProfileQuotaConsumed string `json:"bf2.org/kafkaInstanceProfileQuotaConsumed"`
}
