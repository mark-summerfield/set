// Copyright © 2024 Mark Summerfield. All rights reserved.
// License: GPL-3

package set

import (
    "fmt"
    _ "embed"
    )

//go:embed Version.dat
var Version string

func Hello() string {
    return fmt.Sprintf("Hello set v%s", Version)
}
