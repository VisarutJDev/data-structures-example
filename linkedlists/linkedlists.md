# Floyd's Cycle-Finding Algorithm

Floyd's Cycle-Finding Algorithm, also known as the **Tortoise and Hare** algorithm, is an efficient way to detect cycles in a sequence of values or a linked list. It's named after the two pointers used in the algorithm: one moving slowly (the tortoise) and the other moving quickly (the hare).

### How the Algorithm Works

1. **Initialization**:
   - You start with two pointers: `slow` (the tortoise) and `fast` (the hare). Both pointers start at the head of the linked list.

2. **Movement**:
   - The `slow` pointer moves one step at a time (`slow = slow.Next`).
   - The `fast` pointer moves two steps at a time (`fast = fast.Next.Next`).

3. **Cycle Detection**:
   - As the `slow` and `fast` pointers move through the list, if there is a cycle, the `fast` pointer will eventually meet the `slow` pointer within the cycle.
   - If the `fast` pointer reaches the end of the list (i.e., `fast` becomes `nil`), then there is no cycle.

4. **Proof**:
   - If a cycle exists, the `fast` pointer will enter the cycle sooner and catch up with the `slow` pointer because it moves faster.
   - Specifically, once both pointers are inside the cycle, the distance between them decreases by one step each time they move. Eventually, the `fast` pointer "laps" the `slow` pointer.

### Why It Works

The key idea is that the `fast` pointer, moving twice as fast as the `slow` pointer, will eventually meet the `slow` pointer if a cycle exists. The meeting point within the cycle is a confirmation that the cycle exists.

### Example in a Linked List

Consider a linked list with a cycle:

1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 3 (cycle starts again here)

### Initial positions:

slow and fast start at node 1.
**First iteration:**

slow moves to node 2.
fast moves to node 3.
**Second iteration:**

slow moves to node 3.
fast moves to node 5.
**Third iteration:**

slow moves to node 4.
fast moves to node 3 (re-entering the cycle).
**Fourth iteration:**

slow moves to node 5.
fast moves to node 5.
At this point, slow and fast meet, confirming the presence of a cycle.

Advantages of Floyd's Cycle-Finding Algorithm
Space Efficiency: The algorithm only uses two pointers and doesn't require extra space proportional to the number of nodes, making it very space-efficient (O(1) space complexity).
Time Efficiency: The time complexity is O(n), where n is the number of nodes in the linked list. This is efficient since each node is visited at most a few times.
Summary
Floyd's Cycle-Finding Algorithm is a clever and efficient way to detect cycles in a linked list using two pointers moving at different speeds. If the fast pointer ever equals the slow pointer, a cycle exists; otherwise, if the fast pointer reaches the end, the list is cycle-free.