package constant

import (
	"errors"
	"net/http"
)

const (
	LLvlAccess             = "ACCESS_LOG"
	DateFormatReversed     = "02-01-2006"
	DateFormatStandard     = "2006-01-02"
	DateFormatWithTime     = "2006-01-02 15:04"
	DateFormatWithTime2    = "2006-01-02 15:04:05"
	DateFormatWithTimeZone = "2006-01-02T15:04:05+07:00"
)

var (
	ErrInternalServerError = errors.New("Terjadi error di server. Mohon hubungi administrator.")
	ErrNotFound            = errors.New("your requested Item is not found")
	ErrExpired             = errors.New("expired")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrConflict            = errors.New("item already exist")
	ErrBadParamInput       = errors.New("given Param is not valid")
	ErrDataNotFound        = errors.New("Data not found")
	ErrSubmissionNotFound  = errors.New("Data Submission not found")
	ErrPositionNotFound    = errors.New("Data Position not found")
	ErrCouncilNotFound     = errors.New("Data Council not found")

	// auth
	ErrInvalidLoginCredential    = errors.New("otp tidak valid atau kadaluarsa.")
	ErrInvalidPasswordResetToken = errors.New("token tidak valid atau kadaluarsa.")
	ErrInvalidPassword           = errors.New("kata sandi tidak sesuai")

	// password validation
	ErrPasswordTooShort      = errors.New("Password must be at least 8 characters long")
	ErrPasswordNoDigit       = errors.New("Password must contain at least one digit (0-9)")
	ErrPasswordNoUppercase   = errors.New("Password must contain at least one uppercase letter")
	ErrPasswordNoLowercase   = errors.New("Password must contain at least one lowercase letter")
	ErrPasswordNoSpecialChar = errors.New("Password must contain at least one special character (~, !, @, #, $, %, ^, &, *)")

	ErrOldPasswordDoesNotMatch = errors.New("kata sandi lama tidak sesuai")
	ErrNewPasswordDoesNotMatch = errors.New("kata sandi baru tidak sesuai")
	ErrNewPasswordSameAsOld    = errors.New("kata sandi baru sama dengan kata sandi lama.")

	ErrEmailNotFound     = errors.New("Email tidak ditemukan.")
	ErrEmailAlreadyExist = errors.New("Alamat email sudah terdaftar.")

	ErrUsernameNotFound      = errors.New("Username tidak ditemukan.")
	ErrUsernameAlreadyExists = errors.New("Username sudah terdaftar.")

	ErrRoleNotFound                         = errors.New("Role tidak ditemukan.")
	ErrCoordinatorNotFound                  = errors.New("Koordinator tidak ditemukan.")
	ErrConfirmAccountTokenNotFoundOrExpired = errors.New("token tidak valid.")

	ErrRoleNameAlreadyExist = errors.New("Role sudah ada")

	ErrInvalidToken       = errors.New("Invalid Token")
	ErrUserNotRegistered  = errors.New("Pengurus belum terdaftar")
	ErrFileTypeNotAllowed = errors.New("File type not allowed")
	ErrFileSizeExceeded   = errors.New("File size exceeded")

	// Email
	ErrFailedSendEmail = errors.New("terjadi kesalahan pada sistem pengiriman email")
)

var errorStatusCode = map[error]int{
	ErrInvalidLoginCredential:               http.StatusUnauthorized,
	ErrInvalidPasswordResetToken:            http.StatusUnauthorized,
	ErrInvalidPassword:                      http.StatusUnauthorized,
	ErrPasswordTooShort:                     http.StatusBadRequest,
	ErrPasswordNoDigit:                      http.StatusBadRequest,
	ErrPasswordNoUppercase:                  http.StatusBadRequest,
	ErrPasswordNoLowercase:                  http.StatusBadRequest,
	ErrPasswordNoSpecialChar:                http.StatusBadRequest,
	ErrEmailNotFound:                        http.StatusNotFound,
	ErrOldPasswordDoesNotMatch:              http.StatusUnauthorized,
	ErrNewPasswordDoesNotMatch:              http.StatusUnauthorized,
	ErrNewPasswordSameAsOld:                 http.StatusUnauthorized,
	ErrEmailAlreadyExist:                    http.StatusBadRequest,
	ErrRoleNotFound:                         http.StatusBadRequest,
	ErrConfirmAccountTokenNotFoundOrExpired: http.StatusNotFound,
	ErrRoleNameAlreadyExist:                 http.StatusConflict,
	ErrInvalidToken:                         http.StatusUnauthorized,
	ErrUserNotRegistered:                    http.StatusNotFound,
	ErrDataNotFound:                         http.StatusOK,
	ErrSubmissionNotFound:                   http.StatusInternalServerError,
	ErrFailedSendEmail:                      http.StatusBadRequest,
}

// getStatusCode maps known errors to HTTP status codes.
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if code, exists := errorStatusCode[err]; exists {
		return code
	}
	return http.StatusInternalServerError
}
