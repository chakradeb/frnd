import React, { Component } from "react";

import frndServer from "../apis/frndServer";

class LoginPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: '',
        };
        this.submitAction = this.submitAction.bind(this);
    }

    submitAction() {
        frndServer.post('/api/login', {
            username: this.state.username,
            password: this.state.password
        }, {
            headers: {
                "Content-Type": "application/json",
            }
        }).then(function (res) {
            localStorage.setItem("authToken", res.data.token)
            localStorage.setItem("username", res.data.username)
            frndServer.defaults.headers.common["authToken"] = res.data.token
            window.location.href = "/profile/abc";
        }).catch(function (err) {
            console.error("login failed: ", err)
        })
    }

    render() {
        return (
            <div>
                <h1>Login Page</h1>
                <input
                    id="username"
                    placeholder="enter your username"
                    required={true}
                    onChange={
                        (e) => this.setState({username: e.target.value})
                    }
                />
                <input
                    id="password"
                    type="password"
                    placeholder="enter your password"
                    required={true}
                    onChange={
                        (e) => this.setState({password: e.target.value})
                    }
                />
                <button onClick={this.submitAction}>Submit</button>
            </div>
        )
    }
}

export default LoginPage;
