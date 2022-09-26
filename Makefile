all: install

db:
	@printf "\033[0;32mMake sure that the user has all permissions over the database Area and that the database exists.\033[0m\n"
	@cat area.sql | mysql -p Area

install:
	@echo "Nothing to install."

pr:
	@gh pr create --fill --base dev

.PHONY:	db \
		install
