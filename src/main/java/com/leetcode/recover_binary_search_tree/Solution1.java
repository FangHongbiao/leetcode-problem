/**
 * Leetcode - recover_binary_search_tree
 */
package com.leetcode.recover_binary_search_tree;
import java.util.*;
import com.ciaoshen.leetcode.util.*;

/** 
 * log instance is defined in Solution interface
 * this is how slf4j will work in this class:
 * =============================================
 *     if (log.isDebugEnabled()) {
 *         log.debug("a + b = {}", sum);
 *     }
 * =============================================
 */
// key insight: tree structure uncharted
class Solution1 implements Solution {
    class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }
    public void recoverTree(TreeNode root) {
        List<Integer> traverseList = new ArrayList<>();
        traverseTreeInOrder(root, traverseList);

        Collections.sort(traverseList);
    }
    
    private void traverseTreeInOrderForReset(TreeNode treeNode, List<Integer> list) {
        if(treeNode == null) {
            return;
        }
        
        traverseTreeInOrder(treeNode.left, list);
        TreeNode ret = list.remove(0);
        treeNode.val = ret;
        traverseTreeInOrder(treeNode.right, list);

    }

    private void traverseTreeInOrder(TreeNode treeNode, List<Integer> list) {
        if(treeNode == null) {
            return;
        }
        
        traverseTreeInOrder(treeNode.left, list);
        list.add(treeNode.val);
        traverseTreeInOrder(treeNode.right, list);

    }
}
