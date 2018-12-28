//package math;


class Solution {
    public Node intersect(Node quadTree1, Node quadTree2){

        if (quadTree1.isLeaf) {
            return quadTree1.val ? quadTree1 : quadTree2;
        }

        if (quadTree2.isLeaf) {
            return quadTree2.val ? quadTree2 : quadTree1;
        }


        Node newNode = new Node();
        newNode.topLeft = intersect(quadTree1.topLeft, quadTree2.topLeft);
        newNode.topRight = intersect(quadTree1.topRight, quadTree2.topRight);
        newNode.bottomLeft = intersect(quadTree1.bottomLeft, quadTree2.bottomLeft);
        newNode.bottomRight = intersect(quadTree1.bottomRight, quadTree2.bottomRight);


        if (newNode.topLeft.isLeaf && newNode.topRight.isLeaf
                && newNode.bottomLeft.isLeaf && newNode.bottomRight.isLeaf
                && newNode.topLeft.val == newNode.topRight.val && newNode.topLeft.val == newNode.bottomLeft.val
                && newNode.topLeft.val == newNode.bottomRight.val) {
            newNode.val = newNode.topLeft.val;
            newNode.isLeaf = true;        
        }
        
        return newNode;
    }

    
    //class Node {
    //    public boolean val;
    //    public boolean isLeaf;
    //    public Node topLeft;
    //    public Node topRight;
    //    public Node bottomLeft;
    //    public Node bottomRight;

    //    public Node(){}

    //    public Node(boolean _val, boolean _isLeaf, Node _topLeft, Node _topRight, Node _bottomLeft, Node _bottomRight) {
    //        val = _val;
    //        isLeaf = _isLeaf;
    //        topLeft = _topLeft;
    //        topRight = _topRight;
    //        bottomLeft = _bottomLeft;
    //        bottomRight = _bottomRight;
    //    }
    //}
}
