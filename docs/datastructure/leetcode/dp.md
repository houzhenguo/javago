

## 1. 矩阵 路径问题 leetcode 62

https://leetcode.com/problems/unique-paths/

```java
class Solution {
    public int uniquePaths(int m, int n) {
        if(m<=1 && n<=1) {
            return 1;
        }
        int[][] temp = new int[m][n];
        for (int i=1;i<m;i++) {
            temp[i][0] = 1;
        }
        for (int j=1;j<n;j++) {
            temp[0][j] = 1;
        }
        for  (int i=1;i<m;i++) {
            for (int j =1;j<n;j++) {
                temp[i][j] = temp[i-1][j]+temp[i][j-1];
            }
        }
        return temp[m-1][n-1];
    }
}
```
## 思路：

这是一个动态规划问题，动态规划的思路就是通过一系列的子问题来实现最终问题的求解。那么具体到这道题，可以这么理解，在这个矩形网格框内，第一行和第一列中的每一位置，到达的可能路径都为1。对其他位置，到达的可能路径数量为其正上面位置对应路径的数量加上左边路径的数量（因为只可以向下走或者向右走）。如下表可以看出这一点。


0 |	1 |	1 |	1 |	1 |	1 |	1
- |	- |	- |	- |	- |	- |	-
1 |	2 |	3 |	4 |	5 |	6 |	7
1 | 3 |	6 |	10 | 15 |	21 |	28	


本题目是动态规划的经典题目，一定要理解，可以熟练写出代码。


其他更优化版本:

```java

public int uniquePaths(int m, int n) {
    int[] dp = new int[n];
    Arrays.fill(dp, 1);
    for (int i = 1; i < m; i++) {
        for (int j = 1; j < n; j++) {
            dp[j] = dp[j] + dp[j - 1];
        }
    }
    return dp[n - 1];
}
```

## 爬楼梯问题

[70. Climbing Stairs (Easy)](https://leetcode.com/problems/climbing-stairs/description/)

```java
class Solution {
    public int climbStairs(int n) {
        if (n == 1) {
            return 1;
        }
        if (n <1) {
            return 0;
        }
        int[] dp = new int[n];
        dp[0] = 1; // 第一个台阶有一种
        dp[1] = 2; // 第二个台阶有两种
        for (int i=2;i<n;i++) {
            dp[i] = dp[i-1] +dp[i-2];
        }
        return dp[n-1];
    }
}
```

## 思路

加入我们一共有10个台阶，最后一步到达第10个台阶的有两种方式，
1. 从第8个台阶迈两步到第10个台阶
2. 从第9个台阶迈一步到第10个台阶

F(n) = F(n-1) + F(n-2);所以到达第10个台阶的方式 就是到达第8个台阶的次数与到达第9个台阶次数之和。