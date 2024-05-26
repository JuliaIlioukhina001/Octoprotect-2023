import asyncio

from components.alarm.alarm import play_alarm, stop_alarm


def trigger_alarm():
    """
    Starts playing the alarm
    """
    from components.services.service_registry import State
    if not State.triggered:
        State.triggered = True
        asyncio.ensure_future(_alarm_handler())


def clear_alarm():
    """
    Stops the alarm if it was already running
    """
    from components.services.service_registry import State
    State.triggered = False


async def _alarm_handler():
    from components.services.service_registry import State
    high = True
    while not State.completion.done() and State.triggered:
        await play_alarm(high)
        high = not high
        await asyncio.sleep(0.5)
        await stop_alarm()
        await asyncio.sleep(0.2)
