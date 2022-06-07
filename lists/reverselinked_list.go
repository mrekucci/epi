package lists


type Node struct {
	nextNode *Node
	property rune
}

func ReverseLinkedList(nodeList *Node) *Node {
	var currentNode *Node
	currentNode = nodeList
	var topNode *Node = nil
	for {
		if currentNode == nil {
			break
		}
		var tempNode *Node
		tempNode = currentNode.nextNode
		currentNode.nextNode = topNode
		topNode = currentNode
		currentNode = tempNode
	}
	return topNode
}