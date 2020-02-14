package rerror

var (
	errMap map[int]string
)

//error code
const (
	CustomCode = iota + 1
	UnmarshalCode
	MarshalCode
	BadReqCode
	ForbiddenCode
	UnauthorizedCode
	NotFoundCode
	MiscCode
	ConnCode
	TxCode
	CreateCode
	SelectCode
	InsertCode
	UpdateCode
	DeleteCode
	DropCode
	ExecCode
)

//msg constant
const (
	InvalidReqMsg   = "Invalid Request, Please contact system adminstrator for further clarification."
	ForbiddenMsg    = "Authentication failed due to invalid authentication credentials"
	UnauthorizedMsg = "Missing authorization header"
	BlandReqMsg     = "Blank request, Please provide input to process"
	NotFoundMsg     = "Record not found"
	ServerErrorMsg  = "Sorry unable to process this request. Please Try Again"
)

func init() {
	errMap = map[int]string{
		UnmarshalCode:    InvalidReqMsg,
		MarshalCode:      InvalidReqMsg,
		BadReqCode:       InvalidReqMsg,
		ForbiddenCode:    ForbiddenMsg,
		UnauthorizedCode: ForbiddenMsg,
		NotFoundCode:     NotFoundMsg,
		MiscCode:         ServerErrorMsg,
		ConnCode:         ServerErrorMsg,
		TxCode:           ServerErrorMsg,
		CreateCode:       ServerErrorMsg,
		SelectCode:       ServerErrorMsg,
		InsertCode:       ServerErrorMsg,
		UpdateCode:       ServerErrorMsg,
		DeleteCode:       ServerErrorMsg,
		DropCode:         ServerErrorMsg,
		ExecCode:         ServerErrorMsg,
	}
}
