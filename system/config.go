package system

type appConfig map[string]string

var _config appConfig

func init() {
	_config = make(map[string]string)
}
func Config() appConfig {
	return _config
}
