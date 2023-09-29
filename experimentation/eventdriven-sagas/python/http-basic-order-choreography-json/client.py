import asyncio

import httpx

ORDER_SERVICE = "http://localhost:5001"
CUSTOMER_SERVICE = "http://localhost:5002"


async def main():
    pending_order = await place_order()
    print("pending order view:", pending_order)

    order_changed = False
    while not order_changed:
        order = await get_order(pending_order["id"])
        if order == pending_order:
            print("order still unchanged")
        else:
            order_changed = True
            print(order)
    
    print("done")


async def place_order():
    async with httpx.AsyncClient() as client:
        response = await client.post(
            f"{ORDER_SERVICE}/order",
            json={
                "item_id": 1,
                "customer_id": 1
            }
        )
        return response.json()


async def get_order(id: str):
    async with httpx.AsyncClient() as client:
        response = await client.get(f"{ORDER_SERVICE}/order/{id}")
        return response.json()


if __name__ == "__main__":
    asyncio.run(main())
