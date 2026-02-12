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

func NewKeepassXResolver(file types.ContextPath, password string) (*KepassXResolver, error) {

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
