import asyncio
from asyncio import Future

from components.services.service_registry import initialize


async def main():
    await initialize(Future())
    from components.services.service_registry import State
    await State.error


asyncio.run(main())
