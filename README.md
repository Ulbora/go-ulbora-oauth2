go-ulbora-oauth2
==============

Go OAuth2 library for Ulbora OAuth2 Server

# Installation

```
$ go get github.com/Ulbora/go-ulbora-oauth2

```

## Usage

```
	var auth = new(Oauth)
	auth.Token = aTokenFromUlboraOuth2Server // Passed in the REST service header as Authorization Bearer
	auth.ClientID = 403 // Passed in the REST service header as clientId
	//auth.UserID = "bob" //if controlling service access at user level. Can be passed in the REST service header of 
	auth.ValidationURL = "http://localhost:3000/rs/token/validate"

	var me = new(Claim)
	me.Role = "admin"
	me.URI = "/rs/order/add"
	valid := auth.Authorize(me)
    // if valid, then allow REST service access to continue. Otherwise, throw a 401 code

	
```