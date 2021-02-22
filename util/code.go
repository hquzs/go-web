package util

const (
	// SUCCESS 交易处理成功
	SUCCESS = "SUCCESS"
	// TimeoutError 交易处理超时
	TimeoutError = "TIMEOUT"
	// InvalidRequest InvalidRequest, 交易格式错误
	InvalidRequest = "INVALID_REQUEST"
	// FormatError 交易，请求格式错误
	FormatError = "FORMAT_ERROR"
	// SystemError 系统错误, 未知错误，DB错误，BlockChain错误
	SystemError = "SYSTEM_ERROR"
	// StructError 结构体序列化、反序列化错误
	StructError = "STRUCT_ERROR"
	// UnknownError 其他错误
	UnknownError = "UNKNOWN_ERROR"
	// UnknownLeader leader信息不在本节点内部
	UnknownLeader = "UNKNOWN_LEADER"
)
