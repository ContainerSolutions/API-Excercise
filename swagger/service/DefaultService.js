'use strict';
var queries = require('../queries.js');


/**
 * Add a person to the database
 *
 * person PersonData 
 * returns Person
 **/
exports.people_add = function(person) {
  return queries.create_person(person);
}


/**
 * Get a list of all people
 *
 * returns People
 **/
exports.people_list = function() {
  return queries.get_all_people();
}


/**
 * Delete this person
 *
 * uuid UUID 
 * no response value expected for this operation
 **/
exports.person_delete = function(uuid) {
  return queries.delete_person(uuid)
}


/**
 * Get information about one person
 *
 * uuid UUID 
 * returns Person
 **/
exports.person_get = function(uuid) {
  return queries.get_one_person(uuid);
}


/**
 * Update information about one person
 *
 * uuid UUID 
 * person PersonData 
 * returns Person
 **/
exports.person_update = function(uuid,person) {
  return queries.update_person(uuid, person);
}

