package response

/**
 * Default Error response
 */
type (
	ErrorResponse struct {
		Error ErrorBody `json:"error"`
	}

	ErrorBody struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Status  HttpStatus        `json:"status"`
		Details map[string]string `json:"details"`
	}

	HttpStatus string
)

func NewError(err error, status HttpStatus, details map[string]string) (int, ErrorResponse) {
	errorBody := new(ErrorBody)

	errorBody.Code = status.Code()
	errorBody.Message = err.Error() // TODO: not to show error message directly
	errorBody.Status = status
	errorBody.Details = details

	return status.Code(), ErrorResponse{*errorBody}
}

/**
 * codes below are following https://cloud.google.com/apis/design/errors?hl=ja
 **/
const (
	OK                  = "OK"
	INVALID_ARGUMENT    = "INVALID_ARGUMENT"
	FAILED_PRECONDITION = "FAILED_PRECONDITION"
	OUT_OF_RANGE        = "OUT_OF_RANGE"
	UNAUTHENTICATED     = "UNAUTHENTICATED"
	PERMISSION_DENIED   = "PERMISSION_DENIED"
	NOT_FOUND           = "NOT_FOUND"
	ABORTED             = "ABORTED"
	ALREADY_EXISTS      = "ALREADY_EXISTS"
	RESOURCE_EXHAUSTED  = "RESOURCE_EXHAUSTED"
	CANCELLED           = "CANCELLED"
	DATA_LOSS           = "DATA_LOSS"
	UNKNOWN             = "UNKNOWN"
	INTERNAL            = "INTERNAL"
	NOT_IMPLEMENTED     = "NOT_IMPLEMENTED"
	UNAVAILABLE         = "UNAVAILABLE"
	DEADLINE_EXCEEDED   = "DEADLINE_EXCEEDED"
)

func (c HttpStatus) Code() int {
	switch c {
	case OK:
		return 200
	case INVALID_ARGUMENT:
		return 400
	case FAILED_PRECONDITION:
		return 400
	case OUT_OF_RANGE:
		return 400
	case UNAUTHENTICATED:
		return 401
	case PERMISSION_DENIED:
		return 403
	case NOT_FOUND:
		return 404
	case ABORTED:
		return 409
	case ALREADY_EXISTS:
		return 409
	case RESOURCE_EXHAUSTED:
		return 429
	case CANCELLED:
		return 499
	case DATA_LOSS:
		return 500
	case UNKNOWN:
		return 500
	case INTERNAL:
		return 500
	case NOT_IMPLEMENTED:
		return 501
	case UNAVAILABLE:
		return 503
	case DEADLINE_EXCEEDED:
		return 504
	}

	return 500
}
