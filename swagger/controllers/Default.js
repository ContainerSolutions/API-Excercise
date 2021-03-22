'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

// Takes array of objects with keys in lower snake
// case and converts keys to camel case; then,
// keys 'id' and 'passengerName' are renamed
// to 'uuid' and 'name', respectively.
function format_outgoing_results (results) {
  function rename_keys(result) {
    var renamedKeys = {};
    for (var property in result) {
      if (result.hasOwnProperty(property)) {
        function to_camel_case (key) {
          return key[1].toUpperCase();
        }
        var renamedKey = property.replace(/(\_\w)/g, to_camel_case);
        renamedKeys[renamedKey] = result[property];
      }
    }

    renamedKeys.uuid = renamedKeys.id;
    delete renamedKeys.id;
    renamedKeys.name = renamedKeys.passengerName;
    delete renamedKeys.passengerName;

    function change_type(renamedResult) {
      renamedResult.age = parseInt(renamedResult.age);
      renamedResult.fare = parseFloat(renamedResult.fare);
      return renamedResult;
    }

    return change_type(renamedKeys);
  }
  return results.map(rename_keys);
}

// Takes array of object with keys in camel case
// and converts keys to lower snake case; then,
// keys 'uuid' and 'name' are renamed to
// 'id' and 'passengerName', respectively.
function format_incoming_results (results) {
  function rename_keys(result) {
    var renamedKeys = {};
    for (var property in result) {
      if (result.hasOwnProperty(property)) {
        function to_lower_snake_case (key) {
          return `_${key.toLowerCase()}`;
        }
        var renamedKey = property.replace(/[A-Z]/g, to_lower_snake_case);
        renamedKeys[renamedKey] = result[property];
      }
    }

    renamedKeys.passenger_name = renamedKeys.name;
    delete renamedKeys.name;
    if (renamedKeys.hasOwnProperty('uuid')) {
      renamedKeys.id = renamedKeys.uuid;
      delete renamedKeys.uuid;
    }
    return renamedKeys;
  }
  return results.map(rename_keys);
}

module.exports.people_add = function people_add (req, res, next) {
  var person = req.swagger.params['person'].value;
  Default.people_add(format_incoming_results([person])[0])
    .then(function (response) {
      utils.writeJson(res, format_outgoing_results([response])[0], 201);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.people_list = function people_list (req, res, next) {
  Default.people_list()
    .then(function (response) {
      utils.writeJson(res, format_outgoing_results(response));
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.person_delete = function person_delete (req, res, next) {
  var uuid = req.swagger.params['uuid'].value;
  Default.person_delete(uuid)
    .then(function (response) {
      if (response.rowCount === 0) {
        utils.writeJson(res, {}, 404);
        return;
      }
      utils.writeJson(res, {}, 204);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.person_get = function person_get (req, res, next) {
  var uuid = req.swagger.params['uuid'].value;
  Default.person_get(uuid)
    .then(function (response) {
      utils.writeJson(res, format_outgoing_results([response])[0]);
    })
    .catch(function (response) {
      utils.writeJson(res, {}, 404);
    });
};

module.exports.person_update = function person_update (req, res, next) {
  var uuid = req.swagger.params['uuid'].value;
  var person = req.swagger.params['person'].value;
  Default.person_update(uuid,format_incoming_results([person])[0])
    .then(function (response) {
      utils.writeJson(res, format_outgoing_results([response])[0]);
    })
    .catch(function (response) {
      utils.writeJson(res, {}, 404);
    });
};
