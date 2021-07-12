package goebpf

import (
	"context"
	"fmt"
	"unsafe"
)

/*
#cgo CFLAGS: -I./ebpf
#cgo LDFLAGS: -L${SRCDIR}/ebpf -static -lebpf -lelf -lz
#include <stdint.h>
#include "ebpf.h"
#include "event.h"

void go_handleEvent(void *t_bpfctx, unsigned short evt_type, void *t_data, int t_datalen);

static int wrapper_ebpf_init(int flags) {
  return ebpf_init(go_handleEvent);
}
*/
import "C"

//export go_handleEvent
func go_handleEvent(t_bpfctx unsafe.Pointer, evt_type C.ushort, evt_data unsafe.Pointer, evt_datalen C.int) {

	var event IEvent
	if evt_type == EVENTTYPE_PROCESS_EXIT {
		evt := (*C.struct_Event_ProcessExit)(evt_data)
		pevt := &EventProcessExit{}
		pevt.Pid = (uint32)(evt.pid)
		pevt.Comm = C.GoString(&evt.comm[0])
		event = pevt
	}

	if gEventHandler != nil {
		gEventHandler(event)
	}
}

var gEventHandler EventHandler

func InitEbpf(handler EventHandler) error {
	var ret C.int
	ret = C.wrapper_ebpf_init(0)
	var err error
	if ret != 0 {
		err = fmt.Errorf("fail to init ebpf, %d \n", ret)
		return err
	}

	gEventHandler = handler
	return nil
}

func FreeEbpf() {
	C.ebpf_free()
}

func PollEvents(ctx context.Context, timeout int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		//time.Sleep(time.Millisecond * time.Duration(timeout))
		C.ebpf_poll_event((C.int)(timeout))
		fmt.Println("poll")
	}
}
