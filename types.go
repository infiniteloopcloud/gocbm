package gocbm

type IndexStatsResponse struct {
	PIndexes map[string]PIndex `json:"pindexes"`
	Feeds    map[string]Feed   `json:"feeds"`
}

type Feed struct {
	AgentDCPStats AgentDCPStats `json:"agentDCPStats"`
	DestStats     DestStats     `json:"destStats"`
}

type AgentDCPStats struct {
	TotDCPStreamReqs        int `json:"TotDCPStreamReqs"`
	TotDCPStreamEnds        int `json:"TotDCPStreamEnds"`
	TotDCPRollbacks         int `json:"TotDCPRollbacks"`
	TotDCPSnapshotMarkers   int `json:"TotDCPSnapshotMarkers"`
	TotDCPMutations         int `json:"TotDCPMutations"`
	TotDCPDeletions         int `json:"TotDCPDeletions"`
	TotDCPSeqNoAdvanceds    int `json:"TotDCPSeqNoAdvanceds"`
	TotDCPCreateCollections int `json:"TotDCPCreateCollections"`
}

type DestStats struct {
	TotError              int       `json:"TotError"`
	TimerDataUpdate       TimerData `json:"TimerDataUpdate"`
	TimerDataDelete       TimerData `json:"TimerDataDelete"`
	TimerSnapshotStart    TimerData `json:"TimerSnapshotStart"`
	TimerOpaqueGet        TimerData `json:"TimerOpaqueGet"`
	TimerOpaqueSet        TimerData `json:"TimerOpaqueSet"`
	TimerRollback         TimerData `json:"TimerRollback"`
	TimerCreateCollection TimerData `json:"TimerCreateCollection"`
	TimerSeqNoAdvanced    TimerData `json:"TimerSeqNoAdvanced"`
}

type Rates struct {
	Min  float64 `json:"1-min"`
	Min1 float64 `json:"5-min"`
	Min2 float64 `json:"15-min"`
	Mean float64 `json:"mean"`
}

type PIndex struct {
	PindexStoreStats   PindexStoreStats     `json:"pindexStoreStats"`
	BleveIndexStats    BleveIndexStats      `json:"bleveIndexStats"`
	Basic              PIndexBasic          `json:"basic"`
	Partitions         map[string]Partition `json:"partitions"`
	CopyPartitionStats CopyPartitionStats   `json:"copyPartitionStats"`
}

type PIndexBasic struct {
	DocCount int64 `json:"DocCount"`
}

type PindexStoreStats struct {
	TimerBatchStore TimerData     `json:"TimerBatchStore"`
	Errors          []interface{} `json:"Errors"`
}

type BleveIndexStats struct {
	Index      BleveIndex `json:"index"`
	SearchTime int        `json:"search_time"`
	Searches   int        `json:"searches"`
}

