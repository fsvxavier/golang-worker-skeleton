package interfaces

//GenericLogger is an interface of genericLogger
type GenericLogger interface {
	//GetLogger()
	LogIt(severity, message string, fields map[string]interface{})
	SetModule(name string)
	SetOperation(name string)
	GetHostname() (string, error)
	SetHostname(func() (string, error))
}
