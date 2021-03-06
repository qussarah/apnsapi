package apnsapi

const (
	ErrBadCollapseId               = "BadCollapseId"
	ErrBadDeviceToken              = "BadDeviceToken"
	ErrBadExpirationDate           = "BadExpirationDate"
	ErrBadMessageId                = "BadMessageId"
	ErrBadPriority                 = "BadPriority"
	ErrBadTopic                    = "BadTopic"
	ErrDeviceTokenNotForTopic      = "DeviceTokenNotForTopic"
	ErrDuplicateHeaders            = "DuplicateHeaders"
	ErrIdleTimeout                 = "IdleTimeout"
	ErrMissingDeviceToken          = "MissingDeviceToken"
	ErrMissingTopic                = "MissingTopic"
	ErrPayloadEmpty                = "PayloadEmpty"
	ErrTopicDisallowed             = "TopicDisallowed"
	ErrBadCertificate              = "BadCertificate"
	ErrBadCertificateEnvironment   = "BadCertificateEnvironment"
	ErrExpiredProviderToken        = "ExpiredProviderToken"
	ErrForbidden                   = "Forbidden"
	ErrInvalidProviderToken        = "InvalidProviderToken"
	ErrMissingProviderToken        = "MissingProviderToken"
	ErrBadPath                     = "BadPath"
	ErrMethodNotAllowed            = "MethodNotAllowed"
	ErrUnregistered                = "Unregistered"
	ErrPayloadTooLarge             = "PayloadTooLarge"
	ErrTooManyProviderTokenUpdates = "TooManyProviderTokenUpdates"
	ErrTooManyRequests             = "TooManyRequests"
	ErrInternalServerError         = "InternalServerError"
	ErrServiceUnavailable          = "ServiceUnavailable"
	ErrShutdown                    = "Shutdown"
)
