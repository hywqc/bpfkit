#ifndef __EBPFB_H
#define __EBPFB_H

typedef void (*eBPFHandler)(void* t_bpfctx, unsigned short evt_type, void* t_data, int t_datasize);

enum {
    EOK = 0,
    ERLIMIT_BUMP,
    ESKELETON_OPEN,
    EPROGRAM_LOAD,
    EPROGRAM_ATTACH,
    EPROPBUFFER,
};

int ebpf_init(eBPFHandler handler);
int ebpf_free();

int ebpf_poll_event(int timeout_ms);

#endif