package market

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/salaleser/vk-api/entity"
	"github.com/salaleser/vk-api/method"
	"github.com/salaleser/vk-api/util"
)

// Delete Удаляет товар.
// owner_id - идентификатор владельца товара.
// item_id - идентификатор товара.
// После успешного выполнения возвращает 1.
func Delete(
	ownerID string,
	itemID string) entity.SuccessObject {
	params := url.Values{}
	params.Set("owner_id", "-"+ownerID)
	params.Set("item_id", itemID)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.MarketDelete, params)

	var o entity.SuccessObject
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
