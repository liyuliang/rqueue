package system

var (
	RedisQueuePrefix      = "queue_"
	RedisQueueTotalPrefix = "total_"
	RedisQueueStorage     = "storage"
	RedisUUidKey          = "add_api_token"

	SpiderClientSet  = "client"
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
