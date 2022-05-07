# 执行 makefile 时复制 git hook 脚本
CopyGitHook:=$(shell cp -f githooks/* .git/hooks/)