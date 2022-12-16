package repository

const (
	queryGetById = "SELECT id, name, last_name, resume, oab, tel, profile_picture FROM lawyer_profile lp  where id = ?"
)
