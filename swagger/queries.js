const sql = require('./sql');
const pgp = require('pg-promise')();
var config = {
    password: process.env.POSTGRES_PASSWORD,
    user: process.env.POSTGRES_USER,
    database: process.env.POSTGRES_DB,
    host: process.env.POSTGRES_HOST,
    port: process.env.POSTGRES_PORT
};
const db = pgp(config);

function get_all_people() {
    return db.any('select * from titanic_passengers');
}

function get_one_person(id) {
    return db.one('select * from titanic_passengers where ID = $1', id);
}

function create_person(person) {
    return db.one(sql.add, person);
}

function delete_person(id) {
    return db.result('DELETE FROM titanic_passengers WHERE ID = $1', id);
}

function update_person(id, person) {
    person.id = id;
    return db.one(sql.update, person);
}

module.exports = {
    get_all_people,
    get_one_person,
    create_person,
    delete_person,
    update_person
}