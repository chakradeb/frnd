import React from 'react';
import { Route, Switch } from 'react-router-dom';

import Home from "./pages/home";
import Page404 from "./pages/page404";

export default (
    <Switch>
        <Route path="/" exact> <Home/> </Route>
        <Route path="*"> <Page404 /> </Route>
    </Switch>
);
