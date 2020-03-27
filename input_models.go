package op5

type Command struct {
	CommandName string `json:"command_name"`       //Name / description for the command. Good practise is to use check_ followed by a descriptive name for the test, for example check_http for a simple HTTP check. This, like ALL variables ending in '_name', must be a unique name.
	CommandLine string `json:"command_line"`       // This is the shell command that will be executed. Naemon macros (shortcuts) like $USER$ can be used.
	Register    bool   `json:"register,omitempty"` //Tells op5 Monitor in which file this object should be stored. This can come in handy if you are used to edit the configuration directly in the text files. For security reasons, file creation through the web-interface is not supported.
	FileID      string `json:"file_id,omitempty"`  //If this is set to on/true, the object is considered to be part of the configuration. Otherwise it used as a template.
}

type Contact struct {
	ContactName                 string        `json:"contact_name,omitempty"`
	Alias                       string        `json:"alias,omitempty"`
	HostNotificationOptions     []string      `json:"host_notification_options,omitempty"`
	ServiceNotificationOptions  []string      `json:"service_notification_options,omitempty"`
	HostNotificationCmds        string        `json:"host_notification_cmds,omitempty"`
	ServiceNotificationCmds     string        `json:"service_notification_cmds,omitempty"`
	Register                    bool          `json:"register,omitempty"`
	Template                    string        `json:"template,omitempty"`
	FileID                      string        `json:"file_id,omitempty"`
	HostNotificationsEnabled    bool          `json:"host_notifications_enabled,omitempty"`
	ServiceNotificationsEnabled bool          `json:"service_notifications_enabled,omitempty"`
	CanSubmitCommands           bool          `json:"can_submit_commands,omitempty"`
	RetainStatusInformation     bool          `json:"retain_status_information,omitempty"`
	RetainNonStatusInformation  bool          `json:"retain_nonstatus_information,omitempty"`
	HostNotificationPeriod      string        `json:"host_notification_period,omitempty"`
	ServiceNotificationPeriod   string        `json:"service_notification_period,omitempty"`
	HostNotificationCmdsArgs    string        `json:"host_notification_cmds_args,omitempty"`
	ServiceNotificationCmdsArgs string        `json:"service_notification_cmds_args,omitempty"`
	ContactGroups               []interface{} `json:"contactgroups,omitempty"`
	Email                       string        `json:"email,omitempty"`
	Pager                       string        `json:"pager,omitempty"`
	Address1                    string        `json:"address1,omitempty"`
	Address2                    string        `json:"address2,omitempty"`
	Address3                    string        `json:"address3,omitempty"`
	Address4                    string        `json:"address4,omitempty"`
	Address5                    string        `json:"address5,omitempty"`
	Address6                    string        `json:"address6,omitempty"`
	EnableAccess                string        `json:"enable_access,omitempty"`
	RealName                    string        `json:"realname,omitempty"`
	Password                    string        `json:"password,omitempty"`
	PasswordAlgo                string        `json:"password_algo,omitempty"`
	Modules                     []interface{} `json:"modules,omitempty"`
	Groups                      []interface{} `json:"groups,omitempty"`
	AccessLevels                string        `json:"Access Levels,omitempty"`
	RealNameUser                string        `json:"Realname,omitempty"`
	PasswordUser                string        `json:"Password,omitempty"`
	Role                        []interface{} `json:"Role,omitempty"`
}

