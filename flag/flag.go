package astiflag

import (
	"os"
	"strings"
)

// Subcommand retrieves the subcommand from the input Args
func Subcommand() (o string) {
	if len(os.Args) >= 2 && os.Args[1][0] != '-' {
		o = os.Args[1]
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
	}
	return
}

// Strings represents a flag that can be set several times that will output a []string
type Strings []string

// String implements the flag.Value interface
func (f *Strings) String() string {
	return strings.Join(*f, ",")
}

// Set implements the flag.Value interface
func (f *Strings) Set(i string) error {
	*f = append(*f, i)
	return nil
}
