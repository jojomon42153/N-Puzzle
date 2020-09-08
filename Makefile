# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jojomoon <jojomoon@student.42lyon.fr>      +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/10/30 17:57:13 by gaennuye          #+#    #+#              #
#    Updated: 2020/09/08 18:04:30 by jojomoon         ###   ########lyon.fr    #
#                                                                              #
# **************************************************************************** #

.PHONY: all get install clean test re

GONAME = npuzzle

TEST_FILE = ./ressources/correctInput/taquin_dim3_0.txt

GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOENV = GOPATH=$(GOPATH) GOBIN=$(GOBIN)  
FILES = cmd/algo.go\
		cmd/header.go\
		cmd/heuristics.go\
		cmd/init.go\
		cmd/main.go\
		cmd/parse.go\
		cmd/stats.go\
		cmd/utils.go
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
	@./bin/npuzzle -f $(TEST_FILE)

clean:
	@echo "Cleaning"
	@$(GOENV) go clean
	@rm -rf ./bin/

re: fclean all