type ContactTemplate struct {
	HostNotificationsEnabled    bool          `json:"host_notifications_enabled,omitempty"`
	ServiceNotificationsEnabled bool          `json:"service_notifications_enabled,omitempty"`
	CanSubmitCommands           bool          `json:"can_submit_commands,omitempty"`
	RetainStatusInformation     bool          `json:"retain_status_information,omitempty"`
	RetainNonStatusInformation  bool          `json:"retain_nonstatus_information,omitempty"`
	HostNotificationPeriod      string        `json:"host_notification_period,omitempty"`
	ServiceNotificationPeriod   string        `json:"service_notification_period,omitempty"`
	HostNotificationOptions     []string      `json:"host_notification_options,omitempty"`
	ServiceNotificationOptions  []string      `json:"service_notification_options,omitempty"`
	HostNotificationCmds        string        `json:"host_notification_cmds,omitempty"`
	ServiceNotificationCmds     string        `json:"service_notification_cmds,omitempty"`
	Register                    bool          `json:"register,omitempty"`
	Name                        string        `json:"name,omitempty"`
	FileID                      string        `json:"file_id,omitempty"`
	Template                    string        `json:"template,omitempty"`
	HostNotificationCmdsArgs    string        `json:"host_notification_cmds_args,omitempty"`
	ServiceNotificationCmdsArgs string        `json:"service_notification_cmds_args,omitempty"`
	ContactGroups               []interface{} `json:"contactgroups,omitempty"`
	Email                       string        `json:"email,omitempty"`
	Pager                       string        `json:"pager,omitempty"`
	Address1                    string        `json:"address1,omitempty"`
	Address2                    string        `json:"address2,omitempty"`
	Address3                    string        `json:"address3,omitempty"`
	Address4                    string        `json:"address4,omitempty"`
	Address5                    string        `json:"address5,omitempty"`
	Address6                    string        `json:"address6,omitempty"`
}

type ContactGroup struct {
	ContactGroupName    string        `json:"contactgroup_name,omitempty"`
	Alias               string        `json:"alias,omitempty"`
	Register            bool          `json:"register,omitempty"`
	FileID              string        `json:"file_id,omitempty"`
	Members             []interface{} `json:"members,omitempty"`
	ContactGroupMembers []interface{} `json:"contactgroup_members,omitempty"`
}

type Host struct {
	HostName                   string        `json:"host_name,omitempty"`
	Address                    string        `json:"address,omitempty"`
	MaxCheckAttempts           int           `json:"max_check_attempts,omitempty"`
	CheckInterval              int           `json:"check_interval,omitempty"`
	NotificationInterval       int           `json:"notification_interval,omitempty"`
	Template                   string        `json:"template,omitempty"`
	Register                   bool          `json:"register,omitempty"`
	FileID                     string        `json:"file_id,omitempty"`
	Contacts                   []string      `json:"contacts,omitempty"`
	CheckCommand               string        `json:"check_command,omitempty"`
	RetryInterval              int           `json:"retry_interval,omitempty"`
	ActiveChecksEnabled        bool          `json:"active_checks_enabled,omitempty"`
	PassiveChecksEnabled       bool          `json:"passive_checks_enabled,omitempty"`
	CheckPeriod                string        `json:"check_period,omitempty"`
	EventHandlerEnabled        bool          `json:"event_handler_enabled,omitempty"`
	FlapDetectionEnabled       bool          `json:"flap_detection_enabled,omitempty"`
	ProcessPerfData            bool          `json:"process_perf_data,omitempty"`
	RetainStatusInformation    bool          `json:"retain_status_information,omitempty"`
	RetainNonStatusInformation bool          `json:"retain_nonstatus_information,omitempty"`
	NotificationPeriod         string        `json:"notification_period,omitempty"`
	NotificationOptions        []string      `json:"notification_options,omitempty"`
	NotificationsEnabled       bool          `json:"notifications_enabled,omitempty"`
	Alias                      string        `json:"alias,omitempty"`
	DisplayName                string        `json:"display_name,omitempty"`
	HostGroups                 []interface{} `json:"hostgroups,omitempty"`
	Parents                    []interface{} `json:"parents,omitempty"`
	Children                   []interface{} `json:"children,omitempty"`
	CheckCommandArgs           string        `json:"check_command_args,omitempty"`
	ContactGroups              []interface{} `json:"contact_groups,omitempty"`
	Obsess                     string        `json:"obsess,omitempty"`
	CheckFreshness             string        `json:"check_freshness,omitempty"`
	FreshnessThreshold         string        `json:"freshness_threshold,omitempty"`
	EventHandler               string        `json:"event_handler,omitempty"`
	EventHandlerArgs           string        `json:"event_handler_args,omitempty"`
	LowFlapThreshold           string        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold      string        `json:"high_flap_threshold,omitempty"`
	FlapDetectionOptions   []interface{} `json:"flap_detection_options,omitempty"`
	FirstNotificationDelay string        `json:"first_notification_delay,omitempty"`
	StalkingOptions        []interface{} `json:"stalking_options,omitempty"`
	IconImage              string        `json:"icon_image,omitempty"`
	IconImageAlt           string        `json:"icon_image_alt,omitempty"`
	StatusMapImage         string        `json:"statusmap_image,omitempty"`
	Notes                  string        `json:"notes,omitempty"`
	ActionURL              string        `json:"action_url,omitempty"`
	NotesURL               string        `json:"notes_url,omitempty"`
	TwoDCoords             string        `json:"2d_coords,omitempty"`
	ObsessOverHost         string        `json:"obsess_over_host,omitempty"`
	Services               []Service     `json:"services,omitempty"`
	CustomVars map[string]interface{} `json:"-"`
}

