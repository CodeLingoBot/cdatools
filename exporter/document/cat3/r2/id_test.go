package document_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/document/cat3/r2"
	"github.com/projectcypress/cdatools/models"
)

func TestIdPrint(t *testing.T) {
	var tests = []struct {
		n        document.CDAIdentifiers
		expected string
	}{
		{
			document.NewCDAIdentifiers([]models.CDAIdentifier{models.CDAIdentifier{
				Root: "root", Extension: "extension"},
				models.CDAIdentifier{Extension: "extension"}}),
			fmt.Sprintf(cdaIdentifiersCat3TestTemplate1, "root", "extension", "extension"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("CDAIdentifiers.Print(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var cdaIdentifiersCat3TestTemplate1 = `<id root="%s" extension="%s" />
<id  extension="%s" />`
