package dbutils

import (
	"fmt"
	ct "jira-backend/constants"
	md "jira-backend/models"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckLoginCreds(username string, password string) int64 {
	var count int64
	var user md.User

	DB.Where("username = ? AND password = ?", username, password).Find(&user).Count(&count)
	if count == 1 {
		return int64(user.UserId)
	}
	return -1
}

func RegisterUser(data sk.UserRegister) gin.H {
	username := data.Username
	email := data.Email

	var count int64

	DB.Where("username = ? OR email_id = ?", username, email).Find(&md.User{}).Count(&count)

	if count != 0 {
		return ut.GetErrorResponse(ct.USER_ALREADY_EXISTS)
	}

	DB.Create(&md.User{Username: username, EmailId: email, Firstname: data.FirstName, Lastname: data.LastName, Password: data.Password})

	return ut.GetSuccessResponse(ct.REGISTERATION_SUCCESSFUL, gin.H{})
}

func PrepareInvite(data sk.InviteUserRequest, hostId uint) gin.H {
	var count int64
	var host md.UserRole
	DB.Preload("Project").Where("user_id =? AND project_id =? AND role_id = 1", hostId, data.ProjectId).Find(&host).Count(&count)
	invite_valid_duration := ut.Vconfig.GetInt64("mail_config.duration") * 24
	fmt.Println(invite_valid_duration)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}
	count = 0
	//Check if already joined
	DB.Joins("JOIN user_roles ON users.user_id = user_roles.user_id AND users.email_id = ? AND user_roles.project_id =  ?", data.EmailId, data.ProjectId).Find(&md.User{}).Count(&count)

	if count == 1 {
		return ut.GetErrorResponse(ct.ALREADY_PART_OF_PROJECT)
	}

	count = 0
	var invite md.Invite
	//Check if already has non-expired invite
	DB.Where("email_id = ? AND project_id = ?", data.EmailId, data.ProjectId).Find(&invite).Count(&count)

	if count == 1 && time.Since(invite.CreatedAt).Milliseconds() < (time.Duration(invite_valid_duration)*time.Hour).Milliseconds() {
		return ut.GetErrorResponse(ct.USER_ALREADY_INVITED)
	}

	inviteLink, mailerr := ut.SendEmail(host.Project, data.EmailId)

	if mailerr != nil {
		fmt.Println(mailerr)
		return ut.GetErrorResponse(ct.UNEXPECTED_ERROR_OCCURED)
	}

	DB.Create(&md.Invite{InviteLink: inviteLink, ProjectId: host.ProjectId, EmailId: data.EmailId})
	fmt.Println(inviteLink, mailerr)
	return ut.GetSuccessResponse(ct.INVITATION_SENT, gin.H{})
}

func ValidateInvite(encryptedLink string, userId uint) gin.H {

	var user md.User
	decryptedDetails, err := ut.ParseEmailInvite(encryptedLink)

	if err != nil {
		return ut.GetErrorResponse(ct.LINK_EXPIRED_INVALID)
	}
	DB.Where("user_id = ?", userId).Find(&user)

	if decryptedDetails.EmailId != user.EmailId {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}
	fmt.Println(user)
	DB.Create(&md.UserRole{UserId: userId, ProjectId: decryptedDetails.ProjectId, RoleId: ct.Developer})
	DB.Where("Where invite_link = ?", encryptedLink).Delete(&md.Invite{})

	return ut.GetSuccessResponse(ct.INVITTATION_ACCEPTED, gin.H{})
}
