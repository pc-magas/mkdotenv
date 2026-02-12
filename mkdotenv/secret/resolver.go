package secret

type Resolver interface {
    Resolve(secretVal string) (string, error)
	ResolveWithParam(secretVal string,field string) (string, error)
}