Stack implementation with TDD
===

Requirements:
* A stack is empty on construction.
* A stack has size 0 on construction.
* After n pushes to an empty stack (n > 0), the stack is not empty and its size equals n.
* If one pushes x then pops, the value popped is x, and the size is one less than old size
* If one pushes x then peeks, the value returned is x, and the size stays the same
* If the size is n, then after n pops, the stack is empty and has size 0
* Popping from an empty stack return an error, ErrNoSuchElement
* Peeking into an empty stack return an error, ErrNoSuchElement
