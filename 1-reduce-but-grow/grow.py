def grow(array):
  result = 0
  for number in array:
    result += number
  return result

array = [1, 2, 3, 4]
print(grow(array))