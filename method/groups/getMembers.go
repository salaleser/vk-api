package groups

import (
	"encoding/json"
	"log"
	"net/url"

	"salaleser.ru/vk-api/entity"

	"salaleser.ru/vk-api/method"

	"salaleser.ru/vk-api/util"
)

// GetMembers Возвращает список участников сообщества.
// group_id - идентификатор или короткое имя сообщества.
// sort - сортировка, с которой необходимо вернуть список участников.
// offset - смещение, необходимое для выборки определенного подмножества участников. По умолчанию 0.
// count - количество участников сообщества, информацию о которых необходимо получить.
// fields - список дополнительных полей, которые необходимо вернуть.
//     Доступные значения: sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200,
//     photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile,
//     contacts, connections, site, education, universities, schools, can_post, can_see_all_posts,
//     can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives
// filter -
//     friends — будут возвращены только друзья в этом сообществе.
//     unsure — будут возвращены пользователи, которые выбрали «Возможно пойду» (если сообщество относится к мероприятиям).
//     managers — будут возвращены только руководители сообщества (доступно при запросе с передачей access_token от имени администратора сообщества).
// После успешного выполнения возвращает объект, содержащий число результатов в поле count (integer) и массив идентификаторов пользователей в поле items ([integer]).
//     Если был передан параметр filter=managers, возвращается дополнительное поле role (string), которое содержит уровень полномочий руководителя:
//     moderator — модератор;
//     editor — редактор;
//     administrator — администратор;
//     creator — создатель сообщества.
func GetMembers(
	groupID string,
	sort string,
	offset string,
	count string,
	fields string,
	filter string) entity.MembersObject {
	params := url.Values{}
	params.Set("group_id", groupID)
	params.Set("sort", sort)
	params.Set("offset", offset)
	params.Set("count", count)
	params.Set("fields", fields)
	params.Set("filter", filter)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.GroupsGetMembers, params)

	var o entity.MembersObject
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
