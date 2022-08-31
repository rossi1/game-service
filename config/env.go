package config

type environment string

const (
	staging    environment = "staging"
	local      environment = "local"
	production environment = "production"
)

func (e environment) isvalid() bool {
	switch e {
	case staging, local, production:
		return true
	default:
		return false
	}
}

func (e environment) String() string {
	return string(e)
}
