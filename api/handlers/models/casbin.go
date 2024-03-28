package models

type Policy struct {
	Role     string `json:"role"`
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}

type ListPolePolicyResponse struct {
	Policies []*Policy `json:"policies"`
}

type CreateUserRoleRequest struct {
	UserId string `json:"userId"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
