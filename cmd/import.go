package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"slices"
	"strings"
)


func (a *App) ImportAllUsersBare() {
      var request = NewRequest(GET, a.CanvasUrl, a.CanvasKey)
     users, err := request.GetAllCanvasUsers()
      if err != nil {
      fmt.Println("Error retrieving users")
      }
    users = checkDuplicates(users)
  var unlockedUsers = make([]UnlockEdUser, len(users))
  for idx, user := range users {
    unlockedUsers[idx] = UnlockEdUser{
      Id: user.Id,
      NameFirst: user.Name,
      NameLast: user.SortableName,
      Email: user.LoginId,
      Role: "student",
      EmailVerifiedAt: user.CreatedAt,
      CreatedAt: user.CreatedAt,
      UpdatedAt: user.CreatedAt,
      Username: user.LoginId,
      PasswordReset: 0,
    }
      }
  var unlockedRequest = NewRequest(POST, (a.UnlockedUrl + UNLOCKED_USERS), a.UnlockedKey)
  for _, user := range unlockedUsersJson {
    unlockedRequest.Request.Body = 
  resp, err := unlockedRequest.Send()
  if err != nil {
    fmt.Println("Error sending users")
      }

}

func checkDuplicates(users []CanvasUser) []CanvasUser {
  var uniqueUsers []CanvasUser
  for idx := 0; idx  < len(users) - 1; idx++ {
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
