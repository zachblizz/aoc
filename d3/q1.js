const input = require("./input");

// manhattan distance equation
// Math.abs(p1-q2) + Math.abs(p2-q1)

// always assume central point is (0,0)
function getDistFromCentralPort(point) {

}

function constructPoints(wire) {
  const points = [];

  for (const dirDist of wire) {
    const dir = dirDist[0];
    const dist = dirDist.substr(1, dirDist.length - 1);

    console.log({dir, dist});
    points.push()
  }

  return points;
}

function solve() {
  const wire1Points = constructPoints(input.simple.w1);
}

solve();
