//go:build freebsd

package proc

/*
#cgo LDFLAGS: -lproc
#include <libproc.h>

static int libproc_state(int pid) {
	int ret;
	struct proc_handle* proc;

	ret = proc_attach((pid_t)pid, PATTACH_RDONLY | PATTACH_NOSTOP, &proc);
	if (ret) return -ret;

	ret = proc_state(proc);
	proc_detach(proc, 0);
	proc_free(proc);

	return ret;
}
*/
import "C"
import (
	"errors"
	"fmt"
)

func queryInternal(pid int) (ProcState, error) {
	var ret C.int
	ret = C.libproc_state((C.int)(pid))
	if ret < 0 {
		return 0, errors.New(fmt.Sprintf("libproc: %d", (int)(-ret)))
	}

	return (ProcState)(ret), nil
}
