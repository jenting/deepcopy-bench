package types

import "testing"

var from *EngineStatus = &EngineStatus{
	CurrentSize: int64(10000),
	CurrentReplicaAddressMap: map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	},
	ReplicaModeMap: map[string]ReplicaMode{
		"a": ReplicaModeRW,
		"b": ReplicaModeWO,
		"c": ReplicaModeERR,
	},
	BackupStatus: map[string]*BackupStatus{
		"a": {
			Progress:       100,
			BackupURL:      "s3://localhost",
			Error:          "error",
			SnapshotName:   "snapshotName",
			State:          "state",
			ReplicaAddress: "replicaAddress",
		},
		"b": {
			Progress:       100,
			BackupURL:      "s3://localhost",
			Error:          "error",
			SnapshotName:   "snapshotName",
			State:          "state",
			ReplicaAddress: "replicaAddress",
		},
		"c": {
			Progress:       100,
			BackupURL:      "s3://localhost",
			Error:          "error",
			SnapshotName:   "snapshotName",
			State:          "state",
			ReplicaAddress: "replicaAddress",
		},
	},
	RestoreStatus: map[string]*RestoreStatus{
		"a": {
			IsRestoring:            true,
			LastRestored:           "lastRestored",
			CurrentRestoringBackup: "currentRestoringBackup",
			Progress:               100,
			Error:                  "error",
			Filename:               "filename",
			State:                  "state",
			BackupURL:              "backupURL",
		},
		"b": {
			IsRestoring:            true,
			LastRestored:           "lastRestored",
			CurrentRestoringBackup: "currentRestoringBackup",
			Progress:               100,
			Error:                  "error",
			Filename:               "filename",
			State:                  "state",
			BackupURL:              "backupURL",
		},
		"c": {
			IsRestoring:            true,
			LastRestored:           "lastRestored",
			CurrentRestoringBackup: "currentRestoringBackup",
			Progress:               100,
			Error:                  "error",
			Filename:               "filename",
			State:                  "state",
			BackupURL:              "backupURL",
		},
	},
	PurgeStatus: map[string]*PurgeStatus{
		"a": {
			Error:     "error",
			IsPurging: true,
			Progress:  100,
			State:     "state",
		},
		"b": {
			Error:     "error",
			IsPurging: true,
			Progress:  100,
			State:     "state",
		},
		"c": {
			Error:     "error",
			IsPurging: true,
			Progress:  100,
			State:     "state",
		},
	},
	RebuildStatus: map[string]*RebuildStatus{
		"a": {
			Error:              "error",
			IsRebuilding:       true,
			Progress:           100,
			State:              "state",
			FromReplicaAddress: "fromReplicaAddress",
		},
		"b": {
			Error:              "error",
			IsRebuilding:       true,
			Progress:           100,
			State:              "state",
			FromReplicaAddress: "fromReplicaAddress",
		},
		"c": {
			Error:              "error",
			IsRebuilding:       true,
			Progress:           100,
			State:              "state",
			FromReplicaAddress: "fromReplicaAddress",
		},
	},
	CloneStatus: map[string]*SnapshotCloneStatus{
		"a": {
			IsCloning:          true,
			Error:              "error",
			Progress:           100,
			State:              "state",
			FromReplicaAddress: "fromReplicaAddress",
			SnapshotName:       "snapshotName",
		},
		"b": {
			IsCloning:          true,
			Error:              "error",
			Progress:           100,
			State:              "state",
			FromReplicaAddress: "fromReplicaAddress",
			SnapshotName:       "snapshotName",
		},
		"c": {
			IsCloning:          true,
			Error:              "error",
			Progress:           100,
			State:              "state",
			FromReplicaAddress: "fromReplicaAddress",
			SnapshotName:       "snapshotName",
		},
	},
	Snapshots: map[string]*Snapshot{
		"a": {
			Name:   "a",
			Parent: "head",
			Children: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
			Removed:     true,
			UserCreated: true,
			Created:     "created",
			Size:        "1024",
			Labels: map[string]string{
				"a": "a",
				"b": "b",
				"c": "c",
			},
		},
		"b": {
			Name:   "a",
			Parent: "head",
			Children: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
			Removed:     true,
			UserCreated: true,
			Created:     "created",
			Size:        "1024",
			Labels: map[string]string{
				"a": "a",
				"b": "b",
				"c": "c",
			},
		},
		"c": {
			Name:   "a",
			Parent: "head",
			Children: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
			Removed:     true,
			UserCreated: true,
			Created:     "created",
			Size:        "1024",
			Labels: map[string]string{
				"a": "a",
				"b": "b",
				"c": "c",
			},
		},
	},
}

func BenchmarkEngineStatusManual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		to := &EngineStatus{}
		from.DeepCopyIntoManual(to)
	}
}

func BenchmarkEngineStatusAuto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		to := &EngineStatus{}
		from.DeepCopyIntoAuto(to)
	}
}
