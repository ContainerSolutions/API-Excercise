"""
A Simple Pythoin Script That Tests the Specifications of a Submitted API Test
"""
import unittest
import requests
import pytest
import os
import json

TEST_HOST = os.getenv("TEST_HOST", "http://api.awesome")
TEST_HEADERS={
    "Content-Type": "application/json",
    "Accept": "application/json"
}

class APITester(unittest.TestCase):
    """
    The Test Class That Runs each of the tests
    """

    def test_list_passengers(self):
        """
        Runs a Get against the endpoint to list all passengers
        """
        res = requests.get(
            f"{TEST_HOST}/people",
            headers={
                "Content-Type": "application/json",
                "Accept": "application/json"
            }
        )
        self.assertTrue(
             res.status_code == 200,
             f"Expected GET to return 200 response but got {res.status_code}"
        )
        self.assertTrue(
            len(res.json()) > 0,
            "Expected a response with a list of People"
        )
        self.check_person_types(res.json()[0])

    def test_get_passenger(self):
        """
        Runs a GET for single passenger
        """
        res = requests.get(
            f"{TEST_HOST}/people",
            headers={
                "Content-Type": "application/json",
                "Accept": "application/json"
            }
        )
        user_uuid = res.json()[0].get("uuid")
        res = requests.get(
            f"{TEST_HOST}/people/{user_uuid}",
            headers={
                "Content-Type": "application/json",
                "Accept": "application/json"
            }
        )
        self.assertTrue(
             res.status_code == 200,
             f"Expected GET to return 200 response but got {res.status_code}"
        )
        self.check_person_types(res.json())

    def test_create_passenger(self):
        """
        Runs a POST to create a Passenger
        """
        data = {
            'survived': False,
            'passengerClass': 3,
            'name': 'Mr. Test McExampleFace',
            'sex': 'female',
            'age': 22,
            'siblingsOrSpousesAboard': 1,
            'parentsOrChildrenAboard': 0,
            'fare': 7.25
        }
        res = requests.post(
            f"{TEST_HOST}/people",
            data=json.dumps(data),
            headers=TEST_HEADERS
        )
        self.assertTrue(
             res.status_code == 201,
             f"Expected GET to return 201 response but got {res.status_code}"
        )
        user_uuid = res.json().get('uuid')
        res = requests.get(
            f"{TEST_HOST}/people/{user_uuid}",
            headers={
                "Content-Type": "application/json",
                "Accept": "application/json"
            }
        )
        self.assertTrue(
             res.status_code == 200,
             f"Expected GET to return 201 response but got {res.status_code}"
        )
        self.assertTrue(
             data.get("name") == res.json().get("name"),
            "Name of created User does not match"
        )
        self.check_person_types(res.json())



    def test_put_passenger(self):
        """
        Runs a PUT to update a passenger
        """
        data = {
            'survived': False,
            'passengerClass': 3,
            'name': 'Mrs. Foo Barrington',
            'sex': 'female',
            'age': 22,
            'siblingsOrSpousesAboard': 1,
            'parentsOrChildrenAboard': 0,
            'fare': 7.25
        }
        res = requests.post(
            f"{TEST_HOST}/people",
            data=json.dumps(data),
            headers=TEST_HEADERS
        )
        self.assertTrue(
             res.status_code == 201,
             f"Expected GET to return 201 response but got {res.status_code}"
        )
        user_uuid = res.json().get('uuid')
        data["age"] = 33
        requests.put(
            f"{TEST_HOST}/people/{user_uuid}",
            data=json.dumps(data),
            headers=TEST_HEADERS
        )
        res = requests.get(
            f"{TEST_HOST}/people/{user_uuid}",
            headers={
                "Content-Type": "application/json",
                "Accept": "application/json"
            }
        )
        self.assertTrue(
             res.status_code == 200,
             f"Expected GET to return 201 response but got {res.status_code}"
        )
        self.assertTrue(
             res.json().get("age") == 33,
            "Name of created User does not match"
        )
        self.check_person_types(res.json())

    def test_delete_passenger(self):
        """
        Runs a PUT to update a passenger
        """
        data = {
            'survived': False,
            'passengerClass': 3,
            'name': 'Mrs. Buzzword Bingo',
            'sex': 'female',
            'age': 22,
            'siblingsOrSpousesAboard': 1,
            'parentsOrChildrenAboard': 0,
            'fare': 7.25
        }
        res = requests.post(
            f"{TEST_HOST}/people",
            data=json.dumps(data),
            headers=TEST_HEADERS
        )
        self.assertTrue(
             res.status_code == 201,
             f"Expected GET to return 201 response but got {res.status_code}"
        )
        user_uuid = res.json().get('uuid')
        data["age"] = 33
        res = requests.delete(
            f"{TEST_HOST}/people/{user_uuid}",
            data=json.dumps(data),
            headers=TEST_HEADERS
        )
        self.assertTrue(
             res.status_code == 204,
             f"Expected GET to return 204 response but got {res.status_code}"
        )
        self.check_person_types(res.json())

    def check_person_types(self, person):
        """
        Runs a check on all the fields of a person
        """
        self.assertTrue(
            isinstance(person.get("survived"), bool),
            "Expected Type of field survived to be Boolean"
        )
        self.assertTrue(
            isinstance(person.get("passengerClass"), int),
            "Expected Type of field passengerClass to be Integer"
        )
        self.assertTrue(
            isinstance(person.get("name"), str),
            "Expected Type of field name to be String"
        )
        self.assertTrue(
            isinstance(person.get("sex"), str),
            "Expected Type of field sex to be String"
        )
        self.assertTrue(
            isinstance(person.get("age"), int),
            "Expected Type of field age to be Integer"
        )
        self.assertTrue(
            isinstance(person.get("siblingsOrSpousesAboard"), int),
            "Expected Type of field siblingsOrSpousesAboard to be Integer"
        )
        self.assertTrue(
            isinstance(person.get("parentsOrChildrenAboard"), int),
            "Expected Type of field parentsOrChildrenAboard to be Integer"
        )
        self.assertTrue(
            isinstance(person.get("fare"), float),
            "Expected Type of field fare to be Float"
        )