type HostTemplate struct {
	CheckCommand               string        `json:"check_command,omitempty"`
	MaxCheckAttempts           int      `json:"max_check_attempts,omitempty"`
	CheckInterval              int      `json:"check_interval,omitempty"`
	RetryInterval              int      `json:"retry_interval,omitempty"`
	ActiveChecksEnabled        bool     `json:"active_checks_enabled,omitempty"`
	PassiveChecksEnabled       bool     `json:"passive_checks_enabled,omitempty"`
	CheckPeriod                string   `json:"check_period,omitempty"`
	EventHandlerEnabled        bool     `json:"event_handler_enabled,omitempty"`
	FlapDetectionEnabled       bool     `json:"flap_detection_enabled,omitempty"`
	ProcessPerfData            bool     `json:"process_perf_data,omitempty"`
	RetainStatusInformation    bool     `json:"retain_status_information,omitempty"`
	RetainNonStatusInformation bool     `json:"retain_nonstatus_information,omitempty"`
	NotificationInterval       int      `json:"notification_interval,omitempty"`
	NotificationPeriod         string   `json:"notification_period,omitempty"`
	NotificationOptions        []string `json:"notification_options,omitempty"`
	NotificationsEnabled       bool     `json:"notifications_enabled,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Register                   bool     `json:"register,omitempty"`
	FileID                     string   `json:"file_id,omitempty"`
	Template                   string   `json:"template,omitempty"`
	Alias                      string   `json:"alias,omitempty"`
	DisplayName                string   `json:"display_name,omitempty"`
	Address              string        `json:"address,omitempty"`
	HostGroups           []interface{} `json:"hostgroups,omitempty"`
	Parents              []interface{} `json:"parents,omitempty"`
	Children             []interface{} `json:"children,omitempty"`
	CheckCommandArgs     string        `json:"check_command_args,omitempty"`
	Contacts             []interface{} `json:"contacts,omitempty"`
	ContactGroups        []interface{} `json:"contact_groups,omitempty"`
	Obsess               string        `json:"obsess,omitempty"`
	CheckFreshness       string        `json:"check_freshness,omitempty"`
	FreshnessThreshold   string        `json:"freshness_threshold,omitempty"`
	EventHandler         string        `json:"event_handler,omitempty"`
	EventHandlerArgs     string        `json:"event_handler_args,omitempty"`
	LowFlapThreshold       string        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold      string        `json:"high_flap_threshold,omitempty"`
	FlapDetectionOptions   []interface{} `json:"flap_detection_options,omitempty"`
	FirstNotificationDelay string        `json:"first_notification_delay,omitempty"`
	StalkingOptions        []interface{} `json:"stalking_options,omitempty"`
	IconImage              string        `json:"icon_image,omitempty"`
	IconImageAlt           string        `json:"icon_image_alt,omitempty"`
	StatusMapImage         string        `json:"statusmap_image,omitempty"`
	Notes                  string        `json:"notes,omitempty"`
	ActionURL              string        `json:"action_url,omitempty"`
	NotesURL               string        `json:"notes_url,omitempty"`
	TwoDCoords             string        `json:"2d_coords,omitempty"`
	ObsessOverHost         string        `json:"obsess_over_host,omitempty"`
	CustomVars map[string]interface{} `json:"-"`
}

type HostGroup struct {
	HostGroupName    string        `json:"hostgroup_name,omitempty"`
	Alias            string        `json:"alias,omitempty"`
	Register         bool          `json:"register,omitempty"`
	FileID           string        `json:"file_id,omitempty"`
	Members          []interface{} `json:"members,omitempty"`
	HostGroupMembers []interface{} `json:"hostgroup_members,omitempty"`
	Notes            string        `json:"notes,omitempty"`
	NotesURL         string        `json:"notes_url,omitempty"`
	ActionURL        string        `json:"action_url,omitempty"`
	Services         []Service     `json:"services,omitempty"`
}

type Service struct {
	HostName                   string        `json:"host_name,omitempty"`
	ServiceDescription         string        `json:"service_description,omitempty"`
	CheckCommand               string        `json:"check_command,omitempty"`
	CheckCommandArgs           string        `json:"check_command_args,omitempty"`
	MaxCheckAttempts           int           `json:"max_check_attempts,omitempty"`
	CheckInterval              int           `json:"check_interval,omitempty"`
	ParallelizeCheck           bool          `json:"parallelize_check,omitempty"`
	Obsess                     bool          `json:"obsess,omitempty"`
	FlapDetectionEnabled       bool          `json:"flap_detection_enabled,omitempty"`
	NotificationInterval       int           `json:"notification_interval,omitempty"`
	Template                   string        `json:"template,omitempty"`
	Register                   bool          `json:"register,omitempty"`
	FileID                     string        `json:"file_id,omitempty"`
	Contacts                   []string      `json:"contacts,omitempty"`
	IsVolatile                 bool          `json:"is_volatile,omitempty"`
	RetryInterval              int           `json:"retry_interval,omitempty"`
	ActiveChecksEnabled        bool          `json:"active_checks_enabled,omitempty"`
	PassiveChecksEnabled       bool          `json:"passive_checks_enabled,omitempty"`
	CheckPeriod                string        `json:"check_period,omitempty"`
	EventHandlerEnabled        bool          `json:"event_handler_enabled,omitempty"`
	ProcessPerfData            bool          `json:"process_perf_data,omitempty"`
	RetainStatusInformation    bool          `json:"retain_status_information,omitempty"`
	RetainNonStatusInformation bool          `json:"retain_nonstatus_information,omitempty"`
	NotificationPeriod         string        `json:"notification_period,omitempty"`
	NotificationOptions        []string      `json:"notification_options,omitempty"`
	NotificationsEnabled       bool          `json:"notifications_enabled,omitempty"`
	HostGroupName              string        `json:"hostgroup_name,omitempty"`
	DisplayName                string        `json:"display_name,omitempty"`
	ServiceGroups              []interface{} `json:"servicegroups,omitempty"`
	CheckFreshness             string        `json:"check_freshness,omitempty"`
	FreshnessThreshold         string        `json:"freshness_threshold,omitempty"`
	EventHandler               string        `json:"event_handler,omitempty"`
	EventHandlerArgs           string        `json:"event_handler_args,omitempty"`
	LowFlapThreshold           string        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold          string        `json:"high_flap_threshold,omitempty"`
	FlapDetectionOptions       []interface{} `json:"flap_detection_options,omitempty"`
	FirstNotificationDelay     string        `json:"first_notification_delay,omitempty"`
	ContactGroups              []interface{} `json:"contact_groups,omitempty"`
	StalkingOptions            []interface{} `json:"stalking_options,omitempty"`
	Notes                      string        `json:"notes,omitempty"`
	NotesURL                   string        `json:"notes_url,omitempty"`
	ActionURL                  string        `json:"action_url,omitempty"`
	IconImage                  string        `json:"icon_image,omitempty"`
	IconImageAlt               string        `json:"icon_image_alt,omitempty"`
	ObsessOverService          bool          `json:"obsess_over_service,omitempty"`
	CustomVars map[string]interface{} `json:"-"`
}

type ServiceTemplate struct {
	IsVolatile                 bool          `json:"is_volatile,omitempty"`
	CheckCommand               string        `json:"check_command,omitempty"`
	MaxCheckAttempts           int      `json:"max_check_attempts,omitempty"`
	CheckInterval              int      `json:"check_interval,omitempty"`
	RetryInterval              int      `json:"retry_interval,omitempty"`
	ActiveChecksEnabled        bool     `json:"active_checks_enabled,omitempty"`
	PassiveChecksEnabled       bool     `json:"passive_checks_enabled,omitempty"`
	CheckPeriod                string   `json:"check_period,omitempty"`
	EventHandlerEnabled        bool     `json:"event_handler_enabled,omitempty"`
	FlapDetectionEnabled       bool     `json:"flap_detection_enabled,omitempty"`
	ProcessPerfData            bool     `json:"process_perf_data,omitempty"`
	RetainStatusInformation    bool     `json:"retain_status_information,omitempty"`
	RetainNonStatusInformation bool     `json:"retain_nonstatus_information,omitempty"`
	NotificationInterval       int      `json:"notification_interval,omitempty"`
	NotificationPeriod         string   `json:"notification_period,omitempty"`
	NotificationOptions        []string `json:"notification_options,omitempty"`
	NotificationsEnabled       bool     `json:"notifications_enabled,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Register                   bool     `json:"register,omitempty"`
	FileID                     string   `json:"file_id,omitempty"`
	Template                   string   `json:"template,omitempty"`
	HostgroupName              string   `json:"hostgroup_name,omitempty"`
	DisplayName                string   `json:"display_name,omitempty"`
	CheckCommandArgs           string        `json:"check_command_args,omitempty"`
	Servicegroups              []interface{} `json:"servicegroups,omitempty"`
	ParallelizeCheck           string        `json:"parallelize_check,omitempty"`
	Obsess                     string        `json:"obsess,omitempty"`
	CheckFreshness             string        `json:"check_freshness,omitempty"`
	FreshnessThreshold         string        `json:"freshness_threshold,omitempty"`
	EventHandler               string        `json:"event_handler,omitempty"`
	EventHandlerArgs           string        `json:"event_handler_args,omitempty"`
	LowFlapThreshold           string        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold          string        `json:"high_flap_threshold,omitempty"`
	FlapDetectionOptions       []interface{} `json:"flap_detection_options,omitempty"`
	FirstNotificationDelay     string        `json:"first_notification_delay,omitempty"`
	Contacts                   []interface{} `json:"contacts,omitempty"`
	ContactGroups              []interface{} `json:"contact_groups,omitempty"`
	StalkingOptions            []interface{} `json:"stalking_options,omitempty"`
	Notes                      string        `json:"notes,omitempty"`
	NotesURL                   string        `json:"notes_url,omitempty"`
	ActionURL                  string        `json:"action_url,omitempty"`
	IconImage                  string        `json:"icon_image,omitempty"`
	IconImageAlt               string        `json:"icon_image_alt,omitempty"`
	ObsessOverService          string        `json:"obsess_over_service,omitempty"`
	CustomVars map[string]interface{} `json:"-"`
}

