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

Linked List
-----------

| Problem                                                  | Test         |
|----------------------------------------------------------|:------------:|
| [Merge two sorted lists][28]                             | [tests][29]  |
| [Test for cyclicity][30]                                 | [tests][31]  |
| [Compute the median of a sorted circular linked list][32]| [tests][33]  |

Stacks and Queues
------------------

### Stacks

| Problem                                                  | Test         |
|----------------------------------------------------------|:------------:|
| [Implement a stack with max API][34]                     | [tests][35]  |

[1]: http://elementsofprogramminginterviews.com
[2]: ptypes/parity.go
[3]: ptypes/parity_test.go
[4]: ptypes/swapbits.go
[5]: ptypes/swapbits_test.go
[6]: ptypes/reversebits.go
[7]: ptypes/reversebits_test.go
[8]: ptypes/closestint.go
[9]: ptypes/closestint_test.go
[10]: ptypes/powerset.go
[11]: ptypes/powerset_test.go
[12]: ptypes/intstrconv.go
[13]: ptypes/intstrconv_test.go
[14]: arrays/dutchflag.go
[15]: arrays/dutchflag_test.go
[16]: arrays/maxdiff.go
[17]: arrays/maxdiff_test.go
[18]: arrays/nextperm.go
[19]: arrays/nextperm_test.go
[20]: arrays/spiralmetrix.go
[21]: arrays/spiralmetrix_test.go
[22]: strings/rlecompr.go
[23]: strings/rlecompr_test.go
[24]: strings/reversewords.go
[25]: strings/reversewords_test.go
[26]: strings/index.go
[27]: strings/index_test.go
[28]: lists/mergesorted.go
[29]: lists/mergesorted_test.go
[30]: lists/checkcycle.go
[31]: lists/checkcycle_test.go
[32]: lists/median.go
[33]: lists/median_test.go
[34]: stacks/max.go
[35]: stacks/max_test.go
