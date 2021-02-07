console.log('\x1b[36m%s\x1b[0m', 'NPM preinstall ....')

if (10 < process.versions.node.split('.')[0]) {
    console.log('\x1b[31m', 'Running NODEJS VERSION > 10 --> npm-shrinkwarp required\n')

    const fs = require('fs')

    var str = `{
    "dependencies": {
        "graceful-fs": {
            "version": "^4.0.0"
        }
    }
}
    `

    fs.writeFileSync('npm-shrinkwrap.json', str)
    if (!fs.existsSync('./autorest.common') || fs.readdirSync('./autorest.common').length === 0) {
        const proc = require('child_process')
        proc.execSync('git submodule update --init')
    }
    fs.copyFileSync('npm-shrinkwrap.json', './autorest.common/npm-shrinkwrap.json')
}
else {
    console.log('\x1b[32m', 'Running NODEJS VERSION == 10 (or lower) --> NO npm-shrinkwarp required\n')
}
