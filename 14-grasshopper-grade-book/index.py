from math import floor

def get_grade(s1, s2, s3):
  avg = (s1 + s2 + s3) / 3
  if avg >= 90:
    return 'A'
  elif avg >= 80:
    return 'B'
  elif avg >= 70:
    return 'C'
  elif avg >= 60:
    return 'D'
  else:
    return 'F'
  
print(get_grade(90, 90, 90))

# cara cepat
def get_grade2(s1, s2, s3):
    return {6:'D', 7:'C', 8:'B', 9:'A', 10:'A'}.get(floor((s1 + s2 + s3) / 30), 'F')

print(get_grade2(90, 89, 90))