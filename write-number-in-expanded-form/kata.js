function expandedForm(num) {
  num = String(num);
  const arr = [];
  for (i = 0; i < num.length; i++) {
    if (num[i] != 0) {
      arr.push(num[i]);
    } else {
      arr.push('');
    }
  }

  arr.map((char, i) => {
    if (char) {
      char = +char;
      return char * (1 + Math.pow(10, arr.length - 1 - i));
    }
  });

  console.log(arr);
}

expandedForm(30704);
