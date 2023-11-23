# Request Coalescing Demo (Python)

This repository serves as a simple demonstration of request coalescing in Python (using Asyncio, FastAPI and HTTPX).

## What is request coalescing?

TODO

## Results

### Standard (Not Coalesced)

```bash
> py test.py
Making 100 x 5 concurrent requests (500 total)...
Took 186.244ms
Metrics: {'requests': 500, 'db_calls': 500}
```

### Coalesced

```bash
> py test.py
Making 100 x 5 concurrent requests (500 total)...
Took 42.809ms
Metrics: {'requests': 500, 'db_calls': 100}
```

## License

This repository is licensed under the [MIT License.](/LICENSE)
