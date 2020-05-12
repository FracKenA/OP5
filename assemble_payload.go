package op5

import (
	"encoding/json"
	"fmt"
)

func (p Command) Command() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"command", j, p.CommandName}
	return o
}

func (p Contact) Contact() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"contact", j, p.ContactName}
	return o
}

func (p ContactTemplate) ContactTemplate() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"contacttemplate", j, p.Name}
	return o
}

func (p ContactGroup) ContactGroup() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"contactgroup", j, p.ContactGroupName}
	return o
}

func (p Host) Host() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	var emap map[string]interface{}
	if err = json.Unmarshal(j, &emap); err != nil {
		fmt.Printf("Error %s\n", err)
	}

	for k, v := range p.CustomVars {
		emap[k] = v
	}

	j, err = json.Marshal(emap)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"host", j, p.HostName}

	return o
}

func (p HostTemplate) HostTemplate() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	var emap map[string]interface{}
	if err = json.Unmarshal(j, &emap); err != nil {
		fmt.Printf("Error %s\n", err)
	}

	for k, v := range p.CustomVars {
		emap[k] = v
	}

	j, err = json.Marshal(emap)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"hosttemplate", j, p.Name}
	return o
}

func (p HostGroup) HostGroup() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"hostgroup", j, p.HostGroupName}
	return o
}

func (p Service) Service() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	var emap map[string]interface{}
	if err = json.Unmarshal(j, &emap); err != nil {
		fmt.Printf("Error %s\n", err)
	}

	for k, v := range p.CustomVars {
		emap[k] = v
	}

	j, err = json.Marshal(emap)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"service", j, p.ServiceDescription}
	return o
}

func (p ServiceTemplate) ServiceTemplate() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	var emap map[string]interface{}
	if err = json.Unmarshal(j, &emap); err != nil {
		fmt.Printf("Error %s\n", err)
	}

	for k, v := range p.CustomVars {
		emap[k] = v
	}

	j, err = json.Marshal(emap)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"servicetemplate", j, p.Name}
	return o
}

func (p ServiceGroup) ServiceGroup() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"ServiceGroup", j, p.ServiceGroupName}
	return o
}

func (p TimePeriod) TimePeriod() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"timeperiod", j, p.TimePeriodName}
	return o
}

func (p User) User() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"user", j, p.Username}
	return o
}

func (p UserGroup) UserGroup() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"usergroup", j, p.Name}
	return o
}