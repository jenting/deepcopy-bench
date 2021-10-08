package types

func (e *EngineStatus) DeepCopyIntoManual(to *EngineStatus) {
	*to = *e
	if e.CurrentReplicaAddressMap != nil {
		to.CurrentReplicaAddressMap = make(map[string]string)
		for key, value := range e.CurrentReplicaAddressMap {
			to.CurrentReplicaAddressMap[key] = value
		}
	}
	if e.ReplicaModeMap != nil {
		to.ReplicaModeMap = make(map[string]ReplicaMode)
		for key, value := range e.ReplicaModeMap {
			to.ReplicaModeMap[key] = value
		}
	}
	if e.BackupStatus != nil {
		to.BackupStatus = make(map[string]*BackupStatus)
		for key, value := range e.BackupStatus {
			to.BackupStatus[key] = &BackupStatus{}
			*to.BackupStatus[key] = *value
		}
	}
	if e.RestoreStatus != nil {
		to.RestoreStatus = make(map[string]*RestoreStatus)
		for key, value := range e.RestoreStatus {
			to.RestoreStatus[key] = &RestoreStatus{}
			*to.RestoreStatus[key] = *value
		}
	}
	if e.PurgeStatus != nil {
		to.PurgeStatus = make(map[string]*PurgeStatus)
		for key, value := range e.PurgeStatus {
			to.PurgeStatus[key] = &PurgeStatus{}
			*to.PurgeStatus[key] = *value
		}
	}
	if e.RebuildStatus != nil {
		to.RebuildStatus = make(map[string]*RebuildStatus)
		for key, value := range e.RebuildStatus {
			to.RebuildStatus[key] = &RebuildStatus{}
			*to.RebuildStatus[key] = *value
		}
	}
	if e.CloneStatus != nil {
		to.CloneStatus = make(map[string]*SnapshotCloneStatus)
		for key, value := range e.CloneStatus {
			to.CloneStatus[key] = &SnapshotCloneStatus{}
			*to.CloneStatus[key] = *value
		}
	}
	if e.Snapshots != nil {
		to.Snapshots = make(map[string]*Snapshot)
		for key, source := range e.Snapshots {
			to.Snapshots[key] = &Snapshot{}
			*to.Snapshots[key] = *source
			out := to.Snapshots[key]

			if source.Children != nil {
				out.Children = make(map[string]bool)
				for key, value := range source.Children {
					out.Children[key] = value
				}
			}

			if source.Labels != nil {
				out.Labels = make(map[string]string)
				for key, value := range source.Labels {
					out.Labels[key] = value
				}
			}
		}
	}
}
