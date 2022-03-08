((msg) => {
  console.log('encoded msg: ', msg);
  let len = msg.length;
  /* Fastest way to construct array (fill 0) with small elements */
  let a = new Array(len + 1); for (let i = 0; i < len + 1; ++i) a[i] = 0;
  
  console.time('decode');
  /* Loop through the encrypt string */
  for (let i = 0; i < len; ++i) {
    /* Reverse loop from the position of current encrypt string back to position 0 */
    for (let j = i + 1; j > 0; --j) {
      /* increase the value of prev element if encrypt string is L */
      if (msg[j-1] === 'L') {
        if (a[j-1] <= a[j]) {
          do {
            ++a[j-1];
          } while (a[j-1] <= a[j]);
        }
      /* increase the value of current element if encrypt string is R */
      } else if (msg[j-1] === 'R') {
        if (a[j-1] >= a[j]) {
          do {
            ++a[j];
          } while (a[j-1] >= a[j]);
        }
      /* assign the value of the highest to other if encrypt is = */
      } else {
        if (a[j-1] < a[j]) {
          a[j-1] = a[j];
        } else {
          a[j] = a[j-1];
        }
      }
    }
  }
  console.timeEnd('decode');
  console.log('decoded message: ', a);
  console.log('sum: ', a.reduce((p, c) => p + c))
  
})('LRRL=');