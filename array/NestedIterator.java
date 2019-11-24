import java.util.Stack

public class NestedIterator implements Iterator<Integer> {
    
    private Stack<NestedInteger> nestedIntegerStack;

    public NestedIterator(List<NestedInteger> nestedList) {
	//init stack
	this.nestedIntegerStack = new Stack<>();
	for (int i = nestedList.size() - 1; i >= 0; i--) {
	    this.nestedIntegerStack.push(nestedList.get(i));
	}
    }

    @Override
    public Integer next() {
	return this.nestedIntegerStack.pop().getInteger();
    }

    @Override
    public boolean hasNext() {

	while(!nestedIntegerStack.isEmpty()) {
	    NestedInteger topNestedInteger = this.nestedIntegerStack.peek();
	    if (topNestedInteger.isInteger()) {
	        return true;
	    }
	    
	    this.nestedIntegerStack.pop();
	    for (int i = topNestedInteger.getList().size() - 1; i >= 0; i--) {
	        this.nestedIntegerStack.push(topNestedInteger.getList().get(i));
	    }
	}

	return false;



    }
}
