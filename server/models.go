package server

type NewMeeting struct {
	Topic        string      `json:"topic,omitempty"`
	Type         int         `json:"type,omitempty"`
	StartTime    string      `json:"start_time,omitempty"`
	Duration     int         `json:"duration,omitempty"`
	ScheduledFor string      `json:"schedule_for,omitempty"`
	Timezone     string      `json:"timezone,omitempty"`
	Password     string      `json:"password,omitempty"`
	Agenda       string      `json:"agenda,omitempty"`
	Recurrence   *recurrence `json:"recurrence,omitempty"`
	Settings     *settings   `json:"settings,omitempty"`
}

type Meeting struct {
	UUID              string `json:"uuid,omitempty"`
	ID                int    `json:"id,omitempty"`
	HostID            string `json:"host_id,omitempty"`
	HostEmail         string `json:"host_email,omitempty"`
	Topic             string `json:"topic,omitempty"`
	Type              int    `json:"type,omitempty"`
	Status            string `json:"status,omitempty"`
	Timezone          string `json:"timezone,omitempty"`
	Agenda            string `json:"agenda,omitempty"`
	CreateAt          string `json:"created_at,omitempty"`
	StartURL          string `json:"start_url,omitempty"`
	JoinURL           string `json:"join_url,omitempty"`
	Password          string `json:"password,omitempty"`
	HThreeTwoThree    string `json:"h323_password,omitempty"`
	PSTNPassword      string `json:"pstn_password,omitempty"`
	EncyrptedPassword string `json:"encrypted_password,omitempty"`
}

type MeetingsList struct {
	PageCount  int       `json:"page_count,omitempty"`
	PageNumber int       `json:"page_number,omitempty"`
	PageSize   int       `json:"page_size,omitempty"`
	Meetings   []Meeting `json:"meetings,omitempty"`
}

type recurrence struct {
	Type           int    `json:"type,omitempty"`
	Repeat         int    `json:"repeat_interval,omitempty"`
	Weekly         string `json:"weekly_days,omitempty"`
	MonthlyDay     string `json:"monthly_day,omitempty"`
	MonthlyWeek    string `json:"monthly_week,omitempty"`
	MonthlyWeekDay string `json:"monthly_week_day,omitempty"`
	EndTimes       int    `json:"end_times,omitempty"`
	EndTimeDate    string `json:"end_date_time,omitempty"`
}
type settings struct {
	HostVideo                    bool                    `json:"host_video,omitempty"`
	ParticipantVideo             bool                    `json:"participant_video,omitempty"`
	CNMeeting                    bool                    `json:"cn_meeting,omitempty"` //host meeting in China
	INMeeting                    bool                    `json:"in_meeting,omitempty"` // host meeting in India
	JoinBeforeHost               bool                    `json:"join_before_host,omitempty"`
	MuteUponEntry                bool                    `json:"mute_upon_entry,omitempty"`
	Watermark                    bool                    `json:"watermark,omitempty"`
	PMI                          bool                    `json:"use_pmi,omitempty"`
	ApprovalType                 int                     `json:"approval_type,omitempty"`
	RegistrationType             int                     `json:"registration_type,omitempty"`
	Audio                        string                  `json:"audio,omitempty"`
	AutoRecording                string                  `json:"auto_recording,omitempty"`
	EnforceLogin                 bool                    `json:"enforce_login,omitempty"`
	EnforceLoginDomain           string                    `json:"enforce_login_domains,omitempty"`
	AltHosts                     string                  `json:"alternative_hosts,omitempty"`
	//GlobalDialInCountries        []Globaldialincountries `json:"global_dial_in_countries,omitempty"`
	RegistrantsEmailNotification bool                    `json:"registrants_email_notification,omitempty"`
}
type Globaldialincountries struct {
	Countries []string	`json:"global_dial_in_countries,omitempty"`
}