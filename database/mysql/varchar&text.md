> MySQL 5.0.3版的一项更改包括将VARCHAR字段的最大长度从255个字符增加到65,535个字符。这使得VARCHAR类型比以往任何时候都更类似于TEXT。

## VARCHAR 和 TEXT 的区别
- VARCHAR中的VAR表示您可以将最大大小设置为1到65535之间的任何值。TEXT字段的最大固定大小为65535个字符
- VARCHAR可以是索引的一部分，而TEXT字段要求您指定前缀长度，该长度可以是索引的一部分。

## TEXT
TEXT 实际上有三种类型：TEXT，还有MEDIUMTEXT或LONGTEXT
MEDIUMTEXT最多可存储16 MB的字符串，而LONGTEXT最多可存储4 GB的字符串，应该尽量避免使用
