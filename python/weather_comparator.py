import requests
import inquirer
import json
import datetime
from pprint import pprint

api_key = 'da81a1815b3c06427e12401a34de25d0'

questions = [
  inquirer.List('criteria',
                message="What matters most to you?",
                choices=['Length of day', 'Temperature', 'Rain'],
            ),
]

one_city = input("Enter the one city : ")
two_city = input("Enter the two city : ")
criteria = inquirer.prompt(questions)['criteria']
print(criteria)

one_url = "https://api.openweathermap.org/data/2.5/weather?q="+one_city+"&appid="+api_key
two_url = "https://api.openweathermap.org/data/2.5/weather?q="+two_city+"&appid="+api_key

one_weather_data = requests.get(one_url).json()
two_weather_data = requests.get(two_url).json()

if criteria == 'Temperature':
  one_value = round(one_weather_data['main']['temp'] - 273.15,1)
  two_value = round(two_weather_data['main']['temp'] - 273.15,1)
elif criteria == 'Length of day':
  one_value = str(datetime.timedelta(seconds=one_weather_data['sys']['sunset'] - one_weather_data['sys']['sunrise']))
  two_value = str(datetime.timedelta(seconds=two_weather_data['sys']['sunset'] - two_weather_data['sys']['sunrise']))
elif criteria == 'Rain':
  one_value = one_weather_data['weather'][0]['main']
  two_value = two_weather_data['weather'][0]['main']
else:
  print ("Error in criteria selection")

print(one_city," :",one_value," vs. ",two_city," :",two_value)