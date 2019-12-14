'use strict';

const fs = require('fs');
const path = require('path');
const async = require('async');

const INPUT_PATH = path.join(process.cwd(), process.argv[2]);
const OUTPUT_PATH = path.join(process.cwd(), process.argv[3]);

const separators = ['. ', '? ', '! '];
const lineEndings = ['.', '?', '!'];

const countSentences = (counter, data) => {
  separators.forEach((separator) => {
    for (let index = 0; index < data.length; ) {
      data = data.replace('\r', ' ');
      data = data.replace('\n', ' ');
      index = data.indexOf(separator, index);
      if (index !== -1) {
        index++;
        ++counter;
      } else if (index === -1) {
        index = data.length;
      }
    }
  })
  lineEndings.forEach((lineEnding) => {
    let index = data.length - 2;
    index = data.indexOf(lineEnding, index);
    if (index !== -1) {
     ++counter;
    }
  })
  return counter;
}

const readFile = (fileName, counter) => {
  const filePath = path.join(INPUT_PATH, fileName);
  if (!filePath.startsWith(INPUT_PATH)) {
    console.log(`Can't be read: ${name}`);
    return null;
  }
  const readStream = fs.createReadStream(filePath, { highWaterMark: 128 });
  return readStream;
};

const readMultipleFiles = folderPath => {
  fs.readdir(folderPath, (err, files) => {
    async.each(files, (fileName, callback) => {
      let counter = 0;
      const stream = readFile(fileName, counter);
      stream.on('data', data => {
        counter = countSentences(counter, data.toString());
      });
      stream.on('end', () => {
        writeFile(fileName, counter);
        callback();
      });
    });
    console.log('Total number of processed files: ' + files.length);
  });
}

const writeFile = (fileName, counter) => {
  if (!fs.existsSync(OUTPUT_PATH)) {
    fs.mkdirSync(OUTPUT_PATH);
  }
  const outputFileName = OUTPUT_PATH + '/' + path.parse(fileName).name + '.res';
  fs.writeFile(outputFileName, counter, (err) => {
    if (err) {
      console.error(err);
    }
  })
}

readMultipleFiles(INPUT_PATH);
