package market

import (
	"encoding/json"
	"log"
	"net/url"

	"salaleser.ru/vk-api/entity"

	"salaleser.ru/vk-api/method"
	"salaleser.ru/vk-api/util"
)

// GetAlbums Возвращает список подборок с товарами.
// owner_id - идентификатор владельца товаров.
// offset - смещение относительно первой найденной подборки для выборки определенного подмножества.
// count - количество возвращаемых подборок.
// После успешного выполнения возвращает объект, содержащий число результатов в поле count и массив объектов подборок в поле items.
func GetAlbums(
	ownerID string,
	offset string,
	count string) entity.AlbumObject {
	params := url.Values{}
	params.Set("owner_id", "-"+ownerID)
	params.Set("offset", offset)
	params.Set("count", count)
	params.Set("access_token", util.UserToken)

	r := util.GetHandler(method.MarketGetAlbums, params)

	var o entity.AlbumObject
	err := json.Unmarshal(r, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
