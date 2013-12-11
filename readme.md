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
| 			| Bankers	| LIFO	| FIFO	|
|-----------|-----------|-------|-------|
|test1		| fail		| fail 	| fail	|
|test2		| pass		| pass	| pass 	|
|test3		| fail		| fail	| fail	|
|test4		| pass		| fail	| fail 	|
|test5		| pass		| fail 	| fail 	|
|test6		| pass		| fail	| fail	|
|test7		| fail		| fail	| fail 	|
|test8		| pass		| pass	| fail	|
|test9		| pass		| pass	| fail	|
|test10		| pass		| fail	| fail	|
|Overall	| 70%		| 30%	| 10%	|

Running time (in micro second):  
| 			| Bankers	| LIFO	| FIFO	|
|-----------|-----------|-------|-------|
|test1		| 179		| *166*	| 207	|
|test2		| 171		| 182	| *145*	|
|test3		| 265		| 228	| *217*	|
|test4		| *206*		| 339	| 334	|
|test5		| 677		| 583	| *548*	|
|test6		| 317		| 290	| *289*	|
|test7		| 369		| 333	| *311*	|	
|test8		| 615		| *536*	| 1117	|
|test9		| 601		| *146*	| 167	|
|test10		| 454		| *338*	| 856	|

Bankers algorithm is able to avoid every possible unsafe status, thus it's easy to conclude that for all the test cases Bankers Algorithm fails, both LIFO and FIFO fail as well. These results exemplify the efficiency of Bankers Algorithm in terms of resource management comparing to simple LIFO/FIFO. Bankers Algorithm is more robust.

When comparing the algorithms in terms of running time, Bankers tended to be slightly slower than than its LIFO/FIFO counterparts, most likely due to its more complicated structure. FIFO generally performed the fastest, although this may be due to the high failure rate.

Through the tests, there are a few advantages and disadvantages of the Bankers algorithm implementation shown.  One of the disadvantages is how much information is required by the algorithm in order to schedule the processes.  In certain systems, this information may not be available.


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


Go Programming Language  
Go was used to implement the scheduling algorithms. Similarly to Rust, Go supports concurrency at the language level. The syntax of the language minimizes the length of code, and was similar to other familiar languages.  