function sumArray(arr) {
  if (!arr || arr.length <= 2) return 0;
  return arr.reduce((cur, acc) => cur + acc) - Math.min(...arr) - Math.max(...arr);
}

console.log(sumArray([1, 2, 3, 4, 4]));
