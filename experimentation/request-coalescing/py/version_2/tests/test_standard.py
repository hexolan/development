import asyncio

import pytest


@pytest.mark.anyio
async def test_standard_requests(client) -> None:
    print("Making 100x5 concurrent requests (500 total)...")

    async def make_requests():
        for _ in range(100):
            response = await client.get("/standard/1")
            assert response.status_code == 200

    # Make the requests concurrently
    tasks = []
    for _ in range(5):
        tasks.append(asyncio.create_task(make_requests()))

    # Wait until all concurrent tasks have completed
    for task in tasks:
        await task

    # View the metrics
    result = await client.get("/metrics")
    print("Standard Metrics:", result.json())