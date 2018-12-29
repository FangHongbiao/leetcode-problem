/**
 * Leetcode - course_schedule
 */
package com.leetcode.course_schedule;

import java.util.*;
import com.ciaoshen.leetcode.util.*;

/**
 * log instance is defined in Solution interface this is how slf4j will work in
 * this class: ============================================= if
 * (log.isDebugEnabled()) { log.debug("a + b = {}", sum); }
 * =============================================
 */
class Solution1 implements Solution {
    private List<List<Integer>> adjacentVertexes;
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        List<List<Integer>> adjacentVertexes = new ArrayList<>();
        
        for (int i = 0; i< numCourses; i++) {
            adjacentVertexes.set(i, new ArrayList<>());
        }

        for (int i = 0; i<prerequisites.length; i++) {
            adjacentVertexes.get(prerequisites[i][0]).add(prerequisites[i][1]);
        }

        this.adjacentVertexes = adjacentVertexes;
        Deque<Integer> stack = new ArrayDeque<>();
        Set<Integer> visited = new HashSet<>();
        for (int i = 0; i< numCourses;i++) {
            if (visited.contains(i)) {
                continue;
            }
            topSortUtil(i,stack,visited);
        }
        return stack.size() == numCourses;
    }

    private void topSortUtil(int courseId, Deque<Integer> stack, Set<Integer> visited) {
        visited.add(vertex);
        for (Integer childVertex : this.adjacentVertexes.get(i)) {
            if (visited.contains(childVertex)) {
                continue;
            }
            topSortUtil(childVertex, stack, visited);
        }
        stack.offerFirst(vertex);

    }

}
