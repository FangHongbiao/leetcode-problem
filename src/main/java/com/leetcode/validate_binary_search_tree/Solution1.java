/**
 * Leetcode - Validate_Binary_Search_Tree
 */
package com.leetcode.validate_binary_search_tree;
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
class Solution1 implements Solution {

    public boolean isValidBST(TreeNode root) {
        List<Integer> traverseList = new ArrayList<>();
        traverseTreeInOrder(root, traverseList);
        for(int i = 1; i< traverseList.size();i++) {
            // 必须严格递增
            if (traverseList.get(i-1) >= traverseList.get(i)) {
                return false;
            }
        }
        return true;
    }

    private void traverseTreeInOrder(TreeNode treeNode, List<Integer> list) {
        if(treeNode == null) {
            return;
        }
        
        traverseTreeInOrder(treeNode.left, list);
        list.add(treeNode.val);
        traverseTreeInOrder(treeNode.right, list);

    }

    public static void main(String[] args) {
        System.out.println("ste");

}
