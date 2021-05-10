# Git 学习笔记

> 本笔记是来自阅读Pro Git时所记录 ，[Pro Git中文第二版](https://www.progit.cn/)

## 起步

### Git保证完整性

Git中所有数据在存储前都计算校验和，使用SHA-1散列机制进行计算。

### 三种状态

* 已提交（commited）：文件已保存在本地数据库中
* 已暂存（staged）：对一个已修改文件的做标记，包含在下次提交的快照中
* 已修改（modified）：表示文件已被修改，但还没有保存在数据库中

基本Git的工作流程：

1. 在工作目录中修改文件
2. 将修改的文件快照放入暂存区域
3. 提交更新，将暂存区域的文件快照永久性存储到git本地仓库

### 配置用户信息

``` shell
$ git config --global user.name "xxx"
$ git config --global user.email xxx@example.com
```

## Git 基础

### 获取git仓库

1. 将当前目录初始化为 Git 仓库

   * 为空文件夹时，只需要输入

     ``` shell
     $ git init
     ```

     该命令会创建一个名为 .git 的子目录

   * 为已经存在文件的文件夹中

     在完成初始化操作后你需要将文件夹内的文件进行跟踪并提交，使用add命令进行跟踪，并使用commit进行提交

     ``` shell
     $ git add .
     $ git commit -m "initial project version"
     ```

     add 后面的 '.'表示将文件夹下的所有文件进行跟踪，commit -m 引号内的内容表示提交内容

     

2. 克隆现有的仓库

   想要获得一份已经存在的Git仓库的拷贝，需要使用 clone 命令，当执行 git clone操作时，默认配置下远程 Git 仓库中的**每一个文件**的**每一个版本**都将被拉取下来。

   ``` shell
   $ git clone [url]
   ```

   上面的命令用到了https://传输协议，git 还支持 git:// 、ssh等传输协议。

### git 文件状态图

![Git 下文件生命周期图。](https://www.progit.cn/images/lifecycle.png)

### 检查当前文件状态

检查哪些文件处于什么状态使用 git status 命令查看

``` shell
$ git status
```

### 跟踪新文件

``` shell
$ git add xxx
```

### 暂存已修改文件

将已修改文件存入暂存区也是使用 add 命令，下次提交时就会一并记录到仓库。

若对一个已经暂存（使用 add 命令）的文件再次进行编辑，此时若进行commit提交，提交的将是上一次 add 时的版本，所以若要提交最新的更改，需要再执行一次 add 命令。

``` shell
$ git add README.md
# 对README做了一些修改
$ git add README.md# 需要再进行一次add
$ git commit -m "modify README.md"
```

### 忽略文件

一般我们总会有些文件无需纳入 Git 的管理，也不希望它们总出现在未跟踪文件列表。 通常都是些自动生成的文件，比如日志文件，或者编译过程中创建的临时文件等。 在这种情况下，我们可以创建一个名为 .gitignore 的文件，列出要忽略的文件模式。

``` shell
cat .gitignore
*.[oa]   # 表示忽略所有以.o 或 .a 结尾的文件
*~       # 忽略所有以～结尾的文件
```

星号（**）匹配零个或多个任意字符；[abc] 匹配任何一个列在方括号中的字符（这个例子要么匹配一个 a，要么匹配一个 b，要么匹配一个 c）；问号（?）只匹配一个任意字符；如果在方括号中使用短划线分隔两个字符，表示所有在这两个字符范围内的都可以匹配（比如 [0-9] 表示匹配所有 0 到 9 的数字）。 使用两个星号（*) 表示匹配任意中间目录，比如`a/**/z` 可以匹配 a/z, a/b/z 或 `a/b/c/z`等。

### 查看已暂存和未暂存的修改

``` shell
$ git diff#查看所有修改
$ git diff --cached(或者staged)# 查看已暂存的将要添加到下次提交里的内容
```

### 提交更新

``` shell
$ git commit
$ git commit -m "message"
$ git commit -a #跳过add暂存直接提交
```

### 移除文件

要从 Git 中移除某个文件，就必须要从已跟踪文件清单中移除（确切地说，是从暂存区域移除），然后提交。 可以用 `git rm` 命令完成此项工作，并连带从工作目录中删除指定的文件，这样以后就不会出现在未跟踪文件清单中了。

如果只是简单地从工作目录中手工删除文件，运行 `git status` 时就会在 “Changes not staged for commit” 部分（也就是 *未暂存清单*）看到：

```shell
$ rm PROJECTS.md
$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        deleted:    PROJECTS.md

no changes added to commit (use "git add" and/or "git commit -a")
```

然后再运行 `git rm` 记录此次移除文件的操作：

```shell
$ git rm PROJECTS.md
rm 'PROJECTS.md'
$ git status
On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

    deleted:    PROJECTS.md
```

下一次提交时，该文件就不再纳入版本管理了。 如果删除之前修改过并且已经放到暂存区域的话，则必须要用强制删除选项 `-f`（译注：即 force 的首字母）。 这是一种安全特性，用于防止误删还没有添加到快照的数据，这样的数据不能被 Git 恢复。

另外一种情况是，我们想把文件从 Git 仓库中删除（亦即从暂存区域移除），但仍然希望保留在当前工作目录中。 换句话说，你想让文件保留在磁盘，但是并不想让 Git 继续跟踪。 当你忘记添加 `.gitignore` 文件，不小心把一个很大的日志文件或一堆 `.a` 这样的编译生成文件添加到暂存区时，这一做法尤其有用。 为达到这一目的，使用 `--cached` 选项：

```shell
$ git rm --cached README
```

`git rm` 命令后面可以列出文件或者目录的名字，也可以使用 `glob` 模式。 比方说：

```shell
$ git rm log/\*.log
```

注意到星号 `*` 之前的反斜杠 `\`， 因为 Git 有它自己的文件模式扩展匹配方式，所以我们不用 shell 来帮忙展开。 此命令删除 `log/` 目录下扩展名为 `.log` 的所有文件。 类似的比如：

```shell
$ git rm \*~
```

该命令为删除以 `~` 结尾的所有文件。

