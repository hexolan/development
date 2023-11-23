import asyncio
from datetime import datetime
from typing import Callable

import httpx

API_LINK = "http://localhost:8000"


async def main():
    start_time = datetime.now()
    print("Making 100 x 5 concurrent requests (500 total)...")
    await queue_concurrently(func=make_standard_requests, func_args={"num_requests": 100}, times=5)
    # await queue_concurrently(func=make_coalesced_requests, func_args={"num_requests": 100}, times=5)
    print("Took {delta}ms".format(delta=(datetime.now() - start_time).microseconds / 1000))

    await gather_metrics()


async def queue_concurrently(func: Callable, func_args: dict = {}, times: int = 5) -> None:
    # Queue the concurrent tasks
    tasks = []
    for _ in range(times):
        tasks.append(asyncio.create_task(func(**func_args)))

    # Wait until all tasks have completed
    for task in tasks:
        await task


async def make_standard_requests(num_requests: int, item_id: int = 1) -> None:
    async with httpx.AsyncClient() as client:
        for _ in range(num_requests):
            await client.get(API_LINK + f"/standard/{item_id}")


async def make_coalesced_requests(num_requests: int, item_id: int = 1) -> None:
    async with httpx.AsyncClient() as client:
        for _ in range(num_requests):
            await client.get(API_LINK + f"/coalesced/{item_id}")


async def gather_metrics() -> None:
    async with httpx.AsyncClient() as client:
        response = await client.get(API_LINK + "/metrics")
        print("Metrics:", response.json())


if __name__ == "__main__":
    asyncio.run(main())