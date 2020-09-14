import axios from "axios";

import Config from "../configs/config";

export default axios.create({
    baseURL: `${Config.api.url}`,
})
