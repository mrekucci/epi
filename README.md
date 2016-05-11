epi
===

[![Build Status](https://travis-ci.org/mrekucci/epi.svg)](https://travis-ci.org/mrekucci/epi)
[![Coverage Status](https://coveralls.io/repos/mrekucci/epi/badge.svg?branch=master&service=github)](https://coveralls.io/github/mrekucci/epi?branch=master)
[![GitHub license](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE.txt)

This is a work-in-progress, solutions for [Elements of Programming Interviews][1] problems written in Golang.

Solutions
=========

Primitive Types
---------------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Computing the parity of a word][2]                                      | [tests][3]   |    ✓    |
| [Swap bits][4]                                                           | [tests][5]   |    ✓    |
| [Reverse bits][6]                                                        | [tests][7]   |    ✓    |
| [Find a closest integer with the same weight][8]                         | [tests][9]   |    ✓    |
| [Compute *x × y* without arithmetical operators][10]                     | [tests][11]  |         |
| [Compute *x/y*][12]                                                      | [tests][13]  |         |
| [Compute *x^y*][14]                                                      | [tests][15]  |         |
| [Reverse digits][16]                                                     | [tests][17]  |    ✓    |
| [Check if a decimal integer is a palindrome][18]                         | [tests][19]  |         |
| [Generate uniform random numbers][20]                                    | [tests][21]  |         |
| [Rectangle intersection][22]                                             | [tests][23]  |         |

Arrays
------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [The Dutch national flag problem][24]                                    | [tests][25]  |    ✓    |
| [Increment an arbitrary-precision integer][26]                           | [tests][27]  |         |
| [Multiply two arbitrary-precision integers][28]                          | [tests][29]  |         |
| [Advancing through an array][30]                                         | [tests][31]  |         |
| [Delete a key from an array][32]                                         | [tests][33]  |         |
| [Delete duplicates from a sorted array][34]                              | [tests][35]  |    ✓    |
| [Robot's minimum battery capacity][36]                                   | [tests][37]  |    ✓    | <!-- Rename to: "Buy and sell a stock once"  and fix code according to: https://github.com/epibook/epibook.github.io/commit/44f09980c4039b7a7b3eb5200c6994631e543c7a -->
| [Buy and sell a stock twice][38]                                         | [tests][39]  |         |
| [Enumerate all primes to *n*][40]                                        | [tests][41]  |    ✓    |
| [Permute the elements of an array][42]                                   | [tests][43]  |         |
| [Compute the next permutation][44]                                       | [tests][45]  |    ✓    |
| [Sample offline data][46]                                                | [tests][47]  |         |
| [Sample online data][48]                                                 | [tests][49]  |         |
| [Compute a random permutation][50]                                       | [tests][51]  |         |
| [Compute a random subset][52]                                            | [tests][53]  |         |
| [Generate nonuniform random numbers][54]                                 | [tests][55]  |         |
| [The Sudoku checker problem][56]                                         | [tests][57]  |         |
| [Compute the spiral ordering of a 2D array][58]                          | [tests][59]  |    ✓    |
| [Rotate a 2D array][60]                                                  | [tests][61]  |         |
| [Compute rows in Pascal’s Triangle][62]                                    | [tests][63]  |         |

Strings
-------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Interconvert strings and integers][64]                                  | [tests][65]  |    ✓    |
| [Base conversion][66]                                                    | [tests][67]  |         |
| [Compute the spreadsheet column encoding][68]                            | [tests][69]  |         |
| [Replace and remove][70]                                                 | [tests][71]  |         |
| [Test palindromicity][72]                                                | [tests][73]  |         |
| [Reverse all the words in a sentence][74]                                | [tests][75]  |    ✓    |
| [Compute all mnemonics for a phone number][76]                           | [tests][77]  |    ✓    |
| [The look-and-say problem][78]                                           | [tests][79]  |         |
| [Convert from Roman to decimal][80]                                      | [tests][81]  |         |
| [Compute all valid IP addresses][82]                                     | [tests][83]  |         |
| [Write a string sinusoidally][84]                                        | [tests][85]  |         |
| [Implement run-length encoding][86]                                      | [tests][87]  |    ✓    |
| [Implement the UNIX `tail` command][88]                                  | [tests][89]  |         |
| [Find the first occurrence of a substring][90]                           | [tests][91]  |    ✓    |

Linked List
-----------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Merge two sorted lists][92]                                             | [tests][93]  |    ✓    |
| [Reverse a singly linked list][94]                                       | [tests][95]  |         |
| [Reverse a single sublist][96]                                           | [tests][97]  |         |
| [Test for cyclicity][98]                                                 | [tests][99]  |    ✓    |
| [Test for overlapping lists—lists are cycle-free][100]                     | [tests][101] |         |
| [Test for overlapping lists—lists may have cycles][102]                    | [tests][103] |         |
| [Delete a node from a singly linked list][104]                           | [tests][105] |         |
| [Remove the *k*th last element from a list][106]                         | [tests][107] |         |
| [Remove duplicates from a sorted list][108]                              | [tests][109] |         |
| [Implement cyclic right shift for singly linked lists][110]              | [tests][111] |         |
| [Implement even-odd merge][112]                                          | [tests][113] |         |
| [Test whether a singly linked list is palindromic][114]                  | [tests][115] |         |
| [Implement list pivoting][116]                                           | [tests][117] |         |
| [Add list-based integers][118]                                           | [tests][119] |         |

Stacks and Queues
-----------------

### Stacks

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Implement a stack with max API][120]                                    | [tests][121] |    ✓    |
| [Evaluate RPN expressions][122]                                          | [tests][123] |    ✓    |
| [Test a string over “{,},(,),[,]” for well-formedness][124]                | [tests][125] |    ✓    |
| [Normalize pathnames][126]                                               | [tests][127] |         |
| [BST keys in sort order][128]                                            | [tests][129] |         |
| [Search a postings list][130]                                            | [tests][131] |         |
| [Compute buildings with a sunset view][132]                              | [tests][133] |         |
| [Sort a stack][134]                                                      | [tests][135] |         |

### Queues

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Compute binary tree nodes in order of increasing depth][136]            | [tests][137] |    ✓    |
| [Implement a circular queue][138]                                        | [tests][139] |    ✓    |
| [Implement a queue using stacks][140]                                    | [tests][141] |    ✓    |
| [Implement a queue with max API][142]                                    | [tests][143] |         |

Binary Trees
------------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Test if a binary tree is balanced][144]                                 | [tests][145] |    ✓    |
| [Test if a binary tree is symmetric][146]                                | [tests][147] |    ✓    |
| [Compute the lowest common ancestor in a binary tree][148]               | [tests][149] |    ✓    |
| [Compute the LCA when nodes have parent pointers][150]                   | [tests][151] |         |
| [Sum the root-to-leaf paths in a binary tree][152]                       | [tests][153] |         |
| [Find a root to leaf path with specified sum][154]                       | [tests][155] |         |
| [Compute the *k*th node in an inorder traversal][156]                    | [tests][157] |         |
| [Compute the successor][158]                                             | [tests][159] |         |
| [Implement an inorder traversal with *O(1)* space][160]                  | [tests][161] |    ✓    |
| [Reconstruct a binary tree from traversal data][162]                     | [tests][163] |         |
| [Reconstruct a binary tree from a preorder traversal with markers][164]  | [tests][165] |         |
| [Form a linked list from the leaves of a binary tree][166]               | [tests][167] |         |
| [Compute the exterior of a binary tree][168]                             | [tests][169] |         |
| [Compute the right sibling tree][170]                                    | [tests][171] |         |
| [Implement locking in a binary tree][172]                                | [tests][173] |         |

