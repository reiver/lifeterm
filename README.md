Life Term
=========

A simple implementation of Conway's Game of Life in Go, that runs in the terminal.


Building
--------

To build *lifeterm* (from the source code) run the command:
```
go build -o lifeterm  cell.go  main.go  terminal.go  world.go  world_step.go  world_string.go  world_time.go  world_step_life.go
```


Running
-------

To run *lifeterm* run with a command like:
```
./lifeterm
```


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
