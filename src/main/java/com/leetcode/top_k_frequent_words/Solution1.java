/**
 * Leetcode - top_k_frequent_words
 */
package com.leetcode.top_k_frequent_words;

import java.util.*;
import com.ciaoshen.leetcode.util.*;

/**
 * log instance is defined in Solution interface this is how slf4j will work in
 * this class: ============================================= if
 * (log.isDebugEnabled()) { log.debug("a + b = {}", sum); }
 * =============================================
 */
class Solution1 implements Solution {

    public List<String> topKFrequent(String[] words, int k) {
        // List<String> list = Arrays.asList(word);
        // Collections.sort(list);
        // List<Integer> result = new ArrayList<>();
        // result.add(0);
        // int min = Integer.MAX_VALUE;
        // for (int i = 1; i<list.size();j++) {
        // if (!list.get(i).equals(list.get(i-1))){
        // result.add(i);
        // min = Math.min(min, result.get(result.size()-1)
        // -result.get(result.size()-2));
        // }
        // }
        // result.add(list.size());
        // min = Math.min(min, result.get(result.size()-1)
        // -result.get(result.size()-2));

        // for (int i = 1; i< result.size();i++) {
        // if (result.get(i) - result.get(i-1) == min) {

        // }
        // };
        List<String> list = new ArrayList<>();
        Map<String, Integer> map = new LinkedHashMap<>();
        for (int i = 0; i < words.length; i++) {
            if (map.containsKey(words[i])) {
                map.put(words[i], map.get(words[i]) + 1);
            } else {
                map.put(words[i], 1);
            }
        }

        // 升序比较器
        Comparator<Map.Entry<String, Integer>> valueComparator = new Comparator<Map.Entry<String, Integer>>() {
            @Override
            public int compare(Map.Entry<String, Integer> o1, Map.Entry<String, Integer> o2) {
                if (o2.getValue() - o1.getValue() != 0){

                    return o2.getValue() - o1.getValue();
                }else {
                    return o1.getKey().compareTo(o2.getKey());
                }
            }
        };

        // map转换成list进行排序
        List<Map.Entry<String, Integer>> sortedMap = new ArrayList<Map.Entry<String, Integer>>(map.entrySet());

        // 排序
        Collections.sort(sortedMap, valueComparator);

        for (int i = 0; i< k ; i++) {
            list.add(sortedMap.get(i).getKey());
        }
        return list;
    }

}
