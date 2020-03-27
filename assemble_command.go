package op5

import (
	"encoding/json"
	"fmt"
)

func (p AckHostProblem) AcknowledgeHostProblem() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"ACKNOWLEDGE_HOST_PROBLEM", j, ""}
	return o
}

func (p AckServiceProblem) AcknowledgeServiceProblem() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"ACKNOWLEDGE_SVC_PROBLEM", j, ""}
	return o
}

func (p DelDowntimeByHostName) DeleteDowntimeByHostName() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"DEL_DOWNTIME_BY_HOST_NAME", j, ""}
	return o
}

func (p ProcHostCheckResult) ProcessHostCheckResult() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"PROCESS_HOST_CHECK_RESULT", j, ""}
	return o
}

func (p ProcServiceCheckResult) ProcessServiceCheckResult() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"PROCESS_SERVICE_CHECK_RESULT", j, ""}
	return o
}

func (p SchPropagateHostDowntime) ScheduleAndPropagateHostDowntime() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME", j, ""}
	return o
}

func (p SchPropagateTriggeredHostDowntime) ScheduleAndPropagateTriggeredHostDowntime() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME", j, ""}
	return o
}

func (p SchHostCheck) ScheduleHostCheck() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"SchHostCheck", j, ""}
	return o
}

func (p SchHostDowntime) ScheduleHostDowntime() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"SCHEDULE_HOST_DOWNTIME", j, ""}
	return o
}

func (p SchServiceCheck) ScheduleServiceCheck() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"SCHEDULE_SVC_CHECK", j, ""}
	return o
}

func (p SchServiceDowntime) ScheduleServiceDowntime() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"SCHEDULE_SVC_DOWNTIME", j, ""}
	return o
}

