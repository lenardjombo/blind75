# Blind75 DSA

Welcome to my **Blind75** Data Structures & Algorithms (DSA) journey in **Go**. This repo includes problem solutions, unit tests, pattern-based notes and a tracker for my progress.

---

## DSA Pattern Cheat Sheet

Understanding **patterns** helps in recognizing solutions quickly and writing optimal code confidently. Here's a reference guide to the most common DSA patterns:

---

### 1. Two Pointers
**Use when:** Dealing with sorted arrays, pairs, distances, or partitions.  
**Best for:** Pair/triplet problems, array/window merging.  
**Spot it if:** You're comparing from two ends or closing in toward a condition.

- **Structures:** Arrays, slices  
- **Time Complexity:** `O(n)`, `O(n log n)`  
- **Examples:**  
  - Two Sum II (sorted input)  
  - 3Sum  
  - Container With Most Water

---

### 2. Sliding Window
**Use when:** Tracking subarrays or substrings.  
**Best for:** Max/min/sum/count over dynamic or fixed-size ranges.  
**Spot it if:** You’re moving over a contiguous block.

- **Structures:** Arrays, strings  
- **Time Complexity:** `O(n)`  
- **Examples:**  
  - Maximum Subarray  
  - Longest Substring Without Repeating Characters  
  - Minimum Window Substring

---

### 3. Fast & Slow Pointers
**Use when:** Detecting cycles, or finding midpoints in linked lists.  
**Best for:** Cycle detection, middle element.  
**Spot it if:** You compare elements at different speeds.

- **Structures:** Linked Lists  
- **Time Complexity:** `O(n)`  
- **Examples:**  
  - Linked List Cycle  
  - Happy Number  
  - Find Middle of Linked List

---

### 4. Hashing (Map/Set)
**Use when:** Checking uniqueness, frequency, or fast lookup.  
**Best for:** Frequency counting, detecting duplicates.  
**Spot it if:** You care about existence or counts.

- **Structures:** Map, Set  
- **Time Complexity:** `O(n)`  
- **Examples:**  
  - Two Sum  
  - Group Anagrams  
  - Longest Consecutive Sequence

---

### 5. Binary Search
**Use when:** Data is sorted or searchable via halving.  
**Best for:** Finding targets or boundaries.  
**Spot it if:** You repeatedly cut the search space in half.

- **Structures:** Arrays  
- **Time Complexity:** `O(log n)`  
- **Examples:**  
  - Binary Search  
  - Search in Rotated Sorted Array  
  - First/Last Position of Target

---

### 6. Depth-First Search (DFS)
**Use when:** Exploring trees/graphs deeply.  
**Best for:** Exploring paths, recursion, backtracking.  
**Spot it if:** You go deep before wide.

- **Structures:** Trees, graphs  
- **Time Complexity:** `O(V + E)`  
- **Examples:**  
  - Number of Islands  
  - Subsets  
  - Path Sum

---

### 7. Breadth-First Search (BFS)
**Use when:** Finding shortest path or level traversal.  
**Best for:** Level-order, shortest distance in unweighted graphs.  
**Spot it if:** You explore neighbors layer by layer.

- **Structures:** Queue  
- **Time Complexity:** `O(V + E)`  
- **Examples:**  
  - Binary Tree Level Order Traversal  
  - Word Ladder  
  - Clone Graph

---

### 8. Backtracking
**Use when:** Exploring all combinations with constraints.  
**Best for:** Permutations, N-Queens, configuration search.  
**Spot it if:** You try → undo → try next (DFS + undo).

- **Structures:** Recursion, path  
- **Time Complexity:** `O(2^n)`, `O(n!)`  
- **Examples:**  
  - N-Queens  
  - Word Search  
  - Sudoku Solver

---

### 9. Dynamic Programming (DP)
**Use when:** Overlapping subproblems, optimal substructure.  
**Best for:** Caching results to build up a solution.  
**Spot it if:** You break the problem into smaller, reusable pieces.

- **Structures:** Arrays, maps  
- **Time Complexity:** `O(n)`, `O(n^2)`, etc.  
- **Examples:**  
  - Climbing Stairs  
  - House Robber  
  - Longest Increasing Subsequence

---

### 10. Greedy
**Use when:** Local choices yield a global solution.  
**Best for:** Optimization problems, intervals, selections.  
**Spot it if:** You pick min/max values at each step.

- **Structures:** Arrays, sorting  
- **Time Complexity:** `O(n)` or `O(n log n)`  
- **Examples:**  
  - Jump Game  
  - Merge Intervals  
  - Gas Station

---

### 11. Heap / Priority Queue
**Use when:** Efficient top-k/min/max tracking.  
**Best for:** Task scheduling, merging sorted lists.  
**Spot it if:** You keep retrieving min/max dynamically.

- **Structures:** Heap  
- **Time Complexity:** `O(n log k)`  
- **Examples:**  
  - Kth Largest Element  
  - Merge K Sorted Lists  
  - Task Scheduler

