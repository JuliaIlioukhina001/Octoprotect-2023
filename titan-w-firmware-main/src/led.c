#include "led.h"

bool led_initialized = false;
LEDState led_current_state = Off;

void led_init(){
	if (!gpio_is_ready_dt(&led)) {
        printk("[ERROR] LED is not ready!\n");
		return;
	}
	if (gpio_pin_configure_dt(&led, GPIO_OUTPUT_ACTIVE) < 0) {
        printk("[ERROR] Cannot configure LED!\n");
		return;
	}
    led_initialized = true;
    printk("[INFO] LED Initialized.\n");
}

void led_work_loop(){
    while(1){
        if(!led_initialized){
            k_sleep(K_MSEC(500));
            continue;
        }
        if(led_current_state == On){
            gpio_pin_set_dt(&led, 1);
        }else if(led_current_state == Off){
            gpio_pin_set_dt(&led, 0);
        }else{
            gpio_pin_toggle_dt(&led);
        }
        k_sleep(K_MSEC(500));
    }
}

void led_update(LEDState new){
    led_current_state = new;
}

// Start LED work loop thread
K_THREAD_DEFINE(led_work, 1024, led_work_loop, NULL, NULL, NULL, 2, 0, 0);