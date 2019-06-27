package method

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/salaleser/vk-api/util"
	"golang.org/x/net/html"
)

const (
	authURL      = "https://oauth.vk.com/authorize?%s"
	scope        = "market,photos,friends,wall,groups,messages,offline"
	redirectURI  = "https://oauth.vk.com/blank.html"
	display      = "wap"
	responseType = "token"
)

type AuthResponse struct {
	UserID           int    `json:"user_id"`
	ExpiresIn        int    `json:"expires_in"`
	AccessToken      string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func Auth(login string, password string, clientID string) (string, error) {
	var params = url.Values{
		"client_id":     {clientID},
		"scope":         {scope},
		"redirect_uri":  {redirectURI},
		"display":       {display},
		"v":             {util.V},
		"response_type": {responseType},
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	uri := fmt.Sprintf(authURL, params.Encode())
	resp, err := client.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	args, u := parseForm(resp.Body)

	args.Add("email", login)
	args.Add("pass", password)

	resp, err = client.PostForm(u, args)
	if err != nil {
		return "", err
	}

	if resp.Request.URL.Path != "/blank.html" {
		args, u := parseForm(resp.Body)
		resp, err := client.PostForm(u, args)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.Request.URL.Path != "/blank.html" {
			return "", errors.New("Не удалось получить токен. Возможно,\n" +
				"указанные логин и/или пароль неверные")
		}
	}

	urlArgs, err := url.ParseQuery(resp.Request.URL.Fragment)
	if err != nil {
		return "", err
	}

	return urlArgs["access_token"][0], nil
}

func parseForm(body io.ReadCloser) (url.Values, string) {
	tokenizer := html.NewTokenizer(body)

	u := ""
	formData := map[string]string{}

	end := false
	for !end {
		tag := tokenizer.Next()

		switch tag {
		case html.ErrorToken:
			end = true
		case html.StartTagToken:
			switch token := tokenizer.Token(); token.Data {
			case "form":
				for _, attr := range token.Attr {
					if attr.Key == "action" {
						u = attr.Val
					}
				}
			case "input":
				if token.Attr[1].Val == "_origin" {
					formData["_origin"] = token.Attr[2].Val
				}
				if token.Attr[1].Val == "to" {
					formData["to"] = token.Attr[2].Val
				}
			}
		case html.SelfClosingTagToken:
			switch token := tokenizer.Token(); token.Data {
			case "input":
				if token.Attr[1].Val == "ip_h" {
					formData["ip_h"] = token.Attr[2].Val
				}
				if token.Attr[1].Val == "lg_h" {
					formData["lg_h"] = token.Attr[2].Val
				}
			}
		}
	}

	args := url.Values{}

	for key, val := range formData {
		args.Add(key, val)
	}

	return args, u
}
