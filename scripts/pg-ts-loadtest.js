import http from 'k6/http';
import { SharedArray } from 'k6/data';

const requests = new SharedArray('sample requests query', function(){
    return JSON.parse(open('/data/file_names.json'))
})

export const options = {
    discardResponseBodies: false,
    scenarios: {
      contacts: {
        executor: 'ramping-arrival-rate',
        startRate: 15,
        timeUnit: '1s',
        preAllocatedVUs: 10,
        //maxVUs: 20,
        stages: [
            { target: 20, duration: '1m' },
            { target: 40, duration: '1m' },
            { target: 60, duration: '1m' },
            // { target: 50, duration: '2m' },
        ],
      },
    },
};

export default function () {
    const requestNo = Math.floor(Math.random() * requests.length);
    let queryParam = requests[requestNo]
    let slicedStrQueryParam = queryParam
    //console.log("len", queryParam.length)
    if (queryParam.length > 8){
        slicedStrQueryParam =  queryParam.slice(3, 7)
    }
    
    const url = `http://192.168.1.4:8080/search?query=${slicedStrQueryParam}`
    //console.log("URl", url)
    const params = {
      headers: {
        'Content-Type': 'application/json',
      },
    };
  
    http.get(url);
}