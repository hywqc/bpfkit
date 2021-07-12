#ifndef __FLOW_H
#define __FLOW_H

#define TASK_COMM_LEN 16

enum {
    EVENTTYPE_TEST = 0,
    EVENTTYPE_PROCESS_EXIT,
    EVENTTYPE_TCP_TX,
    EVENTTYPE_TCP_RX,
    EVENTTYPE_TCP_RETRANS,
};

struct Event_Test
{
    uint16_t event_type;
    char msg[100];
};

struct Event_ProcessExit {
    uint16_t event_type;
    uint32_t pid;
    uint32_t uid;
	char comm[TASK_COMM_LEN];
};

struct Event_TCP {
    uint16_t event_type;
    uint32_t pid;
    uint32_t uid;
	char comm[TASK_COMM_LEN];
    uint32_t af;
    uint8_t status;

    union {
		uint32_t saddr_v4;
		uint8_t saddr_v6[16];
	};
	union {
		uint32_t daddr_v4;
		uint8_t daddr_v6[16];
	};

    uint16_t sport;
    uint16_t dport;
    uint32_t data_size;
};







#endif