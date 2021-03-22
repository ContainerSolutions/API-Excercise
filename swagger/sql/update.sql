/*
    Updates existing record in the table titanic_passengers.
*/
UPDATE titanic_passengers SET 
    SURVIVED=${survived},
    PASSENGER_CLASS=${passenger_class}, 
    PASSENGER_NAME=${passenger_name}, 
    SEX=${sex}, 
    AGE=${age}, 
    SIBLINGS_OR_SPOUSES_ABOARD=${siblings_or_spouses_aboard}, 
    PARENTS_OR_CHILDREN_ABOARD=${parents_or_children_aboard}, 
    FARE=${fare}
 WHERE ID=${id}
RETURNING *