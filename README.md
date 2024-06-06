# VORTO Challenge


Based on [this Google Drive folder](https://drive.google.com/drive/folders/1Jb7FmR5Ftrg0jwgIJ-n_oKwOjyJ4gDHI)


## Make File

```make``` 


- -o ./bin/vorto


```make vorto_test```


Will run through all of the Loads and output result.


# Commands

CLI - 


```./bin/vorto v100 -T "./training/Problems/problem20.txt"```

output - 


```
vincentyoumans@vincents-Air vorto % ./bin/vorto v100 -T "./training/Problems/problem20.txt"
[88, 36, 13, 176]
[183, 147, 71, 15, 28, 173]
[69, 7, 143, 167]
[19, 133]
[185, 104, 2, 44]
[125, 18, 145, 144]
[75, 105, 195]
[174, 151, 181, 177]
[45, 193, 81, 198]
[102, 85]
[136, 154, 66, 21]
[192, 77, 178, 53]
[12, 164, 146]
[23, 35]
[1, 150, 100]
[129, 108, 196]
[79, 134, 124]
[40, 98]
[41, 65, 17]
[130, 121, 137, 84, 175]
[11, 138, 67, 49, 109, 197]
[155, 34, 27, 107, 74, 30]
[82, 184, 68, 48, 59, 141]
[78, 112, 135]
[139, 16]
[103, 163, 168, 38, 153]
[179, 114, 73, 58]
[54, 158, 93, 89]
[61, 37]
[142, 50, 32, 157]
[149, 91, 14]
[72, 87, 33]
[20, 83, 113, 22, 5]
[26, 189, 94]
[96, 117, 187]
[110, 127, 31]
[51, 132, 199]
[70, 39, 171]
[148, 123, 111]
[186, 25]
[63, 80, 10]
[9, 24]
[188, 131]
[118, 29, 43]
[6, 55, 120]
```





# Directory Structure

```
├── CHAL_docs
│   ├── drive-download-20240531T045535Z-001
│   │   ├── 2023 10 03 - Vorto Algorithmic Challenge Instructions.docx
│   │   ├── 2023 10 11 - Leaderboard.xlsx
│   │   ├── Training Problems                               // Problem files delivered in challenge
│   │   │   ├── problem1.txt
│   │   │   ├── problem10.txt
│   │   │   ├── problem11.txt
│   │   │   ├── problem12.txt
│   │   │   ├── problem13.txt
│   │   │   ├── problem14.txt
│   │   │   ├── problem15.txt
│   │   │   ├── problem16.txt
│   │   │   ├── problem17.txt
│   │   │   ├── problem18.txt
│   │   │   ├── problem19.txt
│   │   │   ├── problem2.txt
│   │   │   ├── problem20.txt
│   │   │   ├── problem3.txt
│   │   │   ├── problem4.txt
│   │   │   ├── problem5.txt
│   │   │   ├── problem6.txt
│   │   │   ├── problem7.txt
│   │   │   ├── problem8.txt
│   │   │   └── problem9.txt
│   │   ├── evaluateShared.py
│   │   └── evaluationReadMe.txt
│   └── drive-download-20240531T045535Z-001.zip
├── Makefile
├── README.md
├── bin
│   └── vorto
├── go.mod
├── go.sum
├── output
│   ├── P200
│   │   └── p200.json
│   ├── P300
│   │   └── p300.json
│   ├── Problems                                // I saved stages to JSON files for debugging.
│   │   ├── P320                                // The PP util files can be remarked off in the
│   │   │   └── p320.json                       // ./vorto/cmd/v100.go file
│   │   ├── PM500
│   │   │   └── pm500.json
│   │   ├── PM520
│   │   │   └── pm520.json
│   │   ├── PM530
│   │   │   └── pm530.json
│   │   ├── prob20
│   │   │   └── prob20.json
│   │   └── prob300
│   │       └── prob300.json
│   ├── ret
│   ├── slog
│   │   └── P100                                // sLog, which I did not use much
│   │       └── log1000.json
│   └── strucs
├── training
│   └── Problems
│       ├── problem1.txt
│       ├── problem10.txt
│       ├── problem11.txt
│       ├── problem12.txt
│       ├── problem13.txt
│       ├── problem14.txt
│       ├── problem15.txt
│       ├── problem16.txt
│       ├── problem17.txt
│       ├── problem18.txt
│       ├── problem19.txt
│       ├── problem2.txt
│       ├── problem20.txt
│       ├── problem3.txt
│       ├── problem4.txt
│       ├── problem5.txt
│       ├── problem6.txt
│       ├── problem7.txt
│       ├── problem8.txt
│       └── problem9.txt
├── tree.txt
├── vorto
│   ├── LICENSE
│   ├── cmd
│   │   ├── root.go
│   │   └── v100.go
│   ├── main
│   ├── main.go
│   └── pkg
│       ├── slogPkg
│       │   └── slog100.go
│       ├── strucs                          // strucs for challange
│       │   ├── strucs100.go
│       │   ├── strucs150.go
│       │   ├── strucs200.go
│       │   └── strucs500.go
│       ├── util                            // utilities 
│       │   ├── PrettyPrint.go
│       │   ├── savefiles.go
│       │   └── util.go
│       └── v100                            // V100 Steps
│           ├── step100.go
│           ├── step200.go
│           ├── step300.go
│           ├── step500.go
│           └── v100.go
└── vtest.txt                               // Output of all of the Problem files.

28 directories, 79 files
```



# Narrative.

1. I understand that I was suppose to do some google search on this concept, but
I really did not have much time.  I resorted to just common sence.  I suspect there may
be a more interesting solution.

2. The ./output dir is there for debuging by saving JSON files of each step.


## Strategy

### Stage 100. - Open problem file and hydrate
Open problem file that is passed as a parameter in cli


### Stage 200 - Create An ordered list
of Drivers assuming...
- 1 driver.
- no 12 hour violation rule.
- StartPoint is Depot (0,0)
- EndPoint is Depot (0,0)
- Next StartPoint is Next Closest PickUp.


I also made assumption that the Next_PickUp should be the closest Pickup.
Another method would be Next_PickUp = distance(pickUp) + distance(dropOff)
Theory being that the next pickup may not be the most cost effective dependant on dropoff.
The next feature would be to build a ROUTE slice with second method.
Then compare both strategies to find most cost effective path.

### Stage 500 - create Drivers vertex
Create a vertex that will build drivers schedules without violating 12 hour rule.
Also adds only as many drivers as necessary.

### Stage Final - fmt.prints results.




# Notes

from the ./vorto directory

```
go run main.go v100 -T "../training/Problems/problem10.txt"
```

# Disclaimer...

Next feature would be to clean up the code...  I went on several tangents
