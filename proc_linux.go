//go:build linux

package proc

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func queryInternal(pid int) (ProcState, error) {
	status, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", pid))
	if err != nil {
		return 0, err
	}

	for _, line := range strings.Split((string)(status), "\n") {
		kv := strings.Fields(line)
		if len(kv) < 2 {
			continue
		}

		if kv[0] == "State:" {
			switch kv[1] {
			case "D":
				return PROCSTATE_IDLE, nil
			case "I":
				return PROCSTATE_IDLE, nil
			case "R":
				return PROCSTATE_RUN, nil
			case "S":
				return PROCSTATE_IDLE, nil
			case "T":
				return PROCSTATE_STOP, nil
			case "t":
				return PROCSTATE_STOP, nil
			case "X":
				return PROCSTATE_DEAD, nil
			case "Z":
				return PROCSTATE_LOST, nil
			default:
				return 0, errors.New("unknown linux state: " + kv[1])
			}
		}
	}

	return 0, errors.New("no state field in linux proc")
}
