# Sleeping Barber implementation in GO

This is simple project for learning purposes. Just trying to grasp Go concurrency model.

License: MIT

The problem is simple and exposed there: [Sleeping Barber Problem](http://en.wikipedia.org/wiki/Sleeping_barber_problem)

## Rules of SBP:

* a barber manage a shop with a waiting room
* a waiting room has got some seats
* some clients enter the shop, try to find a free seat
* free seat: ok, wait for barber or wake him if he is sleeping
* no seat: go out... (and retry later if you want)
* if the barber has no clients, he decide to sleep
* if client enter and see that the barber is sleeping, wake him

## Usage

```
$ sleeping-barber -h
Usage of /tmp/go-build908394022/command-line-arguments/_obj/exe/barber:
  -barber_delay=15: Time of barber working on customer [ms]
  -customers=10: Number of customers to come to barber shop
  -customers_delay=5: Time between customers incoming to shop [ms]
  -customers_return_delay=20: Time between customers' retry if the barber shop is full [ms]
  -seats=3: Number of seats in barber shop
exit status 2
```

## Example of run
```
$ sleeping-barber 
BARBER: Sleeping.
CUSTOMER 0: Entering barber shop.
BARBER: Woken up by customer 0.
CUSTOMER 1: Entering barber shop.
CUSTOMER 2: Entering barber shop.
BARBER: Customer 0 served.
CUSTOMER 0: Leaving barber shop happy.
CUSTOMER 3: Entering barber shop.
CUSTOMER 4: Entering barber shop.
CUSTOMER 5: Leaving barber shop unserviced.
BARBER: Customer 1 served.
CUSTOMER 1: Leaving barber shop happy.
CUSTOMER 6: Entering barber shop.
CUSTOMER 7: Leaving barber shop unserviced.
CUSTOMER 8: Leaving barber shop unserviced.
BARBER: Customer 2 served.
CUSTOMER 2: Leaving barber shop happy.
CUSTOMER 5: Entering barber shop.
CUSTOMER 9: Leaving barber shop unserviced.
CUSTOMER 7: Leaving barber shop unserviced.
BARBER: Customer 3 served.
CUSTOMER 3: Leaving barber shop happy.
CUSTOMER 8: Entering barber shop.
CUSTOMER 9: Leaving barber shop unserviced.
BARBER: Customer 4 served.
CUSTOMER 4: Leaving barber shop happy.
CUSTOMER 7: Entering barber shop.
CUSTOMER 9: Leaving barber shop unserviced.
BARBER: Customer 6 served.
CUSTOMER 6: Leaving barber shop happy.
BARBER: Customer 5 served.
CUSTOMER 5: Leaving barber shop happy.
CUSTOMER 9: Entering barber shop.
BARBER: Customer 8 served.
CUSTOMER 8: Leaving barber shop happy.
BARBER: Customer 7 served.
CUSTOMER 7: Leaving barber shop happy.
BARBER: Customer 9 served.
BARBER: Closing shop, going home.
CUSTOMER 9: Leaving barber shop happy.
```

