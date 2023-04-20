package core

type Tick interface {
	/**
	 * Initialization method.
	 * @method Initialize
	 * @construCtor
	**/
	Initialize()

	GetTree() *BehaviorTree

	/**
	 * return top subtree node.
	 * return nil when it is runing at major tree
	 *
	**/
	GetLastSubTree() *SubTree

	GetTarget() interface{}

	GetBlackboard() *Blackboard

	/**
	 * Called when entering a node (called by BaseNode).
	 * @method _enterNode
	 * @param {Object} node The node that called this method.
	 * @protected
	**/
	_enterNode(node IBaseNode)

	/**
	 * Callback when opening a node (called by BaseNode).
	 * @method _openNode
	 * @param {Object} node The node that called this method.
	 * @protected
	**/
	_openNode(node *BaseNode)

	/**
	 * Callback when ticking a node (called by BaseNode).
	 * @method _tickNode
	 * @param {Object} node The node that called this method.
	 * @protected
	**/
	_tickNode(node *BaseNode)
	/**
	 * Callback when closing a node (called by BaseNode).
	 * @method _closeNode
	 * @param {Object} node The node that called this method.
	 * @protected
	**/
	_closeNode(node *BaseNode)

	pushSubtreeNode(node *SubTree)
	popSubtreeNode()

	/**
	 * Callback when exiting a node (called by BaseNode).
	 * @method _exitNode
	 * @param {Object} node The node that called this method.
	 * @protected
	**/
	_exitNode(node *BaseNode)

	_getOpenNodes() []IBaseNode

	_getNodeCount() int

	SetTree(tree *BehaviorTree)

	SetDebug(debug interface{})

	SetTarget(target interface{})

	SetBlackBoard(bbd *Blackboard)
}
