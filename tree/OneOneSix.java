import java.util.LinkedList;
import java.util.Queue;

public class OneOneSix {
    private static class Node {
	    public int val;
	    public Node left;
	    public Node right;
	    public Node next;

	    public Node(){}

	    public Node(int _val, Node _left, Node _right, Node _next) {
	        val = _val;
	        left = _left;
	        right = _right;
	        next =_next;
	    }
    }

    private static class LevelNode {
	    public int level;
	    public Node node;

	    public LevelNode(){}

	    public LevelNode(int _level, Node _node) {
	        level = _level;
	        node = _node;
	    }
    }

    public Node connect(Node root) {
	    if (root == null) {
	        return root;
	    }

        Queue<LevelNode> queue = new LinkedList<>();
        int currentLevel = 0;
        LevelNode rootNode = new LevelNode(1, root);
        queue.offer(rootNode);
        Node lastNode = null;
        while(!queue.isEmpty()){
            LevelNode levelNode = queue.poll();
            if (levelNode.level == currentLevel) {
                levelNode.node.next = lastNode;
            }else {
                levelNode.node.next = null;
                currentLevel = levelNode.level;
            }

            lastNode = levelNode.node;
            if (levelNode.node.right != null) {
                LevelNode rightNode = new LevelNode(levelNode.level+1, levelNode.node.right);
                queue.offer(rightNode);
            }

            if (levelNode.node.left != null) {
                LevelNode leftNode = new LevelNode(levelNode.level+1, levelNode.node.left);
                queue.offer(leftNode);
            }
        }

        return root;
    }

    public static void main(String[] args) {
    	System.out.println("hello");
    }
} 
