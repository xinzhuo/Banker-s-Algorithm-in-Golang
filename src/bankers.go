package main
import (
    "fmt"
    "time"
)
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("\n%s took %s", name, elapsed)
}
func run() {
    defer timeTrack(time.Now(), "Banker's Algorithm")
    var safe, exec bool
    var count, i, j, r, p int
    //var i, j, p, r, count int

    fmt.Printf("Enter the number of processes: ")
    var running [5]bool
    fmt.Scanf("%d", &p)
    for i=0; i<p; i++ {
        running[i] = true
        count++
    }

    fmt.Printf("Enter the number of resources: ")
    fmt.Scanf("%d", &r)

    var curr [5][5]int //currently allocated resources
    var maxclaim [5][5]int //maximum resources
    var avl [5]int //available resources
    var alloc [5]int
    var maxres [5]int

    fmt.Printf("Enter the number of resource for instance: ")
    for i=0; i<r; i++ {
        fmt.Scanf("%d", &maxres[i])
    }

    fmt.Printf("Enter maximum resource table: \n")
    for i=0; i<p; i++{
        for j=0; j<r; j++{
            fmt.Scanf("%d", &maxclaim[i][j])
        }
    }

    fmt.Printf("\nEnter allocated resource table: \n")
    for i=0; i<p; i++{
        for j=0; j<r; j++{
            fmt.Scanf("%d", &curr[i][j])
        }
    }

    for i=0; i<p; i++{
        for j=0; j<r; j++{
            alloc[j]+=curr[i][j]
        }
    }

    for i=0; i<r; i++{
        avl[i]=maxres[i]-alloc[i]
    }

    fmt.Printf("\nThe resource of instances:")
    fmt.Println(maxres[0:r]);

    fmt.Printf("\nThe maximum resource table: \n")
    for i=0; i<p; i++{
        fmt.Println(maxclaim[i][0:j])
    }   


    fmt.Printf("\nThe allocated resource table: \n")
    for i=0; i<p; i++{
        fmt.Println(curr[i][0:j])
    }

    fmt.Printf("\nAllocated resources: ")
    fmt.Println(alloc[0:r])

    fmt.Printf("\nAvailable resources: ")
    fmt.Println(avl[0:r])

    fmt.Printf("\nRunning Status: ")
    fmt.Println(running[0:p])
    //fmt.Printf("%d", p);
    for count!=0 {
        for i=0; i<p; i++{
            if(running[i]){
                exec=true;
                for j=0; j<r; j++{
                    if maxclaim[i][j]-curr[i][j] > avl[j]{
                        exec=false
                        break
                    }
                }
                if exec{
                    fmt.Printf("\nProcess %d is executing", i+1)
                    running[i]=false
                    count--
                    safe=true
                    for j=0; j<r; j++ {
                        avl[j]+=curr[i][j]
                    }
                    break
                }
            }
        }
        if !safe {
            fmt.Printf("\nThe processes are in unsafe state (May cause deadlock).")
            break
        }else {
            fmt.Printf("\n\tThe processes are in safe state")
            fmt.Printf("\n\tAvailable resources:")
            fmt.Println(avl[0:r])
        }
    }
}


func main() {
    run()
}