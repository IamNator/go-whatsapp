//source: https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes

package error_code

type MetaError string

func (e MetaError) Error() string {
	return string(e)
}

const (

	//ErrAuthException
	//
	//If no subcode is present, the login status or access
	//token has expired, been
	//revoked, or is otherwise invalid. If a subcode is present, see the subcode.
	//
	//Possible solution: Get a new access token
	ErrAuthException MetaError = "unable to authenticate request"

	//ErrAPIUnknown
	//
	//Possibly a temporary issue due to downtime. Wait and retry the operation.
	//
	//
	//Visit the WhatsApp Business API Status Dashboard and check to make sure there
	//are no typos in your API call.
	ErrAPIUnknown MetaError = "unknown api error"

	//ErrAPIService
	//
	//Temporary issue due to downtime. Wait and retry the operation.
	//
	//
	//Visit the WhatsApp Business API Status Dashboard.
	ErrAPIService MetaError = "service error, please try again later"

	//ErrAPITooManyCalls
	//
	//summary: You have exceeded the rate limit for your account.
	//Wait and retry the operation.
	//
	//Description:
	//
	//Temporary issue due to throttling.
	//Wait and retry the operation, or examine your API request volume.
	ErrAPITooManyCalls MetaError = "too many calls"

	// ErrAPIMethod
	//
	//
	//summary: Capability or permissions issue.
	//
	//description:
	//Capability or permissions issue.
	//
	//Visit the endpoint reference to ensure you are including the needed permission in your call.
	ErrAPIMethod MetaError = "capability or permissions issue."

	//ErrAPIPermissionDenied
	//
	//Permission is either not granted or has been removed.
	//Learn how to handle missing permissions.
	ErrAPIPermissionDenied MetaError = "permission denied"

	//ErrAPIInvalidParameter
	//
	//summary: Invalid parameter.
	//
	//description:
	//The parameter may not be available or may be spelled incorrectly.
	//
	//Possible solution: Visit the endpoint reference to ensure the parameter exists.
	ErrAPIInvalidParameter MetaError = "invalid parameter"

	//ErrAccessTokenHasExpired
	//
	//summary: Access token has expired.
	//
	//description:
	//Your access token has expired.
	//
	//Possible solution: Get a new access token.
	ErrAccessTokenHasExpired MetaError = "access token has expired"

	//ErrAPIPermission
	//
	//summary: Capability or permissions issue.
	ErrAPIPermission MetaError = "permission revoked"
)

var (
	errCodeMap = map[MetaError]int{
		ErrAuthException:       0,
		ErrAPIUnknown:          1,
		ErrAPIService:          2,
		ErrAPIMethod:           3,
		ErrAPITooManyCalls:     4,
		ErrAPIPermissionDenied: 10,
		ErrAPIInvalidParameter: 100,
	}

	codeErrorMap = map[int]MetaError{
		0:   ErrAuthException,
		1:   ErrAPIUnknown,
		2:   ErrAPIService,
		3:   ErrAPIMethod,
		4:   ErrAPITooManyCalls,
		10:  ErrAPIPermissionDenied,
		100: ErrAPIInvalidParameter,
	}
)

func (e MetaError) Code() int {
	return errCodeMap[e]
}

func Error(code int) MetaError {
	//if code > 100 && code < 200 {
	//	return ErrAPIInvalidParameter
	//}
	return codeErrorMap[code]
}

func IsError(err error) bool {
	_, ok := err.(MetaError)
	return ok
}
