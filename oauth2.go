package oauth2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Oauth oauth2 request
type Oauth struct {
	Token         string
	UserID        string
	ClientID      int64
	ValidationURL string
}

// Claim claim
type Claim struct {
	Role  string
	URI   string
	Scope string
}

//AuthReq request to Ulbora OAuth2 Server
type AuthReq struct {
	AccessToken string `json:"accessToken"`
	UserID      string `json:"userId"`
	ClientID    int64  `json:"clientId"`
	Role        string `json:"role"`
	URI         string `json:"uri"`
	Scope       string `json:"scope"`
}

// ValidationResp response from Ulbora OAuth 2 valitate service
type ValidationResp struct {
	Valid bool `json:"valid"`
}

//Authorize the client
func (r *Oauth) Authorize(me *Claim) bool {
	var rtn = false
	var a AuthReq
	a.AccessToken = r.Token
	a.UserID = r.UserID
	a.ClientID = r.ClientID
	a.Role = me.Role
	a.URI = me.URI
	a.Scope = me.Scope

	aJSON, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("before request")
	req, rErr := http.NewRequest("POST", r.ValidationURL, bytes.NewBuffer(aJSON))
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Ulbora OAuth2 proxy err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				vResp := new(ValidationResp)
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&vResp)
				if error != nil {
					log.Println(error.Error())
				} else {
					rtn = vResp.Valid
				}
			}
		}
	}
	return rtn
}
