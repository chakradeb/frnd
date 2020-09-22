import axios from "axios";

import Config from "../configs/config";

let instance = axios.create({
    baseURL: `${Config.api.url}`,
})

instance.interceptors.request.use((config) => {
    const authToken = localStorage.getItem('X-AUTH-TOKEN');
    const refreshToken = localStorage.getItem('X-REFRESH-TOKEN');
    const username = localStorage.getItem('username');
    config.headers.Authorization =  authToken ? authToken : '';
    config.headers.Refresh =  refreshToken ? refreshToken : '';
    config.headers.Username = username ? username : '';
    return config;
},error => {
    Promise.reject(error);
});

instance.interceptors.response.use((response) => {
    return response;
}, function (error) {
    const originalReq =  error.config;
    if (error.response.status === 401 && !originalReq._retry) {
        originalReq._retry = true;
        const refreshToken = localStorage.getItem('X-REFRESH-TOKEN');
        return instance.post('/api/extend', {
            refreshToken: refreshToken,
        }).then(res => {
            if(res.status === 201) {
                localStorage.setItem('X-AUTH-TOKEN', res.data.accessToken)
                return instance(originalReq);
            }
        })
    }
    return Promise.reject(error);
});

export default instance;
