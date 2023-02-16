package utils

var (
	IdVerify            = Rules{"ID": {NotEmpty()}}
	PageInfoVerify      = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty(), Le("30")}}
	PageSizeLimitVerify = Rules{"PageSize": {Le("30")}}
)
