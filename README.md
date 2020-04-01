# Transformers TCG Card Database

This is a project to create program that manages a card database for the Transformers TCG. This is mostly an experiment to learn database design, but I would like to expand it into a deck builder and collection tracker eventually. What follows is some notes regarding game information and design.

## Pip Combinations

There are currently 5 pips in the game:

| Color   | Code |
| ------- | - |
| Orange  | O |
| Blue    | B |
| White   | W |
| Green   | G |
| Black   | B |

With each card having a minimum of 0 pips and a maximum of 3 pips, there are 26 possible combinations without duplicates. Of those, 16 are used in the game so far. There are additionally 4 types of multi-pip cards which are mono-colored: OO, BB, LL, and LLL, bringing our total to 20 used. Currently, there are no WW or GG cards, but that may change in the future now that trait-specific pips have been introduced. 

**NOTE: As of this writing, there are about 20 Battle Cards left to be revealed in wave 5, so these numbers will change**

There are *a lot* of single-pip cards so an accurate count of those will come later.

| Pip Combo | Number of Cards |
| --------- | --------------- |
| Blank     | ?               |
| O         | ?               |
| B         | ?               |
| L         | ?               |
| W         | ?               |
| G         | ?               |
| OO        | 6               |
| BB        | 5               |
| LL        | 10              |
| OB        | 6               |
| OL        | 9               |
| OG        | 8               |
| BL        | 7               |
| BG        | 12              |
| LW        | 1               |
| LG        | 7               |
| WG        | 13              |
| OBW       | 2               |
| OLG       | 1               |
| LLL       | 1               |
| BW        | 0               |
| OW        | 0               |
| OBL       | 0               |
| OBG       | 0               |
| OLW       | 0               |
| OWG       | 0               |
| BLW       | 0               |
| BLG       | 0               |
| BWG       | 0               |
| LWG       | 0               |

