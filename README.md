Sleeping Barber implementation in GO
====================================

This is simple project for learning purposes. Just trying to grasp Go concurrency model.

The problem is simple and exposed there: [Sleeping Barber Problem](http://en.wikipedia.org/wiki/Sleeping_barber_problem)

Rules of SBP:

* a barber manage a shop with a waiting room
* a waiting room has got some seats
* some clients enter the shop, try to find a free seat
* free seat: ok, wait for barber or wake him if he is sleeping
* no seat: go out... (and retry later if you want)
* if the barber has no clients, he decide to sleep
* if client enter and see that the barber is sleeping, wake him

