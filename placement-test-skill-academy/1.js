function findResult(n) {
  if (n === 0) return 0;
  let result = 3;
  let accumulator = 2;
  for (i = 0; i < n - 1; i++) {
    result += accumulator;
    accumulator += 2;
  }
  return result;
}

console.log(findResult(6));
