# -*- coding: utf-8 -*-

# A main script for the iox application template - PaaS - Python

import time
import httplib, urllib,  json

def call_spark():
    #curl  -H "Content-type: application/json" -X POST http://128.107.18.112:8888/function/spark/run -d '{"name":"spark", "runparams":["npateriya@gmail.com"]}'
    print("Calling spark")
    headers={}
    headers["Content-type"] =  "application/json"
    params={}
    params["runparams"] = ["npateriya@gmail.com"]
    conn = httplib.HTTPConnection("128.107.18.112",8888)
    print(json.dumps(params))    
    path = "/function/ioxspark/run"
    conn.request("POST", path, json.dumps(params), headers)
    r2 = conn.getresponse()
    print("status and reason")
    print(r2.status,"+", r2.reason)
    data2 = r2.read()
    print (data2)
    conn.close()
	
def main():
    while True:
        print('hello neelesh')
        call_spark()
        time.sleep(30)

	
if __name__ == '__main__':
    main()
