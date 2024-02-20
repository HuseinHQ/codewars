const areaOrPerimeter = function (l, w) {
  if (l === w) {
    return l * w;
  } else {
    return (l + w) * 2;
  }
};

// cara cepat
// const areaOrPerimeter = (l, w) => (l === w ? l * w : (l + w) * 2);

console.log(areaOrPerimeter(3, 3));
console.log(areaOrPerimeter(6, 10));
