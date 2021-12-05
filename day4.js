const fs = require('fs')

function getNumberDelim(strIn, delim)
{
    split = strIn.split(delim)
    numbers = new Array()
    for(var i=0;i<split.length;++i){
        num = parseInt(split[i])
        if (!isNaN(num))
            numbers.push(num) 
    }
    return numbers
}

function getNumbersCsv(strIn){
    return getNumberDelim(strIn, ',')
}

function getNumbersSpace(strIn){
    arr1 = strIn.split('  ')
    all = new Array()
    for(var i=0; i<arr1.length; ++i){
        all = all.concat(getNumberDelim(arr1[i], ' '))
    }
    return all
}

function parseBoard(data, startIdx){
    idx = startIdx
    row = data[idx]

    board = new Array()

    while(row != ''){
        line = getNumbersSpace(row)
        board.push(line)
        ++idx
        row = data[idx]
    }
    return [board, idx+1]
}

const data = fs.readFileSync('input4test', 'utf8').toString().split('\n');

/// GET INPUT DATA///
numbers = getNumbersCsv(data[0])
console.log(numbers)

var start = 2
var boards = new Array()
while(start < data.length){
    retval = parseBoard(data,start)
    boards.push(retval[0])
    start = retval[1]
}

console.log('done')