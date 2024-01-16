package cmd

import (
	json "encoding/json"
	"io"
	"net/http"
)

type RequestHandler struct {
	Client  *http.Client
	Request *http.Request
	Method  string
}

func NewRequest(method string, base_url string, key string) *RequestHandler {
	req, err := http.NewRequest(method, base_url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("Authorization", "Bearer "+key)
	var request RequestHandler = RequestHandler{
		Request: req,
		Client:  &http.Client{},
	}
	return &request
}

// Sends GET request to the URL and returns the body as a string
func (r *RequestHandler) Send() (string, error) {
	resp, err := r.Client.Do(r.Request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r *RequestHandler) GetAllCanvasUsers() ([]CanvasUser, error) {
	url := r.Request.URL.JoinPath(r.Request.URL.String(), CANVAS_USERS)
	r.Request.URL = url

	req, err := r.Send()
	if err != nil {
		return nil, err
	}
	var users []CanvasUser
	err = json.Unmarshal([]byte(req), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *RequestHandler) GetAllUnlockedUsers() ([]UnlockEdUser, error) {
	url := r.Request.URL.JoinPath(r.Request.URL.String(), UNLOCKED_USERS)
	r.Request.URL = url

	req, err := r.Send()
	if err != nil {
		return nil, err
	}
	var users []UnlockEdUser
	err = json.Unmarshal([]byte(req), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *RequestHandler) SearchCanvasUsers(name string) []CanvasUser {
	url := r.Request.URL.JoinPath(r.Request.URL.String(), CANVAS_USERS+"?search_term="+name)

	r.Request.URL = url
	body, err := r.Send()
	if err != nil {
		return nil
	}
	var users []CanvasUser
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil
	}
	return users
}

func (r *RequestHandler) SearchUnlockedUsers(name string) []UnlockEdUser {
	url := r.Request.URL.JoinPath(r.Request.URL.String(), CANVAS_USERS+"?search_term="+name)
	r.Request.URL = url
	body, err := r.Send()
	if err != nil {
		return nil
	}
	var users []UnlockEdUser
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil
	}
	return users
}

func (r *RequestHandler) AddUserUnlocked(user CanvasUser) {
	url := r.Request.URL.JoinPath(r.Request.URL.String(), CANVAS_USERS)
	r.Request.URL = url
	r.Request.Method = POST

	body, err := json.Marshal(user)
	if err != nil {
		return
	}
	r.Request.Body = body
	r.Request.Method = POST
	r.Send()
}
