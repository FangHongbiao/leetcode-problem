/**
 * Leetcode - merge_k_sorted_lists
 */
package com.leetcode.merge_k_sorted_lists;

import java.util.*;
import com.ciaoshen.leetcode.util.*;

/**
 * log instance is defined in Solution interface this is how slf4j will work in
 * this class: ============================================= if
 * (log.isDebugEnabled()) { log.debug("a + b = {}", sum); }
 * =============================================
 */
class Solution1 implements Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        PriorityQueue<ListNode> heap = new PriorityQueue<>(new Comparator<ListNode>() {
            @Override
            public int compare(ListNode o1, ListNode o2) {                
                return o1.val - o2.val;
            }
        });

        for (int i = 0; i < lists.length; i++) {
            if (lists[i] != null) {
                heap.add(lists[i]);
            }
        }

        List<ListNode> orderedNodes = new ArrayList<>();
        while (heap.size()!=0) {
            ListNode top = heap.poll() ;
            orderedNodes.add(top);
            if (top.next != null) {
                heap.add(top.next);
            }
        }

        for (int i = 0; i< orderedNodes.size();i++) {
            if (i!=orderedNodes.size()-1) {
                orderedNodes.get(i).next = orderedNodes.get(i+1);
            }else{
                orderedNodes.get(i).next = null;
            }
        }
        if (orderedNodes.size() == 0) {
            return null;
        }
        return orderedNodes.get(0);
    }

    public static void main(String[] args) {
        System.out.println("test");
    }
}
