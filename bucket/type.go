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
