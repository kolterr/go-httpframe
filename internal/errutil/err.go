package errutil

import "errors"

var (
	ErrUnknown                   = errors.New("unknown error")
	ErrServerInternal            = errors.New("server internal error")
	ErrIllegalParameter          = errors.New("illegal parameter")
	ErrWrongType                 = errors.New("wrong type")
	ErrPermissionDenied          = errors.New("permission denied")
	ErrCacheOperation            = errors.New("cache operation error")
	ErrNotFound                  = errors.New("not found")
	ErrParameterMissing          = errors.New("parameter missing")
	ErrTokenExpired              = errors.New("token has expired")
	ErrInvalidToken              = errors.New("invalid token")
	ErrTokenNotFound             = errors.New("token not found")
	ErrInvalidRSAPublicKey       = errors.New("invalid public key")
	ErrInvalidRSAPrivateKey      = errors.New("invalid private key")
	ErrUserNotFound              = errors.New("user not found")
	ErrUserHasRegistered         = errors.New("user name has exist")
	ErrWrongPassword             = errors.New("wrong password")
	ErrPermissionExists          = errors.New("permission exists")
	ErrPermissionNotFound        = errors.New("permission not found")
	ErrPermissionGroupNotFound   = errors.New("permission group not found")
	ErrPermissionGroupNameExists = errors.New("permission group name exists")
	ErrNameDuplication           = errors.New("name duplication")
	ErrPermissionGroupIDEmpty    = errors.New("permission group id can not be empty")
	ErrRoleNotFound              = errors.New("role not found")
	ErrWrongExpiredTime          = errors.New("wrong expire time")
	ErrVerifyFailed              = errors.New("verify failed")
	ErrWrongPhoneNumber          = errors.New("wrong phone number")
	ErrDBOperation               = errors.New("db operation error")
	ErrSignFailed                = errors.New("sign failed")
)

const (
	unknown = 5000 + iota
	serverInternal
	illegalParameter
	wrongType
	permissionDenied
	cacheOperation
	notFound
	parameterMissing
	tokenExpired
	invalidToken
	tokenNotFound
	invalidRSAPublicKey
	invalidRSAPrivateKey
	userNotFound
	userHasRegistered
	wrongPassword
	permissionExists
	permissionNotFound
	permissionGroupNotFound
	permissionGroupNameExists
	nameDuplication
	permissionGroupIdEmpty
	roleNotFound
	wrongExpiredTime
	verifyFailed
	wrongPhoneNumber
	dbOperation
	signFailed
)

var code = map[error]int{
	ErrUnknown:                   unknown,
	ErrServerInternal:            serverInternal,
	ErrIllegalParameter:          illegalParameter,
	ErrWrongType:                 wrongType,
	ErrPermissionDenied:          permissionDenied,
	ErrCacheOperation:            cacheOperation,
	ErrNotFound:                  notFound,
	ErrParameterMissing:          parameterMissing,
	ErrTokenExpired:              tokenExpired,
	ErrInvalidToken:              invalidToken,
	ErrTokenNotFound:             tokenNotFound,
	ErrInvalidRSAPublicKey:       invalidRSAPublicKey,
	ErrInvalidRSAPrivateKey:      invalidRSAPrivateKey,
	ErrUserNotFound:              userNotFound,
	ErrUserHasRegistered:         userHasRegistered,
	ErrWrongPassword:             wrongPassword,
	ErrPermissionExists:          permissionExists,
	ErrPermissionNotFound:        permissionNotFound,
	ErrPermissionGroupNotFound:   permissionGroupNotFound,
	ErrPermissionGroupNameExists: permissionGroupNameExists,
	ErrNameDuplication:           nameDuplication,
	ErrPermissionGroupIDEmpty:    permissionGroupIdEmpty,
	ErrRoleNotFound:              roleNotFound,
	ErrWrongExpiredTime:          wrongExpiredTime,
	ErrVerifyFailed:              verifyFailed,
	ErrWrongPhoneNumber:          wrongPhoneNumber,
	ErrDBOperation:               dbOperation,
	ErrSignFailed:                signFailed,
}

func Code(err error) int {
	c, ok := code[err]
	if !ok {
		return unknown
	}
	return c
}
