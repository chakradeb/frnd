import React from "react";
import { Redirect, Route } from "react-router-dom";


const PrivateRoute = function ({children, ...rest}) {
    return (
        <Route
            { ...rest }
            render={
                () => sessionStorage.getItem("auth_key") ? children : <Redirect to="/login"/>
            }
        />
    );
};

export default PrivateRoute;
