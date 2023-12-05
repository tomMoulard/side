package task

type Task struct {
	ID      string `json:"_id"`
	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
		Street  string `json:"street"`
		ZipCode string `json:"zipCode"`
	} `json:"address"`
	Alias      string `json:"alias"`
	Applicants []struct {
		Answers                []string `json:"answers"`
		ApplicantID            string   `json:"applicantId"`
		LastStatusNotified     string   `json:"lastStatusNotified"`
		LastStatusNotifiedDate string   `json:"lastStatusNotifiedDate"`
		Notified               bool     `json:"notified"`
		Reason                 string   `json:"reason"`
		Status                 string   `json:"status"`
		StatusHistory          []struct {
			Date   string `json:"date"`
			Status string `json:"status"`
		} `json:"statusHistory"`
		When string `json:"when"`
	} `json:"applicants"`
	AssigneeID      string `json:"assigneeId"`
	CityID          string `json:"cityId"`
	CompanyNotified bool   `json:"companyNotified"`
	CreatedAt       string `json:"createdAt"`
	CreatorID       string `json:"creatorId"`
	DressCode       string `json:"dressCode"`
	Experiences     string `json:"experiences"`
	HourlyRate      int64  `json:"hourlyRate"`
	IsPreSelection  bool   `json:"isPreSelection"`
	LiveAt          string `json:"liveAt"`
	LocationOptions struct {
		Motorized bool `json:"motorized"`
		Remote    struct {
			Available bool `json:"available"`
			Mandatory bool `json:"mandatory"`
		} `json:"remote"`
	} `json:"locationOptions"`
	LogAsID            string `json:"logAsId"`
	ManagerID          string `json:"managerId"`
	Migrated           bool   `json:"migrated"`
	MissionInformation string `json:"missionInformation"`
	Motive             struct {
		Reason       string `json:"reason"`
		Replacements []struct {
			Justification string `json:"justification"`
			Name          string `json:"name"`
			Position      string `json:"position"`
		} `json:"replacements"`
	} `json:"motive"`
	OrganisationID      string        `json:"organisationId"`
	PostedAt            string        `json:"postedAt"`
	PricingID           string        `json:"pricingId"`
	Purpose             string        `json:"purpose"`
	RequestedSiderIds   []interface{} `json:"requestedSiderIds"`
	RequestedSidersOnly bool          `json:"requestedSidersOnly"`
	SelectionStatus     string        `json:"selectionStatus"`
	Settings            struct {
		ExpectedApplicantNumber int64  `json:"expectedApplicantNumber"`
		HiringEndDate           string `json:"hiringEndDate"`
		MaestroID               string `json:"maestroId"`
	} `json:"settings"`
	ShiftIds             []string `json:"shiftIds"`
	SideNote             string   `json:"sideNote"`
	Status               string   `json:"status"`
	SubmittedAt          string   `json:"submittedAt"`
	SubtaskIds           []string `json:"subtaskIds"`
	Type                 string   `json:"type"`
	TypeID               string   `json:"typeId"`
	UpdatedAt            string   `json:"updatedAt"`
	UsersAlreadyNotified bool     `json:"usersAlreadyNotified"`
	Visible              bool     `json:"visible"`
	WorkConditions       string   `json:"workConditions"`
	WorkLegalStatus      string   `json:"workLegalStatus"`
}
