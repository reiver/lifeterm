Life Term
=========

A simple implementation of Conway's Game of Life in Go, that runs in the terminal.


Screenshot?
-----------

Although *lifeterm* will (probably) look different each time you run it
(and depends on the width and height of your terminal), it looks something like:
```
                                    ██   █  ▓▓▓██   ██                       ▓█
█ █▓                                 ▓▓▓▓▓▓████  █                             
▓  █                                  ███   ▓▓▓██      ██                      
   █                                 ▓█  █             ██                 ▓▓█  
                       ▓             █   █                                █  █ 
                    █    █           █ █          █        ██                ▓ 
                     ██  █                       █ █       ██                  
                      ▓▓▓             ▓█ █▓      █  █    ▓▓▓                █ █
                                        ▓         ██    █ █                 ▓█▓
                         ██                            ▓█ █▓                   
                █        ██       ██▓                                          
               █▓█                ▓█                                 ▓█▓       
              █▓▓▓█                 █                               █   █      
              ▓█▓ █                        ██                ██    ▓           
               ██ ██                       █ █▓             ▓ █     ██ █▓      
          █▓   ▓█ █▓     █                  ▓██              █       ▓▓        
          ██     █     ▓ █            ██                  ███                  
                        █▓           ▓  ▓   ▓██            ▓         ▓█▓       
                                     ▓██  ██ ██                                
                                   ▓█ ▓  ▓ █▓              █         █      ▓  
                                   █ █   █ ▓              █ █        ▓ █   ███ 
                                   █  █ ▓           ██     █               ▓ █▓
```


Dependencies
------------

Go Cellular Automata https://github.com/reiver/go-cellularautomata/square

One way to get this satisfy this dependency is to run:
```
go get github.com/reiver/go-cellularautomata/square
```


Building
--------

To build *lifeterm* (from the source code) run the command:
```
go build -o lifeterm  main.go  terminal.go
```

(This will create the file "lifeterm".)


Running
-------

To run *lifeterm* run with a command like:
```
./lifeterm
```


