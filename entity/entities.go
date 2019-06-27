package entity

// ProfileInfoObject описывает объект, вернувшийся на запрос users.get
type ProfileInfoObject struct {
	R []ProfileInfo `json:"response"`
}

// ProfileInfo описывает профиль
type ProfileInfo struct {
	ID              int      `json:"id"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	MaidenName      string   `json:"maiden_name"`
	ScreenName      string   `json:"screen_name"`
	Status          string   `json:"status"`
	IsClosed        bool     `json:"is_closed"`
	CanAccessClosed bool     `json:"can_access_closed"`
	IsFriend        int      `json:"is_friend"`
	Online          int      `json:"online"`
	LastSeen        LastSeen `json:"last_seen"`
	Blacklisted     int      `json:"blacklisted"`
	BlacklistedByMe int      `json:"blacklisted_by_me"`
	Relation        int      `json:"relation"`
	// RelationPartner RelationPartner    `json:"relation_partner"` // TODO
	RelationPending  int   `json:"relation_pending"`
	RelationRequests []int `json:"relation_requests"`
}

// LastSeen описывает профиль
type LastSeen struct {
	Time     int `json:"time"`
	Platform int `json:"platform"`
}

// AlbumObject описывает объект, вернувшийся на запрос groups.getMembers
type AlbumObject struct {
	R AlbumResponse `json:"response"`
}

// AlbumResponse описывает объект, вернувшийся на запрос groups.getMembers
type AlbumResponse struct {
	Count int     `json:"count"`
	Items []Album `json:"items"`
}

// Album описывает карточку товара
type Album struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Count       int    `json:"count"`
	UpdatedTime int    `json:"updated_time"`
}

type ProductsObject struct {
	R ProductsResponse `json:"response"`
}

type ProductsResponse struct {
	Count int       `json:"count"`
	Items []Product `json:"items"`
}

type Product struct {
	ID           int      `json:"id"`
	OwnerID      int      `json:"owner_id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Price        Price    `json:"price"`
	Category     Category `json:"category"`
	Date         int      `json:"date"`
	ThumbPhoto   string   `json:"thumb_photo"`
	Availability int      `json:"availability"`
}

type Price struct {
	Amount   string   `json:"amount"`
	Currency Currency `json:"currency"`
	Text     string   `json:"text"`
}

type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Section Section `json:"section"`
}

type Section struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CategoriesObject описывает JSON-объект, вернувшийся на запрос market.getCategories
type CategoriesObject struct {
	R CategoriesResponse `json:"response"`
}

type CategoriesResponse struct {
	Count int        `json:"count"`
	Items []Category `json:"items"`
}

type MembersObject struct {
	R MembersResponse `json:"response"`
}

type MembersResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

type UploadURLObject struct {
	R UploadURLResponse `json:"response"`
}

type UploadURLResponse struct {
	UploadURL string `json:"upload_url"`
}

// PostObject описывает JSON-объект "Запись на стене"
// Позволяет создать запись на стене, предложить запись на стене публичной страницы, опубликовать
// существующую отложенную запись.
type PostObject struct {
	R PostResponse `json:"response"`
}

type PostResponse struct {
	PostID int `json:"post_id"`
}

// MarketPhotoFileObject описывает JSON-объект с полями server, photo, hash, crop_data, crop_hash
type MarketPhotoFileObject struct {
	Server   int    `json:"server"`
	Photo    string `json:"photo"`
	Hash     string `json:"hash"`
	CropData string `json:"crop_data"`
	CropHash string `json:"crop_hash"`
}

// MarketPhotoObject описывает JSON-объект Массив, содержащий объект с загруженной фотографией
type MarketPhotoObject struct {
	R []MarketPhotoResponse `json:"response"`
}

// MarketPhotoResponse описывает JSON-объект с загруженной фотографией
type MarketPhotoResponse struct {
	ID       int    `json:"id"`
	AlbumID  int    `json:"album_id"`
	OwnerID  int    `json:"owner_id"`
	UserID   int    `json:"user_id"`
	Photo75  string `json:"photo_75"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Text     string `json:"text"`
	Date     int    `json:"date"`
}

// MarketProductObject описывает JSON-объект Товар (вернувшийся после успешного
// завершения запроса методом market.add)
type MarketProductObject struct {
	R MarketProductResponse `json:"response"`
}

// MarketProductResponse описывает JSON-объект Товар (вернувшийся после успешного
// завершения запроса методом market.add)
type MarketProductResponse struct {
	MarketItemID int `json:"market_item_id"`
}

type SuccessObject struct {
	R int `json:"response"`
}

type MarketAlbumResponse struct {
	R MarketAlbumObject `json:"response"`
}

type MarketAlbumObject struct {
	MarketAlbumID int `json:"market_album_id"`
}

type Response struct {
	Response int `json:"response"`
}
