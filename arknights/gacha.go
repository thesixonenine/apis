package arknights

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

type Gacha struct {
    CharName string
    GachaTs  string
    IsNew    bool
    PoolName string
    Rarity   int
}
type User struct {
    UID          string
    AccountToken string
    RoleToken    string
    Cookie       string
}

func (u *User) GachaPage(pageNo int, pageSize int) []Gacha {
    client := &http.Client{}

    req, err := http.NewRequest("GET", fmt.Sprintf("https://ak.hypergryph.com/user/api/inquiry/gacha/history?uid=%s&category=classic&size=10", u.UID), nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("x-account-token", u.AccountToken)
    req.Header.Set("x-role-token", u.RoleToken)
    req.Header.Set("cookie", u.Cookie)
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    bodyText, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", bodyText)
    return []Gacha{}
}
