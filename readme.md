# Go Foundation

## Magesh Kuppan

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:40 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 3:00 PM (20 mins)
- Wind up       : 5:00 PM

## Methodology
- No powerpoints
- Discussion & Code

## Repository
- https://github.com/tkmagesh/cisco-go-dec-2023

## Software Requirements
- Go tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Go Extension for Visual Studio Code (https://marketplace.visualstudio.com/items?itemName=golang.Go)
### Verification
- go version

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers
    - No classes (only structs)
    - No inheritance (only composition)
    - No reference types (everything is a value)
    - No pointer arithmatic
    - No exceptions (only errors)
    - No try..catch..finally
    - No implicit type conversion
- Better Performance
    - Equivalent to C++
    - Close to the hardware (No VM)
        - Compile to machine code targetting the OS
        - Cross compilation is supported 
- Managed Concurrency
    - Light weight (1 goroutine = 4 KB whereas 1 OS Thread = 2 MB)
    - Language support for concurrency
        - go keyword, channel (data type), channel operator ( <- ), range & select case constructs
    - APIs 
        - sync package
        - sync/atomic package