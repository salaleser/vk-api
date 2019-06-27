package method

// Группа методов Users
const (
	UsersGet = "users.get"
)

// Группа методов Market
const (
	MarketAdd           = "market.add"
	MarketGet           = "market.get"
	MarketGetAlbums     = "market.getAlbums"
	MarketAddAlbum      = "market.addAlbum"
	MarketAddToAlbum    = "market.addToAlbum"
	MarketDelete        = "market.delete"
	MarketDeleteAlbum   = "market.deleteAlbum"
	MarketGetCategories = "market.getCategories"
)

// Группа методов Groups
const (
	GroupsGetMembers = "groups.getMembers"
)

// Группа методов Photos
const (
	PhotosGetMarketUploadServer = "photos.getMarketUploadServer"
	PhotosSaveMarketPhoto       = "photos.saveMarketPhoto"
)

// Группа методов Wall
const (
	WallPost = "wall.post"
)
