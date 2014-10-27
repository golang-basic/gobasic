package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"

//	"github.com/gobasic/basics/cgoexamples"
	"fmt"
)

const (
	dashboardSubDomain = "service"
	DashboardClientID  = "service-agent-dashboard-client"
	DashboardSecret    = "alice10zebra"
)

type Agent struct {
	AgentIPAddress            string `json:"agent_ip_address"`
	AgentVMPassword           string `json:"agent_vm_password"`
	EncryptedPassword          string `json:"-"`
	PoolSize                   int    `json:"pool_size"`
	DashboardRootUri           string `json:"-"`
	Bosh                       `json:"bosh"` // this guy doesnt have a name only the type
	ServicePlanDetails         `json:"service_plan_details"`
	ExternalServicePlanDetails `json:"external_service_plan_details"`
}

type Bosh struct {
	DirectorUUID   string   `json:"-"`
	DNS            []string `json:"dns"`
	ReservedRanges []string `json:"reserved_ranges"`
}
type ServicePlanDetails struct {
	ComponentDetails `json:"component_details"`
}
type ComponentDetails struct {
	NumberOfPHDSlaves int      `json:"number_of_phd_slaves"`
	CheckedComponents []string `json:"checked_components"`
	ComponentSettings `json:"component_settings"`
}
type ComponentSettings struct {
	NamenodeData        Job `json:"namenode_data"`
	ResourceManagerData Job `json:"resource_manager_data"`
}
type Job struct {
	RAM            int `json:"ram_in_mb"`
	CPU            int `json:"cpu"`
}
type ExternalServicePlanDetails struct {
	CheckedComponents   []string `json:"service_plan_components"`
	ExternalHDFSDetails `json:"hdfs"`
}
type ExternalHDFSDetails struct {
	UserDirectory    string `json:"user_directory"`
	UserGroup        string `json:"user_group"`
}



//func LoadSettings(path string, unixCrypt cgoexamples.Crypt) (agent Agent, err error) {
func LoadSettings(path string) (agent Agent, err error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(buffer, &agent)
	if err != nil {
		return
	}
//	agent.EncryptedPassword = unixCrypt.EncryptWithSalt(agent.AgentVMPassword)
	sort.Strings(agent.ServicePlanDetails.ComponentDetails.CheckedComponents)

//	agent.DashboardRootUri = "https://" + dashboardSubDomain + "." + agent.CfSystemDomain
	var lis []struct{}
	fmt.Println(lis)
	return
}

func main() {
	brokerSettings, err := LoadSettings("../../test_helpers/broker_settings.json")
	fmt.Println(brokerSettings)
	if err !=nil {
		fmt.Println(err)
	}
}
