/**
 * Leetcode - course_schedule
 */
package com.leetcode.course_schedule;
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
 // 用DFS来做, 与普通DFS的不同在于, 节点的状态不止有访问过和未访问过, 还多了一个正在访问
class Solution2 implements Solution {

    public static boolean canFinish(int numCourses, int[][] prerequisites) {
        
        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i< numCourses+1;i++) {
            graph.add(new ArrayList());
        }
        for (int i = 0 ;i< prerequisites.length;i++) {
            graph.get(prerequisites[i][0]).add(prerequisites[i][1]);
        }

        // 0 : not visit ; 1: visited ; 2: visiting
        int [] status = new int[numCourses];
        // System.out.println("---"+status.length);
        for (int i = 0; i< numCourses; i++) {
            status[i] = 0;
        }
        for (int i = 0; i< numCourses; i++) {
            boolean ret = dfs(graph, i, status);
            if (!ret) {
                return false;
            }
        }

        return true;
    }

    public static boolean dfs(List<List<Integer>> graph, int course, int [] status) {
        // System.out.println("---"+status.length);
        
        if(status[course] == 1) return true;
        if(status[course] == 2) return false;
        // System.out.println("+++++"+graph.get(course).size());
        // if (graph.get(course).size() == 0) {
        //     return true;
        // } 
        status[course] = 2;
        for (int i = 0; i< graph.get(course).size();i++) {
            // 在这里傻逼了 graph.get(course).get(i) !!!
            if(!dfs(graph, graph.get(course).get(i), status)) return false;
            
        }
        status[course] = 1;
        // for (int i = 0 ;i<status.length;i++) {
        //     System.out.print(status[i] + "\t");
        // }
        return true;
    }

    public static void main(String[] args) {
        System.out.println("test");
        
        int[][] arr1 = new int[][]{{0,1},{0,2},{1,2}};
        boolean ret = canFinish(3, arr1);
        System.out.println(ret);
    }

}
