# (0,1,2,3,1,2,0)
def v1():
    for i in range(3):
        print (i)
        for i in range(3):
            print (i)
            for i in range(3):
                print (i)
                for i in range(3):
                    print (i)
        for i in range(3):
            print (i)
            for i in range(3):
                print (i)
    for i in range(3):
        print (i)

# (0,1,1,2,3,1,1)
def v2():
    for i in range(3):
        print (i)
        for i in range(3):
            print (i)
        for i in range(3):
            print (i)
            for i in range(3):
                print (i)
                for i in range(3):
                    print (i)
        for i in range(3):
            print (i)
        for i in range(3):
            print (i)

# (0,1,1,1,0,1,1)
def v3():
    for i in range(3):
        print (i)
        for i in range(3):
            print (i)
        for i in range(3):
            print (i)
        for i in range(3):
            print (i)
    for i in range(3):
        print (i)
        for i in range(3):
            print (i)
        for i in range(3):
            print (i)


v2()

# P(f) denotes the no. of distinct valid Python programs we can construct using F copies of 
# the code snippet. Code a DP algorithm to compute P(f)

# Use a 2D array A[ 1..F, 0..D ] to store info of subproblems
# F denotes the no. of for loops
# D denotes the max depth, which is F-1
# Each cell in A stores a non-neg integer (rather large)

# After computing A, use it to computer a 1D array B[1..F]
# B[f]=P(f) for all f=1..F  ->  That is, B contains the no. of distinct valid Python programs

# A[f,d] = sum of A[f-1,d-1] + A[f-1,d] + A[f-1,d+1] + ... + A[f-1,D]