type ServiceGroup struct {
	ServiceGroupName    string        `json:"servicegroup_name,omitempty"`
	Register            bool          `json:"register,omitempty"`
	FileID              string        `json:"file_id,omitempty"`
	Members             []string      `json:"members,omitempty"`
	Alias               string        `json:"alias,omitempty"`
	ServiceGroupMembers []interface{} `json:"servicegroup_members,omitempty"`
	Notes               string        `json:"notes,omitempty"`
	NotesURL            string        `json:"notes_url,omitempty"`
	ActionURL           string        `json:"action_url,omitempty"`
}

type TimePeriod struct {
	TimePeriodName string `json:"timeperiod_name,omitempty"`
	Alias          string `json:"alias,omitempty"`
	Sunday         string `json:"sunday,omitempty"`
	Monday         string `json:"monday,omitempty"`
	Tuesday        string `json:"tuesday,omitempty"`
	Wednesday      string `json:"wednesday,omitempty"`
	Thursday       string `json:"thursday,omitempty"`
	Friday         string `json:"friday,omitempty"`
	Saturday       string `json:"saturday,omitempty"`
	Register       bool   `json:"register,omitempty"`
	FileID         string `json:"file_id,omitempty"`
	Exclude        []interface{} `json:"exclude,omitempty"`
}

type User struct {
	Username     string        `json:"username,omitempty"`
	Realname     string        `json:"realname,omitempty"`
	Modules      []string      `json:"modules,omitempty"`
	Groups       []string      `json:"groups,omitempty"`
	PasswordAlgo string        `json:"password_algo,omitempty"`
	AccessLevels []interface{} `json:"Access Levels,omitempty"`
	UsernameUser string        `json:"Username,omitempty"`
	RealNameUser string        `json:"Realname,omitempty"`
	Role         []string      `json:"Role,omitempty"`
}

type UserGroup struct {
	AccessRights []string `json:"access_rights,omitempty"`
	Name         string   `json:"name,omitempty"`
}
