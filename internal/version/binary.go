package version

import (
	"fmt"
	"runtime"
)

const Binary = "0.3.7-HA.1.1.1+fixretry"

func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, Binary, runtime.Version())
}
