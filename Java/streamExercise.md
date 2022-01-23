# Java stream 练习

```java
/**
 * 共有四个交易员
 * transactions 是交易记录，记录了交易员、交易时间、交易金额
 */
@BeforeEach
public void before() {
        Trader raoul = new Trader("Raoul", "Cambridge");
        Trader mario = new Trader("Mario", "Milan");
        Trader alan = new Trader("Alan", "Cambridge");
        Trader brian = new Trader("Brian", "Cambridge");

        transactions = Arrays.asList(
        new Transaction(brian, 2011, 300),
        new Transaction(raoul, 2012, 1000),
        new Transaction(raoul, 2011, 400),
        new Transaction(mario, 2012, 710),
        new Transaction(mario, 2012, 700),
        new Transaction(alan, 2012, 950)
        );
        }
```

### 练习一：找出2011年的记录并按照value大小排序
```java
    @Test
    public void test1() {
        transactions.stream()
                .filter((t) -> t.getYear() == 2011)
                .sorted((t1, t2) -> Integer.compare(t1.getValue(), t2.getValue()))
                .forEach(System.out::println);
    }
```

### 练习二：求出value的平均数
```java
    @Test
    public void test2() {
        Double average = transactions.stream()
                .collect(Collectors.averagingInt(Transaction::getValue));
        System.out.println(average);
    }
```

### 练习三：求出有多少城市
```java
    @Test
    public void test3() {
        transactions.stream().map(x -> x.getTrader().getCity()).distinct().forEach(System.out::println);
    }
```

### 练习四：找出是否有交易员是在 Milan 居住的
```java
    @Test
    public void test4() {
        boolean milan = transactions.stream().anyMatch(x -> x.getTrader().getCity().equals("Milan"));
        System.out.println(milan);
    }
```

### 练习五：所有交易中，最高的交易额是多少
```java
    @Test
    public void test5() {
        // maxBy 中不可用 Integer::max 因为这个函数返回的是两数中大的值并不是比较
        Optional<Integer> maxValue = transactions.stream().map(Transaction::getValue).collect(Collectors.maxBy(Integer::compare));
        System.out.println(maxValue.get());
    }
```

### 练习六：找出每个城市最大的订单
```java
    @Test
    public void Test6() {
        Map<String, Integer> collect = transactions.stream().collect(Collectors.toMap(x -> x.getTrader().getCity(), x -> x.getValue(), Integer::max));
        collect.forEach((x1, x2) -> System.out.println(x1 + ":" + x2.toString()));
    }
```
