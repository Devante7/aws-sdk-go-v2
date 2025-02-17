// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AssistantCapabilityType string

// Enum values for AssistantCapabilityType
const (
	AssistantCapabilityTypeV1 AssistantCapabilityType = "V1"
	AssistantCapabilityTypeV2 AssistantCapabilityType = "V2"
)

// Values returns all known values for AssistantCapabilityType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssistantCapabilityType) Values() []AssistantCapabilityType {
	return []AssistantCapabilityType{
		"V1",
		"V2",
	}
}

type AssistantStatus string

// Enum values for AssistantStatus
const (
	AssistantStatusCreateInProgress AssistantStatus = "CREATE_IN_PROGRESS"
	AssistantStatusCreateFailed     AssistantStatus = "CREATE_FAILED"
	AssistantStatusActive           AssistantStatus = "ACTIVE"
	AssistantStatusDeleteInProgress AssistantStatus = "DELETE_IN_PROGRESS"
	AssistantStatusDeleteFailed     AssistantStatus = "DELETE_FAILED"
	AssistantStatusDeleted          AssistantStatus = "DELETED"
)

// Values returns all known values for AssistantStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssistantStatus) Values() []AssistantStatus {
	return []AssistantStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
	}
}

type AssistantType string

// Enum values for AssistantType
const (
	AssistantTypeAgent AssistantType = "AGENT"
)

// Values returns all known values for AssistantType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssistantType) Values() []AssistantType {
	return []AssistantType{
		"AGENT",
	}
}

type AssociationType string

// Enum values for AssociationType
const (
	AssociationTypeKnowledgeBase AssociationType = "KNOWLEDGE_BASE"
)

// Values returns all known values for AssociationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AssociationType) Values() []AssociationType {
	return []AssociationType{
		"KNOWLEDGE_BASE",
	}
}

type ContentStatus string

// Enum values for ContentStatus
const (
	ContentStatusCreateInProgress ContentStatus = "CREATE_IN_PROGRESS"
	ContentStatusCreateFailed     ContentStatus = "CREATE_FAILED"
	ContentStatusActive           ContentStatus = "ACTIVE"
	ContentStatusDeleteInProgress ContentStatus = "DELETE_IN_PROGRESS"
	ContentStatusDeleteFailed     ContentStatus = "DELETE_FAILED"
	ContentStatusDeleted          ContentStatus = "DELETED"
	ContentStatusUpdateFailed     ContentStatus = "UPDATE_FAILED"
)

// Values returns all known values for ContentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ContentStatus) Values() []ContentStatus {
	return []ContentStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
		"UPDATE_FAILED",
	}
}

type ExternalSource string

// Enum values for ExternalSource
const (
	ExternalSourceAmazonConnect ExternalSource = "AMAZON_CONNECT"
)

// Values returns all known values for ExternalSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExternalSource) Values() []ExternalSource {
	return []ExternalSource{
		"AMAZON_CONNECT",
	}
}

type FilterField string

// Enum values for FilterField
const (
	FilterFieldName FilterField = "NAME"
)

// Values returns all known values for FilterField. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FilterField) Values() []FilterField {
	return []FilterField{
		"NAME",
	}
}

type FilterOperator string

// Enum values for FilterOperator
const (
	FilterOperatorEquals FilterOperator = "EQUALS"
)

// Values returns all known values for FilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FilterOperator) Values() []FilterOperator {
	return []FilterOperator{
		"EQUALS",
	}
}

type ImportJobStatus string

// Enum values for ImportJobStatus
const (
	ImportJobStatusStartInProgress  ImportJobStatus = "START_IN_PROGRESS"
	ImportJobStatusFailed           ImportJobStatus = "FAILED"
	ImportJobStatusComplete         ImportJobStatus = "COMPLETE"
	ImportJobStatusDeleteInProgress ImportJobStatus = "DELETE_IN_PROGRESS"
	ImportJobStatusDeleteFailed     ImportJobStatus = "DELETE_FAILED"
	ImportJobStatusDeleted          ImportJobStatus = "DELETED"
)

// Values returns all known values for ImportJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImportJobStatus) Values() []ImportJobStatus {
	return []ImportJobStatus{
		"START_IN_PROGRESS",
		"FAILED",
		"COMPLETE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
	}
}

type ImportJobType string

// Enum values for ImportJobType
const (
	ImportJobTypeQuickResponses ImportJobType = "QUICK_RESPONSES"
)

