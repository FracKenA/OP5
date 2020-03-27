package op5

type AckHostProblem struct {
	HostName   string `json:"host_name"`
	Sticky     int    `json:"sticky"`
	Notify     bool   `json:"notify"`
	Persistent bool   `json:"persistent"`
	Comment    string `json:"comment"`
}

type AckServiceProblem struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	Sticky             int    `json:"sticky"`
	Notify             bool   `json:"notify"`
	Persistent         bool   `json:"persistent"`
	Comment            string `json:"comment"`
}

type DelDowntimeByHostName struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	DowntimeStartTime  string `json:"downtime_start_time"`
	Comment            string `json:"comment"`
}

type ProcHostCheckResult struct {
	HostName     string `json:"host_name"`
	StatusCode   int    `json:"status_code"`
	PluginOutput string `json:"plugin_output"`
}

type ProcServiceCheckResult struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	StatusCode         int    `json:"status_code"`
	PluginOutput       string `json:"plugin_output"`
}

type SchPropagateHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     bool   `json:"fixed"`
	TriggerID int    `json:"trigger_id"`
	Duration  int    `json:"duration"`
	Comment   string `json:"comment"`
}

type SchPropagateTriggeredHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     bool   `json:"fixed"`
	TriggerID int    `json:"trigger_id"`
	Duration  int    `json:"duration"`
	Comment   string `json:"comment"`
}

type SchHostCheck struct {
	HostName  string `json:"host_name"`
	CheckTime string `json:"check_time"`
}

type SchHostDowntime struct {
	HostName  string `json:"host_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Fixed     bool   `json:"fixed"`
	TriggerID int    `json:"trigger_id"`
	Duration  int    `json:"duration"`
	Comment   string `json:"comment"`
}

type SchServiceCheck struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	CheckTime          string `json:"check_time"`
}

type SchServiceDowntime struct {
	HostName           string `json:"host_name"`
	ServiceDescription string `json:"service_description"`
	StartTime          string `json:"start_time"`
	EndTime            string `json:"end_time"`
	Fixed              bool   `json:"fixed"`
	TriggerID          int    `json:"trigger_id"`
	Duration           int    `json:"duration"`
	Comment            string `json:"comment"`
}