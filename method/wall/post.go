package wall

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"salaleser.ru/vk-api/entity"
	"salaleser.ru/vk-api/method"
	"salaleser.ru/vk-api/util"
)

// Post Позволяет создать запись на стене, предложить запись на стене публичной страницы, опубликовать существующую отложенную запись.
// owner_id — идентификатор пользователя или сообщества, на стене которого должна быть опубликована запись.
// friends_only — 1 — запись будет доступна только друзьям, 0 — всем пользователям. По умолчанию публикуемые записи доступны всем пользователям.
// from_group — данный параметр учитывается, если owner_id < 0 (запись публикуется на стене группы). 1 — запись будет опубликована от имени группы, 0 — запись будет опубликована от имени пользователя (по умолчанию).
// message — текст сообщения (является обязательным, если не задан параметр attachments)
// ... TODO
// После успешного выполнения возвращает идентификатор созданной записи (post_id).
func Post(
	ownerID string,
	friendsOnly string,
	fromGroup string,
	message string) entity.PostObject {
	params := url.Values{}
	params.Set("owner_id", ownerID)
	params.Set("friends_only", friendsOnly)
	params.Set("from_group", fromGroup)
	params.Set("message", message)
	params.Set("access_token", util.UserToken)
	params.Set("v", "5.85")
	util.RemoveEmptyParams(params)
	uri := fmt.Sprintf(util.ApiURL, method.WallPost, params.Encode())

	r := util.PostHandler(uri, "", "")

	var o entity.PostObject
	err := json.Unmarshal(r, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
