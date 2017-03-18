package utils

/**
 * response返回json对象
 */
type ResponseData struct {
	ErrorCode int 		`json:"errorCode"`	// 必需 错误码。正常返回0 异常返回560 错误提示561对应errorInfo
	ErrorInfo interface{} 	`json:"errorInfo"`	// 必需 错误信息。正常返回空”” 异常返回错误信息文本
	Data interface{} 	`json:"data"`		// 可选 返回数据内容。 如果有返回数据，可以是字符串或者数组JSON等等
	Total int 		`json:"total"`		// 可选 分页字段：总条数
	PageNum string 		`json:"pageNum"`	// 可选 分页字段：当前页数
	PageSize string 	`json:"pageSize"`	// 可选 分页字段：当前每页多少条数
}