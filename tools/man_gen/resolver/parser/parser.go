package parser

import (
	"strings"
	"io"
	"bufio"
)

type CommandSpec struct {
	Command string
	ShortDescription string
	Description string
	ValueToResolve Resolve
	Params map[string]Param
	Fields map[string]Field
}

type Param struct {
	Name        string
	Description string
	DummyVal string
	Required    bool
}

type Field struct {
	Name        string
	Description string
}

type Resolve struct {
	Value string
	FormatDescription string
}


func parseParam(parts []string) Param {
	p := Param{}

	if len(parts) > 0 {
		p.Name = parts[0]
	}

	if len(parts) > 1 {
		p.Required = parts[1] == "REQUIRED"
	}

	// extract <default>
	for i, v := range parts {
		if strings.HasPrefix(v, "<") && strings.HasSuffix(v, ">") {
			p.DummyVal = strings.Trim(v, "<>")
			p.Description = strings.Join(parts[i+1:], " ")
			break
		}
	}

	return p
}

func parseField(parts []string) Field {
	f := Field{}

	if len(parts) > 0 {
		f.Name = parts[0]
	}
	if len(parts) > 1 {
		f.Description = strings.Join(parts[1:], " ")
	}

	return f
}

func ParseComment(r io.Reader) CommandSpec {

	scanner := bufio.NewScanner(r)
	var spec CommandSpec

	var current string

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "*")
		line = strings.TrimSpace(line)


		if strings.HasPrefix(line, "@") {
			
			if current == "description" {
				spec.Description += line
				continue
			}

			if current == "resolves" {
				spec.ValueToResolve.FormatDescription +=line
				continue
			}

			line = strings.TrimPrefix(line, "@")

			parts := strings.Fields(line)
			if len(parts) == 0 {
				continue
			}

			key := parts[0]
			rest := parts[1:]

			current = key
			switch key {
				case "command":
					if len(rest) > 0 {
						spec.Command = rest[0]
					}
				case "short-description":
					spec.ShortDescription = strings.Join(rest, " ")
				case "description":
					continue			
				case "resolves":
					spec.ValueToResolve.Value = rest[0]
					continue
				case "param":
					// param name REQUIRED <mydb.kpbx> Keepassx database file name
					p := parseParam(rest)
					spec.Params = append(spec.Params, p)

				case "field":
					// field USERNAME Fetch the username...
					f := parseField(rest)
					spec.Fields = append(spec.Fields, f)
			}
		}
	}
	
	return spec
}