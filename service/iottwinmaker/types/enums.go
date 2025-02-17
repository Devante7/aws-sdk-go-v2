// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ColumnType string

// Enum values for ColumnType
const (
	ColumnTypeNode  ColumnType = "NODE"
	ColumnTypeEdge  ColumnType = "EDGE"
	ColumnTypeValue ColumnType = "VALUE"
)

// Values returns all known values for ColumnType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ColumnType) Values() []ColumnType {
	return []ColumnType{
		"NODE",
		"EDGE",
		"VALUE",
	}
}

type ComponentUpdateType string

// Enum values for ComponentUpdateType
const (
	ComponentUpdateTypeCreate ComponentUpdateType = "CREATE"
	ComponentUpdateTypeUpdate ComponentUpdateType = "UPDATE"
	ComponentUpdateTypeDelete ComponentUpdateType = "DELETE"
)

// Values returns all known values for ComponentUpdateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ComponentUpdateType) Values() []ComponentUpdateType {
	return []ComponentUpdateType{
		"CREATE",
		"UPDATE",
		"DELETE",
	}
}

type DestinationType string

// Enum values for DestinationType
const (
	DestinationTypeS3           DestinationType = "s3"
	DestinationTypeIotsitewise  DestinationType = "iotsitewise"
	DestinationTypeIottwinmaker DestinationType = "iottwinmaker"
)

// Values returns all known values for DestinationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DestinationType) Values() []DestinationType {
	return []DestinationType{
		"s3",
		"iotsitewise",
		"iottwinmaker",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeValidationError           ErrorCode = "VALIDATION_ERROR"
	ErrorCodeInternalFailure           ErrorCode = "INTERNAL_FAILURE"
	ErrorCodeSyncInitializingError     ErrorCode = "SYNC_INITIALIZING_ERROR"
	ErrorCodeSyncCreatingError         ErrorCode = "SYNC_CREATING_ERROR"
	ErrorCodeSyncProcessingError       ErrorCode = "SYNC_PROCESSING_ERROR"
	ErrorCodeSyncDeletingError         ErrorCode = "SYNC_DELETING_ERROR"
	ErrorCodeProcessingError           ErrorCode = "PROCESSING_ERROR"
	ErrorCodeCompositeComponentFailure ErrorCode = "COMPOSITE_COMPONENT_FAILURE"
)

// Values returns all known values for ErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"VALIDATION_ERROR",
		"INTERNAL_FAILURE",
		"SYNC_INITIALIZING_ERROR",
		"SYNC_CREATING_ERROR",
		"SYNC_PROCESSING_ERROR",
		"SYNC_DELETING_ERROR",
		"PROCESSING_ERROR",
		"COMPOSITE_COMPONENT_FAILURE",
	}
}

type GroupType string

// Enum values for GroupType
const (
	GroupTypeTabular GroupType = "TABULAR"
)

// Values returns all known values for GroupType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GroupType) Values() []GroupType {
	return []GroupType{
		"TABULAR",
	}
}

type InterpolationType string

// Enum values for InterpolationType
const (
	InterpolationTypeLinear InterpolationType = "LINEAR"
)

// Values returns all known values for InterpolationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (InterpolationType) Values() []InterpolationType {
	return []InterpolationType{
		"LINEAR",
	}
}

type MetadataTransferJobState string

// Enum values for MetadataTransferJobState
const (
	MetadataTransferJobStateValidating MetadataTransferJobState = "VALIDATING"
	MetadataTransferJobStatePending    MetadataTransferJobState = "PENDING"
	MetadataTransferJobStateRunning    MetadataTransferJobState = "RUNNING"
	MetadataTransferJobStateCancelling MetadataTransferJobState = "CANCELLING"
	MetadataTransferJobStateError      MetadataTransferJobState = "ERROR"
	MetadataTransferJobStateCompleted  MetadataTransferJobState = "COMPLETED"
	MetadataTransferJobStateCancelled  MetadataTransferJobState = "CANCELLED"
)

