package link

public class OneThreeEight {
    class static Node {
	public int val;
	public Node next;
	public Node random;

	public Node() {
	}

	public Node(int _val, Node _next, Node _random) {
	    val = _val;
	    next = _next;
	    random = _random;
	}
    }

    public Node copyRandomList(Node head) {
	if (head == null) {
	    return null;
	}

	Map<Integer, Node> copiedDic = new HashMap<>();

	return deepCopy(head, copiedDic)
    }

    private Node deepCopy(Node node, Map copiedDic) {
	if (null == node) {
	    return null;
	}

	Node copiedNode = new Node();
	copiedDic.put(copiedNode.val, copiedNode);

	if (copiedDic.get(node.next.val) != null) {
	    copiedNode.next = opiedDic.get(node.next.val);
	}else {
	    copiedNode.next = deepCopy(node.next, copiedDic);
	}

	if (copiedDic.get(node.random.val) != null) {
	    copiedNode.random = copiedDic.get(node.random.val);	
	}else {
	    copiedNode.random = deepCopy(node.random, copiedDic);
	}

	return node;
    }

    public static main() {
	System.out.println("haha");
    }
}
