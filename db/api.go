package db

import (
	"probe/internal/types"
)

func InsertProbeResult(result *types.ProbeResult) {
	chanProbeResult <- result
}