// Values returns all known values for MetadataTransferJobState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MetadataTransferJobState) Values() []MetadataTransferJobState {
	return []MetadataTransferJobState{
		"VALIDATING",
		"PENDING",
		"RUNNING",
		"CANCELLING",
		"ERROR",
		"COMPLETED",
		"CANCELLED",
	}
}

type Order string

// Enum values for Order
const (
	OrderAscending  Order = "ASCENDING"
	OrderDescending Order = "DESCENDING"
)

// Values returns all known values for Order. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Order) Values() []Order {
	return []Order{
		"ASCENDING",
		"DESCENDING",
	}
}

type OrderByTime string

// Enum values for OrderByTime
const (
	OrderByTimeAscending  OrderByTime = "ASCENDING"
	OrderByTimeDescending OrderByTime = "DESCENDING"
)

// Values returns all known values for OrderByTime. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrderByTime) Values() []OrderByTime {
	return []OrderByTime{
		"ASCENDING",
		"DESCENDING",
	}
}

type ParentEntityUpdateType string

// Enum values for ParentEntityUpdateType
const (
	ParentEntityUpdateTypeUpdate ParentEntityUpdateType = "UPDATE"
	ParentEntityUpdateTypeDelete ParentEntityUpdateType = "DELETE"
)

// Values returns all known values for ParentEntityUpdateType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ParentEntityUpdateType) Values() []ParentEntityUpdateType {
	return []ParentEntityUpdateType{
		"UPDATE",
		"DELETE",
	}
}

type PricingMode string

// Enum values for PricingMode
const (
	PricingModeBasic        PricingMode = "BASIC"
	PricingModeStandard     PricingMode = "STANDARD"
	PricingModeTieredBundle PricingMode = "TIERED_BUNDLE"
)

// Values returns all known values for PricingMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PricingMode) Values() []PricingMode {
	return []PricingMode{
		"BASIC",
		"STANDARD",
		"TIERED_BUNDLE",
	}
}

type PricingTier string

// Enum values for PricingTier
const (
	PricingTierTier1 PricingTier = "TIER_1"
	PricingTierTier2 PricingTier = "TIER_2"
	PricingTierTier3 PricingTier = "TIER_3"
	PricingTierTier4 PricingTier = "TIER_4"
)

// Values returns all known values for PricingTier. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PricingTier) Values() []PricingTier {
	return []PricingTier{
		"TIER_1",
		"TIER_2",
		"TIER_3",
		"TIER_4",
	}
}

type PropertyGroupUpdateType string

// Enum values for PropertyGroupUpdateType
const (
	PropertyGroupUpdateTypeUpdate PropertyGroupUpdateType = "UPDATE"
	PropertyGroupUpdateTypeDelete PropertyGroupUpdateType = "DELETE"
	PropertyGroupUpdateTypeCreate PropertyGroupUpdateType = "CREATE"
)

// Values returns all known values for PropertyGroupUpdateType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PropertyGroupUpdateType) Values() []PropertyGroupUpdateType {
	return []PropertyGroupUpdateType{
		"UPDATE",
		"DELETE",
		"CREATE",
	}
}

type PropertyUpdateType string

// Enum values for PropertyUpdateType
const (
	PropertyUpdateTypeUpdate PropertyUpdateType = "UPDATE"
	PropertyUpdateTypeDelete PropertyUpdateType = "DELETE"
	PropertyUpdateTypeCreate PropertyUpdateType = "CREATE"
)

// Values returns all known values for PropertyUpdateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PropertyUpdateType) Values() []PropertyUpdateType {
	return []PropertyUpdateType{
		"UPDATE",
		"DELETE",
		"CREATE",
	}
}

type SceneErrorCode string

// Enum values for SceneErrorCode
const (
	SceneErrorCodeMatterportError SceneErrorCode = "MATTERPORT_ERROR"
)

// Values returns all known values for SceneErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SceneErrorCode) Values() []SceneErrorCode {
	return []SceneErrorCode{
		"MATTERPORT_ERROR",
	}
}

type Scope string

// Enum values for Scope
const (
	ScopeEntity    Scope = "ENTITY"
	ScopeWorkspace Scope = "WORKSPACE"
)

