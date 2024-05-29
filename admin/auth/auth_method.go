package auth

type Authorizer interface {
	Obtaining(interface{}) (string, error)
	Authenticating(string) (interface{}, int, error)
}

type authorizerImpl struct{}

func (auth *authorizerImpl) Obtain(a Authorizer, u interface{}) (string, error) {
	return a.Obtaining(u)
}

func (auth *authorizerImpl) Authenticate(a Authorizer, j string) (interface{}, int, error) {
	return a.Authenticating(j)
}

var AuthorizerImpl = new(authorizerImpl)
