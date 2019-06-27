package market

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/salaleser/vk-api/entity"
	"github.com/salaleser/vk-api/method"
	"github.com/salaleser/vk-api/util"
)

// AddAlbum Добавляет новую подборку с товарами.
// owner_id — идентификатор владельца подборки.
// title — eназвание подборки.
// photo_id — идентификатор фотографии-обложки подборки.
// main_album — назначить подборку основной (1 — назначить, 0 — нет).
// После успешного выполнения возвращает идентификатор созданной подборки.
func AddAlbum(
	ownerID string,
	title string,
	photoID string,
	mainAlbum string) entity.MarketAlbumResponse {
	params := url.Values{}
	params.Set("owner_id", "-"+ownerID)
	params.Set("title", title)
	params.Set("photo_id", photoID)
	params.Set("main_album", mainAlbum)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.MarketAddAlbum, params)

	var o entity.MarketAlbumResponse
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
