all: myprogram

myprogram:  
	g++ -Wall -Wextra -g  main.cpp

lab1.o: lab1.cpp lab1.hpp
	g++ -c lab1.cpp

main.o: main.cpp lab1.hpp
	g++ -c main.cpp