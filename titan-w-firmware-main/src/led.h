#ifndef LED_H
#define LED_H

#include <zephyr/kernel.h>
#include <zephyr/drivers/gpio.h>

static const struct gpio_dt_spec led = GPIO_DT_SPEC_GET(DT_ALIAS(led1_green), gpios);

typedef enum LEDState {
    On,
    Blink,
    Off
} LEDState;

// Initialize LED GPIOs
void led_init();

// Update LED State
void led_update(LEDState new);

#endif