function arrayPlusArray(arr1, arr2) {
  arr1 = arr1.reduce((acc, val) => acc + val);
  arr2 = arr2.reduce((acc, val) => acc + val);
  return arr1 + arr2;
}
