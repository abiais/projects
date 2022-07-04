"""Rock Paper Scissors
--------------------------------
"""

import random

score_player = 0
score_computer = 0
options = ['rock', 'paper', 'scissors']
random.seed(42)

def results(a,b) :
    """This functions takes the value played by two players
    and determines who is the winner"""
    
    round_winner = 0
    
    if a == b:
        round_winner = 0
    elif a + b == 5 or a + b == 3:
        if a > b :
            round_winner = 1
        else :
            round_winner = 2
    elif a + b == 4:
        if a < b :
            round_winner = 1
        else :
            round_winner = 2
    
    if round_winner == 1:
        print("You win this round")
    elif round_winner == 2:
        print("You loose this round")
    else:
        print("Draw")

    return round_winner

def computer_play() :
    """This defines what the computer plays in each round"""

    value_computer = ''

    b = random.randint(1,2)

    if b == 1:
        value_computer = 'Rock'
    elif b == 2:
        value_computer = 'Paper'
    elif b == 3:
        value_computer = 'Scissors'

    print('The computer played :', value_computer)

    return b 

## While condition

while score_player < 3 and score_computer < 3:
    
    value_player = input("Rock, Paper, Scissors ?:\n").lower()

    while value_player not in options:
        value_player = input("Please pick from Rock, Paper, Scissors :\n")
    else:
        print('You played ', value_player,' !')
        
    if value_player == 'rock':
        a = 1
    elif value_player == 'paper':
        a = 2
    elif value_player == 'scissors':
        a = 3

    b = computer_play()

    round_winner = results(a,b)

    if round_winner == 1:
        score_player += 1
    elif round_winner == 2:
        score_computer += 1
    
    print('Your score: ',score_player,' vs. Computer score ', score_computer)

if score_player == 3:
    print('You win !!!')
else:
    print('The computer wins :)')