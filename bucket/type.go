package bucket

import (
	"time"
)

type BucketListParameters struct {
	// Namespace for which buckets should be listed.
	Namespace string

	// reference to last object returned.
	Marker string

	// number of objects requested in current fetch
	Limit int

	// Case sensitive prefix of the Bucket name with a wild card(*) Ex : any_prefix_string*
	Name string
}

type Bucket struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
	Link struct {
		Rel  string `json:"rel,omitempty"`
		Href string `json:"href,omitempty"`
	} `json:"link,omitempty"`
	Namespace                  string    `json:"namespace,omitempty"`
	Vpool                      string    `json:"vpool,omitempty"`
	Locked                     bool      `json:"locked,omitempty"`
	FsAccessEnabled            bool      `json:"fs_access_enabled,omitempty"`
	Softquota                  string    `json:"softquota,omitempty"`
	Created                    time.Time `json:"created,omitempty"`
	IsStaleAllowed             bool      `json:"is_stale_allowed,omitempty"`
	IsObjectLockWithAdoAllowed bool      `json:"is_object_lock_with_ado_allowed,omitempty"`
	IsTsoReadOnly              bool      `json:"is_tso_read_only,omitempty"`
	IsObjectLockEnabled        bool      `json:"is_object_lock_enabled,omitempty"`
	DefaultRetention           int       `json:"default_retention,omitempty"`
	BlockSize                  int       `json:"block_size,omitempty"`
	AutoCommitPeriod           int       `json:"auto_commit_period,omitempty"`
	NotificationSize           int       `json:"notification_size,omitempty"`
	BlockSizeInCount           int       `json:"blockSizeInCount,omitempty"`
	NotificationSizeInCount    int       `json:"notificationSizeInCount,omitempty"`
	IsEncryptionEnabled        string    `json:"is_encryption_enabled,omitempty"`
	TagSet                     []struct {
		Key   string `json:"Key,omitempty"`
		Value string `json:"Value,omitempty"`
	} `json:"TagSet,omitempty"`
	Retention                         int    `json:"retention,omitempty"`
	DefaultGroup                      string `json:"default_group,omitempty"`
	DefaultGroupFileReadPermission    bool   `json:"default_group_file_read_permission,omitempty"`
	DefaultGroupFileWritePermission   bool   `json:"default_group_file_write_permission,omitempty"`
	DefaultGroupFileExecutePermission bool   `json:"default_group_file_execute_permission,omitempty"`
	DefaultGroupDirReadPermission     bool   `json:"default_group_dir_read_permission,omitempty"`
	DefaultGroupDirWritePermission    bool   `json:"default_group_dir_write_permission,omitempty"`
	DefaultGroupDirExecutePermission  bool   `json:"default_group_dir_execute_permission,omitempty"`
	MinMaxGovernor                    struct {
		EnforceRetention         bool `json:"enforce_retention,omitempty"`
		MinimumFixedRetention    int  `json:"minimum_fixed_retention,omitempty"`
		MaximumFixedRetention    int  `json:"maximum_fixed_retention,omitempty"`
		MinimumVariableRetention int  `json:"minimum_variable_retention,omitempty"`
		MaximumVariableRetention int  `json:"maximum_variable_retention,omitempty"`
	} `json:"min_max_governor,omitempty"`
	AuditDeleteExpiration        int  `json:"audit_delete_expiration,omitempty"`
	EnableAdvancedMetadataSearch bool `json:"enableAdvancedMetadataSearch,omitempty"`
	IsEmptyBucketInProgress      bool `json:"is_empty_bucket_in_progress,omitempty"`
	SearchMetadata               struct {
		IsEnabled bool `json:"isEnabled,omitempty"`
		Metadata  []struct {
			Type     string `json:"type,omitempty"`
			Name     string `json:"name,omitempty"`
			DataType string `json:"datatype,omitempty"`
		} `json:"metadata,omitempty"`
		MaxKeys        int  `json:"maxKeys,omitempty"`
		MetadataTokens bool `json:"metadata_tokens,omitempty"`
	} `json:"search_metadata,omitempty"`
	APIType                  string `json:"api_type,omitempty"`
	LocalObjectMetadataReads bool   `json:"local_object_metadata_reads,omitempty"`
	Owner                    string `json:"owner,omitempty"`
}

type BucketListResp struct {
	Buckets      []*Bucket `json:"object_bucket,omitempty"`
	MaxBuckets   int       `json:"MaxBuckets,omitempty"`
	NextMarker   string    `json:"NextMarker,omitempty"`
	Filter       string    `json:"Filter,omitempty"`
	NextPageLink string    `json:"NextPageLink,omitempty"`
}

