function solution(string) {
  let result = '';
  for (const char of string) {
    char.toUpperCase() == char ? (result += ' ') : '';
    result += char;
  }
  return result;
}
console.log(solution('camelCasingTest'));

// cara regex
function solution(string) {
  return string.replace(/([A-Z])/g, ' $1');
}
