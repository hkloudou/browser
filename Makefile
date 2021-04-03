git:
	-git autotag -commit 'auto commit' -t -i -f -p
	@echo "current version:`git describe`"