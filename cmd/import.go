package cmd

import (
	"fmt"
	"strings"
	"time"
)

func (a *App) ImportAllUsersBare() {
	request := NewRequest(GET, a.CanvasUrl, a.CanvasKey)
	users, err := request.GetAllCanvasUsers()
	if err != nil {
		fmt.Println("Error retrieving users")
	}
	users = checkDuplicates(users)
	unlockedUsers := make([]PostUnlockedUser, len(users))
	for idx, user := range users {
		name_first := strings.TrimSpace(strings.Split(user.SortableName, ",")[1])
		name_last := strings.TrimSpace(strings.Split(user.SortableName, ",")[0])
		unlockedUsers[idx] = PostUnlockedUser{
			NameFirst:     name_first,
			NameLast:      name_last,
			Email:         user.LoginId,
			Role:          "student",
			Username:      user.LoginId,
			PasswordReset: 1,
		}
	}
	unlockedRequest := NewRequest(POST, (a.UnlockedUrl + UNLOCKED_USERS), a.UnlockedKey)
	for _, user := range unlockedUsers {
		unlockedRequest.AddUserUnlocked(user)
		time.Sleep(200 * time.Millisecond)
	}
}

func checkDuplicates(users []CanvasUser) []CanvasUser {
	var uniqueUsers []CanvasUser
	for idx := 0; idx < len(users)-1; idx++ {
		for i := idx + 1; i < len(users); i++ {
			usera := users[idx]
			userb := users[i]
			if strings.Contains(usera.LoginId, userb.SisImportId) || strings.Contains(usera.ShortName, userb.Name) || strings.Contains(usera.SortableName, userb.Name) || strings.Contains(usera.Name, userb.LoginId) {
				continue
			} else {
				uniqueUsers = append(uniqueUsers, usera)
			}
		}
	}
	return uniqueUsers
}