type BucketCreateReq struct {
	BlockSize         int64  `json:"blockSize,omitempty"`
	NotificationSize  int64  `json:"notificationSize,omitempty"`
	Name              string `json:"name,omitempty"`
	Vpool             string `json:"vpool,omitempty"`
	FilesystemEnabled bool   `json:"filesystem_enabled,omitempty"`
	HeadType          string `json:"head_type,omitempty"`
	Namespace         string `json:"namespace,omitempty"`
	TagSet            []struct {
		Key   string `json:"Key,omitempty"`
		Value string `json:"Value,omitempty"`
	} `json:"TagSet,omitempty"`
	IsEncryptionEnabled               bool   `json:"is_encryption_enabled,omitempty"`
	DefaultGroupFileReadPermission    bool   `json:"default_group_file_read_permission,omitempty"`
	DefaultGroupFileWritePermission   bool   `json:"default_group_file_write_permission,omitempty"`
	DefaultGroupFileExecutePermission bool   `json:"default_group_file_execute_permission,omitempty"`
	DefaultGroupDirReadPermission     bool   `json:"default_group_dir_read_permission,omitempty"`
	DefaultGroupDirWritePermission    bool   `json:"default_group_dir_write_permission,omitempty"`
	DefaultGroupDirExecutePermission  bool   `json:"default_group_dir_execute_permission,omitempty"`
	DefaultGroup                      string `json:"default_group,omitempty"`
	AutocommitPeriod                  int64  `json:"autocommit_period,omitempty"`
	Retention                         int64  `json:"retention,omitempty"`
	IsStaleAllowed                    bool   `json:"is_stale_allowed,omitempty"`
	IsTsoReadOnly                     bool   `json:"is_tso_read_only,omitempty"`
	SearchMetadata                    []struct {
		Type     string `json:"type,omitempty"`
		Name     string `json:"name,omitempty"`
		Datatype string `json:"datatype,omitempty"`
	} `json:"search_metadata,omitempty"`
	Owner          string `json:"owner,omitempty"`
	MinMaxGovernor struct {
		EnforceRetention         bool  `json:"enforce_retention,omitempty"`
		MinimumFixedRetention    int64 `json:"minimum_fixed_retention,omitempty"`
		MaximumFixedRetention    int64 `json:"maximum_fixed_retention,omitempty"`
		MinimumVariableRetention int64 `json:"minimum_variable_retention,omitempty"`
		MaximumVariableRetention int64 `json:"maximum_variable_retention,omitempty"`
	} `json:"min_max_governor,omitempty"`
	AuditedDeleteExpiration int64 `json:"audited_delete_expiration,omitempty"`
}

type BucketCreateResp struct {
	Global   bool `json:"global,omitempty"`
	Remote   bool `json:"remote,omitempty"`
	MetaData struct {
		IsEnabled bool `json:"isEnabled,omitempty"`
		Metadata  []struct {
			Type     string `json:"type,omitempty"`
			Name     string `json:"name,omitempty"`
			DataType string `json:"datatype,omitempty"`
		} `json:"metadata,omitempty"`
		MaxKeys        int  `json:"maxKeys,omitempty"`
		MetadataTokens bool `json:"metadata_tokens,omitempty"`
	} `json:"metaData,omitempty"`
	AdvancedMetadataSearchEnabled bool      `json:"advancedMetadataSearchEnabled,omitempty"`
	Name                          string    `json:"name,omitempty"`
	ID                            string    `json:"id,omitempty"`
	CreationTime                  time.Time `json:"creation_time,omitempty"`
	Inactive                      bool      `json:"inactive,omitempty"`
	Internal                      bool      `json:"internal,omitempty"`
	ErrorMessage                  string    `json:"error_message,omitempty"`
	TagSet                        []struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"TagSet,omitempty"`
}

type BucketQuotaUpdateReq struct {
	BlockSize        int64  `json:"blockSize,omitempty"`
	NotificationSize int64  `json:"notificationSize,omitempty"`
	Namespace        string `json:"namespace,omitempty"`
}

type BucketBillingInfoResp struct {
	Namespace     string    `json:"namespace,omitempty"`
	Name          string    `json:"name,omitempty"`
	VpoolID       string    `json:"vpool_id,omitempty"`
	TotalSize     string    `json:"total_size,omitempty"`
	TotalSizeUnit string    `json:"total_size_unit,omitempty"`
	TotalObjects  int       `json:"total_objects,omitempty"`
	SampleTime    time.Time `json:"sample_time,omitempty"`
	TotalMpuSize  string    `json:"total_mpu_size,omitempty"`
	TotalMpuParts int       `json:"total_mpu_parts,omitempty"`
	TagSet        []struct {
		Tag struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"tag,omitempty"`
	} `json:"TagSet,omitempty"` // update sub struct
	UptodateTill        time.Time `json:"uptodate_till,omitempty"`
	TotalObjectsDeleted string    `json:"total_objects_deleted,omitempty"`
	TotalSizeDeleted    string    `json:"total_size_deleted,omitempty"`
}
