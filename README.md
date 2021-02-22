# Miller-Rabin

This is a simple example of a Miller-Rabin primality test. 
The formula needs to be run multiple times to have a high degree of success,
so I run it 10 times to check for a prime. The tests are a little wonky since they will
sometimes fail. One could improve this by testing after doing the rounds, but I'm not
likely going to make any further commits.

The main block checks for primes between 500,000 and 1,000,000 and outputs to standard out. 
A copy of this output is committed as primes.txt.

A far more refined version of this exists in the Go standard library, and there
are plenty of ways to do this faster, so there is no particular reason to use this.
My main purpose was to explore the math.