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
            return root;
        }    

        Node leftestNode = null;
        Node current = root;

        Node fromNode = null;
        while (current != null)  {
            if (current.left != null) {
                if (fromNode != null) {
                    fromNode.next = current.left;
                }else {
                    leftestNode = current.left;    
                }

                fromNode = current.left;
            }

            if (current.right != null) {
                if (fromNode != null) {
                    fromNode.next = current.right;
                }else {
                    leftestNode = current.right;
                }
                fromNode = current.right;
            }

            current = current.next;
            if (current == null) {
                current = leftestNode;
                leftestNode = null;
                fromNode = null;
            }
        }

        return root;
    }


    public static void main(String[] args) {
        System.out.println("hello");
    }
}
