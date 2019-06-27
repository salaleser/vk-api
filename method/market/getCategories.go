package market

import (
	"encoding/json"
	"log"
	"net/url"

	"salaleser.ru/vk-api/entity"

	"salaleser.ru/vk-api/method"
	"salaleser.ru/vk-api/util"
)

// GetCategories Возвращает список категорий для товаров.
// count - количество категорий, информацию о которых необходимо вернуть.
// offset - смещение, необходимое для выборки определенного подмножества категорий.
// После успешного выполнения возвращает список объектов category.
func GetCategories(
	count string,
	offset string) entity.CategoriesObject {
	params := url.Values{}
	params.Set("count", count)
	params.Set("offset", offset)
	params.Set("access_token", util.UserToken)

	r := util.GetHandler(method.MarketGetCategories, params)

	var o entity.CategoriesObject
	err := json.Unmarshal(r, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
