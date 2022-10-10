t ?=

all: install

db:
	@printf "\033[0;32mMake sure that the user has all permissions over the database Area and that the database exists.\033[0m\n"
	@cat area.sql | mysql -p Area

install:
	@make -C server install
	@make -C web install

run:
	@make -C $(t) start

pr:
	@gh pr create --fill --base dev

start: run

.PHONY: all \
		db \
		install \
		server \
		pr \
