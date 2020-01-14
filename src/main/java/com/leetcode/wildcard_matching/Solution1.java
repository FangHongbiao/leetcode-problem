/**
 * Leetcode - wildcard_matching
 */
package com.leetcode.wildcard_matching;
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

    public boolean isMatch(String s, String p) {
        return isMatch(s, p, 0, 0, s.length(), p.length());
    }

    public boolean isMatch(String s, String p, int si, int pi, int sn, int pn) {
        if (si == sn) {
            if (pi == pn) {
                return true;
            }
            while (pi < pn) {
                if (p.charAt(pi) != '*') {
                    return false;
                }
                pi ++;
            }
            return true;
        }

        if (pi == pn && si < sn) {
            return false;
        }

        if (p.charAt(pi) != '*') {
            if (s.charAt(si) == p.charAt(pi) || s.charAt(si) == '?') {
                return isMatch(s, p, si+1, pi+1, sn, pn);
            } else {
                return false;
            }

        }

        while (si <= sn) {
            if (isMatch(s, p, si, pi+1, sn, pn)) {
                return true;
            }
            si+=1;
        }
        return false;
    }

    public static void main(String[] args) {
        Solution1 solution = new Solution1();
        boolean ab = solution.isMatch("ab", "?*");
        System.out.println(ab);

    }

}
