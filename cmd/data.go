package cmd

const (
	GET                = "GET"
	POST               = "POST"
	PUT                = "PUT"
	PATCH              = "PATCH"
	CANVAS_BASE_URL    = "/api/v1/accounts/self/"
	CANVAS_USERS       = "/api/v1/accounts/self/users"
	UNLOCKED_USERS     = "/api/v1/users"
	UNLOCKED_ALL_USERS = "/api/v1/users?all=true" // remove pagination with?all=true
)

type CanvasUser struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	CreatedAt     string `json:"created_at"`
	SortableName  string `json:"sortable_name"`
	ShortName     string `json:"short_name"`
	SisUserId     string `json:"sis_user_id"`
	IntegrationId string `json:"integration_id"`
	SisImportId   string `json:"sis_import_id"`
	LoginId       string `json:"login_id"`
}

type UnlockEdUser struct {
	Id              int    `json:"id"`
	NameFirst       string `json:"name_first"`
	NameLast        string `json:"name_last"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	EmailVerifiedAt string `json:"email_verified_at"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Username        string `json:"username"`
	PasswordReset   int    `json:"password_reset"`
}

type PostUnlockedUser struct {
	NameFirst     string `json:"name_first"`
	NameLast      string `json:"name_last"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	Username      string `json:"username"`
	PasswordReset int    `json:"password_reset"`
}

type UserMapping struct {
	Id         int
	ExternId   int
	UnlockedId int
}
