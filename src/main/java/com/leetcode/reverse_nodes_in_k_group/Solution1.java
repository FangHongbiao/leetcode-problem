/**
 * Leetcode - reverse_nodes_in_k_group
 */
package com.leetcode.reverse_nodes_in_k_group;

import java.util.*;
import com.ciaoshen.leetcode.util.*;
import com.leetcode.merge_k_sorted_lists.Solution.ListNode;

/**
 * log instance is defined in Solution interface this is how slf4j will work in
 * this class: ============================================= if
 * (log.isDebugEnabled()) { log.debug("a + b = {}", sum); }
 * =============================================
 */
public class Solution1 implements Solution {
    class RetData {
        public ListNode reverseHead;
        public ListNode reverseTail;
        public int length;
        public ListNode newHead;
    }

    public RetData reverseFisrtK(ListNode head, int k) {

        RetData retData = new RetData();
        retData.reverseTail = head;
        int count = 0;
        while (count < k && head != null) {
            count++;

            if (count == k) {
                retData.newHead = head.next;
            }
            ListNode tmp = retData.reverseHead;
            retData.reverseHead = head;
            head = head.next;
            retData.reverseHead.next = tmp;
        }

        retData.length = count;
        return retData;
    }

    public ListNode reverseKGroup(ListNode head, int k) {
        RetData retData = reverseFisrtK(head, k);
        ListNode ret = retData.reverseHead;
        ListNode  lastNode = retData.reverseTail;
        head = retData.newHead;
        int length = retData.length;
        if (retData.newHead == null && retData.length < k) {
            retData = reverseFisrtK(retData.reverseHead, k);
        }
        while(head != null) {
            retData = reverseFisrtK(head, k);
            if (retData.newHead == null && retData.length < k) {
                retData = reverseFisrtK(retData.reverseHead, k);
                ret = retData.reverseHead;
            }
            lastNode.next = retData.reverseHead;
            lastNode = retData.reverseTail;
            length = retData.length;
            head = retData.newHead;
        }
        return ret;
    //     ListNode firstNode = head;
    //     ListNode start = new ListNode(10);
    //     ListNode lastNode = start;
    //     int i = 0;
    //     while (firstNode.next != null) {
    //         i++;
    //         if (i % k == 1) {
    //             lastNode = firstNode;
    //         }
    //         if (i == k) {
    //             ret = firstNode;
    //             firstNode = firstNode.next;
    //             continue;
    //         }
    //         ListNode tmp = firstNode;
    //         firstNode = head.next;
    //         lastNode.next = firstNode;
    //         firstNode.next = tmp;
    //     }
    //     ListNode tmp = lastNode;
    //     ListNode remainNode = lastNode.next;
    //     while (tmp.next != null) {
    //         remainNode = tmp.next;
    //         tmp = tmp.next;

    //     }
    //     return start.next;
    }
}
