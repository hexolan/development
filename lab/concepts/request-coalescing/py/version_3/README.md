unpublished on main request-coalescing-py

modifications:
this approach implements a lock
(
- one commented out approach uses lock instead of a queue
- another approach keeps the queue and uses the lock alongside
)

changes to the time counting to use time.perf_counter()

this is currently a mess and needs work
the timing is definitely askew