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

function markNumbers(board, markers, number){
    for(var y=0; y<board.length; ++y){
        for(var x=0; x<board[y].length; ++x){
            if(board[y][x] == number)
                markers[y][x] = true
        }
    }
}

function checkWinner(markers){
    //check lines
    const allTrue = arr => arr.every( v => v === true )
    for(var y=0; y<markers.length; ++y){
        if(allTrue(markers[y]))
            return true
    }

    //check columns
    for(var x=0; x<markers[0].length; ++x){
        var winner = true
        for(var y=0;y<markers.length; ++y){
            winner &= markers[y][x]
        }
        if (winner)
            return true
    }
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

function calcScore(board,markers){
    score =0
    for(var x=0; x<board[0].length; ++x){
        for(var y=0;y<markers.length; ++y){
            if (!markers[y][x])
                score += board[y][x]
        }
    }
    return score
}

const data = fs.readFileSync('input4', 'utf8').toString().split('\n');

/// GET INPUT DATA///
roundNumbers = getNumbersCsv(data[0])
console.log(numbers)

var start = 2
var boards = new Array()
var markers = new Array()
while(start < data.length){
    retval = parseBoard(data,start)
    boards.push(retval[0])
    const yDim = retval[0].length
    const xDim = retval[0][0].length

    mark = new Array(yDim)
    for (var i=0; i<yDim;++i){
        mark[i] = new Array(xDim).fill(false)
    }
    markers.push(mark)
    start = retval[1]
}

/// Play game
for(var round=0;round<roundNumbers.length;++round){
    for(var board=0; board<boards.length; ++board){
        markNumbers(boards[board],markers[board], roundNumbers[round])
        console.log(roundNumbers[round])
        console.log(markers[board])
    
        if(checkWinner(markers[board]))
        {
            console.log("Score: " + calcScore(boards[board],markers[board])*roundNumbers[round])
            return
        }
            
    }
   
}
