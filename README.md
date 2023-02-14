# The Random Event Simulator
## About this program
This program is a simulater which runs through the given map of events and the probabilities of jumps from one event to another by comparing the current random number with the probabilities of each jump.
## Input Format
The input is stored in the "events" file, and the format is shown as follows.
```
Id, Name, Number of Jumps, Jump Probability, Jump ID, ...
0, Entry,   3, 0.8,   1, 0.1,   2, 0.05, 3
1, Event 1, 1, 0.5,   2
2, Event 2, 2, 0.6,   3, 0.2, 1
3, Event 3, 0, 0,     3
```
Each row is an entry of each event. The length of each row depends on the number of probable jumps.
## Outputs
The program first prints out the parsed map of the random events and then runs a simulation.