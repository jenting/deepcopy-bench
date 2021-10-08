package types

type InstanceState string

type ReplicaMode string

const (
	ReplicaModeRW  = ReplicaMode("RW")
	ReplicaModeWO  = ReplicaMode("WO")
	ReplicaModeERR = ReplicaMode("ERR")
)

type BackupStatus struct {
	Progress       int    `json:"progress"`
	BackupURL      string `json:"backupURL,omitempty"`
	Error          string `json:"error,omitempty"`
	SnapshotName   string `json:"snapshotName"`
	State          string `json:"state"`
	ReplicaAddress string `json:"replicaAddress"`
}

type RestoreStatus struct {
	IsRestoring            bool   `json:"isRestoring"`
	LastRestored           string `json:"lastRestored"`
	CurrentRestoringBackup string `json:"currentRestoringBackup"`
	Progress               int    `json:"progress,omitempty"`
	Error                  string `json:"error,omitempty"`
	Filename               string `json:"filename,omitempty"`
	State                  string `json:"state"`
	BackupURL              string `json:"backupURL"`
}

type PurgeStatus struct {
	Error     string `json:"error"`
	IsPurging bool   `json:"isPurging"`
	Progress  int    `json:"progress"`
	State     string `json:"state"`
}

type RebuildStatus struct {
	Error              string `json:"error"`
	IsRebuilding       bool   `json:"isRebuilding"`
	Progress           int    `json:"progress"`
	State              string `json:"state"`
	FromReplicaAddress string `json:"fromReplicaAddress"`
}

type SnapshotCloneStatus struct {
	IsCloning          bool   `json:"isCloning"`
	Error              string `json:"error"`
	Progress           int    `json:"progress"`
	State              string `json:"state"`
	FromReplicaAddress string `json:"fromReplicaAddress"`
	SnapshotName       string `json:"snapshotName"`
}

type InstanceStatus struct {
	OwnerID             string        `json:"ownerID"`
	InstanceManagerName string        `json:"instanceManagerName"`
	CurrentState        InstanceState `json:"currentState"`
	CurrentImage        string        `json:"currentImage"`
	IP                  string        `json:"ip"`
	Port                int           `json:"port"`
	Started             bool          `json:"started"`
	LogFetched          bool          `json:"logFetched"`
	SalvageExecuted     bool          `json:"salvageExecuted"`
}

type Snapshot struct {
	Name        string            `json:"name"`
	Parent      string            `json:"parent"`
	Children    map[string]bool   `json:"children"`
	Removed     bool              `json:"removed"`
	UserCreated bool              `json:"usercreated"`
	Created     string            `json:"created"`
	Size        string            `json:"size"`
	Labels      map[string]string `json:"labels"`
}

type EngineStatus struct {
	InstanceStatus
	CurrentSize              int64                           `json:"currentSize,string"`
	CurrentReplicaAddressMap map[string]string               `json:"currentReplicaAddressMap"`
	ReplicaModeMap           map[string]ReplicaMode          `json:"replicaModeMap"`
	Endpoint                 string                          `json:"endpoint"`
	LastRestoredBackup       string                          `json:"lastRestoredBackup"`
	BackupStatus             map[string]*BackupStatus        `json:"backupStatus"`
	RestoreStatus            map[string]*RestoreStatus       `json:"restoreStatus"`
	PurgeStatus              map[string]*PurgeStatus         `json:"purgeStatus"`
	RebuildStatus            map[string]*RebuildStatus       `json:"rebuildStatus"`
	CloneStatus              map[string]*SnapshotCloneStatus `json:"cloneStatus"`
	Snapshots                map[string]*Snapshot            `json:"snapshots"`
	SnapshotsError           string                          `json:"snapshotsError"`
	IsExpanding              bool                            `json:"isExpanding"`
	LastExpansionError       string                          `json:"lastExpansionError"`
	LastExpansionFailedAt    string                          `json:"lastExpansionFailedAt"`
}
