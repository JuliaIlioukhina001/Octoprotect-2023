import uuid
from asyncio import Future


class NexusState:
    is_armed: bool = False
    sensitivity: float = 0.4
    stream: bool = False
    paired: [uuid] = []
    triggered: bool = False
    completion: Future = Future()
    error: Future
