# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/10/30 17:57:13 by jmonneri          #+#    #+#              #
#    Updated: 2020/02/25 15:08:21 by jmonneri         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

.PHONY: all get install run fclean test

GONAME = npuzzle

TEST_FILE = ./examples/correctInput.txt

GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOENV = GOPATH=$(GOPATH) GOBIN=$(GOBIN)  
FILES = $(wildcard ls cmd/*.go)
EXECPATH = ./bin/$(GONAME)


all: $(EXECPATH)

$(EXECPATH): $(GOFILES)
	@printf "0️⃣  Building $(FILES) to ./bin \n"
	@$(GOENV) go build -o $(EXECPATH) $(FILES)
	@printf "✅  Built! ✅\n"

get:
	@$(GOENV) go get .

install:
	@$(GOENV) go install $(GOFILES) 

test: all

script: all
	@sh tests.sh

run: all
	@$(EXECPATH) $(TEST_FILE)
	
%:
	@:

fclean:
	@echo "Cleaning"
	@$(GOENV) go clean
	@rm -rf ./bin/

re: fclean all
