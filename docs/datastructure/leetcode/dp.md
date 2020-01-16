

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

## 2. 爬楼梯问题

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

## 3.  强盗抢劫问题

[198. House Robber (Easy)](https://leetcode.com/problems/house-robber/description/)

题目描述：抢劫一排住户，但是不能抢邻近的住户，求最大抢劫量。

定义 dp 数组用来存储最大的抢劫量，其中 dp[i] 表示抢到第 i 个住户时的最大抢劫量。

由于不能抢劫邻近住户，如果抢劫了第 i -1 个住户，那么就不能再抢劫第 i 个住户.注意存在这种问题[2,1,1,2] 这个时候抢到的最大值为 4；

Leetcode最优易懂方法:

```java
public int rob(int[] nums) {
    int [][] dp = new int[num.length +1][2];
    for (int i=1; i<= num.length;i++) {
        dp[i][0] = Math.max(dp[i-1][0],dp[i-1][1]);
        dp[i][1] = num[i-1] + dp[i-1][0];
    }
    return Math.max(dp[nums.length][0], dp[num.length][1]);
}

// dp[i][1] means we rob the current house and dp[i][0] means we don't,

//这个解决方案存在争议,
//https://leetcode.com/problems/house-robber/discuss/55681/Java-O(n)-solution-space-O(1)

// 这是我自己写的，存在的问题是 [2,1,1,2]的问题
class Solution {
    public int rob(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        dp[1] = nums[1];
        if (nums.length == 2) {
            return Math.max(dp[0],dp[1]);
        }
        for (int i=2;i < nums.length;i++) {
            dp[i] = Math.max(dp[i-2] + nums[i], dp[i-1]);
        }
        return dp[dp.length -1];
    }
}


    // 20200106 版本 没有 2 1 1 2的问题
    public static int rob(int[] nums) {
        int n = nums.length;
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return nums[0];
        }
        int[] dp = new int[n];
        dp[0] = nums[0];
        dp[1] = Math.max(nums[0],nums[1]);
        for (int i = 2; i< n;i++) {
            // 偷 与 不偷的问题
            dp[i] = Math.max(dp[i-2] +nums[i], dp[i-1]);
        }
        return dp[n-1];
    }
```

## 4. 斐波那契数列(Easy)

大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）。
n<=39 ; 1、1、2、3、5、8、13、21、34

分析：

F(n) = F(n-1) + F(n-2); if n==1 || 2  ->  1

```java
    // method1: 递归实现
    public static int fib(int n) {
        if (n == 1 || n == 2) {
            return 1;
        }
        return fib(n-1) + fib(n-2);
    }
    // method2: DP
    public static int fibDp(int n) {
        int[] dp = new int[n]; // DP的核心思想就是这个数组缓存结果
        if (n == 1 || n == 2) {
            return 1;
        }
        dp[0] = 1;
        dp[1] = 1;
        for (int i=2;i<n;i++) {
            dp[i] = dp[i-1] + dp[i-2];
        }
        return dp[n-1];
    }
```
## 5. 跳台阶(Easy)

一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

分析：

假如一共有 10 阶，到达第 10阶的方式有两个， a. 从第9阶上去 b. 从第8阶上去,而选择哪一种，又取决与到达 第9 阶和到达第8 阶的方式
F(10) = F(9) + F(8)
F(9) = F(8) + F(7); // 这里面就开始存在重复计算的问题了

```java
    // DP 上楼梯
    public static int stairs(int n) {
        if (n == 1 || n == 2) {
            return n;
        }
        int pre1 = 1; // 代表前 1 阶
        int pre2 = 2; // 代表前 2 阶
        int pre = 0;  // 当前
        for (int i=3;i<=n;i++) {
            pre = pre1 + pre2; // 当前的等于前两种的和
            pre2 = pre1; // 先将 pre1 -> pre2 避免数据丢失
            pre1 = pre; // 将 pre -> pre1
        }
        return pre;
    }
```

## 6. 变态跳台阶(mid)
一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级... 它也可以跳上 n 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

分析 ： 

假如有 10 阶，则第10 阶的方式有 从第1 (fill with 1)跳上去 ，从 2. ...9 阶跳上去
F(10) = F(9) + F(8) + ... + F(1);
```java
    // 复杂版跳台阶，不限制一次跳几个
    public static int stairs(int n) {
        int[] dp = new int[n]; // 创建 dp数组
        Arrays.fill(dp,1); // 全部填充1，可以直接跳上去，不管前面的，所以默认填充1(重要)
        for (int i= 0; i<n; i++) {
            for (int j=0;j<i;j++) { // 处理一下前面的
                dp[i] += dp[j]; // 把前面所有的次数累加起来就是当前的（注意dp[i]已经有一个值为1）
            }
        }
        return dp[n-1];// 返回
    }
```


## 7. 最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

```java
// DP 解法
 public int maxSubArray(int[] nums) {
        int ans = nums[0];
        int sum = 0;
        for(int num: nums) {
            if(sum > 0) { // 前面累加的增长 > 0 那就加上我自己再看看
                sum += num;
            } else {  // 前面累加的增长 < 0  前面的所有累加的增长都白干了，那就从这一次开始看后面的吧
                sum = num;
            }
            ans = Math.max(ans, sum);
        }
        return ans;  // sum就等于从前面某一天到今天的增长
    }
}
```
