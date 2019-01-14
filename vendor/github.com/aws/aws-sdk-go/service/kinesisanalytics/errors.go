// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

const (

	// ErrCodeCodeValidationException for service response error code
	// "CodeValidationException".
	//
	// User-provided application code (query) is invalid. This can be a simple syntax
	// error.
	ErrCodeCodeValidationException = "CodeValidationException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Exception thrown as a result of concurrent modification to an application.
	// For example, two individuals attempting to edit the same application at the
	// same time.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeInvalidApplicationConfigurationException for service response error code
	// "InvalidApplicationConfigurationException".
	//
	// User-provided application configuration is not valid.
	ErrCodeInvalidApplicationConfigurationException = "InvalidApplicationConfigurationException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// Specified input parameter value is invalid.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Exceeded the number of applications allowed.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// Application is not available for this operation.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Specified application can't be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceProvisionedThroughputExceededException for service response error code
	// "ResourceProvisionedThroughputExceededException".
	//
	// Discovery failed to get a record from the streaming source because of the
	// Amazon Kinesis Streams ProvisionedThroughputExceededException. For more information,
	// see GetRecords (http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetRecords.html)
	// in the Amazon Kinesis Streams API Reference.
	ErrCodeResourceProvisionedThroughputExceededException = "ResourceProvisionedThroughputExceededException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is unavailable. Back off and retry the operation.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeUnableToDetectSchemaException for service response error code
	// "UnableToDetectSchemaException".
	//
	// Data format is not valid. Amazon Kinesis Analytics is not able to detect
	// schema for the given streaming source.
	ErrCodeUnableToDetectSchemaException = "UnableToDetectSchemaException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"
)