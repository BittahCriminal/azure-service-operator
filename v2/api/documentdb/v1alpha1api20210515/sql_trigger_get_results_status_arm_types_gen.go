// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

// Deprecated version of SqlTriggerGetResults_STATUS. Use v1beta20210515.SqlTriggerGetResults_STATUS instead
type SqlTriggerGetResults_STATUS_ARM struct {
	Id         *string                             `json:"id,omitempty"`
	Location   *string                             `json:"location,omitempty"`
	Name       *string                             `json:"name,omitempty"`
	Properties *SqlTriggerGetProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string                   `json:"tags,omitempty"`
	Type       *string                             `json:"type,omitempty"`
}

// Deprecated version of SqlTriggerGetProperties_STATUS. Use v1beta20210515.SqlTriggerGetProperties_STATUS instead
type SqlTriggerGetProperties_STATUS_ARM struct {
	Resource *SqlTriggerGetProperties_Resource_STATUS_ARM `json:"resource,omitempty"`
}

// Deprecated version of SqlTriggerGetProperties_Resource_STATUS. Use v1beta20210515.SqlTriggerGetProperties_Resource_STATUS instead
type SqlTriggerGetProperties_Resource_STATUS_ARM struct {
	Body             *string                                                   `json:"body,omitempty"`
	Etag             *string                                                   `json:"_etag,omitempty"`
	Id               *string                                                   `json:"id,omitempty"`
	Rid              *string                                                   `json:"_rid,omitempty"`
	TriggerOperation *SqlTriggerGetProperties_Resource_TriggerOperation_STATUS `json:"triggerOperation,omitempty"`
	TriggerType      *SqlTriggerGetProperties_Resource_TriggerType_STATUS      `json:"triggerType,omitempty"`
	Ts               *float64                                                  `json:"_ts,omitempty"`
}
