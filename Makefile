all: install

db:
	@echo "\033[32mMake sure that the user has all permissions over the database Area and that the database exists.\033[0m"
	@cat area.sql | mysql -p Area

install:
	@echo "Nothing to install."

.PHONY:	db \
		install
