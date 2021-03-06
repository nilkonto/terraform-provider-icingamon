package iapi

/*
Currently to get something working and that can be refactored there is a lot of duplicate and overlapping decleration. In
part this is because when a variable is defined it is set to a default value. This has been problematic with having an attrs
struct that has all the variables. That struct then cannot be used to create the JSON for the create, without modification,
because it would try and set values that are not configurable via the API. i.e. for hosts "LastCheck" So to keep things moving
duplicate or near duplicate defintions of structs are being defined but can be revisited and refactored later and test will
be in place to ensure everything still works.
*/

// ServiceStruct ... stores service results
type ServiceStruct struct {
	Attrs ServiceAttrs `json:"attrs"`
	Joins struct{}     `json:"joins"`
	//	Meta  struct{}     `json:"meta"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ServiceAttrs ...
type ServiceAttrs struct {
	CheckCommand string `json:"check_command"`
	//	CheckInterval float64       `json:"check_interval"`
	//	DisplayName   string        `json:"display_name"`
	//	Groups        []interface{} `json:"groups"`
	//	Name string `json:"name"`
	//	Templates     []string      `json:"templates"`
	//	Type string `json:"type"`
	//	Vars          interface{}   `json:"vars"`
}

// CheckcommandStruct is a struct used to store results from an Icinga2 Checkcommand API call.
type CheckcommandStruct struct {
	Name  string            `json:"name"`
	Type  string            `json:"type"`
	Attrs CheckcommandAttrs `json:"attrs"`
	Joins struct{}          `json:"joins"`
	Meta  struct{}          `json:"meta"`
}

// CheckcommandAttrs ...
type CheckcommandAttrs struct {
	Arguments interface{} `json:"arguments"`
	Command   []string    `json:"command"`
	Templates []string    `json:"templates"`
	//	Env       interface{} `json:"env"`   				// Available to be set but not supported yet
	//	Package   string      `json:"package"`   		// Available to be set but not supported yet
	//	Timeout   float64     `json:"timeout"`   		// Available to be set but not supported yet
	//	Vars      interface{} `json:"vars"`   			// Available to be set but not supported yet
	//	Zone      string      `json:"zone"`   			// Available to be set but not supported yet
}

// HostgroupStruct is a struct used to store results from an Icinga2 HostGroup API Call. The content are also used to generate the JSON for the CreateHost call
type HostgroupStruct struct {
	Name  string         `json:"name"`
	Type  string         `json:"type"`
	Attrs HostgroupAttrs `json:"attrs"`
	Meta  struct{}       `json:"meta"`
	Joins struct{}       `json:"stuct"`
}

// HostgroupAttrs ...
type HostgroupAttrs struct {
	ActionURL   string   `json:"action_url,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Groups      []string `json:"groups,omitempty"`
	Notes       string   `json:"notes,omitempty"`
	NotesURL    string   `json:"notes_url,omitempty"`
	Templates   []string `json:"templates,omitempty"`
}

// HostStruct is a struct used to store results from an Icinga2 Host API Call. The content are also used to generate the JSON for the CreateHost call
type HostStruct struct {
	Templates []string               `json:"templates,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Attrs     map[string]interface{} `json:"attrs,omitempty"`
	Meta      struct{}               `json:"meta,omitempty"`
	Joins     struct{}               `json:"stuct,omitempty"`
}

// HostStructRead is a struct used to store results from an Icinga2 Host API Call. The content are also used to generate the JSON for the CreateHost call
type HostStructRead struct {
	Templates []string      `json:"templates,omitempty"`
	Name      string        `json:"name,omitempty"`
	Type      string        `json:"type,omitempty"`
	Attrs     HostAttrsRead `json:"attrs,omitempty"`
	Meta      struct{}      `json:"meta,omitempty"`
	Joins     struct{}      `json:"stuct,omitempty"`
}

// HostAttrs This is struct lists the attributes that can be set during a CreateHost call. The contents of the struct is converted into JSON
type HostAttrs struct {
	ActionURL           string      `json:"action_url,omitempty"`
	Address             string      `json:"address,omitempty"`
	Address6            string      `json:"address6,omitempty"`
	CheckCommand        string      `json:"check_command,omitempty"`
	CheckInterval       int         `json:"check_interval,omitempty"`
	CheckPeriod         string      `json:"check_period,omitempty"`
	DisplayName         string      `json:"display_name,omitempty"`
	EnableActiveChecks  bool        `json:"enable_active_checks,omitempty"`
	EnableNotifications string      `json:"enable_notifications,omitempty"`
	Groups              []string    `json:"groups,omitempty"`
	Notes               string      `json:"notes,omitempty"`
	NotesURL            string      `json:"notes_url,omitempty"`
	Vars                interface{} `json:"vars,omitempty"`
	Volatile            bool        `json:"volatile,omitempty"`
	Zone                string      `json:"zone,omitempty"`
}

// HostAttrsRead This is struct lists the attributes that can be set during a CreateHost call. The contents of the struct is converted into JSON
type HostAttrsRead struct {
	ActionURL           string      `json:"action_url,omitempty"`
	Address             string      `json:"address,omitempty"`
	Address6            string      `json:"address6,omitempty"`
	CheckCommand        string      `json:"check_command,omitempty"`
	CheckInterval       int         `json:"check_interval,omitempty"`
	CheckPeriod         string      `json:"check_period,omitempty"`
	DisplayName         string      `json:"display_name,omitempty"`
	EnableActiveChecks  bool        `json:"enable_active_checks,omitempty"`
	EnableNotifications bool        `json:"enable_notifications,omitempty"`
	Groups              []string    `json:"groups,omitempty"`
	Notes               string      `json:"notes,omitempty"`
	NotesURL            string      `json:"notes_url,omitempty"`
	Vars                interface{} `json:"vars,omitempty"`
	Volatile            bool        `json:"volatile,omitempty"`
	Zone                string      `json:"zone,omitempty"`
}

// APIResult Stores the results from NewApiRequest
type APIResult struct {
	Error       float64 `json:"error"`
	ErrorString string
	Status      string      `json:"Status"`
	Code        int         `json:"Code"`
	Results     interface{} `json:"results"`
}

// FilteredHostResults ...
type FilteredHostResults struct {
	Results []struct {
		Attrs struct {
			Name string `json:"name"`
		} `json:"attrs"`
		Joins struct {
		} `json:"joins"`
		Meta struct {
		} `json:"meta"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
}

// APIStatus stores the results of an Icinga2 API Status Call
type APIStatus struct {
	Results []struct {
		Name     string   `json:"name"`
		Perfdata []string `json:"perfdata"`
		Status   struct {
			API struct {
				ConnEndpoints       []interface{} `json:"conn_endpoints"`
				Identity            string        `json:"identity"`
				NotConnEndpoints    []interface{} `json:"not_conn_endpoints"`
				NumConnEndpoints    int           `json:"num_conn_endpoints"`
				NumEndpoints        int           `json:"num_endpoints"`
				NumNotConnEndpoints int           `json:"num_not_conn_endpoints"`
				Zones               struct {
					Master struct {
						ClientLogLag int      `json:"client_log_lag"`
						Connected    bool     `json:"connected"`
						Endpoints    []string `json:"endpoints"`
						ParentZone   string   `json:"parent_zone"`
					} `json:"master"`
				} `json:"zones"`
			} `json:"api"`
		} `json:"status"`
	} `json:"results"`
}
