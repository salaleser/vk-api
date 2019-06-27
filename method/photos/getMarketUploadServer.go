package photos

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/salaleser/vk-api/entity"
	"github.com/salaleser/vk-api/method"
	"github.com/salaleser/vk-api/util"
)

// GetMarketUploadServer Возвращает адрес сервера для загрузки фотографии товара.
// После успешной загрузки Вы можете сохранить фотографию, воспользовавшись методом photos.saveMarketPhoto.
// group_id - идентификатор сообщества, для которого необходимо загрузить фотографию товара.
// main_photo - является ли фотография обложкой товара (1 — фотография для обложки, 0 — дополнительная фотография)
// crop_x - координата x для обрезки фотографии (верхний правый угол).
// crop_y - координата y для обрезки фотографии (верхний правый угол).
// crop_width - ширина фотографии после обрезки в px.
// После успешного выполнения возвращает объект с единственным полем upload_url.
func GetMarketUploadServer(
	groupID string,
	mainPhoto string,
	cropX string,
	cropY string,
	cropWidth string) entity.UploadURLObject {
	params := url.Values{}
	params.Set("group_id", groupID)
	params.Set("main_photo", mainPhoto)
	params.Set("crop_x", cropX)
	params.Set("crop_y", cropY)
	params.Set("crop_width", cropWidth)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.PhotosGetMarketUploadServer, params)

	var o entity.UploadURLObject
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
