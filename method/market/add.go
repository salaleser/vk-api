package market

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/salaleser/vk-api/entity"
	"github.com/salaleser/vk-api/method"
	"github.com/salaleser/vk-api/util"
)

// Add Добавляет новый товар.
// owner_id - идентификатор владельца товара.
// name - название товара. Ограничение по длине считается в кодировке cp1251.
// description - описание товара.
// category_id - идентификатор категории товара.
// price - цена товара.
// deleted - статус товара (1 — товар удален, 0 — товар не удален).
// main_photo_id - идентификатор фотографии обложки товара.
// photo_ids - идентификаторы дополнительных фотографий товара.
// После успешного выполнения возвращает идентификатор добавленного товара.
func Add(
	ownerID string,
	name string,
	description string,
	categoryID string,
	price string,
	deleted string,
	mainPhotoID string,
	photoIDs string) entity.MarketProductObject {
	params := url.Values{}
	params.Set("owner_id", ownerID)
	params.Set("name", name)
	params.Set("description", description)
	params.Set("category_id", categoryID)
	params.Set("price", price)
	params.Set("deleted", deleted)
	params.Set("main_photo_id", mainPhotoID)
	params.Set("photo_ids", photoIDs)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.MarketAdd, params)

	var o entity.MarketProductObject
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
