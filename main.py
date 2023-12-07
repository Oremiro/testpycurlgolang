import pycurl_requests as requests
import pycurl
import time

url = "https://aws.okx.com/api/v5/public/time"
num_requests = 201

# Create a session
session = requests.Session()
session.curl.setopt(pycurl.HTTP_VERSION, pycurl.CURL_HTTP_VERSION_2_0)
sum_time = 0
total_duration = 0
for i in range(1, num_requests + 1):
    start_time = time.time()

    # Use the session to make the request
    response = session.get(url)

    elapsed_time = time.time() - start_time
    if (i != 1):
        total_duration += elapsed_time
    print(f"Attempt {i}: HTTP GET request to {url} completed in {elapsed_time} seconds")

# Close the session to free up resources (optional)
average_duration = total_duration / (num_requests - 1)
print(f"\npython: Average time for {num_requests} requests: {average_duration} seconds")
session.close()
