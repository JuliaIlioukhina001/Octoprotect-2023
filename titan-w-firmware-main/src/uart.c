#include "uart.h"

void uart_init(){
    const struct device *const dev = DEVICE_DT_GET(DT_CHOSEN(zephyr_console));
	uint32_t dtr = 0;
	int maxtries = 5;
	if (usb_enable(NULL)) {
		return;
	}
	while (!dtr && maxtries) {
		uart_line_ctrl_get(dev, UART_LINE_CTRL_DTR, &dtr);
		k_sleep(K_MSEC(100));
		// Wait host at most 500ms
		maxtries--;
	}
}