const input = "ABAWDAWDWFWFDWADAWDFAEFAOIDHAWCLCILVUW";

const compressor = (inp) => {
  console.log(`Input: ${inp}`)
  let holder = {};
  let compressed = "";
  
  for(let i in inp) {
    let char = inp[i];
    if(holder[char] >= 1) {
      holder[char] = holder[char] + 1;
      continue;
    }
    holder[char] = 1;
  }

  for(let j in holder) {
    if(holder[j] === 1){
      compressed += j;
      continue;
    }
    compressed += j + holder[j];
  }
  
  return compressed;
}

console.log(`Output: ${compressor(input)}`);