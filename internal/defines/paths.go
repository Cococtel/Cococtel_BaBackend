package defines

const (
	//Ping & consts
	PingPath = "/ping"
	IDPath   = "/:id"

	//Registro y Login
	RegisterPath        = "/register"
	LoginPath           = "/login"
	ValidateLoginPath   = "/validate"
	GetQRDoubleAuthPath = "/qr"
	NotifyQRReadPath    = "/notify" + IDPath
	VerifyPath          = "/verify"
	ProfilePath         = "/profile"
	//Token               = "/token"
)
