/*
    Inserts a new record into the table titanic_passengers.
*/
INSERT INTO titanic_passengers(SURVIVED, PASSENGER_CLASS, PASSENGER_NAME, SEX, AGE, SIBLINGS_OR_SPOUSES_ABOARD, PARENTS_OR_CHILDREN_ABOARD, FARE)
VALUES(${survived}, ${passenger_class}, ${passenger_name}, ${sex}, ${age}, ${siblings_or_spouses_aboard}, ${parents_or_children_aboard}, ${fare})
RETURNING *