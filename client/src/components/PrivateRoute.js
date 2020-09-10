import React from "react";
import { Redirect, Route } from "react-router-dom";

const PrivateRoute = function ({children, ...rest}) {
    return (
        <Route
            { ...rest }
            render={
                () => localStorage.getItem("authToken") ? children : <Redirect to="/login"/>
            }
        />
    );
};

export default PrivateRoute;
