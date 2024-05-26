#ifndef UART_H
#define UART_H

#include <zephyr/kernel.h>
#include <zephyr/usb/usb_device.h>
#include <zephyr/drivers/uart.h>

// Initialize UART for logging
void uart_init();

#endif