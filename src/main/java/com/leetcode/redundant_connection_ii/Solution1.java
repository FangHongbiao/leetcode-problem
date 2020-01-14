/**
 * Leetcode - redundant_connection_ii
 */
package com.leetcode.redundant_connection_ii;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.*;
import com.ciaoshen.leetcode.util.*;

/**
 * log instance is defined in Solution interface this is how slf4j will work in
 * this class: ============================================= if
 * (log.isDebugEnabled()) { log.debug("a + b = {}", sum); }
 * =============================================
 */
class UnionFindSet {
    private int[] parent;
    private int[] depth;

    public UnionFindSet(int num) {
        this.parent = new int[num];
        this.depth = new int[num];
        for (int i = 0; i < num; i++) {
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
            } else if (this.depth[p2] < this.depth[p1]) {
                this.parent[p2] = p1;
            } else {
                this.parent[p1] = p2;
                this.depth[p2]++;
            }
            return true;
        } else {
            return false;
        }
    }
}

class Solution1 implements Solution {
    // 1) Check whether there is a node having two parents.
    // If so, store them as candidates A and B, and set the second edge invalid.
    // 2) Perform normal union find.
    // If the tree is now valid
    // simply return candidate B
    // else if candidates not existing
    // we find a circle, return current edge;
    // else
    // remove candidate A instead of B.
    public int[] findRedundantDirectedConnection(int[][] edges) {
        int[] ret = new int[2];
        UnionFindSet unionFindSet = null;

        int[] count = new int[edges.length + 1];
        for (int i = 0; i < count.length; i++) {
            count[i] = 0;
        }
        for (int i = 0; i < edges.length; i++) {
            count[edges[i][1]]++;
        }

        int index = -1;
        for (int i = 0; i < edges.length+1; i++) {
            if (count[i] > 1) {
                index = i;
            }
        }
        if (index == -1) {
            unionFindSet = new UnionFindSet(edges.length);
            for (int i = 0; i < edges.length; i++) {

                int t1 = unionFindSet.find(edges[i][0] - 1);
                int t2 = unionFindSet.find(edges[i][1] - 1);
                if (t1 == t2) {
                    ret = edges[i];
                } else {
                    unionFindSet.merge(t1, t2);
                }
            }
        } else {
            unionFindSet = new UnionFindSet(edges.length);
            int [] candiate = new int[10];
            int tmp = 0;
            for (int i = 0; i < edges.length; i++) {
                if (edges[i][1] == index) {
                    candiate[tmp] = i;
                    tmp++;
                }
            }
            boolean tag = true;
            for (int i = 0; i < edges.length; i++) {
                if (candiate[1] == i) {
                    continue;
                }
                int t1 = unionFindSet.find(edges[i][0] - 1);
                int t2 = unionFindSet.find(edges[i][1] - 1);
                if (t1 == t2) {
                    tag = false;
                } else {
                    unionFindSet.merge(t1, t2);
                }
            }
            if (tag == true) {
                ret = edges[candiate[1]];
            }else {
                ret = edges[candiate[0]];
            }

        }

        return ret;
    }
}

public class MainClass {
    public static int[] stringToIntegerArray(String input) {
        input = input.trim();
        input = input.substring(1, input.length() - 1);
        if (input.length() == 0) {
          return new int[0];
        }
    
        String[] parts = input.split(",");
        int[] output = new int[parts.length];
        for(int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }
    
    public static int[][] stringToInt2dArray(String input) {
        JsonArray jsonArray = JsonArray.readFrom(input);
        if (jsonArray.size() == 0) {
          return new int[0][0];
        }
    
        int[][] arr = new int[jsonArray.size()][];
        for (int i = 0; i < arr.length; i++) {
          JsonArray cols = jsonArray.get(i).asArray();
          arr[i] = stringToIntegerArray(cols.toString());
        }
        return arr;
    }
    
    public static String integerArrayToString(int[] nums, int length) {
        if (length == 0) {
            return "[]";
        }
    
        String result = "";
        for(int index = 0; index < length; index++) {
            int number = nums[index];
            result += Integer.toString(number) + ", ";
        }
        return "[" + result.substring(0, result.length() - 2) + "]";
    }
    
    public static String integerArrayToString(int[] nums) {
        return integerArrayToString(nums, nums.length);
    }
    
    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            int[][] edges = stringToInt2dArray(line);
            
            int[] ret = new Solution1().findRedundantDirectedConnection(edges);
            
            String out = integerArrayToString(ret);
            
            System.out.print(out);
        }
    }
}