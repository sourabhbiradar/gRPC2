# Max number

Bi-Directional RPC

client sends one number ; server asserts `maximum` among current & privous max number ; returns `set of max numbers`
(No Duplicates)

`Practice Problem` :
client : (1,5,3,6,2,17)
server : (1,5,6,17)

`Explanation` : 
0<1 ,1<5, 5>3, 5<6, 6>2, 6<17 = 1,5,5,6,6,17 = (1,5,6,17)