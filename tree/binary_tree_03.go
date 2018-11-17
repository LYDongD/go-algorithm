package tree

//用数组存储二叉搜索树,该方法还有问题，插入判断不合理
func CreateBinaryTree(source []int) []int {

	if source == nil {
		return nil
	}

	binaryTreeArray := make([]int, 16, 16)

	//binary tree always leave the first elemnt nil
	binaryTreeArray[1] = source[0]
	for i := 1; i < len(source); i++ {
		order := 1

		for binaryTreeArray[order] != 0 {
			if source[i] < binaryTreeArray[order] {
				order = 2 * order
			} else {
				order = 2*order + 1
			}
		}
		//update binaryTreeArray when there is no element to compare
		binaryTreeArray[order] = source[i]
	}

	return binaryTreeArray

}
