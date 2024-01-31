import asyncio
from typing import Optional

from request_coalescing_py.models import Item
from request_coalescing_py.database import DatabaseRepo


class CoalescingRepo:
    def __init__(self, repo: DatabaseRepo):
        self._repo = repo

        self._lock = asyncio.Lock()
        self._queue = asyncio.Queue()
        self._queued = {}  # map of item_id: future
        # self._queued_map = {}  # map of item_id: future

    async def get_by_id(self, item_id: int) -> "asyncio.Future[Optional[Item]]":
        """
        async with self._lock:
            # Check if there is an already pending request for that item.
            fut = self._queued_map.get(item_id)
            if fut:
                return fut
        
            # There is not an existing pending request.
            fut = asyncio.get_event_loop().create_future()
            self._queued_map[item_id] = fut
            return fut
        """
        
        """
        async with self._lock:
            # Check if there is an already pending request for that item.
            fut = self._queued.get(item_id)
            if fut:
                return fut
            
            # There is not a pending request.
            fut = asyncio.get_event_loop().create_future()
            self._queued[item_id] = fut
            await self._queue.put(item_id)
            return fut
        """

        # Check if there is an already pending request for that item.
        fut = self._queued.get(item_id)
        if fut:
            return fut
        
        # There is not a pending request.
        fut = asyncio.get_event_loop().create_future()
        self._queued[item_id] = fut
        await self._queue.put(item_id)
        return fut

    async def process_queue(self) -> None:
        """
        while True:
            await asyncio.sleep(.001)
            async with self._lock:
                for item_id, future in self._queued_map.items():
                    item = await self._repo.get_by_id(item_id)
                    future.set_result(item)

                self._queued_map = {}
        """

        while True:
            item_id = await self._queue.get()
            async with self._lock:
                item = await self._repo.get_by_id(item_id)
                self._queued[item_id].set_result(item)
                del self._queued[item_id]
                self._queue.task_done()