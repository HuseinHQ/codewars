function expandedForm(num) {
  num = String(num).split('');
  result = [];
  num.map((number, indeks) => {
    if (number != '0') {
      result.push(number * Math.pow(10, num.length - 1 - indeks));
    }
  });
  return result.join(' + ');
}

function expandedForm(num) {
  num = String(num).split('');
  const newNum = num.map((val, i) => val != 0 && String(val * Math.pow(10, num.length - 1 - i))).filter(Boolean);
  return newNum.join(' + ');
}

console.log(expandedForm(30704));
