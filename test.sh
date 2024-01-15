#!/bin/bash

##Testing The banner
echo "go run . "banana" standard abc"
go run . "banana" standard abc 
echo
echo "---------------------------------------------------------"

echo "run . "hello" standard"
go run . "hello" standard 
echo
echo "---------------------------------------------------------"

echo "go run . "hello world" shadow"
go run . "hello world" shadow 
echo
echo "---------------------------------------------------------"

echo "go run . "nice 2 meet you" thinkertoy"
go run . "nice 2 meet you" thinkertoy 
echo
echo "---------------------------------------------------------"

echo "go run . "you \& me" standard"
go run . "you & me" standard 
echo
echo "---------------------------------------------------------"

echo "go run . "123" shadow"
go run . "123" shadow 
echo
echo "---------------------------------------------------------"

echo "go run . "/\(\"\)" thinkertoy"
go run . "/(\")" thinkertoy 
echo
echo "---------------------------------------------------------"

echo "go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow"
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow 
echo
echo "---------------------------------------------------------"

echo "go run . '\\!\" #$%&\'\"'\"'()*+,-./' thinkertoy"
go run . '\!" #$%&'"'"'()*+,-./' thinkertoy 
echo
echo "---------------------------------------------------------"

echo "go run . "It\'s Working" thinkertoy"
go run . "It's Working" thinkertoy 
echo
echo "---------------------------------------------------------"
##############

##Tesing the Output Flag 
echo "go run . --output=banner01.txt \"hello\" standard"
go run . --output=banner01.txt "hello" standard 
echo
echo "---------------------------------------------------------"

echo "go run . --output=banner02.txt \"Hello There!\" shadow"
go run . --output=banner02.txt "Hello There!" shadow 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test00.txt \"First\nTest\" shadow"
go run . --output=test00.txt "First\nTest" shadow
echo
echo "---------------------------------------------------------"

echo "go run . --output=test01.txt \"hello\" standard"
go run . --output=test01.txt "hello" standard
echo
echo "---------------------------------------------------------"

echo "go run . --output=test02.txt \"123 -> \#$%\" standard"
go run . --output=test02.txt "123 -> #$%" standard
echo
echo "---------------------------------------------------------"

echo "go run . --output=test03.txt \"432 -> \#$%\&@\" shadow"
go run . --output=test03.txt "432 -> #$%&@" shadow 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test04.txt \"There\" shadow"
go run . --output=test04.txt "There" shadow
echo
echo "---------------------------------------------------------"

echo "go run . --output=test05.txt \"123 -> \"#$%@\" thinkertoy"
go run . --output=test05.txt "123 -> \"#$%@" thinkertoy
echo
echo "---------------------------------------------------------"

echo "go run . -output=test06.txt \"2 you\" thinkertoy"
go run . --output=test06.txt "2 you" thinkertoy
echo
echo "---------------------------------------------------------"

echo "go run . --output=test07.txt 'Testing long output!' standard"
go run . --output=test07.txt 'Testing long output!' standard
echo
echo "---------------------------------------------------------"
##############

##Testing the Color Flag
echo "go run . --color red "banana""
go run . --color red "banana" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=red \"hello world\""
go run . --color=red "hello world" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=green \"1 + 1 = 2\""
go run . --color=green "1 + 1 = 2"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow \"(%&) ??\""
go run . --color=yellow "(%&) ??" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  ello \"Hello\""
go run . --color=yellow  ello "Hello" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  e \"Hello\""
go run .  --color=yellow  e \"Hello\" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  el \"Hello\""
go run . --color=yellow  el \"Hello\" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=orange GuYs \"HeY GuYs\""
go run . --color=orange GuYs "HeY GuYs"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=blue B 'RGB()'"
go run . --color=blue B 'RGB()'  
echo
echo "---------------------------------------------------------"

echo "go run . --color=cyan B \"ReBoOt\""
go run . --color=cyan B "ReBoOt"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple ! \"Reboot 01!\""
go run . --color=purple ! "Reboot 01!" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#Reboot 01!!!\""
go run . --color=purple R "#Reboot 01!!!" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#R eb oot &&&&01!!!\""
go run . --color=purple R "#R eb oot &&&&01!!!"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R --color=cyan ! \"#R eb oot &&&&01!!!\""
go run . --color=purple R --color=cyan ! "#R eb oot &&&&01!!!"  
echo
echo "---------------------------------------------------------"
##############

##Testing the Align Flag
echo "go run . --align right something standard"
go run . --align right something standard
echo
echo "---------------------------------------------------------"

echo "go run . --align=right left standard"
go run . --align=right left standard 
echo
echo "---------------------------------------------------------"

echo "go run . --align=left right standard"
go run . --align=left right standard
echo
echo "---------------------------------------------------------"

echo "go run . --align=center hello shadow"
go run . --align=center hello shadow
echo
echo "---------------------------------------------------------"

echo "go run . --align=justify "1 Two 4" shadow"
go run . --align=justify "1 Two 4" shadow
echo
echo "---------------------------------------------------------"

echo "go run . --align=right 23/32 standard"
go run . --align=right 23/32 standard
echo
echo "---------------------------------------------------------"

echo "go run . --align=right ABCabc123 thinkertoy"
go run . --align=right ABCabc123 thinkertoy
echo
echo "---------------------------------------------------------"

echo "go run .--align=center "#$%\&\"" thinkertoy"
go run . --align=center "#$%&\"" thinkertoy
echo
echo "---------------------------------------------------------"

echo "go run .--align=left '23Hello World!' standard"
go run . --align=left '23Hello World!' standard  
echo
echo "---------------------------------------------------------"

echo "go run . --align=justify 'HELLO there HOW are YOU?!' thinkertoy"
go run . --align=justify 'HELLO there HOW are YOU?!' thinkertoy
echo
echo "---------------------------------------------------------"

echo "go run . --align=right "a -\> A b -\> B c -\> C" shadow"
go run . --align=right "a -> A b -> B c -> C" shadow
echo
echo "---------------------------------------------------------"

echo "go run . --align=right abcd shadow"
go run . --align=right abcd shadow
echo
echo "---------------------------------------------------------"

echo "go run .--align=center ola standard"
go run . --align=center ola standard
echo
echo "---------------------------------------------------------"

##Testing Reverse Flag
echo "go run . --reverse example00.txt"
go run . --reverse example00.txt
echo
echo "---------------------------------------------------------"

go run . --output=example00.txt "Hello World"
echo "go run . --reverse=example00.txt"
go run . --reverse=example00.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example01.txt "123"
echo "go run . --reverse=example01.txt"
go run . --reverse=example01.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example02.txt "#=\["
echo "go run . --reverse=example02.txt"
go run . --reverse=example02.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example03.txt "something&234"
echo "go run . --reverse=example03.txt"
go run . --reverse=example03.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example04.txt "abcdefghijklmnopqrstuvwxyz"
echo "go run . --reverse=example04.txt"
go run . --reverse=example04.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example05.txt "\!\" #$%&'()*+,-./"
echo "go run . --reverse=example05.txt"
go run . --reverse=example05.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example06.txt ":;{=}?@" shadow
echo "go run . --reverse=example06.txt"
go run . --reverse=example06.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example07.txt "ABCDEFGHIJKLMNOPQRSTUVWXYZ" thinkertoy
echo "go run . --reverse=example07.txt"
go run . --reverse=example07.txt  
echo
echo "---------------------------------------------------------"

##Testing Flags togather
echo "go run main.go --color=#663AE2 ! --align=center "Hello, World!""
go run main.go --color=#663AE2 ! --align=center "Hello, World!"
echo
echo "---------------------------------------------------------"

echo "go run main.go --color=red --align=right "Hello, World!""
go run main.go --color=red --align=right "Hello, World!"
echo
echo "---------------------------------------------------------"

echo "go run main.go --align=justify --output=file01.txt "Hello, World!" shadow"
go run main.go --align=justify --output=file01.txt "Hello, World!" shadow
echo
echo "---------------------------------------------------------"

echo "go run main.go --align=left --output=file02.txt "Hello, World!""
go run main.go --align=left --output=file02.txt "Hello, World!"
echo
echo "---------------------------------------------------------"

echo "go run main.go --align=right --output=file03.txt "Hello, World!""
go run main.go --align=right --output=file03.txt "Hello, World!"
echo
echo "---------------------------------------------------------"

echo "go run main.go --align=right --output=file04.txt --color=red "something" shadow"
go run main.go --align=right --output=file04.txt --color=red "something" shadow
echo
echo "---------------------------------------------------------"