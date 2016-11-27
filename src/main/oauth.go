package main

import (
    "fmt"
    "bytes"
    "net/url"
)

var oauthAuthorizationUrl, oauthScopes, oauthRedirectUrl *string

func getOAuthScopes() string {
    if oauthScopes != nil {
        return *oauthScopes
    }
    buff := bytes.Buffer{}
    for _, scope := range conf.OAuth.Scopes {
        if buff.Len() > 0 {
            buff.WriteString("+")
        }
        buff.WriteString(scope)
    }
    temp := buff.String()
    oauthScopes = &temp
    return temp
}

func getOAuthRedirectUrl() string {
    if oauthRedirectUrl != nil {
        return *oauthRedirectUrl
    }
    temp := url.QueryEscape(conf.OAuth.RedirectUrl)
    oauthRedirectUrl = &temp
    return temp
}

func getOAuthAuthorizationUrl() string {
    if oauthAuthorizationUrl != nil {
        return *oauthAuthorizationUrl
    }
    temp := fmt.Sprintf(conf.OAuth.AuthorizationUrl, conf.OAuth.ClientId, getOAuthScopes(), getOAuthRedirectUrl())
    oauthAuthorizationUrl = &temp
    return temp
}

func getOAuthTokenUrl(code string) string {
    return fmt.Sprintf(conf.OAuth.TokenUrl, conf.OAuth.ClientId, conf.OAuth.ClientSecret, getOAuthScopes(), code,
        getOAuthRedirectUrl())
}