# Horcruxes and Shamir Secrets

## Overview

Shamir Secrets share an alarming resemblance to Horcruxes.

## Terminology

* [Horcrux][horcrux]: A part of a human soul
* [Shamir Secret][shamir]: A secret that will be split up

## Similarities

* Precious to the owner
  * Shamir Secret
    * A master secret is very precious to the owner.
  * Horcrux
    * A soul is very precious to the owner.
* Will be split up
  * Shamir Secret
    * A master secret is split up into various shares.
  * Horcrux
    * Horcruxes are created by splitting up the creator's soul.
* Must be kept secret and safe
  * Shamir Secret
    * A secret share is trusted to a person or secretly hidden.  The trustee or location is known only to the creator.
  * Horcrux
    * A horcrux is hidden within a trusted object, animal, or person known only to the creator
* No-part-relations
  * Shamir Secret
    * Each secret share found offers no clue about the other secret shares.  Neither the content nor the total number of other shares can be determined.
  * Horcrux
    * Each horcrux found offers no clue about the other horcruxes.  The total number of horcruxes cannot be deduced.
* Information-theoretic security
  * Shamir Secret
    * No information about the master secret is leaked until N shares are provided.
  * Horcrux
    * The original soul is indestructible until all horcruxes have been destroyed.
* Fault tolerant
  * Shamir Secret
    * As long as N > K, then one or more shares can be lost or destroyed while still being able to recover the master secret.  Lost or destroyed shares may not be known by the master secret owner.
  * Horcrux
    * The original soul continues to function normally until all horcruxes have been destroyed.  Lost or destroyed horcruxes may not be known by the soul's owner.

[horcrux]: https://harrypotter.fandom.com/wiki/Horcrux
[shamir]: https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing
