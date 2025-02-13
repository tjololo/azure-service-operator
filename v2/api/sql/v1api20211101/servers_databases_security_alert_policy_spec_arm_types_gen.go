// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ServersDatabasesSecurityAlertPolicy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ServersDatabasesSecurityAlertPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesSecurityAlertPolicy_Spec_ARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (policy *ServersDatabasesSecurityAlertPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/securityAlertPolicies"
func (policy *ServersDatabasesSecurityAlertPolicy_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/databases/securityAlertPolicies"
}

// Properties of a security alert policy.
type DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM struct {
	// DisabledAlerts: Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection,
	// Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
	DisabledAlerts []string `json:"disabledAlerts,omitempty"`

	// EmailAccountAdmins: Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty"`

	// EmailAddresses: Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// RetentionDays: Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `json:"retentionDays,omitempty"`

	// State: Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the
	// specific database.
	State *DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM `json:"state,omitempty"`

	// StorageAccountAccessKey: Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *string `json:"storageAccountAccessKey,omitempty"`

	// StorageEndpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage
	// will hold all Threat Detection audit logs.
	StorageEndpoint *string `json:"storageEndpoint,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM string

const (
	DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM_Disabled = DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM("Disabled")
	DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM_Enabled  = DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM("Enabled")
)

// Mapping from string to DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM
var databaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM_Values = map[string]DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM{
	"disabled": DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM_Disabled,
	"enabled":  DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_ARM_Enabled,
}
