import asyncio
from typing import Optional

from databases import Database
from pydantic import BaseModel
from fastapi import FastAPI, HTTPException

app = FastAPI()


class Item(BaseModel):
    id: int
    name: str


class ItemRepository:
    def __init__(self, app: FastAPI) -> None:
        self._app = app

        self._queue = asyncio.Queue()
        self._queued_futures = {}

    async def connect_db(self) -> None:
        self._db = Database("sqlite://./test.db")
        await self._db.connect()

    async def disconnect_db(self) -> None:
        await self._db.disconnect()

    async def consume_queue(self) -> None:
        while True:
            id = await self._queue.get()
            item = await self._get_item_by_id(id)
            self._queued_futures[id].set_result(item)
            del self._queued_futures[id]
            self._queue.task_done()

    async def _get_item_by_id(self, id: int) -> Optional[Item]:
        self._app.state.metric_dbcalls += 1

        await asyncio.sleep(.05) # simulate expensive read

        query = "SELECT * FROM 'items' WHERE id = :id"
        row = await self._db.fetch_one(query=query, values={"id": id})
        if row is None:
            return row
        else:
            return Item(**row)

    async def get_item_by_id(self, id: int) -> asyncio.Future[Optional[Item]]:
        fut = self._queued_futures.get(id)
        if fut:
            return fut

        fut = asyncio.get_event_loop().create_future()
        self._queued_futures[id] = fut
        await self._queue.put(id)
        return fut


@app.on_event("startup")
async def startup_event():
    app.state.metric_requests = 0
    app.state.metric_dbcalls = 0

    app.state.item_repository = ItemRepository(app=app)
    await app.state.item_repository.connect_db()
    asyncio.create_task(app.state.item_repository.consume_queue())


@app.on_event("shutdown")
async def shutdown_event():
    await app.state.item_repository.disconnect_db()


@app.get("/metrics")
def view_metrics() -> dict:
    return {
        "requests": app.state.metric_requests,
        "db_calls": app.state.metric_dbcalls
    }


@app.get("/{item_id}")
async def get_item(item_id: int) -> Item:
    app.state.metric_requests += 1

    print(app.state.item_repository._queued_futures)
    pending_item = await app.state.item_repository.get_item_by_id(item_id)
    item = await pending_item
    if item is None:
        raise HTTPException(status_code=404, detail="No item found")
    
    return item
