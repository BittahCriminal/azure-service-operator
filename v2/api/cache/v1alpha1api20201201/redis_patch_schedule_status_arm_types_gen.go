// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

// Deprecated version of RedisPatchSchedule_STATUS. Use v1beta20201201.RedisPatchSchedule_STATUS instead
type RedisPatchSchedule_STATUS_ARM struct {
	Id         *string                     `json:"id,omitempty"`
	Location   *string                     `json:"location,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *ScheduleEntries_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}

// Deprecated version of ScheduleEntries_STATUS. Use v1beta20201201.ScheduleEntries_STATUS instead
type ScheduleEntries_STATUS_ARM struct {
	ScheduleEntries []ScheduleEntry_STATUS_ARM `json:"scheduleEntries,omitempty"`
}

// Deprecated version of ScheduleEntry_STATUS. Use v1beta20201201.ScheduleEntry_STATUS instead
type ScheduleEntry_STATUS_ARM struct {
	DayOfWeek         *ScheduleEntry_DayOfWeek_STATUS `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                         `json:"maintenanceWindow,omitempty"`
	StartHourUtc      *int                            `json:"startHourUtc,omitempty"`
}
