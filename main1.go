/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if(head == nil){
        return nil;
    }
    nextNode := head.Next;
    currNode := head;
    for nextNode!= nil {
        //fmt.Println(currNode.Val);
        if(currNode.Val == nextNode.Val){
            nextNode = nextNode.Next;
            currNode.Next = nextNode;
            continue;
        }
        nextNode = nextNode.Next;
        currNode = currNode.Next
        // fmt.Println(currNode.Val);
    }
    return head;
}
