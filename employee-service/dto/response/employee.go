package responseDto

type EmployeeSummary struct {
	Nip    string `json:"nip"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}
