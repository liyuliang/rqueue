package system

var (
	RedisQueueTotalPrefix   = "total:"
	RedisQueuePrefix        = "queue:"
	RedisTokenExpiredPrefix = "api:token:"

	RedisQueueStorage    = "storage"
	RedisSpiderClientSet = "spiders:client"

	TokenInAddApi    = "token"
	UrlsInAddApi     = "[]url"
	CategoryInAddApi = "type"
	QueueInGetApi    = "queue"

	Method404Code = "NOT_FOUND"
	Method404Msg  = "Not found"

	SystemRedisUri = "redisUri"
	SystemPopNum   = "popNum"
	SystemTplDir   = "tplDir"
	SystemIsDebug  = "debug"
)
