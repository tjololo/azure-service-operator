// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

type StorageAccountsBlobService_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a storage account’s Blob service.
	Properties *StorageAccounts_BlobService_Properties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: Sku name and tier.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type StorageAccounts_BlobService_Properties_STATUS_ARM struct {
	// AutomaticSnapshotPolicyEnabled: Deprecated in favor of isVersioningEnabled property.
	AutomaticSnapshotPolicyEnabled *bool `json:"automaticSnapshotPolicyEnabled,omitempty"`

	// ChangeFeed: The blob service properties for change feed events.
	ChangeFeed *ChangeFeed_STATUS_ARM `json:"changeFeed,omitempty"`

	// ContainerDeleteRetentionPolicy: The blob service properties for container soft delete.
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_STATUS_ARM `json:"containerDeleteRetentionPolicy,omitempty"`

	// Cors: Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Blob service.
	Cors *CorsRules_STATUS_ARM `json:"cors,omitempty"`

	// DefaultServiceVersion: DefaultServiceVersion indicates the default version to use for requests to the Blob service if an
	// incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion *string `json:"defaultServiceVersion,omitempty"`

	// DeleteRetentionPolicy: The blob service properties for blob soft delete.
	DeleteRetentionPolicy *DeleteRetentionPolicy_STATUS_ARM `json:"deleteRetentionPolicy,omitempty"`

	// IsVersioningEnabled: Versioning is enabled if set to true.
	IsVersioningEnabled *bool `json:"isVersioningEnabled,omitempty"`

	// LastAccessTimeTrackingPolicy: The blob service property to configure last access time based tracking policy.
	LastAccessTimeTrackingPolicy *LastAccessTimeTrackingPolicy_STATUS_ARM `json:"lastAccessTimeTrackingPolicy,omitempty"`

	// RestorePolicy: The blob service properties for blob restore policy.
	RestorePolicy *RestorePolicyProperties_STATUS_ARM `json:"restorePolicy,omitempty"`
}

// The blob service properties for change feed events.
type ChangeFeed_STATUS_ARM struct {
	// Enabled: Indicates whether change feed event logging is enabled for the Blob service.
	Enabled *bool `json:"enabled,omitempty"`

	// RetentionInDays: Indicates the duration of changeFeed retention in days. Minimum value is 1 day and maximum value is
	// 146000 days (400 years). A null value indicates an infinite retention of the change feed.
	RetentionInDays *int `json:"retentionInDays,omitempty"`
}

// Sets the CORS rules. You can include up to five CorsRule elements in the request.
type CorsRules_STATUS_ARM struct {
	// CorsRules: The List of CORS rules. You can include up to five CorsRule elements in the request.
	CorsRules []CorsRule_STATUS_ARM `json:"corsRules"`
}

