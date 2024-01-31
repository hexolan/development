import time
import asyncio

import pytest


@pytest.mark.anyio
async def test_coalesced_requests(client) -> None:
    print("Making 5x100 concurrent requests (500 total)...")

    async def make_requests():
        for _ in range(100):
            response = await client.get("/coalesced/1")
            assert response.status_code == 200

    # Make the requests concurrently and wait until they complete
    start_time = time.perf_counter_ns()
    tasks = [asyncio.create_task(make_requests()) for _ in range(5)]
    for task in tasks:
        await task
    end_time = time.perf_counter_ns()
    print("Coalesced Requests: Took {delta}ms".format(delta=(end_time - start_time) / 1000000))

    # View the metrics
    metrics = await client.get("/metrics")
    print("Coalesced Metrics:", metrics.json())
