package config

const (
	//auth
	UserSesion = "user"
	RoleSesion = "roleUser"

	DownloadReport = "/download"
	SendReport     = "/send"
	UserAdmin      = "510068c3-8172-48ce-8d5b-ecb3de591b51"

	AuthGroup        = "/auth"
	AuthRegister     = "/register"
	AuthLogin        = "/login"
	AuthRefreshToken = "/refresh-token"

	//User
	UserGroup  = "/users"
	UserPost   = "/"
	UserGet    = "/:id"
	UserDelete = "/:id"
	UserGetAll = "/"
	UserUpdate = "/"

	//booking
	BookingGroup          = "/booking"
	BookingPost           = "/"
	BookingGet            = "/:id"
	BookingGetAll         = "/"
	BookingGetAllByStatus = "/status/:status"
	Approval              = "/approval"

	//room
	RoomGroup         = "/rooms"
	RoomPost          = "/create"
	RoomGetByroomType = "/" //query
	RoomGetAll        = "/get"
	RoomGetById       = "/:id"
	RoomGetByStatus   = "/status"
	RoomDelete        = "/:id"
	RoomUpdate        = "/:id"
	RoomUpdateStatus  = "/status/:id"
)
