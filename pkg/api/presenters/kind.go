package presenters

import (
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/api"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/errors"
)

func ObjectKind(i interface{}) string {
	switch i.(type) {
	case api.KafkaRequest, *api.KafkaRequest:
		return "Kafka"
	case errors.ServiceError, *errors.ServiceError:
		return "Error"
	default:
		return ""
	}
}
