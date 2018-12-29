/**
 * Leetcode - remove_duplicates_from_sorted_array
 */
package com.leetcode.remove_duplicates_from_sorted_array;

import java.util.*;
import com.ciaoshen.leetcode.util.*;

/**
 * log instance is defined in Solution interface this is how slf4j will work in
 * this class: ============================================= if
 * (log.isDebugEnabled()) { log.debug("a + b = {}", sum); }
 * =============================================
 */
class Solution1 implements Solution {
    public int removeDuplicates(int[] nums) {
        // if (nums.length == 0) {
        // return 0;
        // }
        // int start = nums[0];
        // int end = nums[nums.length-1];
        // // nums = new int[end-start+1];
        // for (int i = 0; i<end-start+1; i++) {
        // nums[i] = i+start;
        // }
        // return end-start+1;

        Set<Integer> set = new LinkedHashSet();

        for (int i = 0; i < nums.length; i++) {
            set.add(nums[i]);
        }
        int count = 0;
        for (Integer i : set) {
            nums[count++] = i;
        }
        return set.size();
    }

    public static void main(String[] args) {
        System.out.println("test");
    }
}
