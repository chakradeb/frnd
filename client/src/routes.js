import React from 'react';
import { Route, Switch } from 'react-router-dom';

import Home from "./pages/home";

export default (
    <Switch>
        <Route path="/"> <Home/> </Route>
    </Switch>
);
