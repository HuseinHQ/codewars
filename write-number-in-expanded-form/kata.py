from math import pow

def expanded_form(num):
  numStr = str(num)
  numArr = list(numStr)
  result = []
  for indeks, number in enumerate(numArr):
    if number != '0':
      result.append(str(int(number) * int(pow(10, len(numArr) - 1 - indeks))))
  return ' + '.join(result)

def expanded_form(num):
  num = str(num)
  return ' + '.join(x + '0' * (len(num) - i - 1) for i, x in enumerate(num) if x != '0')

print(expanded_form(20304))  # Outputs: '20000 + 300 + 4'