type BleveIndex struct {
	CurFilesIneligibleForRemoval       int   `json:"CurFilesIneligibleForRemoval"`
	CurOnDiskBytes                     int   `json:"CurOnDiskBytes"`
	CurOnDiskFiles                     int   `json:"CurOnDiskFiles"`
	CurRootEpoch                       int   `json:"CurRootEpoch"`
	LastMergedEpoch                    int   `json:"LastMergedEpoch"`
	LastPersistedEpoch                 int   `json:"LastPersistedEpoch"`
	MaxBatchIntroTime                  int   `json:"MaxBatchIntroTime"`
	MaxFileMergeZapIntroductionTime    int   `json:"MaxFileMergeZapIntroductionTime"`
	MaxFileMergeZapTime                int   `json:"MaxFileMergeZapTime"`
	MaxMemMergeZapTime                 int   `json:"MaxMemMergeZapTime"`
	TotAnalysisTime                    int   `json:"TotAnalysisTime"`
	TotBatchIntroTime                  int   `json:"TotBatchIntroTime"`
	TotBatches                         int   `json:"TotBatches"`
	TotBatchesEmpty                    int   `json:"TotBatchesEmpty"`
	TotDeletes                         int   `json:"TotDeletes"`
	TotEventTriggerCompleted           int   `json:"TotEventTriggerCompleted"`
	TotEventTriggerStarted             int   `json:"TotEventTriggerStarted"`
	TotFileMergeForceOpsCompleted      int   `json:"TotFileMergeForceOpsCompleted"`
	TotFileMergeForceOpsStarted        int   `json:"TotFileMergeForceOpsStarted"`
	TotFileMergeIntroductions          int   `json:"TotFileMergeIntroductions"`
	TotFileMergeIntroductionsDone      int   `json:"TotFileMergeIntroductionsDone"`
	TotFileMergeIntroductionsObsoleted int   `json:"TotFileMergeIntroductionsObsoleted"`
	TotFileMergeIntroductionsSkipped   int   `json:"TotFileMergeIntroductionsSkipped"`
	TotFileMergeLoopBeg                int   `json:"TotFileMergeLoopBeg"`
	TotFileMergeLoopEnd                int   `json:"TotFileMergeLoopEnd"`
	TotFileMergeLoopErr                int   `json:"TotFileMergeLoopErr"`
	TotFileMergePlan                   int   `json:"TotFileMergePlan"`
	TotFileMergePlanErr                int   `json:"TotFileMergePlanErr"`
	TotFileMergePlanNone               int   `json:"TotFileMergePlanNone"`
	TotFileMergePlanOk                 int   `json:"TotFileMergePlanOk"`
	TotFileMergePlanTasks              int   `json:"TotFileMergePlanTasks"`
	TotFileMergePlanTasksDone          int   `json:"TotFileMergePlanTasksDone"`
	TotFileMergePlanTasksErr           int   `json:"TotFileMergePlanTasksErr"`
	TotFileMergePlanTasksSegments      int   `json:"TotFileMergePlanTasksSegments"`
	TotFileMergePlanTasksSegmentsEmpty int   `json:"TotFileMergePlanTasksSegmentsEmpty"`
	TotFileMergeSegments               int   `json:"TotFileMergeSegments"`
	TotFileMergeSegmentsEmpty          int   `json:"TotFileMergeSegmentsEmpty"`
	TotFileMergeWrittenBytes           int   `json:"TotFileMergeWrittenBytes"`
	TotFileMergeZapBeg                 int   `json:"TotFileMergeZapBeg"`
	TotFileMergeZapEnd                 int   `json:"TotFileMergeZapEnd"`
	TotFileMergeZapIntroductionTime    int   `json:"TotFileMergeZapIntroductionTime"`
	TotFileMergeZapTime                int64 `json:"TotFileMergeZapTime"`
	TotFileSegmentsAtRoot              int   `json:"TotFileSegmentsAtRoot"`
	TotIndexTime                       int64 `json:"TotIndexTime"`
	TotIndexedPlainTextBytes           int   `json:"TotIndexedPlainTextBytes"`
	TotIntroduceLoop                   int   `json:"TotIntroduceLoop"`
	TotIntroduceMergeBeg               int   `json:"TotIntroduceMergeBeg"`
	TotIntroduceMergeEnd               int   `json:"TotIntroduceMergeEnd"`
	TotIntroducePersistBeg             int   `json:"TotIntroducePersistBeg"`
	TotIntroducePersistEnd             int   `json:"TotIntroducePersistEnd"`
	TotIntroduceRevertBeg              int   `json:"TotIntroduceRevertBeg"`
	TotIntroduceRevertEnd              int   `json:"TotIntroduceRevertEnd"`
	TotIntroduceSegmentBeg             int   `json:"TotIntroduceSegmentBeg"`
	TotIntroduceSegmentEnd             int   `json:"TotIntroduceSegmentEnd"`
	TotIntroducedItems                 int   `json:"TotIntroducedItems"`
	TotIntroducedSegmentsBatch         int   `json:"TotIntroducedSegmentsBatch"`
	TotIntroducedSegmentsMerge         int   `json:"TotIntroducedSegmentsMerge"`
	TotItemsToPersist                  int   `json:"TotItemsToPersist"`
	TotMemMergeBeg                     int   `json:"TotMemMergeBeg"`
	TotMemMergeDone                    int   `json:"TotMemMergeDone"`
	TotMemMergeErr                     int   `json:"TotMemMergeErr"`
	TotMemMergeSegments                int   `json:"TotMemMergeSegments"`
	TotMemMergeZapBeg                  int   `json:"TotMemMergeZapBeg"`
	TotMemMergeZapEnd                  int   `json:"TotMemMergeZapEnd"`
	TotMemMergeZapTime                 int   `json:"TotMemMergeZapTime"`
	TotMemorySegmentsAtRoot            int   `json:"TotMemorySegmentsAtRoot"`
	TotOnErrors                        int   `json:"TotOnErrors"`
	TotPersistLoopBeg                  int   `json:"TotPersistLoopBeg"`
	TotPersistLoopEnd                  int   `json:"TotPersistLoopEnd"`
	TotPersistLoopErr                  int   `json:"TotPersistLoopErr"`
	TotPersistLoopProgress             int   `json:"TotPersistLoopProgress"`
	TotPersistLoopWait                 int   `json:"TotPersistLoopWait"`
	TotPersistLoopWaitNotified         int   `json:"TotPersistLoopWaitNotified"`
	TotPersistedItems                  int   `json:"TotPersistedItems"`
	TotPersistedSegments               int   `json:"TotPersistedSegments"`
	TotPersisterMergerNapBreak         int   `json:"TotPersisterMergerNapBreak"`
	TotPersisterNapPauseCompleted      int   `json:"TotPersisterNapPauseCompleted"`
	TotPersisterSlowMergerPause        int   `json:"TotPersisterSlowMergerPause"`
	TotPersisterSlowMergerResume       int   `json:"TotPersisterSlowMergerResume"`
	TotSnapshotsRemovedFromMetaStore   int   `json:"TotSnapshotsRemovedFromMetaStore"`
	TotTermSearchersFinished           int   `json:"TotTermSearchersFinished"`
	TotTermSearchersStarted            int   `json:"TotTermSearchersStarted"`
	TotUpdates                         int   `json:"TotUpdates"`
	AnalysisTime                       int   `json:"analysis_time"`
	Batches                            int   `json:"batches"`
	Deletes                            int   `json:"deletes"`
	Errors                             int   `json:"errors"`
	IndexTime                          int64 `json:"index_time"`
	NumBytesUsedDisk                   int   `json:"num_bytes_used_disk"`
	NumBytesUsedDiskByRoot             int   `json:"num_bytes_used_disk_by_root"`
	NumBytesUsedDiskByRootReclaimable  int   `json:"num_bytes_used_disk_by_root_reclaimable"`
	NumFilesOnDisk                     int   `json:"num_files_on_disk"`
	NumItemsIntroduced                 int   `json:"num_items_introduced"`
	NumItemsPersisted                  int   `json:"num_items_persisted"`
	NumPersisterNapMergerBreak         int   `json:"num_persister_nap_merger_break"`
	NumPersisterNapPauseCompleted      int   `json:"num_persister_nap_pause_completed"`
	NumPlainTextBytesIndexed           int   `json:"num_plain_text_bytes_indexed"`
	NumRecsToPersist                   int   `json:"num_recs_to_persist"`
	NumRootFilesegments                int   `json:"num_root_filesegments"`
	NumRootMemorysegments              int   `json:"num_root_memorysegments"`
	TermSearchersFinished              int   `json:"term_searchers_finished"`
	TermSearchersStarted               int   `json:"term_searchers_started"`
	TotalCompactionWrittenBytes        int   `json:"total_compaction_written_bytes"`
	Updates                            int   `json:"updates"`
}

