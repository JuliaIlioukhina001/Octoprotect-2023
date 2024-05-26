#include <zephyr/bluetooth/gatt.h>
#include <zephyr/settings/settings.h>

#include "uart.h"
#include "led.h"
#include "accelerometer.h"

// Bluetooth Service UUID
#define TITAN_W_SERVICE_VAL \
	BT_UUID_128_ENCODE(0xE0E62F49, 0x9E34, 0x462C, 0x8240, 0x720C01FC5374)

// Acceleration Characteristic UUID
#define TITAN_W_ACCELEROMETER_CHAR_VAL \
	BT_UUID_128_ENCODE(0xDC77CEBA, 0xDEE1, 0x4467, 0x9814, 0xE2EA1DA1098A)

static struct bt_uuid_128 titan_uuid = BT_UUID_INIT_128(TITAN_W_SERVICE_VAL);

static struct bt_uuid_128 acc_char_uuid = BT_UUID_INIT_128(TITAN_W_ACCELEROMETER_CHAR_VAL);

// Acceleration value buffer
static uint8_t acc_value[6] = {0,0,0,0,0,0};

static ssize_t read_acc(struct bt_conn *conn, const struct bt_gatt_attr *attr,
			void *buf, uint16_t len, uint16_t offset)
{
	const char *value = attr->user_data;
	return bt_gatt_attr_read(conn, attr, buf, len, offset, value, 6);
}

static ssize_t deny_write(struct bt_conn *conn, const struct bt_gatt_attr *attr,
			    const void *buf, uint16_t len, uint16_t offset,
			    uint8_t flags)
{
	return 0;
}

BT_GATT_SERVICE_DEFINE(titan_svc,
	BT_GATT_PRIMARY_SERVICE(&titan_uuid),
	BT_GATT_CHARACTERISTIC(&acc_char_uuid.uuid,
			       BT_GATT_CHRC_READ |
			       BT_GATT_CHRC_NOTIFY,
			       BT_GATT_PERM_READ,
			       read_acc, deny_write, acc_value),
	BT_GATT_CCC(NULL,
		    BT_GATT_PERM_READ | BT_GATT_PERM_WRITE_ENCRYPT)
);

static const struct bt_data ad[] = {
	BT_DATA_BYTES(BT_DATA_FLAGS, (BT_LE_AD_GENERAL | BT_LE_AD_NO_BREDR)),
	BT_DATA_BYTES(BT_DATA_UUID128_ALL, TITAN_W_SERVICE_VAL),
};

static void connected(struct bt_conn *conn, uint8_t err)
{
	char addr[BT_ADDR_LE_STR_LEN];
	bt_addr_le_to_str(bt_conn_get_dst(conn), addr, sizeof(addr));
	if (err) {
		printk("(MAC: %s) Connection failed (err 0x%02x)\n", addr, err);
	} else {
		printk("(MAC: %s) Connected\n", addr);
	}
}

static void disconnected(struct bt_conn *conn, uint8_t reason)
{
	char addr[BT_ADDR_LE_STR_LEN];
	bt_addr_le_to_str(bt_conn_get_dst(conn), addr, sizeof(addr));
	printk("Disconnected (MAC: %s,reason 0x%02x)\n", addr, reason);
}

BT_CONN_CB_DEFINE(conn_callbacks) = {
	.connected = connected,
	.disconnected = disconnected,
};

static void bt_ready(void)
{
	int err;
	printk("Bluetooth initialized\n");
	if (IS_ENABLED(CONFIG_SETTINGS)) {
		settings_load();
	}
	err = bt_le_adv_start(BT_LE_ADV_CONN_NAME, ad, ARRAY_SIZE(ad), NULL, 0);
	if (err) {
		printk("Advertising failed to start (err %d)\n", err);
		return;
	}
	printk("Advertising successfully started\n");
}

int main(void)
{
	uart_init();
	accelerometer_init();
	led_init();
	led_update(Blink);
	int err;

	err = bt_enable(NULL);
	if (err) {
		printk("Bluetooth init failed (err %d)\n", err);
		return 0;
	}
	bt_set_name("NVDIIA Titan W GPU over BLE");
	bt_unpair(BT_ID_DEFAULT, BT_ADDR_LE_ANY);
	bt_ready();

	while (1) {
		k_sleep(K_MSEC(100));
		accelerometer_read(acc_value);
		int ret = bt_gatt_notify(NULL, &titan_svc.attrs[1], &acc_value, sizeof(acc_value));
		if(ret == -ENOTCONN)
			led_update(Off);
		else
			led_update(Blink);
	}
	return 0;
}