Heaps
-----

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Merge sorted files][174]                                                | [tests][175] |    ✓    |
| [Sort an increasing-decreasing array][176]                               | [tests][177] |    ✓    |
| [Sort an almost-sorted array][178]                                       | [tests][179] |         |
| [Compute the *k* closest stars][180]                                     | [tests][181] |         |
| [Compute the median of online data][182]                                 | [tests][183] |    ✓    |
| [Compute the *k* largest elements in a max-heap][184]                    | [tests][185] |         |
| [Implement a stack API using a heap][186]                                | [tests][187] |         |

Searching
---------

### Binary Search

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Search a sorted array for first occurrence of *k*][188]                 | [tests][189] |    ✓    |
| [Search a sorted array for the first element greater than *k*][190]      | [tests][191] |    ✓    |
| [Search a sorted array for entry equal to its index][192]                | [tests][193] |    ✓    |
| [Search a cyclically sorted array][194]                                  | [tests][195] |         |
| [Compute the integer square root][196]                                   | [tests][197] |         |
| [Compute the real square root][198]                                      | [tests][199] |    ✓    |

### Generalized Search

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Search in a 2D sorted array][200]                                       | [tests][201] |    ✓    |
| [Find the min and max simultaneously][202]                               | [tests][203] |    ✓    |
| [Find the *k*th largest element][204]                                    | [tests][205] |    ✓    |
| [Compute the optimum mailbox placement][206]                             | [tests][207] |         |
| [Find the missing IP address][208]                                       | [tests][209] |         |
| [Find the duplicate and missing elements][210]                           | [tests][211] |         |

