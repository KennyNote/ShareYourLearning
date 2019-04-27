# Git
## 安装Git
### Linux平台
首先，你可以试着输入git，看看系统有没有安装Git
```shell
$ git
The program 'git' is currently not installed. You can install it by typing:
sudo apt-get install git
```

如果没有安装，则可以通 sudo apt-get install git 进行安装

### Mac OS X
1.通过homebrew安装，http://brew.sh/
2.通过AppStore安装XCode，选择菜单“Xcode”->“Preferences”，在弹出窗口中找到“Downloads”，选择“Command Line Tools”，点“Install”就可以完成安装了。

### Windows平台
下载地址 https://git-scm.com/downloads

安装完成后，在开始菜单里找到“Git”->“Git Bash”



安装完成
安装完成后，还需要最后一步设置，在命令行输入：
```shell
$ git config --global user.name "Your Name"
$ git config --global user.email "email@example.com"
```

查看git的用户名和邮箱
```shell
$ git config user.name
$ git config user.email
```

注意git config命令的--global参数，用了这个参数，表示你这台机器上所有的Git仓库都会使用这个配置，当然也可以对某个仓库指定不同的用户名和Email地址。

## 创建版本库
选择一个合适的地方，创建一个空目录：
```shell
$ mkdir learngit
$ cd learngit
$ pwd
/Users/michael/learngit
```

<font color=red>如果你使用Windows系统，为了避免遇到各种莫名其妙的问题，请确保目录名（包括父目录）不包含中文。</font>

通过git init命令把这个目录变成Git可以管理的仓库：
```shell
$ git init
Initialized empty Git repository in /Users/michael/learngit/.git/
```


把文件添加到版本库
所有的版本控制系统，其实只能跟踪文本文件的改动，比如TXT文件，网页，所有的程序代码等等
不幸的是，Microsoft的Word格式是二进制格式，因此，版本控制系统是没法跟踪Word文件的改动的
因为文本是有编码的，比如中文有常用的GBK编码，日文有Shift_JIS编码，如果没有历史遗留问题，强烈建议使用标准的UTF-8编码

<font color=red>使用Windows的童鞋要特别注意：</font>

千万不要使用Windows自带的记事本编辑任何文本文件。原因是Microsoft开发记事本的团队使用了一个非常弱智的行为来保存UTF-8编码的文件，他们自作聪明地在每个文件开头添加了0xefbbbf（十六进制）的字符，你会遇到很多不可思议的问题，比如，网页第一行可能会显示一个“?”，明明正确的程序一编译就报语法错误，等等，都是由记事本的弱智行为带来的。建议你下载Notepad++代替记事本，不但功能强大，而且免费！记得把Notepad++的默认编码设置为UTF-8 without BOM即可：


在learngit中或者其子文件夹下创建一个readme.txt文件，内容如下
```shell
Git is a version control system.
Git is free software.
```

1.用命令git add告诉Git，把文件添加到仓库：
```shell
$ git add readme.txt
```
2.用命令git commit告诉Git，把文件提交到仓库：
```shell
$ git commit -m "wrote a readme file"
[master (root-commit) eaadf4e] wrote a readme file
 1 file changed, 2 insertions(+)
 create mode 100644 readme.txt
 ```

解释一下git commit命令，-m后面输入的是本次提交的说明

为什么Git添加文件需要add，commit一共两步呢？因为commit可以一次提交很多文件，所以你可以多次add不同的文件，比如：
```shell
$ git add file1.txt
$ git add file2.txt file3.txt
$ git commit -m "add 3 files."
```

这样文件文件就被提交到本地仓库中了