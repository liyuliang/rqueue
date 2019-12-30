package system

import "github.com/liyuliang/utils/format"

var _config format.MapData

func init() {
	_config = format.Map()
}

func Config() format.MapData {
	return _config
}


