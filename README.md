# Go-ATM
This repository contains a basic ATM system that provides a seamless and secure experience for customers. The ATM system allows customers to easily and efficiently access their bank accounts, withdraw cash, deposit money, regenerate the pinand generate the bank transaction statements.

## Requirements
1)Golang <br />
2)MongoDB <br />
3)Gin Framework <br />
4)Thunder Client/Postman


## Operations Performed & Understanding Of Functions

#### Structure Of User Table
Name      string   
Age       int      
Gender    string   
Pin       string   
AccountNo int      
Balance   string   
Statement []string 

#### Start The Program
1)Move to Go-ATM directory ($ cd Go-ATM) <br />
2)go run main.go

#### Create Account
URL : (POST) http://localhost:9090/v1/user/create <br />
In order to create account you need to pass data through thunder client which will be taken in backend through gin context in function of createuser in controller package and then passed to service package in createuser and then is saved in database(Pin in kept encrypted in database).
![Image text](https://github.com/deepakyadav810/Go-ATM/blob/main/Images/createacc.png)

#### Deposit And Wthdraw Money
URL : (PATCH) http://localhost:9090/v1/user/update <br />
After passing data from quick charge it goes to DepositWithdraw in controller package and the it is passed to DepositWithdraw in service package and then it is updated in database.
![Image text](https://github.com/deepakyadav810/Go-ATM/blob/main/Images/depositwithdraw.png)

#### Transfer Amount From One Account To Another Account
URL : (PATCH) http://localhost:9090/v1/user/transfer/<Account No. to transfer amount> <br />
In this transfer account number should be written in Param and person from whose amount is debited in JSON body. Next again same process data will be passed from gin to Transfer function in controller. After this data is passed on to service package in DepostWithdraw function and changes are updated in database
![Image text](https://github.com/deepakyadav810/Go-ATM/blob/main/Images/transfer.png)

#### Change Pin
URL : (PATCH) http://localhost:9090/v1/user/updatepin <br />
Open quickcharge and write in json body new pin and account number and it will follow same process as above but function used here is ChangePin in both controller package and service package.
![Image text](https://github.com/deepakyadav810/Go-ATM/blob/main/Images/changepin.png)

#### Display Statement
URL : (GET) http://localhost:9090/v1/user/get/<Account No. to fetch details> <br />
Here gettransaction function is used to view the statement. In quick charge just need to write the account number in param.
![Image text](https://github.com/deepakyadav810/Go-ATM/blob/main/Images/statement.png)
