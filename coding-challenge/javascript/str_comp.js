let input = "ABAWDAWDWFWFDWADAWDFAEFAOIDHAWCLCILVUW";

const compressor = (inp) => {
  let output = "";
  let i = 0;
  
  while(inp.length !== 0) {
    let char = inp[i];
    let regex = new RegExp(char, "g");
    let matches = inp.match(regex);
    output += `${char}${matches.length === 1 ? '' : matches.length}`;
    inp = inp.replace(regex, '');
    i = (i + 1) % inp.length;
  }
  return output;
}

console.log(compressor(input))