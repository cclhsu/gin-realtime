package model

// Comment represents the Comment message.
type Comment struct {
	ID		  string `json:"ID"`
	UUID	  string `json:"UUID"`
	Content	  string `json:"content"`
	UpdatedBy string `json:"updatedBy"`
	UpdatedAt string `json:"updatedAt"`
}

// CommonDate represents the CommonDate message.
type CommonDate struct {
	CreatedAt	string `json:"createdAt"`
	CreatedBy	string `json:"createdBy"`
	UpdatedAt	string `json:"updatedAt"`
	UpdatedBy	string `json:"updatedBy"`
	StartedAt	string `json:"startedAt,omitempty"`
	StartedBy	string `json:"startedBy,omitempty"`
	StartDate	string `json:"startDate,omitempty"`
	EndDate		string `json:"endDate,omitempty"`
	CompletedAt string `json:"completedAt,omitempty"`
	CompletedBy string `json:"completedBy,omitempty"`
}

// Duration represents the Duration message.
type Duration struct {
	ID		  string `json:"ID"`
	UUID	  string `json:"UUID"`
	StartDate string `json:"startDate"`
	EndDate	  string `json:"endDate"`
}

// IdUuidStatus represents the IdUuidStatus message.
type IdUuidStatus struct {
	ID	   string				`json:"ID"`
	UUID   string				`json:"UUID"`
	Status GENERAL_STATUS_TYPES `json:"status"`
}

// IdUuid represents the IdUuid message.
type IdUuid struct {
	ID	 string `json:"ID"`
	UUID string `json:"UUID"`
}

// NameUrl represents the NameUrl message.
type NameUrl struct {
	Name string `json:"name"`
	URL	 string `json:"url"`
}

// Pagination represents the Pagination message.
type Pagination struct {
	Skip int32 `json:"skip"`
	Take int32 `json:"take"`
}

// Relation represents the Relation message.
type Relation struct {
	UUID		 string			`json:"UUID"`
	RelationType RELATION_TYPES `json:"relationType"`
	SourceUUID	 string			`json:"sourceUUID"`
	TargetUUID	 string			`json:"targetUUID"`
	CreatedAt	 string			`json:"createdAt"`
	UpdatedAt	 string			`json:"updatedAt"`
}

// GENERAL_STATUS_TYPES represents the enumeration for general status types.
type GENERAL_STATUS_TYPES int32

const (
	GENERAL_STATUS_TYPES_UNSPECIFIED GENERAL_STATUS_TYPES = 0
	GENERAL_STATUS_TYPES_INACTIVE	 GENERAL_STATUS_TYPES = 1
	GENERAL_STATUS_TYPES_ACTIVE		 GENERAL_STATUS_TYPES = 2
	GENERAL_STATUS_TYPES_PLANNED	 GENERAL_STATUS_TYPES = 3
	GENERAL_STATUS_TYPES_TODO		 GENERAL_STATUS_TYPES = 4
	GENERAL_STATUS_TYPES_IN_PROGRESS GENERAL_STATUS_TYPES = 5
	GENERAL_STATUS_TYPES_DONE		 GENERAL_STATUS_TYPES = 6
	GENERAL_STATUS_TYPES_COMPLETED	 GENERAL_STATUS_TYPES = 7
	GENERAL_STATUS_TYPES_CANCELLED	 GENERAL_STATUS_TYPES = 8
)

// PROJECT_ROLE_TYPES represents the enumeration for project role types.
type PROJECT_ROLE_TYPES int32

const (
	PROJECT_ROLE_TYPES_UNSPECIFIED PROJECT_ROLE_TYPES = 0
	PROJECT_ROLE_TYPES_PM		   PROJECT_ROLE_TYPES = 1
	PROJECT_ROLE_TYPES_EM		   PROJECT_ROLE_TYPES = 2
	PROJECT_ROLE_TYPES_DEV		   PROJECT_ROLE_TYPES = 3
	PROJECT_ROLE_TYPES_QA		   PROJECT_ROLE_TYPES = 4
	PROJECT_ROLE_TYPES_BA		   PROJECT_ROLE_TYPES = 5
	PROJECT_ROLE_TYPES_UX		   PROJECT_ROLE_TYPES = 6
	PROJECT_ROLE_TYPES_O		   PROJECT_ROLE_TYPES = 7
)

// RELATION_TYPES represents the enumeration for relation types.
type RELATION_TYPES int32

const (
	RELATION_TYPES_UNSPECIFIED	RELATION_TYPES = 0
	RELATION_TYPES_PARENT		RELATION_TYPES = 1
	RELATION_TYPES_SUBTASKS		RELATION_TYPES = 2
	RELATION_TYPES_PREDECESSORS RELATION_TYPES = 3
	RELATION_TYPES_SUCCESSORS	RELATION_TYPES = 4
	RELATION_TYPES_RELATES_TO	RELATION_TYPES = 5
	RELATION_TYPES_BLOCKED_BY	RELATION_TYPES = 6
)

// SCRUM_ROLE_TYPES represents the enumeration for scrum role types.
type SCRUM_ROLE_TYPES int32

const (
	SCRUM_ROLE_TYPES_UNSPECIFIED SCRUM_ROLE_TYPES = 0
	SCRUM_ROLE_TYPES_PO			 SCRUM_ROLE_TYPES = 1
	SCRUM_ROLE_TYPES_SM			 SCRUM_ROLE_TYPES = 2
	SCRUM_ROLE_TYPES_MEMBER		 SCRUM_ROLE_TYPES = 3
	SCRUM_ROLE_TYPES_O			 SCRUM_ROLE_TYPES = 4
)
