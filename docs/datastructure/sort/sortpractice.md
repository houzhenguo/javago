

sort 练习

```java
// sort swap less
public abstract class Sort<T extends Comparable<T>{
    public abstract void sort(T[] nums);
    // 交换两个值的位置
    protected void swap(T[] nums, int i, int j)
    {
        int temp = nums[i];
        nums[i]  = nums[j];
        nums[j]  = temp;
    }
    // t1 < t2 return true
    protected boolean less(T t1, T t2)
    {
        return t1.compareTo(t2) < 0;
    }
}

```



1. 选择排序 选择最小的放在当前 子数组的第一个位置

```java
public class SelectSort<T extends Comparable<T>> extends Sort<T>
{
    @Override
    public void sort(T[] nums)
    {

        for(int i = 0; i< nums.length-1;i++)
        {
            int minIndex = i;
            for(int j = i+1;j<nums.length;j++)
            {
                if(less(nums[i], nums[j]))
                {
                    minIndex = j;
                }
            }
            swap(nums, i, minIndex);
        }
    }
}
```

2. 冒泡排序 将大的浮到左侧
```java
public void sort(T[] nums)
{
    for(int i = nums.length-1;i>0 && !flag;i--)
    {
        boolean flag = false;
        for(int j=0;j<i;j++)
        {
            if(nums[j]>nums[j+1])
            {
                swap();
                flag = true;
            }
        }
    }
}
```

3. Insert Sort 每次交换相邻的位置

```java
    for(int i = 1;i<nums.length;i++)
    {
        for(int j = i;j>0&& nums[j]<nums[j-1];j--)
        {
            swap();
        }
    }
```

4. 交换两个有序数组

原地归并的抽象方法

```java
public static void merge(T[] a, int lo, int mid, int hi)
{
    int i = lo;
    int j = mid+1;
    for(int k = lo;k<hi;k++)
    {
        aux[k] = a[k];
    }
    for(int k = lo;k<hi;k++)
    {
        if(i> mid) aux[k] =a[j++];
        if(j > hi) aux[k] = a[i++];
        if(aux[j]>aux[i]) aux[k] = a[i++];
        else              aux[k] = a[j++]; 
    }
}
```