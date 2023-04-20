package core

import (
	_ "fmt"
)

/**
 * A new Tick object is instantiated every tick by BehaviorTree. It is passed
 * as parameter to the nodes through the tree during the traversal.
 *
 * The role of the Tick class is to store the instances of tree, debug,
 * target and blackboard. So, all nodes can access these informations.
 *
 * For internal uses, the Tick also is useful to store the open node after
 * the tick signal, in order to let `BehaviorTree` to keep track and close
 * them when necessary.
 *
 * This class also makes a bridge between nodes and the debug, passing the
 * node state to the debug if the last is provided.
 *
 * @module b3
 * @class Tick
**/
type TickImp struct {
	/**
	 * The tree reference.
	 * @property {b3.BehaviorTree} tree
	 * @readOnly
	**/
	tree *BehaviorTree
	/**
	 * The debug reference.
	 * @property {Object} debug
	 * @readOnly
	 */
	debug interface{}
	/**
	 * The target object reference.
	 * @property {Object} target
	 * @readOnly
	**/
	target interface{}
	/**
	 * The blackboard reference.
	 * @property {b3.Blackboard} blackboard
	 * @readOnly
	**/
	Blackboard *Blackboard
	/**
	 * The list of open nodes. Update during the tree traversal.
	 * @property {Array} _openNodes
	 * @protected
	 * @readOnly
	**/
	_openNodes []IBaseNode

	/**
	 * The list of open subtree node.
	 * push subtree node before execute subtree.
	 * pop subtree node after execute subtree.
	**/
	_openSubtreeNodes []*SubTree

	/**
	 * The number of nodes entered during the tick. Update during the tree
	 * traversal.
	 *
	 * @property {Integer} _nodeCount
	 * @protected
	 * @readOnly
	**/
	_nodeCount int
}

func NewTick() Tick {
	tick := &TickImp{}
	tick.Initialize()
	return tick
}

/**
 * Initialization method.
 * @method Initialize
 * @construCtor
**/
func (this *TickImp) Initialize() {
	// set by BehaviorTree
	this.tree = nil
	this.debug = nil
	this.target = nil
	this.Blackboard = nil

	// updated during the tick signal
	this._openNodes = nil
	this._openSubtreeNodes = nil
	this._nodeCount = 0
}

func (this *TickImp) GetTree() *BehaviorTree {
	return this.tree
}

/**
 * Called when entering a node (called by BaseNode).
 * @method _enterNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *TickImp) _enterNode(node IBaseNode) {
	this._nodeCount++
	this._openNodes = append(this._openNodes, node)

	// TODO: call debug here
}

/**
 * Callback when opening a node (called by BaseNode).
 * @method _openNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *TickImp) _openNode(node *BaseNode) {
	// TODO: call debug here
}

/**
 * Callback when ticking a node (called by BaseNode).
 * @method _tickNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *TickImp) _tickNode(node *BaseNode) {
	// TODO: call debug here
	//fmt.Println("Tick _tickNode :", this.debug, " id:", node.GetID(), node.GetTitle())
}

/**
 * Callback when closing a node (called by BaseNode).
 * @method _closeNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *TickImp) _closeNode(node *BaseNode) {
	// TODO: call debug here

	ulen := len(this._openNodes)
	if ulen > 0 {
		this._openNodes = this._openNodes[:ulen-1]
	}

}

func (this *TickImp) pushSubtreeNode(node *SubTree) {
	this._openSubtreeNodes = append(this._openSubtreeNodes, node)
}
func (this *TickImp) popSubtreeNode() {
	ulen := len(this._openSubtreeNodes)
	if ulen > 0 {
		this._openSubtreeNodes = this._openSubtreeNodes[:ulen-1]
	}
}

/**
 * return top subtree node.
 * return nil when it is runing at major tree
 *
**/
func (this *TickImp) GetLastSubTree() *SubTree {
	ulen := len(this._openSubtreeNodes)
	if ulen > 0 {
		return this._openSubtreeNodes[ulen-1]
	}
	return nil
}

/**
 * Callback when exiting a node (called by BaseNode).
 * @method _exitNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *TickImp) _exitNode(node *BaseNode) {
	// TODO: call debug here
}

func (this *TickImp) GetTarget() interface{} {
	return this.target
}

func (this *TickImp) GetBlackboard() *Blackboard {
	return this.Blackboard
}

func (this *TickImp) _getOpenNodes() []IBaseNode {
	return this._openNodes
}

func (this *TickImp) _getNodeCount() int {
	return this._nodeCount
}

func (this *TickImp) SetTree(tree *BehaviorTree) {
	this.tree = tree
}

func (this *TickImp) SetDebug(debug interface{}) {
	this.debug = debug
}

func (this *TickImp) SetTarget(target interface{}) {
	this.target = target
}

func (this *TickImp) SetBlackBoard(bbd *Blackboard) {
	this.Blackboard = bbd
}
