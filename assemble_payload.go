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
	o := Payload{"Command", j, p.CommandName}
	return o
}

func (p Contact) Contact() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"Contact", j, p.ContactName}
	return o
}

func (p ContactTemplate) ContactTemplate() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"ContactTemplate", j, p.Name}
	return o
}

func (p ContactGroup) ContactGroup() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"ContactGroup", j, p.ContactGroupName}
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
	o := Payload{"Host", j, p.HostName}
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
	o := Payload{"HostTemplate", j, p.Name}
	return o
}

func (p HostGroup) HostGroup() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"HostGroup", j, p.HostGroupName}
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
	o := Payload{"Service", j, p.ServiceDescription}
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
	o := Payload{"ServiceTemplate", j, p.Name}
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
	o := Payload{"TimePeriod", j, p.TimePeriodName}
	return o
}

func (p User) User() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"User", j, p.Username}
	return o
}

func (p UserGroup) UserGroup() Payload {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	o := Payload{"UserGroup", j, p.Name}
	return o
}