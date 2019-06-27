package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/salaleser/poster/logger"

	"github.com/salaleser/vk-api/entity"
)

const (
	V      = "5.85"
	ApiURL = "https://api.vk.com/method/%s?%s"
)

var (
	UserToken string
	DebugMode = false
)

// Error описывает JSON-объект Ошибка, вернувшийся от vk.com
type Error struct {
	Error Content `json:"error"`
}

// Content описывает содержимое JSON-объекта Ошибка
type Content struct {
	ErrorCode     int                `json:"error_code"`
	ErrorMsg      string             `json:"error_msg"`
	RequestParams []RequestParameter `json:"request_params"`
}

// RequestParameter описывает параметры запроса JSON-объекта Ошибка
type RequestParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PostHandler отправляет POST-запрос
func PostHandler(uri string, method string, filename string) []byte {
	if DebugMode {
		log.Println("[DEBUG] POST-request:", uri)
	}

	var contentType string
	body := &bytes.Buffer{}
	if len(filename) > 0 {
		fileDir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		filePath := path.Join(fileDir, filename)

		file, err := os.Open(filePath)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		writer := multipart.NewWriter(body)
		contentType = writer.FormDataContentType()
		part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
		if err != nil {
			log.Println(err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			log.Println(err)
		}
		writer.Close()
	}

	time.Sleep(500 * time.Millisecond)
	resp, err := http.Post(uri, contentType, body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return convert(*resp)
}

// GetHandler отправляет GET-запрос
func GetHandler(method string, params url.Values) []byte {
	RemoveEmptyParams(params)
	params.Set("v", V)
	uri := fmt.Sprintf(ApiURL, method, params.Encode())
	if DebugMode {
		log.Println("[DEBUG] GET-request:", uri)
	}

	time.Sleep(500 * time.Millisecond)
	resp, err := http.Get(uri)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	return convert(*resp)
}

// RemoveEmptyParams удаляет пустые параметры из "net/url".Values
func RemoveEmptyParams(params url.Values) {
	for key, val := range params {
		for _, v := range val {
			if len(v) == 0 {
				params.Del(key)
			}
		}
	}
}

func convert(resp http.Response) []byte {
	var b []byte
	for {
		bs := make([]byte, 1000000)
		n, err := resp.Body.Read(bs)
		b = bs[:n]
		if n == 0 || err != nil {
			break
		}
	}

	// var o Error
	// err := json.Unmarshal(b, &o)
	// if err == nil {
	// 	log.Fatalf("[ERROR] %d --- %s", o.Error.ErrorCode, o.Error.ErrorMsg)
	// }

	if DebugMode {
		log.Println("[DEBUG] Response:", string(b))
	}
	return b
}

// DownloadFile скачивает файл с указанного адреса на локальный диск
func DownloadFile(url string, filepath string) {
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

// UploadFile загружает файл на сервер VK
// После успешной загрузки сервер возвращает в ответе JSON-объект с полями server, photo, hash, crop_data, crop_hash
func UploadFile(uri string, filename string) entity.MarketPhotoFileObject {
	b := PostHandler(uri, "", filename)

	var o entity.MarketPhotoFileObject
	err := json.Unmarshal(b, &o)
	if err != nil {
		logger.Appendf("Ошибка при попытке загрузить файл %q на сервер (%s)", filename, err)
	}
	return o
}
