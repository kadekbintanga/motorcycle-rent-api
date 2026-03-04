package constant

import "net/http"

const RequestIDKey = "API_CALL_ID"

type ResponseMap struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	ResponseStatusUnauthorized = "UNAUTHORIZED"
	ResponseStatusSuccess      = "SUCCESS"
	ResponseStatusFailed       = "FAILED"
	ResponseStatusBlocked      = "BLOCKED"
	ResponseStatusRejected     = "REJECTED"
	ResponseStatusBadRequest   = "BAD_REQUEST"
)

var (
	Res200Success                       = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Success"}
	Res200Save                          = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Success save data"}
	Res200Update                        = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Success update data"}
	Res200Get                           = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Success get data"}
	Res200Request                       = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Success request data"}
	Res200Delete                        = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Success delete data"}
	Res200SuccessVerifyOTP              = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "OTP Verified Successfully"}
	Res200SuccessRegister               = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Success register user"}
	Res200SuccessProfileCompletion      = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Profile Completed Successfully"}
	Res200SuccessRequestDoctor          = ResponseMap{Code: http.StatusOK, Status: ResponseStatusSuccess, Message: "Doctor has been successfully requested for approval"}
	Res201CreateScheduleUpdateRequest   = ResponseMap{Code: http.StatusCreated, Status: ResponseStatusSuccess, Message: "Success request update special schedule"}
	Res201DoctorTimeOffRequestSubmitted = ResponseMap{Code: http.StatusCreated, Status: ResponseStatusSuccess, Message: "Doctor time off request submitted"}
)

var (
	Res201Create        = ResponseMap{Code: http.StatusCreated, Status: ResponseStatusSuccess, Message: "Success create data"}
	Res201CreateRequest = ResponseMap{Code: http.StatusCreated, Status: ResponseStatusSuccess, Message: "Doctor has been successfully requested for approval"}
)

var (
	Res400Failed                = ResponseMap{Code: http.StatusBadRequest, Status: ResponseStatusFailed, Message: "Failed"}
	Res400InvalidPayload        = ResponseMap{Code: http.StatusBadRequest, Status: ResponseStatusBadRequest, Message: "Invalid payload data"}
	Res400PlateNumberExists     = ResponseMap{Code: http.StatusBadRequest, Status: ResponseStatusFailed, Message: "Plate number already exists"}
	Res404MotorcycleNotFound    = ResponseMap{Code: http.StatusNotFound, Status: ResponseStatusFailed, Message: "Motorcycle not found"}
	Res400InvalidMotorcycleUUID = ResponseMap{Code: http.StatusBadRequest, Status: ResponseStatusFailed, Message: "Invalid motorcycle ID"}
)

var (
	Res400InvalidEmailorPassword = ResponseMap{Code: http.StatusBadRequest, Status: ResponseStatusFailed, Message: "Invalid email or password"}
)

var Res401Unauthorized = ResponseMap{Code: http.StatusUnauthorized, Status: ResponseStatusUnauthorized, Message: "Unauthorized"}

var (
	Res422SomethingWentWrong = ResponseMap{Code: http.StatusUnprocessableEntity, Status: ResponseStatusFailed, Message: "Something Went Wrong"}
)

var ResMap422SomethingWentWrong = map[string]any{"status": ResponseStatusFailed, "message": "Something Went Wrong", "data": nil, "meta": nil}