Hash Tables
-----------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Partition into anagrams][212]                                           | [tests][213] |    ✓    |
| [Test for palindromic permutations][214]                                 | [tests][215] |    ✓    |
| [Is an anonymous letter constructible?][216]                             | [tests][217] |    ✓    |
| [Implement an ISBN cache][218]                                           | [tests][219] |         |
| [Compute the LCA, optimizing for close ancestors][220]                   | [tests][221] |    ✓    |
| [Compute the *k* most frequent queries][222]                             | [tests][223] |         |
| [Find the nearest repeated entries in an array][224]                     | [tests][225] |         |
| [Find the smallest subarray covering all values][226]                    | [tests][227] |    ✓    |
| [Find smallest subarray sequentially covering all values][228]           | [tests][229] |         |
| [Find the longest subarray with distinct entries][230]                   | [tests][231] |         |
| [Find the length of a longest contained range][232]                      | [tests][233] |         |
| [Compute the average of the top three scores][234]                       | [tests][235] |         |
| [Compute all string decompositions][236]                                 | [tests][237] |         |
| [Find a highest affinity pair][238]                                      | [tests][239] |         |
| [Test the Collatz conjecture][240]                                       | [tests][241] |         |
| [Implement a hash function for chess][242]                               | [tests][243] |         |

Sorting
-------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Compute the intersection of two sorted arrays][244]                     | [tests][245] |    ✓    |
| [Implement mergesort in-place][246]                                      | [tests][247] |    ✓    |
| [Count the frequencies of characters in a sentence][248]                 | [tests][249] |    ✓    |
| [Find unique elements][250]                                              | [tests][251] |         |
| [Render a calendar][252]                                                 | [tests][253] |         |
| [Sets of disjoint intervals][254]                                        | [tests][255] |         |
| [Compute the union of intervals][256]                                    | [tests][257] |    ✓    |
| [Partitioning and sorting an array with many repeated entries][258]      | [tests][259] |         |
| [Team photo day—1][260]                                                    | [tests][261] |         |
| [Implement a fast sorting algorithm for lists][262]                      | [tests][263] |    ✓    |
| [Compute a salary threshold][264]                                        | [tests][265] |         |

Binary Search Trees
-------------------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Test if a binary tree satisfies the BST property][266]                  | [tests][267] |    ✓    |
| [Find the first occurrence of a key in a BST][268]                       | [tests][269] |    ✓    |
| [Find the first key larger than a given value in a BST][270]             | [tests][271] |    ✓    |
| [Find the *k* largest elements in a BST][272]                            | [tests][273] |         |
| [Compute the LCA in a BST][274]                                          | [tests][275] |         |
| [Reconstruct a BST from traversal data][276]                             | [tests][277] |         |
| [Find the closest entries in three sorted arrays][278]                   | [tests][279] |         |
| [Enumerate numbers of the form *a + b√2*][280]                           | [tests][281] |         |
| [The most visited pages problem][282]                                    | [tests][283] |         |
| [Build a minimum height BST from a sorted array][284]                    | [tests][285] |         |
| [Insertion and deletion in a BST][286]                                   | [tests][287] |         |
| [Test if three BST nodes are totally ordered][288]                       | [tests][289] |         |
| [The range lookup problem][290]                                          | [tests][291] |         |
| [Add credits][292]                                                       | [tests][293] |         |
| [Count the number of entries in an interval][294]                        | [tests][295] |         |

Recursion
---------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [The Tower of Hanoi problem][296]                                        | [tests][297] |    ✓    |
| [Generate all nonattacking placements of *n*-Queens][298]                | [tests][299] |    ✓    |
| [Generate permutations][300]                                             | [tests][301] |    ✓    |
| [Generate the power set][302]                                            | [tests][303] |    ✓    |
| [Generate all subsets of size *k*][304]                                  | [tests][305] |    ✓    |
| [Generate strings of matched parens][306]                                | [tests][307] |         |
| [Generate palindromic decompositions][308]                               | [tests][309] |         |
| [Generate binary trees][310]                                             | [tests][311] |         |
| [Implement a Sudoku solver][312]                                         | [tests][313] |    ✓    |
| [Compute a Gray code][314]                                               | [tests][315] |         |
| [Compute the diameter of a tree][316]                                    | [tests][317] |         |

Dynamic Programming
-------------------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Count the number of score combinations][318]                            | [tests][319] |         |
| [Compute the Levenshtein distance][320]                                  | [tests][321] |         |
| [Count the number of ways to traverse a 2D array][322]                   | [tests][323] |         |
| [Plan a fishing trip][324]                                               | [tests][325] |         |
| [Search for a sequence in a 2D array][326]                               | [tests][327] |         |
| [The knapsack problem][328]                                              | [tests][329] |         |
| [Divide the spoils fairly][330]                                          | [tests][331] |         |
| [The **bedbathandbeyond.com** problem][332]                              | [tests][333] |         |
| [Find the minimum weight path in a triangle][334]                        | [tests][335] |         |
| [Pick up coins for maximum gain][336]                                    | [tests][337] |         |
| [Count the number of moves to climb stairs][338]                         | [tests][339] |         |
| [Compute the probability of a Republican majority][340]                  | [tests][341] |         |
| [The pretty printing problem][342]                                       | [tests][343] |         |
| [Find the longest nondecreasing subsequence][344]                        | [tests][345] |         |

