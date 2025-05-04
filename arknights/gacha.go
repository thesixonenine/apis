package arknights

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)
type Category string

const (
    SpringFest  Category = "spring_fest"
    AnniverFest Category = "anniver_fest"
    Normal      Category = "normal"
    Classic     Category = "classic"
)

var CategoryList = []Category{SpringFest, AnniverFest, Normal, Classic}

type GachaData struct {
    HasMore bool
    List    []Gacha
}
type Gacha struct {
    CharId   string `json:"charId"`
    CharName string `json:"charName"`
    GachaTs  string `json:"gachaTs"`
    IsNew    bool   `json:"isNew"`
    PoolId   string `json:"poolId"`
    PoolName string `json:"poolName"`
    Pos      int    `json:"pos"`
    Rarity   int    `json:"rarity"`
}
type GachaCate struct {
    Id string `json:"id"`
    Name string `json:"name"`
}
type User struct {
    UID          string
    AccountToken string
    RoleToken    string
    Cookie       string
}

// NewUser create user from url
func NewUser(url string) *User {
    // https://web-api.hypergryph.com/account/info/hg
    return &User{}
}
func (u *User) DoReq(req *http.Request)(*http.Response, error) {
    req.Header.Set("x-account-token", u.AccountToken)
    req.Header.Set("x-role-token", u.RoleToken)
    req.Header.Set("cookie", u.Cookie)
    return (&http.Client{}).Do(req)
}
// Cate fetch all category
func (u *User) Cate() []GachaCate {
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://ak.hypergryph.com/user/api/inquiry/gacha/cate?uid=%s", u.UID), nil)
    resp, err := u.DoReq(req)
    if err != nil {
        log.Fatal(err)
        return []GachaCate{}
    }
    defer resp.Body.Close()
    bodyText, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
        return []GachaCate{}
    }
    fmt.Printf("%s\n", bodyText)
    respJson := Resp[[]GachaCate]{}
    _ = json.Unmarshal(bodyText, &respJson)
    if respJson.Code != 0 {
        log.Fatalf("查询失败code[%d]", respJson.Code)
        return []GachaCate{}
    }
    return respJson.Data
}
// GachaPage fetch gacha
func (u *User) GachaPage(cate string, pos int, gachaTs string) ([]Gacha, bool) {
    more := ""
    if gachaTs != "" {
        more = fmt.Sprintf("&pos=%d&gachaTs=%s", pos, gachaTs)
    }
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://ak.hypergryph.com/user/api/inquiry/gacha/history?uid=%s&category=%s%s&size=10", u.UID, cate, more), nil)
    resp, err := u.DoReq(req)
    if err != nil {
        log.Fatal(err)
        return []Gacha{}, false
    }
    defer resp.Body.Close()
    bodyText, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
        return []Gacha{}, false
    }
    fmt.Printf("%s\n", bodyText)
    respJson := Resp[GachaData]{}
    _ = json.Unmarshal(bodyText, &respJson)
    if respJson.Code != 0 {
        log.Fatalf("查询失败code[%d]", respJson.Code)
        return []Gacha{}, false
    }
    return respJson.Data.List, respJson.Data.HasMore
}
// GachaAll fetch all gacha
func (u *User) GachaAll(cate string) []Gacha {
    page, hasMore := u.GachaPage(cate, 0, "")
    for hasMore {
        time.Sleep(500 * time.Millisecond)
        gachas, b := u.GachaPage(cate, page[len(page)-1].Pos, page[len(page)-1].GachaTs)
        hasMore = b
        page = append(page, gachas...)
    }
    return page
}
