package presenters

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/public"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
)

func PresentSupportedKafkaInstanceType(supportedInstanceType *api.SupportedKafkaInstanceType) public.SupportedKafkaInstanceType {
	reference := PresentReference(supportedInstanceType.Id, supportedInstanceType)
	return public.SupportedKafkaInstanceType{
		Kind:  reference.Kind,
		Id:    reference.Id,
		Sizes: GetSupportedKafkaSizes(supportedInstanceType.SupportedKafkaSizes),
	}
}

func GetSupportedKafkaSizes(supportedKafkaSizes []api.SupportedKafkaSize) []public.SupportedKafkaSize {
	sizes := make([]public.SupportedKafkaSize, 0)
	for _, c := range supportedKafkaSizes {
		sizes = append(sizes, public.SupportedKafkaSize{Kind: c.Kind, Id: c.Id, IngressThroughputPerSec: c.IngressThroughputPerSec, EgressThroughputPerSec: c.EgressThroughputPerSec, TotalMaxConnections: c.TotalMaxConnections, MaxDataRetentionSize: c.MaxDataRetentionSize, MaxPartitions: c.MaxPartitions, MaxDataRetentionPeriod: c.MaxDataRetentionPeriod, MaxConnectionAttemptsPerSec: c.MaxConnectionAttemptsPerSec, QuotaConsumed: c.QuotaConsumed, QuotaType: c.QuotaType, CapacityConsumed: c.CapacityConsumed})
	}
	return sizes
}
