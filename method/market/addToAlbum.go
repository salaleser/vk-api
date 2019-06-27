package market

import (
	"encoding/json"
	"log"
	"net/url"

	"salaleser.ru/vk-api/entity"
	"salaleser.ru/vk-api/method"
	"salaleser.ru/vk-api/util"
)

// AddToAlbum Добавляет товар в одну или несколько выбранных подборок.
// owner_id — идентификатор владельца товара.
// item_id — идентификатор товара.
// album_ids — идентификаторы подборок, в которые нужно добавить товар.
// После успешного выполнения возвращает 1.
func AddToAlbum(
	ownerID string,
	itemID string,
	albumIDs string) entity.Response {
	params := url.Values{}
	params.Set("owner_id", "-"+ownerID)
	params.Set("item_id", itemID)
	params.Set("album_ids", albumIDs)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.MarketAddToAlbum, params)

	var o entity.Response
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
