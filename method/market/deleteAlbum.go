package market

import (
	"encoding/json"
	"log"
	"net/url"

	"salaleser.ru/vk-api/entity"
	"salaleser.ru/vk-api/method"
	"salaleser.ru/vk-api/util"
)

// DeleteAlbum Удаляет подборку с товарами.
// owner_id - идентификатор владельца подборки.
// album_id - идентификатор подборки.
// После успешного выполнения возвращает 1.
func DeleteAlbum(
	ownerID string,
	albumID string) entity.SuccessObject {
	params := url.Values{}
	params.Set("owner_id", "-"+ownerID)
	params.Set("album_id", albumID)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.MarketDeleteAlbum, params)

	var o entity.SuccessObject
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