Greedy Algorithms and Invariants
--------------------------------

### Greedy Algorithms

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Implement Huffman coding][346]                                          | [tests][347] |    ✓    |
| [Compute an optimum assignment of tasks][348]                            | [tests][349] |    ✓    |
| [Implement a schedule which minimizes waiting time][350]                 | [tests][351] |    ✓    |
| [The interval covering problem][352]                                     | [tests][353] |         |

### Invariants

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [The 3-sum problem][354]                                                 | [tests][355] |    ✓    |
| [Find the majority element][356]                                         | [tests][357] |         |
| [The gasup problem][358]                                                 | [tests][359] |         |
| [Compute the maximum water trapped by a pair of vertical lines][360]     | [tests][361] |         |
| [Compute the largest rectangle under the skyline][362]                   | [tests][363] |         |

Graphs
------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Identify the celebrity][364]                                            | [tests][365] |    ✓    |
| [Search a maze][366]                                                     | [tests][367] |    ✓    |
| [Paint a Boolean matrix][368]                                            | [tests][369] |    ✓    |
| [Compute enclosed regions][370]                                          | [tests][371] |         |
| [Degrees of connectedness—1][372]                                          | [tests][373] |    ✓    |
| [Clone a graph][374]                                                     | [tests][375] |         |
| [Making wired connections][376]                                          | [tests][377] |         |
| [Transform one string to another][378]                                   | [tests][379] |         |
| [The shortest straight-line program for *x^n*][380]                      | [tests][381] |         |
| [Team photo day—2][382]                                                    | [tests][383] |         |
| [Compute a shortest path with fewest edges][384]                         | [tests][385] |         |

Parallel Computing
------------------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Implement caching for a multithreaded dictionary][386]                  | [tests][387] |         |
| [Analyze two unsynchronized interleaved threads][388]                    | [tests][389] |         |
| [Implement synchronization for two interleaving threads][390]            | [tests][391] |         |
| [Implement a thread pool][392]                                           | [tests][393] |         |
| [Implement asynchronous callbacks][394]                                  | [tests][395] |         |
| [Implement a Timer class][396]                                           | [tests][397] |         |
| [The readers-writers problem][398]                                       | [tests][399] |         |
| [The readers-writers problem with write preference][400]                 | [tests][401] |         |
| [Test the Collatz conjecture in parallel][402]                           | [tests][403] |         |
| [Design TeraSort and PetaSort][404]                                      | [tests][405] |         |
| [Implement distributed throttling][406]                                  | [tests][407] |         |

Design Problems
---------------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Design a spell checker][408]                                            | [tests][409] |         |
| [Design a solution to the stemming problem][410]                         | [tests][411] |         |
| [Plagiarism detector][412]                                               | [tests][413] |         |
| [Pair users by attributes][414]                                          | [tests][415] |         |
| [Design a system for detecting copyright infringement][416]              | [tests][417] |         |
| [Design *TEX*][418]                                                      | [tests][419] |         |
| [Design a search engine][420]                                            | [tests][421] |         |
| [Implement PageRank][422]                                                | [tests][423] |         |
| [Design a scalable priority system][424]                                 | [tests][425] |         |
| [Create photomosaics][426]                                               | [tests][427] |         |
| [Implement Mileage Run][428]                                             | [tests][429] |         |
| [Implement Connexus][430]                                                | [tests][431] |         |
| [Design an online advertising system][432]                               | [tests][433] |         |
| [Design a recommendation system][434]                                    | [tests][435] |         |
| [Design an optimized way of distributing large files][436]               | [tests][437] |         |
| [Design the World Wide Web][438]                                         | [tests][439] |         |
| [Estimate the hardware cost of a photo sharing app][440]                 | [tests][441] |         |

Honors Class
------------

