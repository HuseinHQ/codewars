def sum_mix(arr):
  total = 0
  for i in arr:
    total += int(i)
  return total

# Cara pendek
def sum_mix(arr):
  return sum(int(i) for i in arr)

# cara pendek 2
def sum_mix(arr):
  return sum(map(int, arr))

print(sum_mix([9, 3, '7', '3']))