package proc

type ProcState int

const (
	PROCSTATE_IDLE   = 1
	PROCSTATE_STOP   = 2
	PROCSTATE_RUN    = 3
	PROCSTATE_UNDEAD = 4
	PROCSTATE_DEAD   = 5
	PROCSTATE_LOST   = 6
)

func Query(pid int) (ProcState, error) {
	return queryInternal(pid)
}