// The service properties for soft delete.
type DeleteRetentionPolicy_STATUS_ARM struct {
	// AllowPermanentDelete: This property when set to true allows deletion of the soft deleted blob versions and snapshots.
	// This property cannot be used blob restore policy. This property only applies to blob service and does not apply to
	// containers or file share.
	AllowPermanentDelete *bool `json:"allowPermanentDelete,omitempty"`

	// Days: Indicates the number of days that the deleted item should be retained. The minimum specified value can be 1 and
	// the maximum value can be 365.
	Days *int `json:"days,omitempty"`

	// Enabled: Indicates whether DeleteRetentionPolicy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// The blob service properties for Last access time based tracking policy.
type LastAccessTimeTrackingPolicy_STATUS_ARM struct {
	// BlobType: An array of predefined supported blob types. Only blockBlob is the supported value. This field is currently
	// read only
	BlobType []string `json:"blobType"`

	// Enable: When set to true last access time based tracking is enabled.
	Enable *bool `json:"enable,omitempty"`

	// Name: Name of the policy. The valid value is AccessTimeTracking. This field is currently read only
	Name *LastAccessTimeTrackingPolicy_Name_STATUS_ARM `json:"name,omitempty"`

	// TrackingGranularityInDays: The field specifies blob object tracking granularity in days, typically how often the blob
	// object should be tracked.This field is currently read only with value as 1
	TrackingGranularityInDays *int `json:"trackingGranularityInDays,omitempty"`
}

// The blob service properties for blob restore policy
type RestorePolicyProperties_STATUS_ARM struct {
	// Days: how long this blob can be restored. It should be great than zero and less than DeleteRetentionPolicy.days.
	Days *int `json:"days,omitempty"`

	// Enabled: Blob restore is enabled if set to true.
	Enabled *bool `json:"enabled,omitempty"`

	// LastEnabledTime: Deprecated in favor of minRestoreTime property.
	LastEnabledTime *string `json:"lastEnabledTime,omitempty"`

	// MinRestoreTime: Returns the minimum date and time that the restore can be started.
	MinRestoreTime *string `json:"minRestoreTime,omitempty"`
}

// Specifies a CORS rule for the Blob service.
type CorsRule_STATUS_ARM struct {
	// AllowedHeaders: Required if CorsRule element is present. A list of headers allowed to be part of the cross-origin
	// request.
	AllowedHeaders []string `json:"allowedHeaders"`

	// AllowedMethods: Required if CorsRule element is present. A list of HTTP methods that are allowed to be executed by the
	// origin.
	AllowedMethods []CorsRule_AllowedMethods_STATUS_ARM `json:"allowedMethods"`

	// AllowedOrigins: Required if CorsRule element is present. A list of origin domains that will be allowed via CORS, or "*"
	// to allow all domains
	AllowedOrigins []string `json:"allowedOrigins"`

	// ExposedHeaders: Required if CorsRule element is present. A list of response headers to expose to CORS clients.
	ExposedHeaders []string `json:"exposedHeaders"`

	// MaxAgeInSeconds: Required if CorsRule element is present. The number of seconds that the client/browser should cache a
	// preflight response.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}

type LastAccessTimeTrackingPolicy_Name_STATUS_ARM string

const LastAccessTimeTrackingPolicy_Name_STATUS_ARM_AccessTimeTracking = LastAccessTimeTrackingPolicy_Name_STATUS_ARM("AccessTimeTracking")

// Mapping from string to LastAccessTimeTrackingPolicy_Name_STATUS_ARM
var lastAccessTimeTrackingPolicy_Name_STATUS_ARM_Values = map[string]LastAccessTimeTrackingPolicy_Name_STATUS_ARM{
	"accesstimetracking": LastAccessTimeTrackingPolicy_Name_STATUS_ARM_AccessTimeTracking,
}

type CorsRule_AllowedMethods_STATUS_ARM string

const (
	CorsRule_AllowedMethods_STATUS_ARM_CONNECT = CorsRule_AllowedMethods_STATUS_ARM("CONNECT")
	CorsRule_AllowedMethods_STATUS_ARM_DELETE  = CorsRule_AllowedMethods_STATUS_ARM("DELETE")
	CorsRule_AllowedMethods_STATUS_ARM_GET     = CorsRule_AllowedMethods_STATUS_ARM("GET")
	CorsRule_AllowedMethods_STATUS_ARM_HEAD    = CorsRule_AllowedMethods_STATUS_ARM("HEAD")
	CorsRule_AllowedMethods_STATUS_ARM_MERGE   = CorsRule_AllowedMethods_STATUS_ARM("MERGE")
	CorsRule_AllowedMethods_STATUS_ARM_OPTIONS = CorsRule_AllowedMethods_STATUS_ARM("OPTIONS")
	CorsRule_AllowedMethods_STATUS_ARM_PATCH   = CorsRule_AllowedMethods_STATUS_ARM("PATCH")
	CorsRule_AllowedMethods_STATUS_ARM_POST    = CorsRule_AllowedMethods_STATUS_ARM("POST")
	CorsRule_AllowedMethods_STATUS_ARM_PUT     = CorsRule_AllowedMethods_STATUS_ARM("PUT")
	CorsRule_AllowedMethods_STATUS_ARM_TRACE   = CorsRule_AllowedMethods_STATUS_ARM("TRACE")
)

// Mapping from string to CorsRule_AllowedMethods_STATUS_ARM
var corsRule_AllowedMethods_STATUS_ARM_Values = map[string]CorsRule_AllowedMethods_STATUS_ARM{
	"connect": CorsRule_AllowedMethods_STATUS_ARM_CONNECT,
	"delete":  CorsRule_AllowedMethods_STATUS_ARM_DELETE,
	"get":     CorsRule_AllowedMethods_STATUS_ARM_GET,
	"head":    CorsRule_AllowedMethods_STATUS_ARM_HEAD,
	"merge":   CorsRule_AllowedMethods_STATUS_ARM_MERGE,
	"options": CorsRule_AllowedMethods_STATUS_ARM_OPTIONS,
	"patch":   CorsRule_AllowedMethods_STATUS_ARM_PATCH,
	"post":    CorsRule_AllowedMethods_STATUS_ARM_POST,
	"put":     CorsRule_AllowedMethods_STATUS_ARM_PUT,
	"trace":   CorsRule_AllowedMethods_STATUS_ARM_TRACE,
}
