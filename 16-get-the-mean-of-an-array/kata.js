const getAverage = (marks) => Math.floor(marks.reduce((cur, acc) => cur + acc) / marks.length);
