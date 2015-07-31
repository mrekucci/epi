epi
===

[![Build Status](https://travis-ci.org/mrekucci/epi.svg)](https://travis-ci.org/mrekucci/epi)
[![Coverage Status](https://coveralls.io/repos/mrekucci/epi/badge.svg?branch=master)](https://coveralls.io/r/mrekucci/epi?branch=master)
[![GitHub license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE.txt)

This is a work-in-progress, solutions for [Elements of Programming Interviews][1] problems written in Golang.

Solutions
=========

Primitive Types
---------------

| Problem                                                  | Test         |
|----------------------------------------------------------|:------------:|
| [Compute parity][2]                                      | [tests][3]   |
| [Swap bits][4]                                           | [tests][5]   |
| [Reverse bits][6]                                        | [tests][7]   |
| [Find a closest integer with the same weight][8]         | [tests][9]   |
| [Generate the power set][10]                             | [tests][11]  |
| [Interconvert strings and integers][12]                  | [tests][13]  |

Arrays and Strings
------------------

### Arrays

| Problem                                                  | Test         |
|----------------------------------------------------------|:------------:|
| [The Dutch national flag problem][14]                    | [tests][15]  |
| [Robot's minimum battery capacity][16]                   | [tests][17]  |
| [Compute the next permutation][18]                       | [tests][19]  |
| [Compute the spiral ordering of a 2D array][20]          | [tests][21]  |

### Strings

| Problem                                                  | Test         |
|----------------------------------------------------------|:------------:|
| [Implement run-length encoding][22]                      | [tests][23]  |
| [Reverse all the words in a sentence][24]                | [tests][25]  |
| [Find the first occurrence of a substring][26]           | [tests][27]  |


[1]: http://elementsofprogramminginterviews.com
[2]: parity.go
[3]: parity_test.go
[4]: swapbits.go
[5]: swapbits_test.go
[6]: reversebits.go
[7]: reversebits_test.go
[8]: closestint.go
[9]: closestint_test.go
[10]: powerset.go
[11]: powerset_test.go
[12]: intstrconv.go
[13]: intstrconv_test.go
[14]: dutchflag.go
[15]: dutchflag_test.go
[16]: maxdiff.go
[17]: maxdiff_test.go
[18]: nextperm.go
[19]: nextperm_test.go
[20]: spiralmetrix.go
[21]: spiralmetrix_test.go
[22]: rlecompr.go
[23]: rlecompr_test.go
[24]: reversewords.go
[25]: reversewords_test.go
[26]: strindex.go
[27]: strindex_test.go
