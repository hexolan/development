import pytest
from httpx import AsyncClient
from asgi_lifespan import LifespanManager

from request_coalescing_py.main import app


@pytest.fixture(scope="session")
def anyio_backend():
    return "asyncio"


@pytest.fixture(scope="function")
async def test_app():
    async with LifespanManager(app) as manager:
        yield manager.app


@pytest.fixture(scope="function")
async def client(test_app):
    async with AsyncClient(app=test_app, base_url="http://test") as client:
        yield client