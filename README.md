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

Life Term only has one dependency:

* Go Cellular Automata https://github.com/reiver/go-cellularautomata

One way to satisfy this dependency is to install it with a command like the following.
```
go get github.com/reiver/go-cellularautomata/square
```

Note that you may need to be root to do this, if that is the case for you then you
may be able to do that with the command
```
sudo go get github.com/reiver/go-cellularautomata/square
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


