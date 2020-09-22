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
            localStorage.setItem("X-AUTH-TOKEN", res.data.accessToken)
            localStorage.setItem("X-REFRESH-TOKEN", res.data.refreshToken)
            localStorage.setItem("username", res.data.username)
            window.location.href = `/profile/${res.data.username}`;
        }).catch(function (err) {
            console.error("login failed: ", err)
        })
    }

    render() {
        return (
            <div>
                <div className="form-group">
                    <label>Username</label>
                    <input
                        className="form-control is-valid"
                        id="username"
                        placeholder="enter your username"
                        required={true}
                        onChange={
                            (e) => this.setState({username: e.target.value})
                        }
                    />
                </div>

                <div className="form-group">
                    <label>Password</label>
                    <input
                        className="form-control is-valid"
                        id="password"
                        type="password"
                        placeholder="enter your password"
                        required={true}
                        onChange={
                            (e) => this.setState({password: e.target.value})
                        }
                    />
                </div>
                <button className="btn btn-primary" style={{ backgroundColor: "#030406" }} onClick={this.submitAction}>Submit</button>
            </div>
        )
    }
}

export default LoginPage;
