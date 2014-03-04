package goxcar

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ENDPOINT = "https://new.boxcar.io/api/notifications"
)

func Notify(credentials string, title string, long_message string, sound string) (string, error) {

	values := make(url.Values)
	values.Set("user_credentials", credentials)
	values.Set("notification[title]", title)
	values.Set("notification[long_message]", long_message)
	values.Set("notification[sound]", sound)

	r, err := http.PostForm(ENDPOINT, values)
	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return string(body), nil
}