// Values returns all known values for ImportJobType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImportJobType) Values() []ImportJobType {
	return []ImportJobType{
		"QUICK_RESPONSES",
	}
}

type KnowledgeBaseStatus string

// Enum values for KnowledgeBaseStatus
const (
	KnowledgeBaseStatusCreateInProgress KnowledgeBaseStatus = "CREATE_IN_PROGRESS"
	KnowledgeBaseStatusCreateFailed     KnowledgeBaseStatus = "CREATE_FAILED"
	KnowledgeBaseStatusActive           KnowledgeBaseStatus = "ACTIVE"
	KnowledgeBaseStatusDeleteInProgress KnowledgeBaseStatus = "DELETE_IN_PROGRESS"
	KnowledgeBaseStatusDeleteFailed     KnowledgeBaseStatus = "DELETE_FAILED"
	KnowledgeBaseStatusDeleted          KnowledgeBaseStatus = "DELETED"
)

// Values returns all known values for KnowledgeBaseStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseStatus) Values() []KnowledgeBaseStatus {
	return []KnowledgeBaseStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
	}
}

type KnowledgeBaseType string

// Enum values for KnowledgeBaseType
const (
	KnowledgeBaseTypeExternal       KnowledgeBaseType = "EXTERNAL"
	KnowledgeBaseTypeCustom         KnowledgeBaseType = "CUSTOM"
	KnowledgeBaseTypeQuickResponses KnowledgeBaseType = "QUICK_RESPONSES"
)

// Values returns all known values for KnowledgeBaseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseType) Values() []KnowledgeBaseType {
	return []KnowledgeBaseType{
		"EXTERNAL",
		"CUSTOM",
		"QUICK_RESPONSES",
	}
}

type Order string

// Enum values for Order
const (
	OrderAsc  Order = "ASC"
	OrderDesc Order = "DESC"
)

// Values returns all known values for Order. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Order) Values() []Order {
	return []Order{
		"ASC",
		"DESC",
	}
}

type Priority string

// Enum values for Priority
const (
	PriorityHigh   Priority = "HIGH"
	PriorityMedium Priority = "MEDIUM"
	PriorityLow    Priority = "LOW"
)

// Values returns all known values for Priority. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Priority) Values() []Priority {
	return []Priority{
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

type QueryConditionComparisonOperator string

// Enum values for QueryConditionComparisonOperator
const (
	QueryConditionComparisonOperatorEquals QueryConditionComparisonOperator = "EQUALS"
)

// Values returns all known values for QueryConditionComparisonOperator. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryConditionComparisonOperator) Values() []QueryConditionComparisonOperator {
	return []QueryConditionComparisonOperator{
		"EQUALS",
	}
}

type QueryConditionFieldName string

// Enum values for QueryConditionFieldName
const (
	QueryConditionFieldNameResultType QueryConditionFieldName = "RESULT_TYPE"
)

// Values returns all known values for QueryConditionFieldName. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryConditionFieldName) Values() []QueryConditionFieldName {
	return []QueryConditionFieldName{
		"RESULT_TYPE",
	}
}

type QueryResultType string

// Enum values for QueryResultType
const (
	QueryResultTypeKnowledgeContent QueryResultType = "KNOWLEDGE_CONTENT"
	QueryResultTypeGenerativeAnswer QueryResultType = "GENERATIVE_ANSWER"
)

// Values returns all known values for QueryResultType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QueryResultType) Values() []QueryResultType {
	return []QueryResultType{
		"KNOWLEDGE_CONTENT",
		"GENERATIVE_ANSWER",
	}
}

type QuickResponseFilterOperator string

// Enum values for QuickResponseFilterOperator
const (
	QuickResponseFilterOperatorEquals QuickResponseFilterOperator = "EQUALS"
	QuickResponseFilterOperatorPrefix QuickResponseFilterOperator = "PREFIX"
)

// Values returns all known values for QuickResponseFilterOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QuickResponseFilterOperator) Values() []QuickResponseFilterOperator {
	return []QuickResponseFilterOperator{
		"EQUALS",
		"PREFIX",
	}
}

type QuickResponseQueryOperator string

// Enum values for QuickResponseQueryOperator
const (
	QuickResponseQueryOperatorContains          QuickResponseQueryOperator = "CONTAINS"
	QuickResponseQueryOperatorContainsAndPrefix QuickResponseQueryOperator = "CONTAINS_AND_PREFIX"
)

