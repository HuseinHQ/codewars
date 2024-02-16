//return the total number of smiling faces in the array
function countSmileys(arr) {
  const eyes = [':', ';'];
  const noses = ['-', '~'];
  const mouths = [')', 'D'];

  let count = 0;
  for (const i of arr) {
    if (!eyes.includes(i[0])) {
      continue;
    }
    if (i.length === 2) {
      if (!mouths.includes(i[1])) {
        continue;
      }
    } else {
      if (!noses.includes(i[1])) {
        continue;
      } else if (!mouths.includes(i[2])) {
        continue;
      }
    }
    count++;
  }

  return count;
}

console.log(countSmileys([':D', ':~)', ';~D', ':)']));
