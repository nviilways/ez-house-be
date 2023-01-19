package error

var ErrorCode = map[int]string{
	400: "BAD_REQUEST",
	401: "STATUS_UNAUTHORIZED",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	409: "CONFLICT",
	500: "INTERNAL_SERVER_ERROR",
}
