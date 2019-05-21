# Secure Rolls

A service allowing a DM to request a roll from a Player and the player to respond with a secure roll that is guaranteed
not to be tampered with.  This is done by a series of signatures applied to json payloads using the JWS protocol.

This package also defines a set of libraries for composing rolls and using the secure rand facility on the current
platform to produce cryptographically secure random numbers.

# Rolls

A roll is a single die roll or a combination of die rolls manipulated in a few different ways. The following are the
types of rolls supported by this service:

## Constants (e.g. 1)

A constant in roll expression that is not used by the single die roll term or the Multiple Homogeneous Die Roll term.

## Single Die Roll (e.g. d20)

The traditional d4, d6, d8, d10, d12, d20, or d100.  But since we are using a computer to roll the dice, we can use as
many sides as we want.  Therefore anything that can match `d[0-9]+` is allowed.  That means d7 or d81 ...

## Multiple Homogeneous Die Rolls (e.g. 3d6)

Rolling a character for your new DnD 5e compaign.  Then roll 6 separate `3d6` for your stats. 

## Discarding Die Rolls (e.g. 4d6D1)

While the discarding feature can be applied to any die roll that produces multiple base results it is most often used
with the Multiple Homogeneous Die Rolls described above.  For instance, say you want to let your players have a better
chance at high stats.  You might choose to allow them to roll `4d6` for their stats and discard the lowest roll.  To do
this you would request a `4d6D1`.  For the opposite affect you would ask for a `4d6D>1`, the `>` tells the roller to
discard the higher values.

## Rerolling Die Rolls (e.g. 8d6r<2)

Let's say one of your characters has a class feature that allows them to reroll all 1s in their damage roll.  You might
do a `10d6r<2`  

## Mulipliers

Any two rolls or constants can be multiplied.  For instance `d6 * d8` would roll a `d6` and multiply the result by the 
result of rolling a `d8`.  Also, `d6 * 2` is legal and would result in an even number between 2 and 12 inclusive.

## Summation

Any two rolls or constants can be summed.  For instance `d6 + d8` would roll a `d6` and add the result to the result of
rolling a `d8`.

