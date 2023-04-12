// source: https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes
package errors

type MetaError string

func (e MetaError) Error() string {
	return string(e)
}

const (

	// ################ Authorization Errors ################

	// ErrAuthException
	//
	// HTTP 401
	//
	// code: 0
	//
	// If no subcode is present, the login status or access
	// token has expired, been
	// revoked, or is otherwise invalid. If a subcode is present, see the subcode.
	//
	// Possible solution: Get a new access token
	ErrAuthException MetaError = "unable to authenticate request"

	// ErrAPIMethod
	//
	// HTTP 500
	//
	// code: 3
	//
	// summary: Capability or permissions issue.
	//
	// description:
	// Capability or permissions issue.
	//
	// Visit the endpoint reference to ensure you are including the needed permission in your call.
	ErrAPIMethod MetaError = "capability or permissions issue."

	// ErrAPIPermissionDenied
	//
	// HTTP 403
	//
	// code: 10
	//
	// Permission is either not granted or has been removed.
	// Learn how to handle missing permissions.
	ErrAPIPermissionDenied MetaError = "permission denied"

	// ErrAccessTokenHasExpired
	//
	// HTTP 401
	//
	// code: 190
	//
	// summary: Access token has expired.
	//
	// description:
	// Your access token has expired or is invalid.
	//
	// Possible solution: Get a new access token.
	ErrAccessTokenHasExpired MetaError = "access token has expired or is invalid"

	// ErrAPIPermission
	//
	// HTTP 403
	//
	// code: 200-299
	//
	// summary: Capability or permissions issue.
	ErrAPIPermission MetaError = "permission revoked"

	// ###################### Throttling Errors ######################

	// ErrAPITooManyCalls
	//
	// HTTP 429
	//
	// code: 4
	//
	// summary: The app has reached its API call rate limit..
	// Wait and retry the operation.
	//
	// Description:
	//
	// Temporary issue due to throttling.
	// Wait and retry the operation, or examine your API request volume.
	ErrAPITooManyCalls MetaError = "too many calls"

	// ErrRateLimitIssues
	//
	// HTTP 429
	//
	// code: 80007
	//
	// summary: The WhatsApp Business Account has reached its rate limit.
	//
	// possible solution: See WhatsApp Business Account Rate Limits.
	// Try again later or reduce the frequency or amount of API queries the app is making.
	ErrRateLimitIssues MetaError = "rate limit issues"

	// ErrRateLimitHit
	//
	// HTTP 429
	//
	// code: 130429
	//
	// summary: Cloud API message throughput has been reached.
	//
	// possible solution: The app has reached the API's throughput limit.
	// See Throughput. Try again later or reduce the frequency with which the app sends messages.
	ErrRateLimitHit MetaError = "rate limit hit"

	// ErrSpamRateLimitHit
	//
	// HTTP 429
	//
	// code: 131048
	//
	// summary: Message failed to send because there are restrictions on how many messages can be sent from this phone number.
	// This may be because too many previous messages were blocked or flagged as spam.
	ErrSpamRateLimitHit MetaError = "spam rate limit hit"

	// ErrPairRateLimitHit
	//
	// HTTP 429
	//
	// code: 131056
	//
	// description: Too many messages sent from the sender phone number to the same recipient phone number in a short period of time.
	//
	// possible solutions:
	// Wait and retry the operation, if you intend to send messages to the same phone number.
	// You can still send messages to a different phone number without waiting
	ErrPairRateLimitHit MetaError = "(business account, consumer account) pair rate limit hit"

	//  ###################### Integrity Errors ######################

	// ErrTemporarilyBlocked
	//
	// HTTP 403
	//
	// code 368
	//
	// The WhatsApp Business Account associated with the app has been restricted or disabled
	// for violating a platform policy. Visit Policy Enforcement document to learn about policy violations and how to resolve them.
	ErrTemporarilyBlocked MetaError = "temporarily blocked for policies violations"

	// ErrAccountLocked
	//
	// HTTP 403
	//
	// code 131031
	//
	// The WhatsApp Business Account associated with the app has been restricted or disabled
	// for violating a platform policy, or the request data could not be verified against data set on the
	// WhatsApp Business Account (e.g., the two-step pin included in the request is incorrect).
	// Visit Policy Enforcement document to learn about policy violations and how to resolve them.
	ErrAccountLocked MetaError = "account has been locked"

	//  ###################### Other Errors ######################

	// ErrAPIUnknown
	//
	// HTTP 400
	//
	// code: 1
	//
	// Possibly a temporary issue due to downtime. Wait and retry the operation.
	//
	//
	// Visit the WhatsApp Business API Status Dashboard and check to make sure there
	// are no typos in your API call.
	ErrAPIUnknown MetaError = "unknown api error"

	// ErrAPIService
	//
	// HTTP 503
	//
	// code 2
	//
	// Temporary issue due to downtime. Wait and retry the operation.
	//
	//
	// Visit the WhatsApp Business API Status Dashboard.
	ErrAPIService MetaError = "service error, please try again later"

	// ErrAPIInvalidParameter
	//
	// HTTP 400
	//
	// code 100
	//
	// summary: Invalid parameter.
	//
	// description:
	// The parameter may not be available or may be spelled incorrectly.
	//
	// Possible solution: Visit the endpoint reference to ensure the parameter exists.
	ErrAPIInvalidParameter MetaError = "invalid parameter"

	// ErrSomethingWentWrong
	//
	// HTTP 500
	//
	// code 131000
	//
	// sumary: Message failed to send due to an unknown error.
	//
	// possible solutions: Try again. If the error persists, open a Direct Support ticket.
	ErrSomethingWentWrong MetaError = "something went wrong"

	// ErrAccessDenied
	//
	// HTTP 403
	//
	// code 131005
	//
	// summary: Permission is either not granted or has been removed.
	//
	// possible solutions: Use the access token debugger to verify that your app has been granted the permissions required by the endpoint.
	ErrAccessDenied MetaError = "access denied"

	// ErrRequiredParameterIsMissing
	//
	//
	// code 131008
	//
	// summary: A required parameter is missing.
	ErrRequiredParameterIsMissing MetaError = "required parameter is missing"

	// ErrParameterValueIsNotValid
	//
	// code 131009
	//
	// summary: The parameter value is not valid.
	ErrParameterValueIsNotValid MetaError = "parameter value is not valid"

	// ErrServiceUnavailable
	//
	// code 131016
	//
	// summary: The service is unavailable.
	ErrServiceUnavailable MetaError = "service unavailable"

	// ErrRecipientCannotBeSender
	//
	// code 131021
	//
	// summary: The recipient cannot be the sender.
	ErrRecipientCannotBeSender MetaError = "recipient cannot be sender"

	// ErrMessageUndeliverable
	//
	// code 131026
	//
	// summary: The message is undeliverable.
	ErrMessageUndeliverable MetaError = "message undeliverable"

	// ErrBusinessEligibilityPaymentIssue
	//
	// code 131042
	//
	// summary: The business is not eligible to send messages.
	ErrBusinessEligibilityPaymentIssue MetaError = "business eligibility payment issue"

	//ErrIncorrectCertificate
	//
	// code 131045
	//
	// summary: The certificate is incorrect.
	ErrIncorrectCertificate MetaError = "incorrect certificate"

	// ErrReEngagementMessage
	//
	// code 131047
	//
	// summary: The message is a re-engagement message.
	ErrReEngagementMessage MetaError = "re-engagement message"

	// ErrUnsupportedMessageType
	//
	//
	// code 131051
	//
	// summary: The message type is not supported.
	ErrUnsupportedMessageType MetaError = "unsupported message type"

	// ErrMediaDownloadError
	//
	//
	// code 131052
	//
	// summary: An error occurred while downloading media.
	ErrMediaDownloadError MetaError = "media download error"

	// ErrMediaUploadError
	//
	// code 131053
	//
	// summary: An error occurred while uploading media.
	ErrMediaUploadError MetaError = "media upload error"

	// ErrTemplateParamCountMismatch
	//
	// code 132000
	//
	// summary: The number of parameters provided for the template does not match the expected number.
	ErrTemplateParamCountMismatch MetaError = "template parameter count mismatch"

	// ErrTemplateDoesNotExist
	//
	// code 132001
	//
	// summary: The specified template does not exist.
	ErrTemplateDoesNotExist MetaError = "template does not exist"

	// ErrTemplateHydratedTextTooLong
	//
	// code 132005
	//
	// summary: The hydrated text of the template exceeds the maximum allowed length.
	ErrTemplateHydratedTextTooLong MetaError = "template hydrated text too long"

	// ErrTemplateFormatCharacterPolicyViolated
	//
	// code 132007
	//
	// summary: The template contains formatting characters that violate policy.
	ErrTemplateFormatCharacterPolicyViolated MetaError = "template format character policy violated"

	// ErrTemplateParameterFormatMismatch
	//
	// code 132012
	//
	// summary: The format of one or more template parameters is incorrect.
	ErrTemplateParameterFormatMismatch MetaError = "template parameter format mismatch"

	// ErrTemplateParameterInvalid
	//
	// code 132015
	//
	// summary: One or more of the template parameters provided are invalid.
	ErrTemplateParameterInvalid MetaError = "template parameter invalid"

	// ErrIncompleteDeregistration
	//
	// code 133000
	//
	// summary: Deregistration incomplete.
	ErrIncompleteDeregistration MetaError = "incomplete deregistration"

	// ErrServerTemporarilyUnavailable
	//
	// code 133004
	//
	// summary: The server is temporarily unavailable.
	ErrServerTemporarilyUnavailable MetaError = "server temporarily unavailable"

	// ErrTwoStepVerificationPINMismatch
	//
	// code 133005
	//
	// summary: The two-step verification PIN provided is incorrect.
	ErrTwoStepVerificationPINMismatch MetaError = "two-step verification PIN mismatch"

	// ErrPhoneNumberReVerificationNeeded
	//
	// code 133006
	//
	// summary: The phone number needs to be re-verified.
	ErrPhoneNumberReVerificationNeeded MetaError = "phone number re-verification needed"

	// ErrTooManyTwoStepVerificationPINGuesses
	//
	// code 133008
	//
	// summary: The maximum number of allowed two-step verification PIN guesses has been exceeded.
	ErrTooManyTwoStepVerificationPINGuesses MetaError = "too many two-step verification PIN guesses"

	// ErrTwoStepVerificationPinGuessedTooFast
	//
	// code 133009
	//
	// summary: Two-step verification PIN guessed too quickly.
	ErrTwoStepVerificationPinGuessedTooFast MetaError = "two-step verification PIN guessed too fast"

	// ErrPhoneNumberNotRegistered
	//
	// code 133010
	//
	// summary: The phone number is not registered.
	ErrPhoneNumberNotRegistered MetaError = "phone number not registered"

	// ErrGenericUserError
	//
	// code 135000
	//
	// summary: Generic user error.
	ErrGenericUserError MetaError = "generic user error"
)

var (
	errCodeMap = map[MetaError]int{
		// Authorization Errors [ 5 ]
		ErrAuthException:         0,   // 401
		ErrAPIMethod:             3,   // 403
		ErrAPIPermissionDenied:   10,  // 403
		ErrAccessTokenHasExpired: 190, // 403
		ErrAPIPermission:         200, // status: 200-299 but http code: 403

		// Throttling Errors [ 5 ]
		ErrAPITooManyCalls:  4,      // 429
		ErrRateLimitIssues:  80007,  // 429
		ErrRateLimitHit:     130429, // 429
		ErrSpamRateLimitHit: 131048, // 429
		ErrPairRateLimitHit: 131056, // 429

		// Integrity Errors [ 2 ]
		ErrTemporarilyBlocked: 368,    // 403
		ErrAccountLocked:      131031, // 403

		// Other Errors. [ 20 ]
		ErrAPIUnknown:                            1,
		ErrAPIService:                            2,
		ErrAPIInvalidParameter:                   100,
		ErrSomethingWentWrong:                    131000,
		ErrAccessDenied:                          131005,
		ErrRequiredParameterIsMissing:            131008,
		ErrParameterValueIsNotValid:              131009,
		ErrServiceUnavailable:                    131016,
		ErrRecipientCannotBeSender:               131021,
		ErrMessageUndeliverable:                  131026,
		ErrBusinessEligibilityPaymentIssue:       131042,
		ErrIncorrectCertificate:                  131045,
		ErrReEngagementMessage:                   131047,
		ErrUnsupportedMessageType:                131051,
		ErrMediaDownloadError:                    131052,
		ErrMediaUploadError:                      131053,
		ErrTemplateParamCountMismatch:            132000,
		ErrTemplateDoesNotExist:                  132001,
		ErrTemplateHydratedTextTooLong:           132005,
		ErrTemplateFormatCharacterPolicyViolated: 132007,
		ErrTemplateParameterFormatMismatch:       132012,
		ErrTemplateParameterInvalid:              132015,
		ErrIncompleteDeregistration:              133000,
		ErrServerTemporarilyUnavailable:          133004,
		ErrTwoStepVerificationPINMismatch:        133005,
		ErrPhoneNumberReVerificationNeeded:       133006,
		ErrTooManyTwoStepVerificationPINGuesses:  133008,
		ErrTwoStepVerificationPinGuessedTooFast:  133009,
		ErrPhoneNumberNotRegistered:              133010,
		ErrGenericUserError:                      135000,
	}

	codeErrorMap = func() map[int]MetaError {
		codeErrorMap := make(map[int]MetaError)
		for err, code := range errCodeMap {
			codeErrorMap[code] = err
		}

		// add ErrAPIPermission for 200-299
		for i := 200; i < 300; i++ {
			codeErrorMap[i] = ErrAPIPermission
		}

		return codeErrorMap
	}()

	NoError = MetaError("")
)

// IsValid returns true if the error is a valid meta error.
func (e MetaError) IsValid() bool {
	_, ok := errCodeMap[e]
	return ok
}

// Error returns the error string.
func Error(code int) MetaError {
	return codeErrorMap[code]
}

// IsError returns true if the error is a valid meta error.
func IsError(err error) bool {
	_, ok := err.(MetaError)
	return ok
}

func IsErrorCode(code int, statusCode int) bool {

	if code == 0 && statusCode == 403 {
		return true
	}

	_, ok := codeErrorMap[code]
	if ok {
		return true
	}

	return false
}

// Code returns the error code.
func (e MetaError) Code() int {
	return errCodeMap[e]
}
