function sumArray(arr) {
  return arr.reduce((cur, acc) => cur + acc) - Math.min(...arr) - Math.max(...arr);
}

console.log(sumArray([1, 2, 3, 4, 4]));
