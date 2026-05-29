/**
 * @command plain
 * @short-description Place a raw value upon the variable
 * 
 * @resolves
 * 
 * The resolved value is a raw text placed upon the environmanta variable. For example if you have:
 * ```
 * mkdotenv()::resolve("RAW")::plain
 * MYVAR=
 * ```
 * It would generate:
 * ```
 * MYVAR=RAW
 * ```
 * 
 */

package secret

type PlaintextResolver struct {
}

func NewPlaintextResolver() *PlaintextResolver {
    return &PlaintextResolver{}
}

func (r *PlaintextResolver) Resolve(secret_val string) (string, error) {
	return r.ResolveWithParam(secret_val,"");
}


func (r *PlaintextResolver) ResolveWithParam(secretVal string,field string) (string, error) {
	return secretVal,nil
}