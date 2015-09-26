package spec

type List struct {
	Id   string
	Name string
}

type ListsResponse struct {
	Lists []List
}

type MembersResponse struct {
	Members []Member
}

type MergeFields struct {
	FirstName string `json:"FNAME"`
	LastName  string `json:"LNAME"`
}
type MemberRequest struct {
	EmailType    string       `json:"email_type"`
	EmailAddress string       `json:"email_address"`
	StatusIfNew  string       `json:"status_if_new"`
	Status       string       `json:"status"`
	MergeFields  *MergeFields `json:"merge_fields"`
	Language     string       `json:"language"`
	Vip          bool         `json:"vip"`
}

type Member struct {
	EmailType    string       `json:"email_type"`
	EmailAddress string       `json:"email_address"`
	StatusIfNew  string       `json:"status_if_new"`
	Status       string       `json:"status"`
	MergeFields  *MergeFields `json:"merge_fields"`
	Language     string       `json:"language"`
	Vip          bool         `json:"vip"`
}
