/**
 * @command keepassx
 * @short-description Resolve secrets from kepassx database file and place them upon .env
 * 
 * @description
 *
 * Keepassx resolver is used to read a keepassx datapase provided from @param{name} locked with @param{password}.
 * Then it is traversed upon the entry provided from resolve(), afterwards the appropriate field provided from @field{USERNAME}
 * 
 * @resolves parent/child/entry 
 * 
 * KeePassX databases are structured as a tree:
 *
 * - parent
 *   - child
 *     - entry
 *
 * Each node in the hierarchy is represented as a string segment in the path.
 *
 * @param name REQUIRED <mydb.kpbx> Keepassx database file name
 * @param password REQUIRED <***>  keepassx password
 *
 * @field USERNAME Fetch the username of an entry
 * @field PASSWORD Fetch the password of the entry
 * @field URL Fetch the url  of the entry
 * @field NOTES Fetch the notes of the entry
 */

package secret

import (
	"os"
	"fmt"
	"strings"
	"github.com/tobischo/gokeepasslib/v3"
	"github.com/pc-magas/mkdotenv/core/context/types"
	"github.com/pc-magas/mkdotenv/msg"
)

type KepassXResolver struct {
	File string
	Password string
}


//@resolver{name} keepassx
//@resolver{description} Keepassx secret resolver
func NewKeepassXResolver(
		file types.ContextPath, //@resolver{arg} 
		password string //@resolver{arg}
	) (*KepassXResolver, error) {

	dbfile := file.Value()

	_,err := os.Stat(dbfile)
	msg.HandleFileError(err,dbfile)

	return &KepassXResolver{
		File:     dbfile,
		Password: password,
	}, nil
}

func (r *KepassXResolver) Resolve(secret_val string) (string, error) {
	return r.ResolveWithParam(secret_val,"PASSWORD");
}

func findEntry(groups []gokeepasslib.Group, pathParts []string) *gokeepasslib.Entry {
	if len(pathParts) == 0 {
		return nil
	}

	groupName := pathParts[0]

	for _, g := range groups {
		if g.Name == groupName {
			if len(pathParts) == 2 { // last part is entry
				entryName := pathParts[1]
				for _, e := range g.Entries {
					if e.GetTitle() == entryName {
						return &e
					}
				}
			} else if len(pathParts) > 2 { // deeper subgroup
				return findEntry(g.Groups, pathParts[1:])
			}
		}
	}
	return nil
}

func (r *KepassXResolver) ResolveWithParam(secretVal string,field string) (string, error) {

	file,_:= os.Open(r.File)
	db := gokeepasslib.NewDatabase()
    db.Credentials = gokeepasslib.NewPasswordCredentials(r.Password)
	_ = gokeepasslib.NewDecoder(file).Decode(db)
	db.UnlockProtectedEntries()

	pathParts := strings.Split(secretVal, "/")
	entry := findEntry(db.Content.Root.Groups, pathParts)
	
	if entry == nil {
		return "", fmt.Errorf("entry not found: %s", secretVal)
	}
	
	switch strings.ToUpper(field) {
		case "PASSWORD":
			return entry.GetPassword(), nil
		case "USERNAME":
			return entry.GetContent("UserName"), nil
		case "URL":
			return entry.GetContent("URL"), nil
		case "NOTES":
			return entry.GetContent("Notes"), nil
		default:
			return "", fmt.Errorf("unsupported field: %s", field)
	}
}
