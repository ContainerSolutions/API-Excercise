const {QueryFile} = require('pg-promise');
const {join: joinPath} = require('path');

module.exports = {
    add: sql('add.sql'),
    update: sql('update.sql')
};

// Helper for linking to external query files;
function sql(file) {

    const fullPath = joinPath(__dirname, file);

    const options = {
        minify: true
    };

    const qf = new QueryFile(fullPath, options);

    if (qf.error) {
        // Something is wrong with our query file
        // Testing all files through queries can be cumbersome,
        // so we also report it here, while loading the module:
        console.error(qf.error);
    }

    return qf;
}