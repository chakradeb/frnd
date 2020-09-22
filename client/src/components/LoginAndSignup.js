import React, {Component} from "react";
import {Tab, TabList, TabPanel, Tabs} from "react-tabs";

import './LoginAndSignup.css';

import LoginPage from "../pages/login";
import SignupPage from "../pages/signup";

class LandingPage extends Component {
    componentDidMount() {
        if(localStorage.getItem("authToken")) window.location.href = "/";
    }

    render() {
        return (
            <React.Fragment>
                <div className="tab">
                <Tabs className="form-content">
                    <TabList className="nav nav-pills">
                        <Tab className="nav-item">
                            <div className="nav-link">Login</div>
                        </Tab>
                        <Tab className="nav-item">
                            <div className="nav-link">Sign Up</div>
                        </Tab>
                    </TabList>

                    <div className="tab-content">
                        <TabPanel>
                            <LoginPage/>
                        </TabPanel>
                        <TabPanel>
                            <SignupPage/>
                        </TabPanel>
                    </div>
                </Tabs>
                </div>
            </React.Fragment>
        )
    }
}

export default LandingPage;
