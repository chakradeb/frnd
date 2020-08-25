import React from "react";
import { Redirect, Route } from "react-router-dom";


const PrivateRoute = function ({children, ...rest}) {
    return (
        <Route
            { ...rest }
            render={
                () => sessionStorage.getItem("token") ? children : <Redirect to="/login"/>
            }
        />
    );
};

export default PrivateRoute;
