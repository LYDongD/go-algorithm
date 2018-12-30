package tree;

import java.util.ArrayList;
import java.util.List;


class Solution {

    static class Node {
        public int val;
        public List<Node> children;

        public Node() {}

        public Node(int _val,List<Node> _children) {
            val = _val;
            children = _children;
        }
    };


    public List<Integer> preorder(Node root) {
        List<Integer> result = new LinkedList<>();
        preOrderRecursively(root, result);
        return result;
    }


    private void preOrderRecursively(Node root, List<Integer> result) {
        if (root == null) {
            return;
        }

       result.add(root.val);
       if (root.children != null) {
           for (Node node : root.children) {
                preOrderRecursively(node, result);
           }
       }

    }

    //public static void main(String[] args) {
    //    Node root = new Node();
    //    root.val = 1;
    //    List<Node> firstLevel = new ArrayList<>();
    //    Node child01 = new Node();
    //    child01.val = 3;
    //    Node child02 = new Node();
    //    child02.val = 2;
    //    Node child03 = new Node();
    //    child03.val = 4;
    //    firstLevel.add(child01);
    //    firstLevel.add(child02);
    //    firstLevel.add(child03);
    //    root.children = firstLevel;

    //    List<Node> secondLevel = new ArrayList<>();
    //    Node child04 = new Node();
    //    child04.val = 5;
    //    Node child05 = new Node();
    //    child05.val = 6;
    //    secondLevel.add(child04);
    //    secondLevel.add(child05);
    //    child01.children = secondLevel;

    //    Solution solution = new Solution();
    //    List<Integer> result = solution.preorder(root);
    //    for (int ele : result) {
    //        System.out.println(ele);
    //    }
    //}
}


