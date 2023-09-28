# Event Driven Project

## Service Brainstorm

* Customers/Users/Profile Service

* Products

* Inventory

* Orders

* Payments

* Notifications

* Authentication

* Fulfilment?/Shipping? Service

* Shopping Cart

## Planned Features

* List of Products (possibly into seperate 'listing' service or something)
  * Featured Products
    * potentially implement request coalescing for fetching featured (as something to show)
  * Recommended Products
  * Recently Purchased? ("Buy Again")

* Add Products to Shopping Cart
  * View Shopping Cart
  * Buy Items in Shopping Cart
  * Buy Items Individually (skip cart)

* View Products
  * Check Stock of Products (out-of-stock)

* Ordering Products
  * Customer should have a sufficient balance (use a SAGA; cancel if not)
  * When product is ordered should be added to fulfilment queue? (staff can mark the order as shipped, etc)
  * When status of order is updated - send notification to user (email, sms, etc)

* Customers
  * Customers can login
  * Maybe customers -> profile service (as there can be staff accounts, etc)

## Clasification (Extra)

* Project Name
  * Generic Frontend with standardised (openapi swagger.json) HTTP route schema
  * That frontend can be plugged into any backend (allowing me to exper and provide examples of event driven, grpc, monolith, etc)
  * Project can sort of act like the real-world example project for different technologies/communication patterns
    * https://github.com/gothinkster/realworld
    * Potentially build my own realworld project?

* Generic Standardised Project (temp codename: Isotoped)
  * Dev Codename Ideas/Inspiration:
    * Isotope
    * Isomorphic
    * Isomorphous (capable of crystallizing in a form similar to that of another compound or mineral - used esp. of substances so closely related that they form end members of a series of solid solutions)
    * Derived
    * Isosteric (having the same number of valence electrons in the same configuration, but differing in the kinds and numbers of atoms)