package interface

import (
        "encoding/json"
        "fmt"
        "net/http"
        "strings"

        "zura.org/oauth2-api/usecase"
        "golang.org/x/oauth2"
       )

type Handler struct {
    OAuthConfig *oauth2.Config
        UserService *usecase.UserService
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
url:=h.OAuthConfig.AuthodeURL("state",oauth2.AccessTypeOffline)
        http.Redirect(w,r,http.StatusFound)
}

func (h *Handler) CallbackHandler(w http.ResponseWriter, r *http.Request) {
code:=r.URL.Query().Get("code")
         token,err:=h.oAuthConfig.Exchange(r.Context(),code)
         if err!=nil {
             http.Error(w,"OAuth exchange failed", http.StatusUnauthorized)
                 return
         }
client:=h.OAuthConfig.Client(r.Context(),token)
           resp,_:=client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
           defer resp.Body.Close()

           var data map[string]string
           json.NewDecoder(resp.Body).Decode(&data)

           user, err:=h.UserService.GetOrCreateUser(data["email"], data["name"])
           if err!=nil{
               http.Error(w,"User save failed", http.StatusInternalServierError)
                   return
           }
       http.SetCookie(w,&http.Cookie{Name:"user".Name,Value:user.Email,Path:"/"}
               http.Redirect(w,r,"/protected",http.StatusFound)
}

func (h *Handler) ProtectedHandler(w http.ResponseWriter, r *http.Request){
cookie,err:=r.Cookie("user")
if err!=nil || strings.TrimSpace(cookie.Value)==""{
http.Error(w,"Forbidden",http.StatusForbidden)
return
}
fmt.Fprintf(w,"Welcome %s!",cookie.Value)
       }

