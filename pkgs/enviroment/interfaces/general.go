package interfaces

//ConfigEnviroment ...
type ConfigEnviroment interface {
	GetTag(tag string) (string, error)
	SetFileConfig(file string)
}
