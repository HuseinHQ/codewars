function RecursiveRandomNumberGenerator(n, seed = 1) {
  const a = 1103;
  const c = 1234;
  const m = Math.pow(2, 31);

  if (n === 0) return seed;

  const previousNumber = RecursiveRandomNumberGenerator(n - 1, seed);

  const newNumber = (a * previousNumber + c) % m;
  return newNumber;
}

console.log(RecursiveRandomNumberGenerator(100));
