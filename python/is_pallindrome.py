def is_palindrome(a):

    list_a=[]
    list_b=[]

    for i in range(0,len(a)):
        list_a.append(a[i])
        list_b.append(a[len(a)-i-1])
    print(list_a==list_b)

text = (input("Please enter a string to check if it is a palindrome : "))

print(is_palindrome(text))
