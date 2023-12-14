# Branch Sums

Given a Binary Search Tree return an array which shows us the sum of all the branches present in the tree

## Example
            10
           /   \
          5     15
         / \   /  \
        2   5 13  22
       /        \
      1          14

We have 4 branches in total
        1: 10-5-2-1
        2: 10-5-5
        3: 10-15-13-14
        4: 10-15-22

return array is: [18,20,52,57]