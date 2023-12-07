import pycurl_requests as requests
import pycurl
import time

url = "https://aws.okx.com/api/v5/public/time"
num_requests = 201
timeout = 0.5  # seconds
# Create a session
session = requests.Session()
session.curl.setopt(pycurl.HTTP_VERSION, pycurl.CURL_HTTP_VERSION_2_0)
sum_time = 0
total_duration = 0
i = 0
while True:
    i += 1
    start_time = time.time()

    # Use the session to make the request
    response = session.get(url)

    elapsed_time = time.time() - start_time
    if (i != 1):
        total_duration += elapsed_time
        average_duration = total_duration / float(i - 1)
        print(f"[PY] Average time after {i - 1} requests: {average_duration} seconds\n")

    time.sleep(timeout)

# Close the session to free up resources (optional)
session.close()
