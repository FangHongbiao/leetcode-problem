/**
 * Leetcode - average_of_levels_in_binary_tree
 */
package com.leetcode.average_of_levels_in_binary_tree;
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
    
    public List<Double> averageOfLevels(TreeNode root) {
        
        List<List<TreeNode>> list = new ArrayList();
        
        traverseTree(root, 0, list);
        List<Double> results = new ArrayList<>();
        for (int i = 0; i<list.size();i++) {
            double sum = 0.0;
            List<TreeNode> treeNodes = list.get(i);
            for(TreeNode treeNode : treeNodes) {
                sum += treeNode.val;
            }
            results.add(sum / treeNodes.size());
        }


        return results;
    }

    private void traverseTree(TreeNode treeNode, int depth, List<List<TreeNode>> list) {
        if(treeNode == null) {
            return;
        }
        if(list.size() <= depth) {
            list.add(new ArrayList<TreeNode>());
        }
        list.get(depth).add(treeNode);
        traverseTree(treeNode.left, depth+1, list);
        traverseTree(treeNode.right, depth+1, list);

    }

    public static void main(String[] args) {
        System.out.println("test");
    }

}
