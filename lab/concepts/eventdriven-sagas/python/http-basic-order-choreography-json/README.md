Order Service
* Creates order (pending state) and dispatches created event
* Updates order upon reciept of a credit reserved/credit exceeded event

Customer Service
* Listens to order created events
* Attempts to reserve credit from the customer for that order (then dispatches its own event of the result)