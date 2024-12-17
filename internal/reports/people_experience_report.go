package reports

type PeopleExperienceReport struct {
	PeopleId   uint64 `json:"people_id"`
	Fcs        string `json:"fcs"`
	Posts      string `json:"posts"`
	Experience uint64 `json:"experience"`
}
