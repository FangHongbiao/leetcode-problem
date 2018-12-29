/**
 * Leetcode - array_of_doubled_pairs
 */
package com.leetcode.array_of_doubled_pairs;
import java.util.*;
import com.ciaoshen.leetcode.util.*;
import com.sun.org.apache.regexp.internal.recompile;

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

    public boolean canReorderDoubled(int[] A) {
        List<Integer> pos = new ArrayList<>();
        List<Integer> neg = new ArrayList<>();
        int count = 0;
        for(int i =0 ; i<A.length;i++) {
           if(A[i]<0) {
                neg.add(A[i]);
            }else if(A[i]>0){
                pos.add(A[i]);
            }else {
                count++;
            }
        }
        if (count %2 !=0 || pos.size() %2 !=0 || neg.size() %2 !=0) {
            return false;
        }

        Collections.sort(pos);
        Collections.sort(neg);

        for (int i = 1; i< pos.size(); i+=2) {
            if (pos.get(i) != 2*pos.get(i-1)) {
                return false;
            }
        }

        for (int i = 1; i< neg.size(); i+=2) {
            if (2*neg.get(i) != neg.get(i-1)) {
                return false;
            }
        }
        return true;
    }
}
