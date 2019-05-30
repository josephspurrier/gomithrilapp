/* eslint-disable no-console */
const { exec } = require('child_process')

module.exports = function globalSetup() {
  exec('bash $GOPATH/bash/resetdb.sh', (err, stdout, stderr) => {
    if (err) {
      // node couldn't execute the command
      return
    }

    // the *entire* stdout and stderr (buffered)
    console.log(`stdout: ${stdout}`)
    console.log(`stderr: ${stderr}`)
  })
}
