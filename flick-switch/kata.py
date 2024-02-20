def flick_switch(lst):
    if len(lst) == 0: return []
    result = []
    flick = True
    for i in range(len(lst)):
        if lst[i] == 'flick':
            flick = not flick
        result.append(flick)
    return result

print(flick_switch(['codewars', 'flick', 'code', 'wars']))