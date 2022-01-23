## Java Lambda 表达式使用例子

### 1、List 映射为 map
```java
public class BillingLogs {
    // 团队 ID
    Integer teamId;
    // 扣费时间
    LocalTimeDate billDate; 
}
// ———————————————————————————————————————————
// 筛选出每个 team 中最新的扣费日志 
// toMap 的前两个参数是 k 和 v，最后一个是 k 重复时的选择策略
Map<Integer, LocalTimeDate> lastLogs = billingLogs
        .stream()
        .collect(Collectors.toMap(LostCustomerUpdateDTO::getTeamId, LostCustomerUpdateDTO::getBillingDate, (l1, l2) -> l1.compareTo(l2) > 0 ? l1 : l2));
```