// Values returns all known values for Scope. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Scope) Values() []Scope {
	return []Scope{
		"ENTITY",
		"WORKSPACE",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeS3           SourceType = "s3"
	SourceTypeIotsitewise  SourceType = "iotsitewise"
	SourceTypeIottwinmaker SourceType = "iottwinmaker"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"s3",
		"iotsitewise",
		"iottwinmaker",
	}
}

type State string

// Enum values for State
const (
	StateCreating State = "CREATING"
	StateUpdating State = "UPDATING"
	StateDeleting State = "DELETING"
	StateActive   State = "ACTIVE"
	StateError    State = "ERROR"
)

// Values returns all known values for State. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"CREATING",
		"UPDATING",
		"DELETING",
		"ACTIVE",
		"ERROR",
	}
}

type SyncJobState string

// Enum values for SyncJobState
const (
	SyncJobStateCreating     SyncJobState = "CREATING"
	SyncJobStateInitializing SyncJobState = "INITIALIZING"
	SyncJobStateActive       SyncJobState = "ACTIVE"
	SyncJobStateDeleting     SyncJobState = "DELETING"
	SyncJobStateError        SyncJobState = "ERROR"
)

// Values returns all known values for SyncJobState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SyncJobState) Values() []SyncJobState {
	return []SyncJobState{
		"CREATING",
		"INITIALIZING",
		"ACTIVE",
		"DELETING",
		"ERROR",
	}
}

type SyncResourceState string

// Enum values for SyncResourceState
const (
	SyncResourceStateInitializing SyncResourceState = "INITIALIZING"
	SyncResourceStateProcessing   SyncResourceState = "PROCESSING"
	SyncResourceStateDeleted      SyncResourceState = "DELETED"
	SyncResourceStateInSync       SyncResourceState = "IN_SYNC"
	SyncResourceStateError        SyncResourceState = "ERROR"
)

// Values returns all known values for SyncResourceState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SyncResourceState) Values() []SyncResourceState {
	return []SyncResourceState{
		"INITIALIZING",
		"PROCESSING",
		"DELETED",
		"IN_SYNC",
		"ERROR",
	}
}

type SyncResourceType string

// Enum values for SyncResourceType
const (
	SyncResourceTypeEntity        SyncResourceType = "ENTITY"
	SyncResourceTypeComponentType SyncResourceType = "COMPONENT_TYPE"
)

// Values returns all known values for SyncResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SyncResourceType) Values() []SyncResourceType {
	return []SyncResourceType{
		"ENTITY",
		"COMPONENT_TYPE",
	}
}

type Type string

// Enum values for Type
const (
	TypeRelationship Type = "RELATIONSHIP"
	TypeString       Type = "STRING"
	TypeLong         Type = "LONG"
	TypeBoolean      Type = "BOOLEAN"
	TypeInteger      Type = "INTEGER"
	TypeDouble       Type = "DOUBLE"
	TypeList         Type = "LIST"
	TypeMap          Type = "MAP"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"RELATIONSHIP",
		"STRING",
		"LONG",
		"BOOLEAN",
		"INTEGER",
		"DOUBLE",
		"LIST",
		"MAP",
	}
}

type UpdateReason string

// Enum values for UpdateReason
const (
	UpdateReasonDefault           UpdateReason = "DEFAULT"
	UpdateReasonPricingTierUpdate UpdateReason = "PRICING_TIER_UPDATE"
	UpdateReasonEntityCountUpdate UpdateReason = "ENTITY_COUNT_UPDATE"
	UpdateReasonPricingModeUpdate UpdateReason = "PRICING_MODE_UPDATE"
	UpdateReasonOverwritten       UpdateReason = "OVERWRITTEN"
)

// Values returns all known values for UpdateReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UpdateReason) Values() []UpdateReason {
	return []UpdateReason{
		"DEFAULT",
		"PRICING_TIER_UPDATE",
		"ENTITY_COUNT_UPDATE",
		"PRICING_MODE_UPDATE",
		"OVERWRITTEN",
	}
}
