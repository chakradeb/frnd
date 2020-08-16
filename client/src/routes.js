import React from 'react';
import { Route, Switch } from 'react-router-dom';

import LoginAndSignup from "./hoc/LoginAndSignup";
import Home from "./pages/home";
import Profile from "./pages/profile";
import Page404 from "./pages/page404";

export default (
    <Switch>
        <Route path="/login"> <LoginAndSignup/> </Route>
        <Route path="/" exact> <Home/> </Route>
        <Route path="/profile/:id"> <Profile/> </Route>
        <Route path="*"> <Page404 /> </Route>
    </Switch>
);
