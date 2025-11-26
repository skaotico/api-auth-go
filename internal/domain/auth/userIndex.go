package auth

type UserIndex struct {
	ActiveJwt     string `json:"activeJwt"`
	ActiveRefresh string `json:"activeRefresh"`
	LastLogin     int64  `json:"lastLogin"`
}
