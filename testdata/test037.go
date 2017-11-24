package testdata

// following imports were changed by Max Chechel (github.com/hexdigest)
// due to modifications in internal/render/templates/function.tmpl
import "github.com/hexdigest/gotests"

type someIndirectImportedStruct gotests.Options

func (smtg *someIndirectImportedStruct) Foo037() {}
