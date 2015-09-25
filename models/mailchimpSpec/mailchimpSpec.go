package mailchimpSpec

type List struct {
	Id   string
	Name string
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
