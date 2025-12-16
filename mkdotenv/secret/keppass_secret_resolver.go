package secret

import (
	"github.com/tobischo/gokeepasslib/v3"
	"os"
	"fmt"
	"strings"
)

type KepassXResolver struct {
	File string
	Password string
}

func NewKeepassXResolver(file, password string) (*KepassXResolver, error) {
	// Check if the file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, fmt.Errorf("keepass file does not exist: %s", file)
	} else if err != nil {
		// Some other filesystem error (e.g., permission denied)
		return nil, fmt.Errorf("error accessing file %s: %w", file, err)
	}

	return &KepassXResolver{
		File:     file,
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
