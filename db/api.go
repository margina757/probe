package db

import (
	"probe/internal/types"
)

func InsertProbeResult(result *types.ProbeResult) {
	chanProbeResult <- result
}

func QueryProbeResultsByDest(dests []*string) {

}
