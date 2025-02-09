package rbac

// Action identifies an action a subject carries out on a resource for
// authorization purposes.
type Action int

const (
	WatchAction Action = iota
	CreateOrganizationAction
	UpdateOrganizationAction
	GetOrganizationAction
	ListOrganizationsAction
	GetEntitlementsAction
	DeleteOrganizationAction

	CreateVCSProviderAction
	GetVCSProviderAction
	ListVCSProvidersAction
	DeleteVCSProviderAction

	CreateAgentTokenAction
	ListAgentTokensAction
	DeleteAgentTokenAction

	CreateOrganizationTokenAction
	DeleteOrganizationTokenAction

	CreateRunTokenAction

	CreateModuleAction
	CreateModuleVersionAction
	UpdateModuleAction
	ListModulesAction
	GetModuleAction
	DeleteModuleAction
	DeleteModuleVersionAction

	CreateVariableAction
	UpdateVariableAction
	ListVariablesAction
	GetVariableAction
	DeleteVariableAction

	GetRunAction
	ListRunsAction
	ApplyRunAction
	CreateRunAction
	DiscardRunAction
	DeleteRunAction
	CancelRunAction
	EnqueuePlanAction
	StartPhaseAction
	FinishPhaseAction
	PutChunkAction
	TailLogsAction

	GetPlanFileAction
	UploadPlanFileAction

	GetLockFileAction
	UploadLockFileAction

	ListWorkspacesAction
	GetWorkspaceAction
	CreateWorkspaceAction
	DeleteWorkspaceAction
	SetWorkspacePermissionAction
	UnsetWorkspacePermissionAction
	UpdateWorkspaceAction

	ListTagsAction
	DeleteTagsAction
	TagWorkspacesAction
	AddTagsAction
	RemoveTagsAction
	ListWorkspaceTags

	LockWorkspaceAction
	UnlockWorkspaceAction
	ForceUnlockWorkspaceAction

	CreateStateVersionAction
	ListStateVersionsAction
	GetStateVersionAction
	DeleteStateVersionAction
	RollbackStateVersionAction
	DownloadStateAction
	GetStateVersionOutputAction

	CreateConfigurationVersionAction
	ListConfigurationVersionsAction
	GetConfigurationVersionAction
	DownloadConfigurationVersionAction
	DeleteConfigurationVersionAction

	CreateUserAction
	ListUsersAction
	GetUserAction
	DeleteUserAction

	CreateTeamAction
	UpdateTeamAction
	GetTeamAction
	ListTeamsAction
	DeleteTeamAction
	AddTeamMembershipAction
	RemoveTeamMembershipAction

	CreateNotificationConfigurationAction
	UpdateNotificationConfigurationAction
	ListNotificationConfigurationsAction
	GetNotificationConfigurationAction
	DeleteNotificationConfigurationAction
)
