# N Meetings in One Room

> Striver SDE Sheet 043 · Day 8 Greedy — pick meetings by earliest end time; the finisher that frees the room soonest can never hurt. Go, stdlib only, tested.

## Problem

One meeting room, n meetings with start and end times; a meeting can be scheduled only if it starts *strictly after* the previous one ends. Maximize the number of meetings held.

## Approach

Sort by end time and greedily accept every meeting that starts after the last accepted end. The exchange argument that justifies it: among all compatible choices, the one ending earliest leaves the largest remaining window, so any optimal schedule can be rewritten to start with it without losing meetings. Sorting an index array (rather than zipping pairs) keeps the inputs untouched. Note the strict `>` — this problem's convention says back-to-back meetings clash; its sibling, activity selection (#048), allows touching intervals.

**Complexity:** O(n log n) time, O(n) index space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
