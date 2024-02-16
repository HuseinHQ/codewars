var min = function (list) {
  let min = list[0];
  list.map((num) => {
    num < min ? (min = num) : null;
  });
  return min;
};

var max = function (list) {
  let max = list[0];
  list.map((num) => {
    num > max ? (max = num) : null;
  });
  return max;
};

// cara cepat
const min = (list) => Math.min(...list);
const max = (list) => Math.max(...list);
