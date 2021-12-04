var lastVal = 9999
var count = 0

const fs = require('fs')

try {
    const data = fs.readFileSync('input1', 'utf8').toString().split('\n');
    for (const line of data) {
        const num = parseInt(line)
        if (num>lastVal){
            count = count + 1 
        }
        lastVal = num
    }
  //console.log(data)
} catch (err) {
  console.error(err)
}
console.log(count)


/////

prev2 = 0
prev1 = 0
curr = 0

count =0
currLine = 0
lastAvg = 9999

try {
    const data = fs.readFileSync('input1', 'utf8').toString().split('\n');
    for (const line of data) {
        const num = parseInt(line)
        prev2 = prev1
        prev1 = curr
        curr = num
        var avg = (prev2+prev1+curr)/3

        if(currLine > 1){
            if (avg>lastAvg){
                count = count + 1 
            }
            lastAvg = avg
        }
        currLine = currLine +1
    }
  //console.log(data)
} catch (err) {
  console.error(err)
}
console.log(count)