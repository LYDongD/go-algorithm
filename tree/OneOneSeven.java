public class OneOneSeven {
    private static class Node {
        public int val;
        public Node left;
        public Node right;
        public Node next;

        public Node() {}
    }


    public Node connect(Node root) {
        if (root == null) {
            return;
        }    

        Node leftestNode = null;
        Node current = root;

        Node fromNode = null;
        while (current != null)  {
            if (current.Left) != null {
                if (fromNode != null) {
                    fromNode.Next = current.Left;
                    fromNode = current.Left;
                }
            }

            if (current.Right) != null {
                fromNode.Next = current.Right;
                fromNode.Next = current.Right;
            }

            current = current.Next;
            if current == null {
                current = leftestNode
            }
        }
    }
}
