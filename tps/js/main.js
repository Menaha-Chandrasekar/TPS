

// import http from 'k6/http';
// import { sleep } from 'k6';
// export const options ={
//   vus : 100,
//   iterations : 10000
// }
// export default function () {
//   // Send an HTTP GET request to your Go server
//   const response = http.get('http://localhost:5000/api'); // Update the URL

 
//   if (response.status !== 200) {
//     console.error(`Request failed with status code: ${response.status}`);
//   }

//   // Sleep for 1 second (adjust as needed)
//  // sleep(1);
// }


import http from 'k6/http';
import { check } from 'k6';
export const options ={
  vus : 120,
  iterations : 1200
}
export default function () {
const url = 'http://localhost:5000/tokens'; // Update the URL to your API endpoint
const headers = {
'Content-Type': 'application/json',
Authorization: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.zWQD6InfmyvVA1foShp-DAq7PyoF80BrMHrAViGUi_I', // Replace with a valid token
};

const payload = JSON.stringify({
// Your JSON payload here
});

const response = http.post(url, payload, { headers });

check(response, {
'Status is 200': (r) => r.status === 200,
});
}