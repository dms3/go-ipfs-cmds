package cli

import (
	"strings"
	"testing"

	"github.com/dms3-fs/go-fs-cmdkit"
	"github.com/dms3-fs/go-fs-cmds"
)

func TestSynopsisGenerator(t *testing.T) {
	command := &cmds.Command{
		Arguments: []cmdkit.Argument{
			cmdkit.StringArg("required", true, false, ""),
			cmdkit.StringArg("variadic", false, true, ""),
		},
		Options: []cmdkit.Option{
			cmdkit.StringOption("opt", "o", "Option"),
		},
		Helptext: cmdkit.HelpText{
			SynopsisOptionsValues: map[string]string{
				"opt": "OPTION",
			},
		},
	}
	syn := generateSynopsis(command, "cmd")
	t.Logf("Synopsis is: %s", syn)
	if !strings.HasPrefix(syn, "cmd ") {
		t.Fatal("Synopsis should start with command name")
	}
	if !strings.Contains(syn, "[--opt=<OPTION> | -o]") {
		t.Fatal("Synopsis should contain option descriptor")
	}
	if !strings.Contains(syn, "<required>") {
		t.Fatal("Synopsis should contain required argument")
	}
	if !strings.Contains(syn, "<variadic>...") {
		t.Fatal("Synopsis should contain variadic argument")
	}
	if !strings.Contains(syn, "[<variadic>...]") {
		t.Fatal("Synopsis should contain optional argument")
	}
	if !strings.Contains(syn, "[--]") {
		t.Fatal("Synopsis should contain options finalizer")
	}
}
