import sys
import http.client, urllib.parse

if len(sys.argv) > 1 :
    params = urllib.parse.urlencode({'action': 'create', 'token': '776a7a6a4e6d67674c7658557641506a57444b4c5369654f586b634b7175424f5852544a4c62796975687173', 'numberToDial': sys.argv[1]})
    headers = {"Content-type": "application/x-www-form-urlencoded", "Accept" : "*/*"}
    conn = http.client.HTTPSConnection("api.tropo.com")
    print(params)    
    path = "/1.0/sessions" + "?" + params
    conn.request("GET", path)
    r2 = conn.getresponse()
    print("status and reason")
    print(r2.status,"+", r2.reason)
    data2 = r2.read()
    print (data2)
    conn.close()
else :
    print('Please providide cell number  format 1XXXYYYZZZZ to send message, in cli use -p or --param')

