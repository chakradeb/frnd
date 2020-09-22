import React from 'react';
import { Route, Switch } from 'react-router-dom';

import Home from "./pages/home";
import Profile from "./pages/profile";
import Page404 from "./pages/page404";
import LoginAndSignup from "./components/LoginAndSignup";
import PrivateRoute from "./components/PrivateRoute";

export default (
    <Switch>
        <Route path="/login"> <LoginAndSignup/> </Route>
        <PrivateRoute path="/profile/:id"> <Profile/> </PrivateRoute>
        <PrivateRoute path="/" exact> <Home/> </PrivateRoute>
        <Route path="*"> <Page404 /> </Route>
    </Switch>
);
