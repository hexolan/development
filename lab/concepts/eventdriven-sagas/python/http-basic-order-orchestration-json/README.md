Order Service
* Creates order (pending state) and dispatches a reserve credit command
* Updates order based on response to the reserve credit command

Customer Service
* Listens to customer command channel
* Attempts to reserve credit in response to reserve credit command (and returns response to the create order saga channel)