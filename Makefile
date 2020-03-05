# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/10/30 17:57:13 by jmonneri          #+#    #+#              #
#    Updated: 2020/03/04 17:55:35 by jmonneri         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

.PHONY: all get install run fclean test

GONAME = npuzzle

TEST_FILE = ./ressources/correctInput/taquin_dim3_0.txt

GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOENV = GOPATH=$(GOPATH) GOBIN=$(GOBIN)  
FILES = $(wildcard ls cmd/*.go)
EXECPATH = ./bin/$(GONAME)


all: $(EXECPATH)

$(EXECPATH): $(FILES)
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
	@$(EXECPATH) -f $(TEST_FILE) -hm	
	
%:
	@:

fclean:
	@echo "Cleaning"
	@$(GOENV) go clean
	@rm -rf ./bin/

re: fclean all
