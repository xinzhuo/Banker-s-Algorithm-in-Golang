#This is an implementation of the Banker's Algorithm in Golang  
The Bankers Algorithm is a resource allocation and deadlock provention algorithm developed by Edsger Dijkstra that tests for safety by simulating the allocation of predetermined maximum possible amounts of all resources, and then makes a "s-state" check to test for possible deadlock conditions for all other pending activities, before deciding whether allocation should be allowed to continue. The algorithm was developed in the design process for the operating system.  

The Bankers Algorithm need three things to work:  
- How much of each resource each process could request  
- How much of each resource each process is allocated
- How much of each resource the system is currently holding
The algorithm only allocate resources to processes after it determined the processes are in safe state.

We also implemented a scheduling program in traditional FIFO/LIFO in similar fashion. We would use a number of test cases to compare the running time and pass/fail rate in FIFO and Banker's Algorithm.   

This implementation will determine if certain state is safe or unsafe (may cause deadlock)  
If the given state is safe, which means its all subsequent states are safe, the program will give the safe sequence of the processes and the corresponding available resources in OS  
The program also have a timer to measure its performance with different test cases  

Result:

The test cases are provided in the subdirectory "tests"   
The test cases has the format:  
P: Number of Processes  
R: Number of Resources  
int[R]: Available Resources in OS  
int[P][R]: Matrix of Maximum Request Resources  
int[P][R]: Matrix of Allocated Resources  

Sample Test Case:  
5  
4  
8 5 9 7  
3 2 1 4  
0 2 5 2  
5 1 0 5  
1 5 3 0  
3 0 3 3  
2 0 1 1  
0 1 2 1  
4 0 0 3  
0 2 1 0  
1 0 3 0  