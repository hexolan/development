import asyncio

import httpx

ORDER_SERVICE = "http://localhost:5001"
CUSTOMER_SERVICE = "http://localhost:5002"


async def main():
    pending_order = await place_order()
    print("pending", pending_order)
    # todo: create a retry method that keeps polling get order under data is different from pending_order (then exper with diff methods)
    order = await get_order(pending_order["id"])
    print("order", order)


async def place_order():
    async with httpx.AsyncClient() as client:
        response = await client.post(
            f"{ORDER_SERVICE}/order",
            data={
                "item_id": 1,
                "customer_id": 1
            }
        )
        print(response)
        return response.data


async def get_order(id: str):
    async with httpx.AsyncClient() as client:
        response = await client.get(f"{ORDER_SERVICE}/order/{id}")
        print(response)
        return response.data


if __name__ == "__main__":
    asyncio.run(main())
