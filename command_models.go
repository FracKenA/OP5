package op5

type AcknowledgeHostProblem struct {
	HostName   string `json:"host_name"`
	Sticky     string `json:"sticky"`
	Notify     string `json:"notify"`
	Persistent string `json:"persistent"`
	Comment    string `json:"comment"`
}

type AcknowledgeServiceProblem struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	Sticky             string `json:"sticky"`
	Notify             string `json:"notify"`
	Persistent         string `json:"persistent"`
	Comment            string `json:"comment"`
}

type DeleteDowntimeByHostName struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	DowntimeStartTime  string `json:"downtime_start_time"`
	Comment            string `json:"comment"`
}

type ProcessHostCheckResult struct {
	HostName     string `json:"host_name"`
	StatusCode   string `json:"status_code"`
	PluginOutput string `json:"plugin_output"`
}

type ProcessServiceCheckResult struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	StatusCode         string `json:"status_code"`
	PluginOutput       string `json:"plugin_output"`
}

type ScheduleAndPropagateHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     string `json:"fixed"`
	TriggerID string `json:"trigger_id"`
	Duration  string `json:"duration"`
	Comment   string `json:"comment"`
}

type ScheduleAndPropagateTriggeredHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     string `json:"fixed"`
	TriggerID string `json:"trigger_id"`
	Duration  string `json:"duration"`
	Comment   string `json:"comment"`
}

type ScheduleHostCheck struct {
	HostName  string `json:"host_name"`
	CheckTime string `json:"check_time"`
}

type ScheduleHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     string `json:"fixed"`
	TriggerID string `json:"trigger_id"`
	Duration  string `json:"duration"`
	Comment   string `json:"comment"`
}

type ScheduleServiceCheck struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	CheckTime          string `json:"check_time"`
}

type ScheduleServiceDowntime struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	StartTime          string `json:"start_time"`
	EndTime            string `json:"end_time"`
	Fixed              string `json:"fixed"`
	TriggerID          string `json:"trigger_id"`
	Duration           string `json:"duration"`
	Comment            string `json:"comment"`
}