type Partition struct {
	Seq         int    `json:"seq"`
	SeqReceived int    `json:"seqReceived"`
	SourceSeq   int    `json:"sourceSeq"`
	Uuid        string `json:"uuid"`
}

type CopyPartitionStats struct {
	TotCopyPartitionStart     int `json:"TotCopyPartitionStart"`
	TotCopyPartitionFinished  int `json:"TotCopyPartitionFinished"`
	TotCopyPartitionTimeInMs  int `json:"TotCopyPartitionTimeInMs"`
	TotCopyPartitionFailed    int `json:"TotCopyPartitionFailed"`
	TotCopyPartitionRetries   int `json:"TotCopyPartitionRetries"`
	TotCopyPartitionErrors    int `json:"TotCopyPartitionErrors"`
	TotCopyPartitionSkipped   int `json:"TotCopyPartitionSkipped"`
	TotCopyPartitionCancelled int `json:"TotCopyPartitionCancelled"`
	TotCopyPartitionOnHttp2   int `json:"TotCopyPartitionOnHttp2"`
}

type TimerData struct {
	Count       int         `json:"count"`
	Min         int         `json:"min"`
	Max         int         `json:"max"`
	Mean        float64     `json:"mean"`
	Stddev      float64     `json:"stddev"`
	Percentiles Percentiles `json:"percentiles"`
	Rates       Rates       `json:"rates"`
}

type Percentiles struct {
	Median float64 `json:"median"`
	Field2 float64 `json:"75%"`
	Field3 float64 `json:"95%"`
	Field4 float64 `json:"99%"`
	Field5 float64 `json:"99.9%"`
}
