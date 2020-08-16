import React, { Component } from "react";

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
        window.alert("Login Successful");
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
