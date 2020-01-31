
```java
// 1. 任意一个数组中重复的数字
//  这种方式不太对，需要中间加一层 while 
    public boolean duplicate(int numbers[],int length,int [] duplication) {
        if (length <= 0) {
            return false;
        }
        for (int i = 0; i < length; i++) {
            int temp = numbers[i];
            if (temp == i) {
                continue;
            }
            int currIndexNum = numbers[temp];
            if (currIndexNum != temp) {
                swap(numbers, currIndexNum, temp);
            }else {
                duplication[0] = currIndexNum;
                return true;
            }
        }
        return false;
    }
    public void swap(int[] array, int i, int j) {
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp; 
    }

    public boolean duplicate(int numbers[],int length,int [] duplication) {
        if (length <=0 ) {
            return false;
        }
        for (int i = 0; i < length ; i++) {
            while (i != numbers[i]) {
                if (numbers[numbers[i]] != numbers[i]) {
                    swap(numbers, i, numbers[i]);
                }else {
                    duplication[0] = numbers[i];
                    return true;
                }
                
            }
            
        }
        return false;
    }

    // 2. 二维数组找数
    public boolean Find(int target, int [][] array) {
        if (array == null || array.length == 0 || array[0].length == 0) {
            return false;
        }
        int rows = array.length;
        int cols = array[0].length;
        int r = 0;
        int c = cols - 1;
        while (r < rows && c >=0) { // 注意这个地方
            if (target == array[r][c]) {
                return true;
            }
            if (target > array[r][c]) {
                r++;
            }else {
                c--;
            }
        }
        return false;
    }    

    // 3. 
    // We Are Happy. -> We%20Are%20Happy
    public String replaceSpace(StringBuffer str) {
    	int p1 = str.length() - 1;
        for (int i = 0;i <= p1; i++) {   // 注意等于号
            if (str.charAt(i) == ' ') {
                str.append("  ");
            }
        }
        int p2 = str.length() - 1;
        while (p1 >= 0 && p2 > p1) {
            char c = str.charAt(p1--);
            if (c == ' ') {
                str.setCharAt(p2--,'0');
                str.setCharAt(p2--,'2');
                str.setCharAt(p2--,'%');
            }else {
                str.setCharAt(p2--,c);
            }
        }
        return str.toString();
    }

    //4. // 从尾到头打印链表 

    public ArrayList<Integer> printListFromTailToHead(ListNode listNode) {
        ArrayList<Integer> res = new ArrayList<>();
        if (listNode == null) {
            return res;
        }
        while (listNode != null) {
            res.add(listNode.val);
            listNode = listNode.next;
        }
        Collections.reverse(res);
        return res;
    }

    // 5. 重建二叉树 
    // 前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}
import java.util.*;
public class Solution {
    public Map<Integer, Integer> map = new HashMap<>(); // key = num, val = index in order
    public TreeNode reConstructBinaryTree(int [] pre,int [] in) {
        int preN = pre.length;
        int inN = in.length;
        if (preN == 0 || inN == 0) {
            return null;
        }
        for (int i = 0; i < inN;i++) {
            map.put(in[i],i);
        }
        return reBuildTree(pre,0, preN-1, 0);
    }

    public TreeNode reBuildTree(int[] pre, int preL, int preR, int inL) {
        if (preL > preR) {
            return null;
        }
        TreeNode node = new TreeNode(pre[preL]);
        int inIndex = map.get(pre[preL]);
        int leftSize = inIndex - inL;
        node.left = reBuildTree(pre, preL+1, preL+ leftSize,inL);
        node.right = reBuildTree(pre, preL+leftSize+1,preR,inL+leftSize+1);
        return node;
    }
}

// 6. 二叉树的下一个结点
    // 给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回

    // 中序遍历 左 根 右
    public TreeLinkNode GetNext(TreeLinkNode pNode)
    {
        
    }

```