// Values returns all known values for QuickResponseQueryOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QuickResponseQueryOperator) Values() []QuickResponseQueryOperator {
	return []QuickResponseQueryOperator{
		"CONTAINS",
		"CONTAINS_AND_PREFIX",
	}
}

type QuickResponseStatus string

// Enum values for QuickResponseStatus
const (
	QuickResponseStatusCreateInProgress QuickResponseStatus = "CREATE_IN_PROGRESS"
	QuickResponseStatusCreateFailed     QuickResponseStatus = "CREATE_FAILED"
	QuickResponseStatusCreated          QuickResponseStatus = "CREATED"
	QuickResponseStatusDeleteInProgress QuickResponseStatus = "DELETE_IN_PROGRESS"
	QuickResponseStatusDeleteFailed     QuickResponseStatus = "DELETE_FAILED"
	QuickResponseStatusDeleted          QuickResponseStatus = "DELETED"
	QuickResponseStatusUpdateInProgress QuickResponseStatus = "UPDATE_IN_PROGRESS"
	QuickResponseStatusUpdateFailed     QuickResponseStatus = "UPDATE_FAILED"
)

// Values returns all known values for QuickResponseStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (QuickResponseStatus) Values() []QuickResponseStatus {
	return []QuickResponseStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"CREATED",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED",
	}
}

type RecommendationSourceType string

// Enum values for RecommendationSourceType
const (
	RecommendationSourceTypeIssueDetection RecommendationSourceType = "ISSUE_DETECTION"
	RecommendationSourceTypeRuleEvaluation RecommendationSourceType = "RULE_EVALUATION"
	RecommendationSourceTypeOther          RecommendationSourceType = "OTHER"
)

// Values returns all known values for RecommendationSourceType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationSourceType) Values() []RecommendationSourceType {
	return []RecommendationSourceType{
		"ISSUE_DETECTION",
		"RULE_EVALUATION",
		"OTHER",
	}
}

type RecommendationTriggerType string

// Enum values for RecommendationTriggerType
const (
	RecommendationTriggerTypeQuery      RecommendationTriggerType = "QUERY"
	RecommendationTriggerTypeGenerative RecommendationTriggerType = "GENERATIVE"
)

// Values returns all known values for RecommendationTriggerType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationTriggerType) Values() []RecommendationTriggerType {
	return []RecommendationTriggerType{
		"QUERY",
		"GENERATIVE",
	}
}

type RecommendationType string

// Enum values for RecommendationType
const (
	RecommendationTypeKnowledgeContent   RecommendationType = "KNOWLEDGE_CONTENT"
	RecommendationTypeGenerativeResponse RecommendationType = "GENERATIVE_RESPONSE"
	RecommendationTypeGenerativeAnswer   RecommendationType = "GENERATIVE_ANSWER"
)

// Values returns all known values for RecommendationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationType) Values() []RecommendationType {
	return []RecommendationType{
		"KNOWLEDGE_CONTENT",
		"GENERATIVE_RESPONSE",
		"GENERATIVE_ANSWER",
	}
}

type Relevance string

// Enum values for Relevance
const (
	RelevanceHelpful    Relevance = "HELPFUL"
	RelevanceNotHelpful Relevance = "NOT_HELPFUL"
)

// Values returns all known values for Relevance. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Relevance) Values() []Relevance {
	return []Relevance{
		"HELPFUL",
		"NOT_HELPFUL",
	}
}

type RelevanceLevel string

// Enum values for RelevanceLevel
const (
	RelevanceLevelHigh   RelevanceLevel = "HIGH"
	RelevanceLevelMedium RelevanceLevel = "MEDIUM"
	RelevanceLevelLow    RelevanceLevel = "LOW"
)

// Values returns all known values for RelevanceLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RelevanceLevel) Values() []RelevanceLevel {
	return []RelevanceLevel{
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

type SourceContentType string

// Enum values for SourceContentType
const (
	SourceContentTypeKnowledgeContent SourceContentType = "KNOWLEDGE_CONTENT"
)

// Values returns all known values for SourceContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SourceContentType) Values() []SourceContentType {
	return []SourceContentType{
		"KNOWLEDGE_CONTENT",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeRecommendation TargetType = "RECOMMENDATION"
	TargetTypeResult         TargetType = "RESULT"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"RECOMMENDATION",
		"RESULT",
	}
}
