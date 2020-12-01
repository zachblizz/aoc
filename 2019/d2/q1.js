// opcodes
// 99 - finished
// unknown - something went wrong...
// 1 - add
// 2 - multiply
// 4 digit opcodes
// step by 4

const input = [1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,5,23,2,13,23,27,1,10,27,31,2,6,31,35,1,9,35,39,2,10,39,43,1,43,9,47,1,47,9,51,2,10,51,55,1,55,9,59,1,59,5,63,1,63,6,67,2,6,67,71,2,10,71,75,1,75,5,79,1,9,79,83,2,83,10,87,1,87,6,91,1,13,91,95,2,10,95,99,1,99,6,103,2,13,103,107,1,107,2,111,1,111,9,0,99,2,14,0,0];

function basicIntCodeComp(input) {
  for (i = 0; i < input.length; i += 4) {
    const op = input[i];
    const a = input[input[i+1]];
    const b = input[input[i+2]];
    const p = input[i+3];

    if (op === 99) {
      break;
    }

    input[p] = op === 1 ? a + b : a * b;
  }

  return input;
}

console.log(basicIntCodeComp(input)[0]);
