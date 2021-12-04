const fs = require('fs')

var horizontal = 0
var vertical = 0

try {
    const data = fs.readFileSync('input2', 'utf8').toString().split('\n');
    for (const line of data) {
        const instruction = line.split(' ')
        const val = parseInt(instruction[1])
        switch (instruction[0]){
            case 'forward':
                horizontal = horizontal + val
                break
            case 'down':
                vertical = vertical + val
                break
            case 'up':
                vertical = vertical - val
                break
        }
    }
} catch (err) {
  console.error(err)
}
console.log(horizontal)
console.log(vertical)
console.log(horizontal*vertical)


////////////

var horizontal = 0
var vertical = 0
var aim = 0

try {
    const data = fs.readFileSync('input2', 'utf8').toString().split('\n');
    for (const line of data) {
        const instruction = line.split(' ')
        const val = parseInt(instruction[1])
        switch (instruction[0]){
            case 'forward':
                horizontal = horizontal + val
                vertical = vertical + aim*val
                break
            case 'down':
                aim = aim + val
                break
            case 'up':
                aim = aim - val
                break
        }
    }
} catch (err) {
  console.error(err)
}
console.log(horizontal)
console.log(vertical)
console.log(horizontal*vertical)