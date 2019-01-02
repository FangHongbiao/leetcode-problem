/**
 * Leetcode - verifying_an_alien_dictionary
 */
package com.leetcode.verifying_an_alien_dictionary;
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

    public boolean isAlienSorted(String[] words, String order) {
        Map<Character, Character> map = new HashMap<>();
        for (int i = 0; i< order.length();i++) {
            map.put( order.charAt(i), (char)('a' + i));
        }
        List<String> list = new ArrayList<>();
        for (int i = 0; i< words.length; ++i) {
            
            StringBuffer sb = new StringBuffer();
            for(int j = 0; j<words[i].length(); j++) {
                sb.append(map.get(words[i].charAt(j)));   
            }
            list.add(sb.toString());
        }
        
        List<String> list1= new ArrayList<String>();
        list1.addAll(list);
        Collections.sort(list);
       
        for (int i = 0; i< list.size(); ++i) {
            if(!list.get(i).equals(list1.get(i))){
                return false;
            }
        }
        return true;
    }
    public static void main(String[] args) {
        Solution1 s = new Solution1();

        boolean ret = s.isAlienSorted(new String [] {"ubg","kwh"}, "qcipyamwvdjtesbghlorufnkzx");
        System.out.println(ret);
    }
}
