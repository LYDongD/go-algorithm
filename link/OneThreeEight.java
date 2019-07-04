package link;
import java.util.*;

public class OneThreeEight {
    private static class Node {
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

        Map<Node, Node> copiedDic = new HashMap<>();
        Node node = head;
        while(node != null) {
            copiedDic.put(node, new Node(node.val, null, null));
            node = node.next;
        }

        node = head;
        while(node != null) {
            copiedDic.get(node).next = copiedDic.get(node.next);
            copiedDic.get(node).random = copiedDic.get(node.random);
            node = node.next;
        }

        return copiedDic.get(head);
    }


    public static void main(String[] args) {
	    System.out.println("haha");
    }
}
