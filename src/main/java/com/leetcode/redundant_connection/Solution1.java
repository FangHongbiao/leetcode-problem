/**
 * Leetcode - redundant_connection
 */
package com.leetcode.redundant_connection;
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

class UnionFindSet {
    private int [] parent;
    private int [] depth;
    public UnionFindSet(int num) {
        this.parent = new int[num];
        this.depth = new int[num];
        for (int i = 0; i< num; i++) {
            this.parent[i] = i;
            this.depth[i] = 0;
        }
    }

    public int find(int target) {
        if (target != this.parent[target]) {
            this.parent[target] = find(this.parent[target]);
        }
        return parent[target];
    }

    public boolean merge(int t1, int t2) {
        int p1 = this.find(t1);
        int p2 = this.find(t2);
        if (p1 != p2) {
            if (this.depth[p2] > this.depth[p1]) {
                this.parent[p1] = p2;
            }else if (this.depth[p2] < this.depth[p1]){
                this.parent[p2] = p1;
            } else {
                this.parent[p1] = p2;
                this.depth[p2]++;
            }
            return true;
        }else {
            return false;
        }
    }
}

class Solution1 implements Solution {
    public int[] findRedundantConnection(int[][] edges) {
        int[] ret = new int[2];
        UnionFindSet unionFindSet = new UnionFindSet(edges.length+1);

        for(int i = 0 ; i< edges.length; i++) {

                int t1 = unionFindSet.find(edges[i][0]-1);
                int t2 = unionFindSet.find(edges[i][1]-1);
                if (t1 == t2) {
                    ret = edges[i];
                }else{
                    unionFindSet.merge(t1, t2);
                }   
        }
        return ret;
    }
}