import React, {Component} from "react";
import {Tab, TabList, TabPanel, Tabs} from "react-tabs";

import 'react-tabs/style/react-tabs.css';
import './LoginAndSignup.css';

import LoginPage from "../pages/login";
import SignupPage from "../pages/signup";

class LandingPage extends Component {
    componentDidMount() {
        if(sessionStorage.getItem("token")) window.location.href = "/";
    }

    render() {
        return (
            <div className="container">
                <h1>Welcome!</h1>
                <Tabs>
                    <TabList>
                        <Tab>Login</Tab>
                        <Tab>Sign Up</Tab>
                    </TabList>

                    <TabPanel>
                        <LoginPage/>
                    </TabPanel>
                    <TabPanel>
                        <SignupPage/>
                    </TabPanel>
                </Tabs>
            </div>
        )
    }
}

export default LandingPage;