| Problem                                                                  | Test         | Solved  |
|--------------------------------------------------------------------------|:------------:|:-------:|
| [Compute the greatest common divisor][442]                               | [tests][443] |         |
| [Find the first missing positive entry][444]                             | [tests][445] |         |
| [Buy and sell a stock *k* times][446]                                    | [tests][447] |         |
| [Compute the maximum product of all entries but one][448]                | [tests][449] |         |
| [Compute the longest contiguous increasing subarray][450]                | [tests][451] |         |
| [Rotate an array][452]                                                   | [tests][453] |         |
| [Identify positions attacked by rooks][454]                              | [tests][455] |         |
| [Justify text][456]                                                      | [tests][457] |         |
| [Reverse sublists *k* at a time][458]                                    | [tests][459] |         |
| [Implement list zipping][460]                                            | [tests][461] |         |
| [Copy a postings list][462]                                              | [tests][463] |         |
| [Compute the median of a sorted circular linked list][464]               | [tests][465] |    ✓    |
| [Compute the longest substring with matching parens][466]                | [tests][467] |         |
| [Compute the maximum of a sliding window][468]                           | [tests][469] |         |
| [Implement preorder and postorder traversals without recursion][470]     | [tests][471] |         |
| [Compute fair bonuses][472]                                              | [tests][473] |         |
| [Find *k* elements closest to the median][474]                           | [tests][475] |         |
| [Search a sorted array of unknown length][476]                           | [tests][477] |         |
| [Search in two sorted arrays][478]                                       | [tests][479] |         |
| [Find the *k*th largest element—large *n*, small *k*][480]               | [tests][481] |         |
| [Find an element that appears only once][482]                            | [tests][483] |         |
| [Find the line through the most points][484]                             | [tests][485] |         |
| [Find the shortest unique prefix][486]                                   | [tests][487] |         |
| [Compute the smallest nonconstructible change][488]                      | [tests][489] |         |
| [Find the most visited pages in a window][490]                           | [tests][491] |         |
| [Convert a sorted doubly linked list into a BST][492]                    | [tests][493] |         |
| [Convert a BST to a sorted doubly linked list][494]                      | [tests][495] |         |
| [Merge two BSTs][496]                                                    | [tests][497] |         |
| [Test if a binary tree is an almost BST][498]                            | [tests][499] |         |
| [The view from above][500]                                               | [tests][501] |         |
| [Searching a min-first BST][502]                                         | [tests][503] |         |
| [Implement regular expression matching][504]                             | [tests][505] |         |
| [Synthesize an expression][506]                                          | [tests][507] |         |
| [Count inversions][508]                                                  | [tests][509] |         |
| [Draw the skyline][510]                                                  | [tests][511] |         |
| [Find the two closest points][512]                                       | [tests][513] |         |
| [Measure with defective jugs][514]                                       | [tests][515] |         |
| [Compute the maximum subarray sum in a circular array][516]              | [tests][517] |         |
| [Determine the critical height][518]                                     | [tests][519] |         |
| [Voltage selection in a logic circuit][520]                              | [tests][521] |         |
| [Find the maximum 2D subarray][522]                                      | [tests][523] |         |
| [Trapping water][524]                                                    | [tests][525] |         |
| [Load balancing][526]                                                    | [tests][527] |         |
| [Search for a pair-sum in an abs-sorted array][528]                      | [tests][529] |         |
| [The heavy hitter problem][530]                                          | [tests][531] |         |
| [Find the longest subarray whose sum ≤ *k*][532]                          | [tests][533] |         |
| [Degrees of connectedness—2][534]                                          | [tests][535] |         |
| [Compute a minimum delay schedule, unlimited resources][536]             | [tests][537] |         |
| [Road network][538]                                                      | [tests][539] |         |
| [Test if arbitrage is possible][540]                                     | [tests][541] |         |
| [The readers-writers problem with fairness][542]                         | [tests][543] |         |
| [Implement a producer-consumer queue][544]                               | [tests][545] |         |

