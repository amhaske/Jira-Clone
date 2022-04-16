package constants

import "errors"

const (
	DEV  = "dev"
	PROD = "prod"
	TEST = "test"
)

//api response constant
const (
	Status   = "status"
	Message  = "message"
	Response = "resp"
)

const (
	Access_token = "access_token"
	USER_ID      = "User_Id"
)

//User Roles
const (
	Owner     = 1
	Developer = 2
	Tester    = 3
)

//issue status
const (
	Created     = 1
	In_progress = 2
	Solved      = 3
	Closed      = 4
	Rejected    = 5
)

//
const (
	NOT_PART_OF_THE_PROJECT  = "NOT_PART_OF_THE_PROJECT"
	PROJECT_DELETE_SUCCESS   = "PROJECT_DELETE_SUCCESS"
	PROJECT_DELETE_ERROR     = "PROJECT_DELETE_ERROR"
	ACTION_NOT_AUTHORIZED    = "ACTION_NOT_AUTHORIZED"
	BAD_REQUEST              = "BAD_REQUEST"
	SPRINT_ALREADY_EXISTS    = "SPRINT_EXISTS"
	PROJECT_ALREADY_EXISTS   = "PROJECT_EXISTS"
	WRONG_DATE_FORMAT        = "WRONG_DATE_FORMAT"
	NOT_VALID_SPRINT         = "INVALID_SPRINT"
	SPRINT_DELETE_SUCCESS    = "SPRINT_DELETE_SUCCESS"
	USER_DOESNT_EXIST        = "INVALID_USER"
	ISSUE_DELETE_SUCCESS     = "ISSUE_DELETE_SUCCESS"
	SPRINT_UPDATE_SUCCESS    = "SPRINT_UPDATE_SUCCESS"
	ISSUE_UPDATE_SUCCESS     = "ISSUE_UPDATE_SUCCESS"
	LOGOUT_SUCCESSFUL        = "LOGOUT_SUCCESSFUL"
	INVALID_CREDENTIALS      = "INVALID_CREDENTIALS"
	LOGIN_FAILURE            = "LOGIN_FAILURE"
	LOGIN_SUCCESSFUL         = "LOGIN_SUCCESSFUL"
	EXPIRED_TOKEN            = "EXPIRED_TOKEN"
	USER_ALREADY_EXISTS      = "USER_ALREADY_EXISTS"
	REGISTERATION_SUCCESSFUL = "REGISTERATION_SUCCESSFUL"
)

const (
	SPRINT_DATE_FORMAT = "2006-01-02"
)

var (
	ErrTokenInvalid = errors.New("token is invalid")
)