---

### 12. Union Find / Disjoint Set
**Use when:** Grouping and connectivity.  
**Best for:** Connected components.  
**Spot it if:** You join/discover components.

- **Structures:** Arrays with path compression  
- **Time Complexity:** ~ `O(α(n))`  
- **Examples:**  
  - Graph Valid Tree  
  - Accounts Merge  
  - Number of Connected Components

---

### 13. Matrix Traversal (Island Pattern)
**Use when:** Navigating 2D grids.  
**Best for:** Connected components in a matrix.  
**Spot it if:** You move in 4 directions over a grid.

- **Structures:** 2D arrays  
- **Time Complexity:** `O(m * n)`  
- **Examples:**  
  - Number of Islands  
  - Flood Fill  
  - Max Area of Island

---

### 14. In-place Linked List Reversal
**Use when:** Reversing linked lists without extra memory.  
**Best for:** Full or partial linked list reversal.  
**Spot it if:** You’re modifying `next` pointers in-place.

- **Structures:** Linked Lists  
- **Time Complexity:** `O(n)`  
- **Examples:**  
  - Reverse Linked List  
  - Reverse Sublist  
  - Reverse Every K Elements

---

### 15. Two Heaps
**Use when:** Balancing two halves dynamically.  
**Best for:** Running median, partitioning.  
**Spot it if:** You maintain min-max balance.

- **Structures:** Min-heap + Max-heap  
- **Time Complexity:** `O(log n)`  
- **Examples:**  
  - Median of Data Stream  
  - Sliding Window Median  
  - Maximize Capital

---

### 16. Subsets
**Use when:** Finding power sets or combinations.  
**Best for:** Generating `2^n` configurations.  
**Spot it if:** You explore inclusion/exclusion.

- **Structures:** Recursion, queue  
- **Time Complexity:** `O(2^n)`  
- **Examples:**  
  - Subsets  
  - Subsets with Duplicates  
  - Permutations

---

### 17. Modified Binary Search
**Use when:** Searching for conditions instead of values.  
**Best for:** Range finding, peak element, transition points.  
**Spot it if:** You use binary search with a custom rule.

- **Structures:** Arrays  
- **Time Complexity:** `O(log n)`  
- **Examples:**  
  - Ceiling of Number  
  - Order-agnostic Search  
  - Next Letter

---

### 18. Bitwise XOR
**Use when:** Finding single/missing numbers.  
**Best for:** XOR cancellation tricks.  
**Spot it if:** You apply bit logic to toggle states.

- **Structures:** Arrays  
- **Time Complexity:** `O(n)`  
- **Examples:**  
  - Single Number  
  - Two Single Numbers  
  - Complement of Base 10

---

### 19. K-way Merge
**Use when:** Merging multiple sorted inputs.  
**Best for:** Efficient merging, top-k from k lists.  
**Spot it if:** You merge multiple sorted sources with a heap.

- **Structures:** Min-heap  
- **Time Complexity:** `O(N log K)`  
- **Examples:**  
  - Merge K Sorted Lists  
  - Kth Smallest Element in Sorted Matrix  
  - Smallest Range Covering K Lists

---

### 20. Topological Sort
**Use when:** Scheduling tasks with dependencies (DAG).  
**Best for:** Ordering prerequisites.  
**Spot it if:** You have directional dependency graphs.

- **Structures:** Graphs, in-degree map  
- **Time Complexity:** `O(V + E)`  
- **Examples:**  
  - Course Schedule  
  - Task Scheduler  
  - Topological Ordering

---

### 21. Trie (Prefix Tree)
**Use when:** Searching by word prefixes.  
**Best for:** Autocomplete, word dictionaries.  
**Spot it if:** You reuse common prefixes efficiently.

- **Structures:** Trie  
- **Time Complexity:** `O(L)` (L = word length)  
- **Examples:**  
  - Longest Common Prefix  
  - Word Search  
  - Insert/Search Word

---

### 22. Monotonic Stack
**Use when:** Finding next greater/smaller elements.  
**Best for:** Histograms, spans, temperature problems.  
**Spot it if:** You look at previous/next greater/smaller.

- **Structures:** Stack  
- **Time Complexity:** `O(n)`  
- **Examples:**  
  - Next Greater Element  
  - Largest Rectangle in Histogram  
  - Daily Temperatures

---

### 23. 0/1 Knapsack
**Use when:** Selecting items under constraints.  
**Best for:** Resource optimization under capacity.  
**Spot it if:** You choose to include/exclude each item.

- **Structures:** 2D/1D DP array  
- **Time Complexity:** `O(nW)`  
- **Examples:**  
  - 0/1 Knapsack  
  - Equal Subset Sum Partition  
  - Subset Sum

---

> **"Pattern recognition is your most powerful weapon in problem solving.  
> Train your eye to see structure, and your code will write itself."**

---

## Test Coverage

This repo includes Go test files for each solution using `testing` package.  
Run all tests with:

```bash
go test ./...
