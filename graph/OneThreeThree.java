

public class OneThreeThree{
    
    private static class Node{
	public int val;
	public List<Node> neighbors;

	public Node() {}
	public Node(int _val, List<Node> _neighbors) {
	    val = _val;
	    neighbors = _neighbors;
	}
    }

    public Node cloneGraph(Node node) {
	if (node == null) {
	    return node;
	}

	Map<Integer, Node> copied = new HashMap<>();
	clone(node, copied);
	return node;
    }

    private void clone(Node node, Map<Integer, Node> copied) {
	Node cloneNode = new Node();
	cloneNode.val = node.val;
	copied.put(node.val, cloneNode);
	List<Node> cloneNeightbors = new ArrayList<>();
	for (i := 0;  i < node.neighbors.size(); i++) {
	   Node neightborNode = node.neighbors[i];
	   if (copied.get(neightborNode.val) != null) {
		cloneNeightbors.add(copied.get(neightborNode.val));
	   }else {
		cloneNeightbors.add(clone(neightborNode, copied));
	   }
	}
    }
}
