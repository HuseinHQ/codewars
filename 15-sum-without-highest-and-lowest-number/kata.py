def sum_array(arr):
  result = 0
  minFlag = False
  maxFlag = False

  if arr == None:
    return result
  
  for value in arr:
    if value == min(arr):
      if minFlag == False:
        minFlag = True
      else:
        result += value
    elif value == max(arr):
      if maxFlag == False:
        maxFlag = True
      else:
        result += value
    else:
      result += value
  
  return result

print(sum_array([6, 0, 1, 10, 10]))
