package photos

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/salaleser/vk-api/entity"
	"github.com/salaleser/vk-api/method"
	"github.com/salaleser/vk-api/util"
)

// SaveMarketPhoto Сохраняет фотографии после успешной загрузки на URI, полученный
// методом photos.getMarketUploadServer.
// group_id - идентификатор группы, для которой нужно загрузить фотографию.
// photo - параметр, возвращаемый в результате загрузки фотографии на сервер.
// server - параметр, возвращаемый в результате загрузки фотографии на сервер.
// hash - параметр, возвращаемый в результате загрузки фотографии на сервер.
// crop_data - параметр, возвращаемый в результате загрузки фотографии на сервер.
// crop_hash - параметр, возвращаемый в результате загрузки фотографии на сервер.
// После успешного выполнения возвращает массив, содержащий объект с загруженной фотографией.
func SaveMarketPhoto(
	groupID string,
	photo string,
	server string,
	hash string,
	cropData string,
	cropHash string) entity.MarketPhotoObject {
	params := url.Values{}
	params.Set("group_id", groupID)
	params.Set("photo", photo)
	params.Set("server", server)
	params.Set("hash", hash)
	params.Set("crop_data", cropData)
	params.Set("crop_hash", cropHash)
	params.Set("access_token", util.UserToken)

	b := util.GetHandler(method.PhotosSaveMarketPhoto, params)

	var o entity.MarketPhotoObject
	err := json.Unmarshal(b, &o)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
