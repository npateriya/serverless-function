import sys
import http.client, urllib.parse,  json

auth_header = {'Authorization':'Bearer YjFjMmQ2YTEtMzk2OC00YzVjLTk0NzItMzM0NTcwOWMxNjkyNDA5YzYyNWYtNTU5'}

if len(sys.argv) > 1 :
	headers={}
	headers["Content-type"] =  "application/json"
	headers["Authorization"] = "Bearer YjFjMmQ2YTEtMzk2OC00YzVjLTk0NzItMzM0NTcwOWMxNjkyNDA5YzYyNWYtNTU5"
	params={}
	params["text"] = "Hello from Cisco serverless"
	params["toPersonEmail"] = sys.argv[1]
	#params_enc = urllib.urlencode(params)
	conn = http.client.HTTPSConnection("api.ciscospark.com")
	print(json.dumps(params))    
	path = "/v1/messages"
	conn.request("POST", path, json.dumps(params), headers)
	r2 = conn.getresponse()
	print("status and reason")
	print(r2.status,"+", r2.reason)
	data2 = r2.read()
	print (data2)
	conn.close()
else :
	print('Please provide spark useremail xxx@cisco.com to send message, in cli use -p or --param')
