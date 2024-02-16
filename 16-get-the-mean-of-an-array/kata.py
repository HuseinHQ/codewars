from math import floor

def get_average(marks):
  total = 0
  for value in marks:
    total += value

  avg = total / len(marks)
  return floor(avg)

print(get_average([1, 2, 3, 4, 5]))