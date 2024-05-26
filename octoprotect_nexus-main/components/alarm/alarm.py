import asyncio

import pigpio
import time

# Prevent multiple threads from playing at the same time.
_semaphore = asyncio.Semaphore(1)
pi = pigpio.pi()


async def play_alarm(high):
    await _semaphore.acquire()

    pi.hardware_PWM(18, 5000 if high else 4000, 900000)
    _semaphore.release()


async def stop_alarm():
    await _semaphore.acquire()

    pi.hardware_PWM(18, 0, 0)
    _semaphore.release()
