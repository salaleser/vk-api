package users

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/salaleser/vk-api/entity"
	"github.com/salaleser/vk-api/method"

	"github.com/salaleser/vk-api/util"
)

// Get Возвращает расширенную информацию о пользователях.
// user_ids - перечисленные через запятую идентификаторы пользователей или их короткие имена (screen_name). По умолчанию — идентификатор текущего пользователя.
// fields - список дополнительных полей профилей, которые необходимо вернуть.
// name_case - падеж для склонения имени и фамилии пользователя.
// После успешного выполнения возвращает массив объектов пользователей.
func Get(
	userIDs string,
	fields string,
	nameCase string) entity.ProfileInfoObject {
	params := url.Values{}
	params.Set("user_ids", userIDs)
	params.Set("fields", fields)
	params.Set("name_case", nameCase)
	params.Set("access_token", util.UserToken)

	r := util.GetHandler(method.UsersGet, params)

	var o entity.ProfileInfoObject
	err := json.Unmarshal(r, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
