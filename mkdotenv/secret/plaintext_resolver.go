package secret

type PlaintextResolver struct {
	File string
	Password string
}

func (r PlaintextResolver) Resolve(secret_val string) (string, error) {
	return r.ResolveWithParam(secret_val,"");
}


func (r PlaintextResolver) ResolveWithParam(secretVal string,field string) (string, error) {
	return secretVal,nil
}