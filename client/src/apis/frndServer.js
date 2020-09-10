import axios from "axios";

const Config = require('../configs/config')();

export default axios.create({
    baseURL: `http://${Config.api.host}:${Config.api.port}`,
})
