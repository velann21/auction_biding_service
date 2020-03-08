Auction Biding System:
-------------------------------
 1. Auction Biding system is to perform the auction on the product which is available on Inventory system.

Currently Auction System supports:
———————————————————————------------
1. Users can bid on the product
2. User can see all the bids which is done by them.
3. We can see all the bids on specific product
4. We can see the winner bid on specific product
5. We cannot bid the same price which is already been bided

How Auction system announce the winner now:
————————————————————————————————------------
If No Bid is perfomed on system for more than 2 min the system removes the product from Auction system, And announce the result for that Product.


Features which should be added:
———————————————————————-----------
1. We should implement the Inventory system to maintain the product which should be available for Auction.
2. Auth service should be added to authenticate all the API call to the platform.
3. We should make the Real time system using the Websockets, Rather than just using the Rest API
4. Currently all the data is stored only in memory data structures, We should persist the data for reliable system
5. User should be able to search for the specific products.


Technical Description:
————————————————---------
1. User will get Authenticated first
2. User will get the list products available in the inventory for the biding.
3. Once he gets the list of products he start the biding (Concurrently multiple people can bid for the same product), We should be able to support all the concurrent bids, To manage the concurrent users I have used the Optimistic lock.
4. Once the Bid is placed first it will write the data into User Bids Map (Key is User_id in map and value is Slice of Object) (Map[string][]Object), This is very useful to get the all the bids placed by user.
5. Once the Bid is placed it also writes into separate map for Item Based Tracking (Key is product_id and Value is slice of Object which Heapify for ease of search to announce the Winner Bid)
6. We can now get the list bids placed by user using the user_id
7. We can get the list of bids on specific Item to see all the bids placed on the Item.
8. We can get the winner bid on Item.

DataStructure Used:
—————————————-------
1. Map, With slice as an value:
    The map I have used since the search is O(1) to find the key, I have chosen the slice as an value since I can Just append the value to it which is again an constant time
2. Max heap:
    I used Max heap so that write will be log n,  but to find the winner of the bid it just take an constant time
3. Log Files:
   I used the log files to persist the data to make it Durability.

Design Patterns Used:
———————————————------
1. Dependency Injection to avoid the tight couple between the objects
2. Factory Pattern to make ease of object creation
3. Singleton pattern with Double check mechanism to control the concurrency access of object.
4. Command Pattern
5. Chain of responsibility pattern
6. Single Responsibility principle
7. Interface segregation principle

Concurrency control:
——————————————-------
1. Used the channels to make the communication between two go routine, I used the go routine to make Producer, consumer pattern/ Command pattern.
2. Used the Optimistic Lock to make write fast.
3. Used the RW mutex to Avoid the unnecessary lock to the whole object.
4. Used the pessimistic Lock as well.

Project Structure Explanation:
————————————————————-----------
1. Cmd:
    This is the place app bootstrap starts 
2. Routes:
    All the app routes are placed under this directory
3. Controller: 
   All the routes entry point and its acts as an orchestrator for an endpoint
4. Entities:
    This has 2 parts 
        1. Request
                All the API request body are deserialised with help of this and validate the request body here first
        2. Response:
               All the response structs are kept here
5. Service:
    All the business logos are written here
6. Dao:
   All the data model and database queries are done here
7. Deployments:
    This has the Docker File, Make file, This Make file is very handy to release the our distribution to docker hub and to release the tag to GitHub
    Make file commands:
      Note: If you want to run the make commands go to the directory as follows and then run the commands: ./deployments/makefile/Makefile
      Make release will release the entire process.
8. Helpers:
     This is to keep all the helper code or shared code for the project;

Distributed Source:
—————————————--------
1. GitHub:
     https://github.com/velann21/auction_biding_service

2.Docker:
   docker pull singaravelan21/auction_biding:v0.0.3
    

GO VERSION:
------------
go 1.13


API Documents:
—————————-----

1.USERBID:
-----------
Method: POST

Request:

curl --location --request POST 'http://localhost:8080/api/v1/biding/usersBid' \
--header 'Content-Type: application/json' \
--data-raw '{
	"userID":"2",
	"productID":"1",
	"bidPrice":800.00,
	"timeStamp":"2020-02-01T21:39:02+00:00"
}'

Response:
{
    "status": "OK",
    "data": null
}

2.GETBIDBYUSER:
-------------
Method: GET

Request:

curl --location --request GET 'http://localhost:8080/api/v1/biding/getBidByUser?user_id=2'

Response: {
    "status": "OK",
    "data": [
        {
            "UserBids": [
                {
                    "userID": "2",
                    "productID": "2",
                    "bidPrice": 300,
                    "timeStamp": "2020-02-01T21:39:02+00:00"
                }
            ]
        }
    ]
}




3.GETBIDBYITEM
----------------
METHOD: GET
Request:

curl --location --request GET 'http://localhost:8080/api/v1/biding/getBidByItem?product_id=1'

Response:
{
    "status": "OK",
    "data": [
        {
            "ProductBids": [
                null,
                {
                    "userID": "2",
                    "productID": "2",
                    "bidPrice": 600,
                    "timeStamp": "2020-02-01T21:39:02+00:00",
                    "UUID": "e0be4fd6-562e-43d6-9b56-3abf3a86d3c0"
                },
                {
                    "userID": "2",
                    "productID": "2",
                    "bidPrice": 500,
                    "timeStamp": "2020-02-01T21:39:02+00:00",
                    "UUID": "d7c20781-1542-49d6-b332-b5fa2267f184"
                },
                 {
                    "userID": "2",
                    "productID": "2",
                    "bidPrice": 150,
                    "timeStamp": "2020-02-01T21:39:02+00:00",
                    "UUID": "e0f39b24-038b-429b-a727-c00df91e6b6e"
                }
            ]
        }
    ]
}

4.GETWINNINGBID:
-----------------
METHOD: GET
Request:
 
curl --location --request GET 'http://localhost:8080/api/v1/biding/getWinningBid?product_id=1'

Response:
{
    "status": "OK",
    "data": [
        {
            "WinnerBid": {
                "userID": "2",
                "productID": "2",
                "bidPrice": 800,
                "timeStamp": "2020-02-01T21:39:02+00:00",
                "UUID": "c1256bc8-093e-4659-bac6-34e206ab0264"
            }
        }
    ]
}


4.ListProducts:
----------------
Method: GET
Request:

curl --location --request GET 'http://localhost:8080/api/v1/inventory/products'


