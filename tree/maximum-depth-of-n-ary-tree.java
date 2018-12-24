package tree;

import java.util.List;
import java.util.ArrayList;

class Solution {

    Integer maxDepth = Integer.MIN_VALUE;


    public int maxDepth(Node root) {

        if (root == null) return 0;
        
        maxDepth(root, 1);

        return maxDepth;
    }


    public void maxDepth(Node root, int depth) {
        
        if (root == null) {
            return;
        }

        if (depth > maxDepth) {
            maxDepth = depth;
        }

        depth++;

        for (Node node : root.children) {
            maxDepth(node, depth);
        }
        
    }


//    static class Node {
//        public int val;
//        public List<Node> children;
//
//        public Node() {}
//
//        public Node(int _val, List<Node> _children) {
//            val = val;
//            children = _children;
//        }
//    }
//
//    public static void main(String[] args) {
//        
//        Node root = new Node();
//        root.val = 3;
//        List<Node> children = new ArrayList<>();
//        Node node1 = new Node();
//        node1.val = 3;
//        Node node2 = new Node();
//        node2.val = 2;
//        Node node3 = new Node();
//        node3.val = 4;
//        children.add(node1);
//        children.add(node2);
//        children.add(node3);
//        root.children = children;
//
//      //  List<Node> children2 = new ArrayList<>();
//      //  Node node4 = new Node();
//      //  node4.val = -1;
//      //  Node node5 = new Node();
//      //  node5.val = 8;
//      //  children2.add(node4);
//      //  children2.add(node5);
//      //  node1.children = children2;
//      //  
//      
//        Solution solution =  new Solution();
//        int result = solution.maxDepth(root);
//        System.out.println(result) ;
//    }

}

