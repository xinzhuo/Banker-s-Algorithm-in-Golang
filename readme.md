This is an implementation of the Banker's Algorithm in Golang  
It will determine if certain state is safe or unsafe (may cause deadlock)  
If the given state is safe, which means its all subsequent states are safe, the program will give the safe sequence of the processes and the corresponding available resources in OS  
The program also have a timer to measure its performance with different test cases  

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