SELECT
       count(1)
FROM
	norman_user_property_mark_result 
WHERE
	property_id = 2 
	AND JSON_CONTAINS ( property_result -> '$[*]', '"是"', '$' );