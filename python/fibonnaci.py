def fibo(n : int):
    if n < 0:
        print("n must be positive or or nul")
    elif n == 0 or n == 1:
        return n
    else:
        return fibo(n - 1) + fibo(n - 2)

i = int(input("Please pick an integer that is positive or nul:"))

print(fibo(i))

