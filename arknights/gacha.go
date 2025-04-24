package arknights

type Gacha struct {
	Ts   int64
	Pool string
}

func GachaPage(pageNo int, pageSize int) []Gacha {
	return []Gacha{}
}
