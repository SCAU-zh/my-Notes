# java 常见日期格式转换
> 总结一下日常使用的几种日期时间格式的互相转换 ：LocalDate、LocalDateTime、Timestamp、Date
## Date
- 转 LocalDate
```java
    Date date = new Date();
    Instant instant = date.toInstant();
    ZoneId zoneId = ZoneId.systemDefault();

    // atZone()方法返回在指定时区从此Instant生成的ZonedDateTime。
    LocalDate localDate = instant.atZone(zoneId).toLocalDate();
    System.out.println("Date = " + date);
    System.out.println("LocalDate = " + localDate);

```
- 转 LocalDateTime
```java
    Date date = new Date();
    Instant instant = date.toInstant();
    ZoneId zoneId = ZoneId.systemDefault();

    LocalDateTime localDateTime = instant.atZone(zoneId).toLocalDateTime();
    System.out.println("Date = " + date);
    System.out.println("LocalDateTime = " + localDateTime);
    // 使用工厂方法
    LocalDateTime localDateTime = LocalDateTime.ofInstant(date.toInstant(), zoneId);
```
- 转 Timestamp
```java
    Timestamp timestamp = Timestamp.from(new Date().toInstant());

    long timestamp = new Date().toInstant().toEpochMilli();
```

## LocalDateTime
- 转 Date
```java
    Date.from(LocalDateTime.now().toInstant(ZoneOffset.UTC));
```

- 转 LocalDate
```java
    LocalDateTime.now().toLocalDate()
```

- 转 Timestamp
```java
    Timestamp timestamp = Timestamp.valueOf(LocalDateTime.now());

    LocalDateTime.now().toInstant(ZoneOffset.UTC).toEpochMilli();
```

## LocalDate
- 转 LocalDate
```java
    LocalDate localDate = LocalDate.now();
    localDate.atStartOfDay();
```
- 转 Date
```java
    Date.from(LocalDate.now().atStartOfDay().toInstant(ZoneOffset.UTC));
```
- 转 Timestamp
```java
    Timestamp.from(localDate.atStartOfDay().toInstant(ZoneOffset.UTC));
```

## Timestamp
- 转 LocalDateTime
```java
    timestamp.toLocalDateTime();
```
- 转 Date
```java
    Date.from(timestamp.toInstant());
```
- 转LocalDate
```java
    timestamp.toLocalDateTime().toLocalDate();
```
