
create table offers.hotel(
cm_offer_id varchar(50),
hotel_id varchar(30),
name varchar(50),
country varchar(10),
address varchar(100),
latitude float,
longitude float,
telephone varchar(50),
amenities JSON,        
description varchar(500),
room_count int,
currency varchar(20))

create TABLE offers.offer(
cm_offer_id varchar(50),
original_data  json,
capacity json,
number int,
price float,
currency varchar(10),
check_in varchar(10),
check_out varchar(10),
fees json)


create table offers.rate_plan(
cm_offer_id varchar(50),
 hotel_id varchar(30),
rate_plan_id varchar(10),
cancellation_policy json,
name varchar(30),
other_conditions json,
meal_plan varchar(20)
)


create table offers.room(
cm_offer_id varchar(50),
hotel_id varchar(30),
room_id varchar(20),
description varchar(100),
name varchar(100),
capacity json
 )
 -------------------------------

 --- insert into offers.mytable(name,data)values("shuhbam",'["name","frname"]');
                
 ---   SELECT JSON_EXTRACT(other_conditions, '$[1]') FROM offers.rate_plan where hotel_id= "BH~46456";           

--- SELECT JSON_EXTRACT(data, '$') FROM offers.mytable where id = 4;