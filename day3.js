const fs = require('fs')

function binStrToNum(binString){
    var num = 0
    for(var i=0;i<binString.length; ++i){
        if(binString[i] == '1')
            num = num + Math.pow(2,binString.length-i-1)
    }
    return num
}

function getColumnSum(data, idx, column){
    columnSum = 0
    for(var lineIdx=0; lineIdx<idx.length; lineIdx++)
    {
        if(data[idx[lineIdx]].substr(column,1) == '1'){
            ++columnSum
        }
    }
    return columnSum
}

function printLines(data,idx){
    console.log('--------------')
    for(var i=0; i<idx.length; ++i)
        console.log(data[idx[i]])
    console.log('--------------')
}

try {
    const data = fs.readFileSync('input3', 'utf8').toString().split('\n');
    const lineCount = data.length
    var columns = new Array(data[0].length).fill(0)

    //count column values
    for (const line of data) {
        for(var ii=0; ii<line.length;++ii)
        {
            if (line.substr(ii,1) == '1'){
                columns[ii] = columns[ii] + 1
            }    
        }
    }

    var gammaRate = 0
    var epsilonRate = 0
    for(var ii=0; ii<columns.length;++ii){
        if(columns[ii]>lineCount/2){
            gammaRate = gammaRate + Math.pow(2,columns.length-ii-1)
        }
        else{
            epsilonRate = epsilonRate + Math.pow(2,columns.length-ii-1)
        }
    }

} catch (err) {
  console.error(err)
}


console.log(columns)
console.log(gammaRate)
console.log(epsilonRate)
console.log(gammaRate*epsilonRate)

////////////

try {
    const data = fs.readFileSync('input3', 'utf8').toString().split('\n');
    
    const lineCount = data.length

    var idxOxy = Array.from(Array(lineCount).keys())
    var idxCo2 = Array.from(Array(lineCount).keys())

    for(var column=0; column<columns.length; ++column)
    {
        columnSumOxy = getColumnSum(data,idxOxy, column)
        columnSumCo2 = getColumnSum(data,idxCo2, column)
        const mostCommonBitOxy = columnSumOxy>=idxOxy.length/2
        const leastCommonBitCo2 = columnSumCo2<idxCo2.length/2

        //Oxy
        if (idxOxy.length>1){
            var newIdxOxy = new Array()
            for(var lineIdx=0; lineIdx<idxOxy.length; lineIdx++)
            {
                const val = data[idxOxy[lineIdx]].substr(column,1) == '1'
                if( mostCommonBitOxy == val){
                    newIdxOxy.push(idxOxy[lineIdx])
                }
            }
            idxOxy = newIdxOxy
        }

        //CO2
        if (idxCo2.length>1){
            var newIdxCo2 = new Array()
            for(var lineIdx=0; lineIdx<idxCo2.length; lineIdx++)
            {
                const val = data[idxCo2[lineIdx]].substr(column,1) == '1'
                if( leastCommonBitCo2 == val){
                    newIdxCo2.push(idxCo2[lineIdx])
                }
            }
            idxCo2 = newIdxCo2
            //printLines(data,idxCo2)

        }
    }
    const oxyVal = binStrToNum(data[idxOxy[0]])
    console.log("Oxy len: " + idxOxy.length)
    console.log("Oxy val: " + oxyVal)

    console.log("Co2 len:" + idxCo2.length)
    console.log("Co2 val:" + binStrToNum(data[idxCo2[0]]))

    console.log(binStrToNum(data[idxOxy[0]]) * binStrToNum(data[idxCo2[0]]))
} catch (err) {
  console.error(err)
}


