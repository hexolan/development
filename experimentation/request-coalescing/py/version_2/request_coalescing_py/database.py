import asyncio
from typing import Optional

from fastapi import FastAPI
from databases import Database

from request_coalescing_py.models import Item


class DatabaseRepo:
    def __init__(self, app: FastAPI) -> None:
        self.app = app

    async def start_db(self) -> None:
        self._db = Database("sqlite:///:memory:")
        # self._db = Database("sqlite:///:memory:")
        await self._db.connect()
        
        try:
            await self._db.execute(query="CREATE TABLE IF NOT EXISTS 'items' (id INTEGER PRIMARY KEY, name TEXT)")
            await self._db.execute(query="INSERT INTO 'items' (id, name) VALUES (1, 'test')")
        except Exception:
            pass

    async def stop_db(self) -> None:
        await self._db.disconnect()

    async def get_by_id(self, item_id: int) -> Optional[Item]:
        self.app.state.metrics["db_calls"] += 1

        # Simulate expensive read (50ms/db read)
        # await asyncio.sleep(.05)

        # Simulate expensive read (100ms)
        await asyncio.sleep(.1)

        # todo: comment out once db in place
        """
        return Item(
            id=item_id,
            name=f"Item {item_id}"
        )
        """

        row = await self._db.fetch_one(
            query="SELECT * FROM 'items' WHERE id = :id",
            values={"id": item_id}
        )
        if row:
            return Item(**row)
        
        return None