/**
 * Leetcode - binary_tree_maximum_path_sum
 */
package com.leetcode.binary_tree_maximum_path_sum;
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
    int maxValue = Integer.MIN_VALUE;
    public int maxPathSum(TreeNode root) {
        pathSum(root);
        return maxValue;
    }

    public int pathSum(TreeNode root) {
        
        if (root == null) {
            return 0;
        }
        int left = Math.max(0,pathSum(root.left));
        int right = Math.max(0,pathSum(root.right));
        // 统计最大值时, 把当前节点算进去, 即考虑, 以每一个节点为根的时候,获得的最大路径
        // 但是在作为子树为上层节点提供服务时, 只能去左树或者右树, 相对上层节点,在当前节点就不可以分叉了
        maxValue = Math.max(maxValue, left + right + root.val);
        return Math.max(left, right) + root.val;
    }


    public static void main(String[] args) {
        System.out.println("exit");
    }

}
