package market

import (
	"encoding/json"
	"log"
	"net/url"

	"salaleser.ru/vk-api/entity"

	"salaleser.ru/vk-api/method"
	"salaleser.ru/vk-api/util"
)

// Get Возвращает список товаров в сообществе.
// owner_id - идентификатор владельца товаров.
// album_id - идентификатор подборки, товары из которой нужно вернуть.
// count - количество возвращаемых товаров.
// offset - смещение относительно первого найденного товара для выборки определенного подмножества.
// extended - 1 — будут возвращены дополнительные поля likes, can_comment, can_repost, photos, views_count. По умолчанию эти поля не возвращается.
// После успешного выполнения возвращает объект, содержащий число результатов в поле count и массив объектов товаров в поле items.
func Get(
	ownerID string,
	albumID string,
	count string,
	offset string,
	extended string) entity.ProductsObject {
	params := url.Values{}
	params.Set("owner_id", "-"+ownerID)
	params.Set("album_id", albumID)
	params.Set("count", count)
	params.Set("offset", offset)
	params.Set("extended", extended)
	params.Set("access_token", util.UserToken)

	r := util.GetHandler(method.MarketGet, params)

	var o entity.ProductsObject
	err := json.Unmarshal(r, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
