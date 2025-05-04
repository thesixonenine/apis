package arknights

import (
    "encoding/json"
    "fmt"
    "testing"
)

func TestGachaPage(t *testing.T) {
    user := NewUser("")
    user.UID = ""
    user.RoleToken=""
    user.AccountToken=""
    user.Cookie=""

    for _, cate := range user.Cate() {
        page := user.GachaAll(cate.Id)
        pageStr, _ := json.Marshal(page)
        fmt.Println(string(pageStr))
        fmt.Println("######")
    }

    //for _,cate := range CategoryList {
    //    page := user.GachaAll(string(cate))
    //    pageStr, _ := json.Marshal(page)
    //    fmt.Println(string(pageStr))
    //    fmt.Println("######")
    //}
}