[1]:   http://elementsofprogramminginterviews.com
[2]:   ptypes/parity.go
[3]:   ptypes/parity_test.go
[4]:   ptypes/swapbits.go
[5]:   ptypes/swapbits_test.go
[6]:   ptypes/reversebits.go
[7]:   ptypes/reversebits_test.go
[8]:   ptypes/closestint.go
[9]:   ptypes/closestint_test.go
[10]:  in_progress.md
[11]:  in_progress.md
[12]:  in_progress.md
[13]:  in_progress.md
[14]:  in_progress.md
[15]:  in_progress.md
[16]:  ptypes/reverseint.go
[17]:  ptypes/reverseint_test.go
[18]:  in_progress.md
[19]:  in_progress.md
[20]:  in_progress.md
[21]:  in_progress.md
[22]:  in_progress.md
[23]:  in_progress.md
[24]:  arrays/dutchflag.go
[25]:  arrays/dutchflag_test.go
[26]:  in_progress.md
[27]:  in_progress.md
[28]:  in_progress.md
[29]:  in_progress.md
[30]:  in_progress.md
[31]:  in_progress.md
[32]:  in_progress.md
[33]:  in_progress.md
[34]:  arrays/duplicates.go
[35]:  arrays/duplicates_test.go
[36]:  arrays/maxdiff.go
[37]:  arrays/maxdiff_test.go
[38]:  in_progress.md
[39]:  in_progress.md
[40]:  arrays/enumprimes.go
[41]:  arrays/enumprimes_test.go
[42]:  in_progress.md
[43]:  in_progress.md
[44]:  arrays/nextperm.go
[45]:  arrays/nextperm_test.go
[46]:  in_progress.md
[47]:  in_progress.md
[48]:  in_progress.md
[49]:  in_progress.md
[50]:  in_progress.md
[51]:  in_progress.md
[52]:  in_progress.md
[53]:  in_progress.md
[54]:  in_progress.md
[55]:  in_progress.md
[56]:  in_progress.md
[57]:  in_progress.md
[58]:  arrays/spiralmetrix.go
[59]:  arrays/spiralmetrix_test.go
[60]:  in_progress.md
[61]:  in_progress.md
[62]:  in_progress.md
[63]:  in_progress.md
[64]:  strings/intstrconv.go
[65]:  strings/intstrconv_test.go
[66]:  in_progress.md
[67]:  in_progress.md
[68]:  in_progress.md
[69]:  in_progress.md
[70]:  in_progress.md
[71]:  in_progress.md
[72]:  in_progress.md
[73]:  in_progress.md
[74]:  strings/reversewords.go
[75]:  strings/reversewords_test.go
[76]:  strings/phonemnemo.go
[77]:  strings/phonemnemo_test.go
[78]:  in_progress.md
[79]:  in_progress.md
[80]:  in_progress.md
[81]:  in_progress.md
[82]:  in_progress.md
[83]:  in_progress.md
[84]:  in_progress.md
[85]:  in_progress.md
[86]:  strings/rlecompr.go
[87]:  strings/rlecompr_test.go
[88]:  in_progress.md
[89]:  in_progress.md
[90]:  strings/index.go
[91]:  strings/index_test.go
[92]:  lists/mergesorted.go
[93]:  lists/mergesorted_test.go
[94]:  in_progress.md
[95]:  in_progress.md
[96]:  in_progress.md
[97]:  in_progress.md
[98]:  lists/checkcycle.go
[99]:  lists/checkcycle_test.go
[100]: in_progress.md
[101]: in_progress.md
[102]: in_progress.md
[103]: in_progress.md
[104]: in_progress.md
[105]: in_progress.md
[106]: in_progress.md
[107]: in_progress.md
[108]: in_progress.md
[109]: in_progress.md
[110]: in_progress.md
[111]: in_progress.md
[112]: lists/evenoddmerge.go
[113]: lists/evenoddmerge_test.go
[114]: in_progress.md
[115]: in_progress.md
[116]: in_progress.md
[117]: in_progress.md
[118]: in_progress.md
[119]: in_progress.md
[120]: stacks/max.go
[121]: stacks/max_test.go
[122]: stacks/eval.go
[123]: stacks/eval_test.go
[124]: stackt/wellformed.go
[125]: stackt/wellformed_test.go
[126]: in_progress.md
[127]: in_progress.md
[128]: in_progress.md
[129]: in_progress.md
[130]: in_progress.md
[131]: in_progress.md
[132]: in_progress.md
[133]: in_progress.md
[134]: in_progress.md
[135]: in_progress.md
[136]: queues/btorder.go
[137]: queues/btorder_test.go
[138]: queues/circqueue.go
[139]: queues/circqueue_test.go
[140]: queues/stackqueue.go
[141]: queues/stackqueue_test.go
[142]: in_progress.md
[143]: in_progress.md
[144]: btrees/balanced.go
[145]: btrees/balanced_test.go
[146]: btrees/symmetric.go
[147]: btrees/symmetric_test.go
[148]: btrees/lca.go
[149]: btrees/lca_test.go
[150]: in_progress.md
[151]: in_progress.md
[152]: in_progress.md
[153]: in_progress.md
[154]: in_progress.md
[155]: in_progress.md
[156]: in_progress.md
[157]: in_progress.md
[158]: in_progress.md
[159]: in_progress.md
[160]: btrees/traversal.go
[161]: btrees/traversal_test.go
[162]: in_progress.md
[163]: in_progress.md
[164]: in_progress.md
[165]: in_progress.md
[166]: in_progress.md
[167]: in_progress.md
[168]: in_progress.md
[169]: in_progress.md
[170]: in_progress.md
[171]: in_progress.md
[172]: in_progress.md
[173]: in_progress.md
[174]: heaps/mergesorted.go
[175]: heaps/mergesorted_test.go
[176]: heaps/sortk.go
[177]: heaps/sortk_test.go
[178]: in_progress.md
[179]: in_progress.md
[180]: in_progress.md
[181]: in_progress.md
[182]: heaps/median.go
[183]: heaps/median_test.go
[184]: in_progress.md
[185]: in_progress.md
[186]: in_progress.md
[187]: in_progress.md
[188]: bsearch/firstk.go
[189]: bsearch/firstk_test.go
[190]: bsearch/greaterk.go
[191]: bsearch/greaterk_test.go
[192]: bsearch/equals.go
[193]: bsearch/equals_test.go
[194]: in_progress.md
[195]: in_progress.md
[196]: in_progress.md
[197]: in_progress.md
[198]: bsearch/sqrtreal.go
[199]: bsearch/sqrtreal_test.go
[200]: search/matrix.go
[201]: search/matrix_test.go
[202]: search/minmax.go
[203]: search/minmax_test.go
[204]: search/largestkth.go
[205]: search/largestkth_test.go
[206]: in_progress.md
[207]: in_progress.md
[208]: in_progress.md
[209]: in_progress.md
[210]: in_progress.md
[211]: in_progress.md
[212]: htables/anagram.go
[213]: htables/anagram_test.go
[214]: htables/palindrom.go
[215]: htables/palindrom_test.go
[216]: htables/letter.go
[217]: htables/letter_test.go
[218]: in_progress.md
[219]: in_progress.md
[220]: htables/lca.go
[221]: htables/lca_test.go
[222]: in_progress.md
[223]: in_progress.md
[224]: in_progress.md
[225]: in_progress.md
[226]: htables/smallestsubarray.go
[227]: htables/smallestsubarray_test.go
[228]: in_progress.md
[229]: in_progress.md
[230]: in_progress.md
[231]: in_progress.md
[232]: in_progress.md
[233]: in_progress.md
[234]: in_progress.md
[235]: in_progress.md
[236]: in_progress.md
[237]: in_progress.md
[238]: in_progress.md
[239]: in_progress.md
[240]: in_progress.md
[241]: in_progress.md
[242]: in_progress.md
[243]: in_progress.md
[244]: sorting/intersection.go
[245]: sorting/intersection_test.go
[246]: sorting/merge.go
[247]: sorting/merge_test.go
[248]: sorting/charfreq.go
[249]: sorting/charfreq_test.go
[250]: in_progress.md
[251]: in_progress.md
[252]: in_progress.md
[253]: in_progress.md
[254]: in_progress.md
[255]: in_progress.md
[256]: sorting/union.go
[257]: sorting/union_test.go
[258]: in_progress.md
[259]: in_progress.md
[260]: in_progress.md
[261]: in_progress.md
[262]: sorting/sortlist.go
[263]: sorting/sortlist_test.go
[264]: in_progress.md
[265]: in_progress.md
[266]: bstrees/property.go
[267]: bstrees/property_test.go
[268]: bstrees/firstk.go
[269]: bstrees/firstk_test.go
[270]: bstrees/greaterk.go
[271]: bstrees/greaterk_test.go
[272]: in_progress.md
[273]: in_progress.md
[274]: in_progress.md
[275]: in_progress.md
[276]: in_progress.md
[277]: in_progress.md
[278]: in_progress.md
[279]: in_progress.md
[280]: in_progress.md
[281]: in_progress.md
[282]: in_progress.md
[283]: in_progress.md
[284]: in_progress.md
[285]: in_progress.md
[286]: in_progress.md
[287]: in_progress.md
[288]: in_progress.md
[289]: in_progress.md
[290]: in_progress.md
[291]: in_progress.md
[292]: in_progress.md
[293]: in_progress.md
[294]: in_progress.md
[295]: in_progress.md
[296]: recursion/hanoitowers.go
[297]: recursion/hanoitowers_test.go
[298]: recursion/nqueens.go
[299]: recursion/nqueens_test.go
[300]: recursion/perm.go
[301]: recursion/perm_test.go
[302]: recursion/powerset.go
[303]: recursion/powerset_test.go
[304]: recursion/subsets.go
[305]: recursion/subsets_test.go
[306]: in_progress.md
[307]: in_progress.md
[308]: in_progress.md
[309]: in_progress.md
[310]: in_progress.md
[311]: in_progress.md
[312]: recursion/sudoku.go
[313]: recursion/sudoku_test.go
[314]: in_progress.md
[315]: in_progress.md
[316]: in_progress.md
[317]: in_progress.md
[318]: in_progress.md
[319]: in_progress.md
[320]: in_progress.md
[321]: in_progress.md
[322]: in_progress.md
[323]: in_progress.md
[324]: in_progress.md
[325]: in_progress.md
[326]: in_progress.md
[327]: in_progress.md
[328]: in_progress.md
[329]: in_progress.md
[330]: in_progress.md
[331]: in_progress.md
[332]: in_progress.md
[333]: in_progress.md
[334]: in_progress.md
[335]: in_progress.md
[336]: in_progress.md
[337]: in_progress.md
[338]: in_progress.md
[339]: in_progress.md
[340]: in_progress.md
[341]: in_progress.md
[342]: in_progress.md
[343]: in_progress.md
[344]: in_progress.md
[345]: in_progress.md
[346]: greedy/huffman.go
[347]: greedy/huffman_test.go
[348]: greedy/pairtasks.go
[349]: greedy/pairtasks_test.go
[350]: greedy/schedule.go
[351]: greedy/schedule_test.go
[352]: in_progress.md
[353]: in_progress.md
[354]: invariants/threesum.go
[355]: invariants/threesum_test.go
[356]: in_progress.md
[357]: in_progress.md
[358]: in_progress.md
[359]: in_progress.md
[360]: in_progress.md
[361]: in_progress.md
[362]: in_progress.md
[363]: in_progress.md
[364]: graphs/celebrity.go
[365]: graphs/celebrity_test.go
[366]: graphs/maze.go
[367]: graphs/maze_test.go
[368]: graphs/flipcolor.go
[369]: graphs/flipcolor_test.go
[370]: in_progress.md
[371]: in_progress.md
[372]: graphs/minconnected.go
[373]: graphs/minconnected_test.go
[374]: in_progress.md
[375]: in_progress.md
[376]: in_progress.md
[377]: in_progress.md
[378]: in_progress.md
[379]: in_progress.md
[380]: in_progress.md
[381]: in_progress.md
[382]: in_progress.md
[383]: in_progress.md
[384]: in_progress.md
[385]: in_progress.md
[386]: in_progress.md
[387]: in_progress.md
[388]: in_progress.md
[389]: in_progress.md
[390]: in_progress.md
[391]: in_progress.md
[392]: in_progress.md
[393]: in_progress.md
[394]: in_progress.md
[395]: in_progress.md
[396]: in_progress.md
[397]: in_progress.md
[398]: in_progress.md
[399]: in_progress.md
[400]: in_progress.md
[401]: in_progress.md
[402]: in_progress.md
[403]: in_progress.md
[404]: in_progress.md
[405]: in_progress.md
[406]: in_progress.md
[407]: in_progress.md
[408]: in_progress.md
[409]: in_progress.md
[410]: in_progress.md
[411]: in_progress.md
[412]: in_progress.md
[413]: in_progress.md
[414]: in_progress.md
[415]: in_progress.md
[416]: in_progress.md
[417]: in_progress.md
[418]: in_progress.md
[419]: in_progress.md
[420]: in_progress.md
[421]: in_progress.md
[422]: in_progress.md
[423]: in_progress.md
[424]: in_progress.md
[425]: in_progress.md
[426]: in_progress.md
[427]: in_progress.md
[428]: in_progress.md
[429]: in_progress.md
[430]: in_progress.md
[431]: in_progress.md
[432]: in_progress.md
[433]: in_progress.md
[434]: in_progress.md
[435]: in_progress.md
[436]: in_progress.md
[437]: in_progress.md
[438]: in_progress.md
[439]: in_progress.md
[440]: in_progress.md
[441]: in_progress.md
[442]: in_progress.md
[443]: in_progress.md
[444]: in_progress.md
[445]: in_progress.md
[446]: in_progress.md
[447]: in_progress.md
[448]: in_progress.md
[449]: in_progress.md
[450]: in_progress.md
[451]: in_progress.md
[452]: in_progress.md
[453]: in_progress.md
[454]: in_progress.md
[455]: in_progress.md
[456]: in_progress.md
[457]: in_progress.md
[458]: in_progress.md
[459]: in_progress.md
[460]: in_progress.md
[461]: in_progress.md
[462]: in_progress.md
[463]: in_progress.md
[464]: honorsclass/median.go
[465]: honorsclass/median_test.go
[466]: in_progress.md
[467]: in_progress.md
[468]: in_progress.md
[469]: in_progress.md
[470]: in_progress.md
[471]: in_progress.md
[472]: in_progress.md
[473]: in_progress.md
[474]: in_progress.md
[475]: in_progress.md
[476]: in_progress.md
[477]: in_progress.md
[478]: in_progress.md
[479]: in_progress.md
[480]: in_progress.md
[481]: in_progress.md
[482]: in_progress.md
[483]: in_progress.md
[484]: in_progress.md
[485]: in_progress.md
[486]: in_progress.md
[487]: in_progress.md
[488]: in_progress.md
[489]: in_progress.md
[490]: in_progress.md
[491]: in_progress.md
[492]: in_progress.md
[493]: in_progress.md
[494]: in_progress.md
[495]: in_progress.md
[496]: in_progress.md
[497]: in_progress.md
[498]: in_progress.md
[499]: in_progress.md
[500]: in_progress.md
[501]: in_progress.md
[502]: in_progress.md
[503]: in_progress.md
[504]: in_progress.md
[505]: in_progress.md
[506]: in_progress.md
[507]: in_progress.md
[508]: in_progress.md
[509]: in_progress.md
[510]: in_progress.md
[511]: in_progress.md
[512]: in_progress.md
[513]: in_progress.md
[514]: in_progress.md
[515]: in_progress.md
[516]: in_progress.md
[517]: in_progress.md
[518]: in_progress.md
[519]: in_progress.md
[520]: in_progress.md
[521]: in_progress.md
[522]: in_progress.md
[523]: in_progress.md
[524]: in_progress.md
[525]: in_progress.md
[526]: in_progress.md
[527]: in_progress.md
[528]: in_progress.md
[529]: in_progress.md
[530]: in_progress.md
[531]: in_progress.md
[532]: in_progress.md
[533]: in_progress.md
[534]: in_progress.md
[535]: in_progress.md
[536]: in_progress.md
[537]: in_progress.md
[538]: in_progress.md
[539]: in_progress.md
[540]: in_progress.md
[541]: in_progress.md
[542]: in_progress.md
[543]: in_progress.md
[544]: in_progress.md
[545]: in_progress.md